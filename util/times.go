package util

import (
	"fmt"
	"time"
)

type DateTime struct {
	time.Time
}

func (t *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	s = s[1 : len(s)-1]
	result, err := time.Parse(time.DateTime, s)
	if err != nil {
		return err
	}
	t.Time = result
	return nil
}

func (t *DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format(time.DateTime))), nil
}
