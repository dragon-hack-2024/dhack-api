package util

import "time"

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	parsedTime, err := time.Parse("\"2006-01-02\"", str)
	if err != nil {
		return err
	}
	*ct = CustomTime(parsedTime)
	return nil
}
