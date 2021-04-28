package main

import (
	"errors"
	"fmt"
	"time"
)

func validateInput(dateTime string) error {
	layout := "2006-01-02"
	_, err := time.Parse(layout, dateTime)
	return err
}
func convertDatetimeString(ngay, thang, nam *int) (dateTime string, err error) {
	if ngay == nil && thang == nil && nam == nil {
		return
	}

	if nam == nil {
		err = errors.New("Year is not entered")
		return
	}

	if thang == nil {
		err = errors.New("Month is not entered")
		return
	}
	if ngay == nil {
		err = errors.New("Day is not entered")
		return
	}
	dateTime = fmt.Sprintf("%04d-%02d-%02d", *nam, *thang, *ngay)
	err = validateInput(dateTime)
	if err != nil {
		err = errors.New("invalid input")
	}
	return
}

func main() {
	ngay, thang, nam := 28, 2, 10000

	dateTime, err := convertDatetimeString(&ngay, &thang, &nam)

	if err == nil {
		fmt.Println(dateTime)
	} else {
		fmt.Println(err)
	}

}
