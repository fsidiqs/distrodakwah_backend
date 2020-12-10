package userclass

import (
	"encoding/json"
	"strings"
	"time"
)

type Birthdate time.Time

const (
	layoutISO = "2006-01-02"
)

var nilTime = (time.Time{}).UnixNano()

func (birthdate *Birthdate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(layoutISO, s)
	if err != nil {
		return err
	}
	*birthdate = Birthdate(t)
	return nil
}

func (birthdate Birthdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(birthdate)
}

// func (birthdate *Birthdate) MarshalJSON() ([]byte, error) {
// 	return []byte(fmt.Sprintf("\"%s\"", ubd.Time.Format(BirthdateLayout))), nil
// }

func (bd Birthdate) IsSet() bool {
	t := time.Time(bd)
	return t.UnixNano() != nilTime
}

func (bd Birthdate) Format() string {
	t := time.Time(bd)
	return t.Format(time.RFC3339)
}

// func (bd Birthdate) retur() time.Time {
// 	t := time.Time(bd)
// 	return t.Format(layoutISO)
// }
