package constant

import "time"

type OnlyTime struct {
	time.Time
}

func (t *OnlyTime) UnmarshalJSON(data []byte) error {
	parsedTime, err := time.Parse(`"15:04:05"`, string(data))
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}

func (t OnlyTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format("15:04:05") + `"`), nil
}
