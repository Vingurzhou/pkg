package satellite

import (
	"fmt"
	"math"

	satellite "github.com/joshuaferrara/go-satellite"
)

const GM = 398600.4418 // 地球引力常数，单位 km^3/s^2

// nadc.china-vo.org/astrodict
// https://satellitemap.space/zh/tle-calculator
// https://gitcode.com/open-source-toolkit/ac324
// type Satellite struct {
// 	Line1             string    // 1 65055U 25164A   25329.77566319 -.00034908  00000-0 -22699-2 0 00003
// 	Line2             string    // 2 65055 041.1321 355.6156 0004144 325.3028 090.7476 15.07206270017711
// 	MeanMotion        float64   // Mean Motion 平均运动 rev/day
// 	SemiMajorAxis     float64   // Semi-Major Axis 半长轴 单位km
// 	Eccentricity      string    // Eccentricity 偏心率
// 	Inclination       string    // Inclination 轨道倾角
// 	RAAN              string    // Right Ascension of the Ascending Node(RAAN) 升交点赤经
// 	ArgumentOfPerigee string    // Argument of Perigee 近地点幅角
// 	MeanAnomaly       string    // Mean Anomaly 平近点角
// 	EpochTime         time.Time // Epoch Date (UTC)历元时间
// }

// func NewSatellite(line1 string, line2 string) (*Satellite, error) {
// 	///////////////////
// 	meanMotion, err := strconv.ParseFloat(line2[52:63], 64)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// 开普勒第三定律推导
// 	semiMajorAxis := math.Pow(
// 		(GM*math.Pow(24*60*60/meanMotion, 2))/
// 			(4*math.Pow(math.Pi, 2)),
// 		1.0/3.0)
// 	////////////////////
// 	eccentricity := "0." + line2[26:33]
// 	inclination := line2[8:16]
// 	raan := line2[17:25]
// 	argumentOfPerigee := line2[34:42]
// 	meanAnomaly := line2[43:51]
// 	////////////////////
// 	epochYear, err := strconv.Atoi(line1[18:20])
// 	if err != nil {
// 		return nil, err
// 	}
// 	if epochYear < 57 {
// 		epochYear += 2000
// 	} else {
// 		epochYear += 1900
// 	}
// 	//////////////////////////////
// 	epochDay, err := strconv.ParseFloat(line1[20:32], 64)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// 拆整数天和小数天
// 	dayInt, dayFrac := math.Modf(epochDay)
// 	dayInt = dayInt - 1 // 因为1月1日算第1天
// 	// 将小数天精确转换成纳秒
// 	totalNanoseconds := dayFrac * 24 * 60 * 60 * 1e9
// 	duration := time.Duration(totalNanoseconds) * time.Nanosecond
// 	// 构建时间
// 	epochTime := time.Date(epochYear, 1, 1, 0, 0, 0, 0, time.UTC).
// 		AddDate(0, 0, int(dayInt)).
// 		Add(duration)
// 	////////////////////////
// 	return &Satellite{
// 		SemiMajorAxis:     semiMajorAxis,
// 		Eccentricity:      eccentricity,
// 		Inclination:       inclination,
// 		RAAN:              raan,
// 		ArgumentOfPerigee: argumentOfPerigee,
// 		MeanAnomaly:       meanAnomaly,
// 		Line1:             line1,
// 		Line2:             line2,
// 		MeanMotion:        meanMotion,
// 		EpochTime:         epochTime,
// 	}, nil
// }

type ClassicalOrbitalElements struct {
	MeanMotion        float64 // Mean Motion 平均运动 rev/day
	SemiMajorAxis     float64 // Semi-Major Axis 半长轴 单位km
	Eccentricity      float64 // Eccentricity 偏心率
	Inclination       float64 // Inclination 轨道倾角
	RAAN              float64 // Right Ascension of the Ascending Node(RAAN) 升交点赤经
	ArgumentOfPerigee float64 // Argument of Perigee 近地点幅角
	MeanAnomaly       float64 // Mean Anomaly 平近点角
}

func NewClassicalOrbitalElements(line1, line2 string) (*ClassicalOrbitalElements, error) {
	sat := satellite.TLEToSat(line1, line2, satellite.GravityWGS84)
	pos, vel := satellite.Propagate(sat, 2026, 1, 21, 1, 34, 9)
	fmt.Println(pos, vel)
	coe := ECIToCOE(pos, vel)
	// fmt.Println("a =", coe.SemiMajorAxis, "km")
	// fmt.Println("e =", coe.Eccentricity)
	// fmt.Println("i =", coe.Inclination*180/math.Pi, "deg")
	// fmt.Println("Ω =", coe.RAAN*180/math.Pi, "deg")
	// fmt.Println("ω =", coe.ArgumentOfPerigee*180/math.Pi, "deg")
	// fmt.Println("ν =", coe.MeanAnomaly*180/math.Pi, "deg")
	fmt.Printf("%+v", coe)
	return nil, nil
}

func ECIToCOE(r, v satellite.Vector3) ClassicalOrbitalElements {
	const mu = GM

	rmag := norm(r)
	vmag := norm(v)

	// 比角动量
	h := cross(r, v)
	hmag := norm(h)

	// 倾角
	i := math.Acos(h.Z / hmag)

	// 节点向量
	n := satellite.Vector3{X: -h.Y, Y: h.X, Z: 0}
	nmag := norm(n)

	// 偏心率向量
	evec := satellite.Vector3{
		X: (v.Y*h.Z-v.Z*h.Y)/mu - r.X/rmag,
		Y: (v.Z*h.X-v.X*h.Z)/mu - r.Y/rmag,
		Z: (v.X*h.Y-v.Y*h.X)/mu - r.Z/rmag,
	}
	e := norm(evec)

	// 半长轴
	a := 1 / (2/rmag - (vmag*vmag)/mu)

	// RAAN
	raan := math.Acos(n.X / nmag)
	if n.Y < 0 {
		raan = 2*math.Pi - raan
	}

	// 近地点幅角
	argp := math.Acos(dot(n, evec) / (nmag * e))
	if evec.Z < 0 {
		argp = 2*math.Pi - argp
	}

	// 真近点角
	nu := math.Acos(dot(evec, r) / (e * rmag))
	if dot(r, v) < 0 {
		nu = 2*math.Pi - nu
	}

	return ClassicalOrbitalElements{
		SemiMajorAxis:     a,
		Eccentricity:      e,
		Inclination:       i,
		RAAN:              raan,
		ArgumentOfPerigee: argp,
		MeanAnomaly:       nu,
	}
}
func dot(a, b satellite.Vector3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func cross(a, b satellite.Vector3) satellite.Vector3 {
	return satellite.Vector3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func norm(v satellite.Vector3) float64 {
	return math.Sqrt(dot(v, v))
}
