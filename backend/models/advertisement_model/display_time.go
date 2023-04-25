package advertisement_model

import "time"

type DisplayTime struct {
	StartDate string `json:"start_date" gorm:"column:start_date"`
	EndDate   string `json:"end_date" gorm:"column:end_date"`
}

func (d *DisplayTime) Validate() error {
	_, err := time.Parse(time.RFC3339, d.StartDate)
	if err != nil {
		return err
	}
	_, err = time.Parse(time.RFC3339, d.EndDate)
	if err != nil {
		return err
	}
	return nil
}
