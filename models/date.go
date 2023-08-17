package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// Date	has to be in the format YYYY-MM-DD.
type Date string

func (date *Date) Scan(value interface{}) (err error) {
	stringValue := value.(string)
	*date = Date(stringValue)
	return nil
}

func (date Date) Value() (driver.Value, error) {
	return string(date), nil
}

func (d *Date) UnmarshalJSON(bytes []byte) error {
	var date string
	err := json.Unmarshal(bytes, &date)
	if err != nil {
		return err
	}

	layout := "2006-01-02"
	_, err = time.Parse(layout, string(date))
	if err != nil {
		fmt.Println("Invalid date format:", err)
		return err
	}

	*d = Date(date)

	return nil
}
