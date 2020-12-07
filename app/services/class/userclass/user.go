package userclass

import (
	"fmt"
	"strings"
	"time"
)

type UserBirthDate struct {
	time.Time
}

const (
	UserBirthDateLayout = "1900-13-12 00:00:00"
)

var nilTime = (time.Time{}).UnixNano()

func (ubd *UserBirthDate) UnmarsahlJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ubd.Time = time.Time{}
		return
	}
	ubd.Time, err = time.Parse(UserBirthDateLayout, s)
	return
}

func (ubd *UserBirthDate) MarshalJSON() ([]byte, error) {
	if ubd.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ubd.Time.Format(UserBirthDateLayout))), nil
}

func (ubd *UserBirthDate) IsSet() bool {
	return ubd.UnixNano() != nilTime
}
