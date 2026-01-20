package satellite

import (
	"math"
	"strconv"
	"time"
)

const GM = 398600.4418 // 地球引力常数，单位 km^3/s^2

// nadc.china-vo.org/astrodict
// https://satellitemap.space/zh/tle-calculator
type Satellite struct {
	Line1             string    // 1 65055U 25164A   25329.77566319 -.00034908  00000-0 -22699-2 0 00003
	Line2             string    // 2 65055 041.1321 355.6156 0004144 325.3028 090.7476 15.07206270017711
	MeanMotion        float64   // Mean Motion 平均运动 rev/day
	SemiMajorAxis     float64   // Semi-Major Axis 半长轴 单位km
	Eccentricity      string    // Eccentricity 偏心率
	Inclination       string    // Inclination 轨道倾角
	RAAN              string    // Right Ascension of the Ascending Node(RAAN) 升交点赤经
	ArgumentOfPerigee string    // Argument of Perigee 近地点幅角
	MeanAnomaly       string    // Mean Anomaly 平近点角
	EpochTime         time.Time // Epoch Date (UTC)历元时间
}

func NewSatellite(line1 string, line2 string) (*Satellite, error) {
	///////////////////
	meanMotion, err := strconv.ParseFloat(line2[52:63], 64)
	if err != nil {
		return nil, err
	}
	// 开普勒第三定律推导
	semiMajorAxis := math.Pow(
		(GM*math.Pow(24*60*60/meanMotion, 2))/
			(4*math.Pow(math.Pi, 2)),
		1.0/3.0)
	////////////////////
	eccentricity := "0." + line2[26:33]
	inclination := line2[8:16]
	raan := line2[17:25]
	argumentOfPerigee := line2[34:42]
	meanAnomaly := line2[43:51]
	////////////////////
	epochYear, err := strconv.Atoi(line1[18:20])
	if err != nil {
		return nil, err
	}
	if epochYear < 57 {
		epochYear += 2000
	} else {
		epochYear += 1900
	}
	//////////////////////////////
	epochDay, err := strconv.ParseFloat(line1[20:32], 64)
	if err != nil {
		return nil, err
	}
	// 拆整数天和小数天
	dayInt, dayFrac := math.Modf(epochDay)
	dayInt = dayInt - 1 // 因为1月1日算第1天
	// 将小数天精确转换成纳秒
	totalNanoseconds := dayFrac * 24 * 60 * 60 * 1e9
	duration := time.Duration(totalNanoseconds) * time.Nanosecond
	// 构建时间
	epochTime := time.Date(epochYear, 1, 1, 0, 0, 0, 0, time.UTC).
		AddDate(0, 0, int(dayInt)).
		Add(duration)
	////////////////////////
	return &Satellite{
		SemiMajorAxis:     semiMajorAxis,
		Eccentricity:      eccentricity,
		Inclination:       inclination,
		RAAN:              raan,
		ArgumentOfPerigee: argumentOfPerigee,
		MeanAnomaly:       meanAnomaly,
		Line1:             line1,
		Line2:             line2,
		MeanMotion:        meanMotion,
		EpochTime:         epochTime,
	}, nil
}
