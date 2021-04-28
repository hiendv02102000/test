package main

import (
	"errors"
	"fmt"
	"time"
)

const Layout = "2006-01-02"

func validateInput(dateTime string) error {
	_, err := time.Parse(Layout, dateTime)
	return err
}
func validateData(ngay, thang, nam *int) error {
	if ngay == nil && thang == nil && nam == nil {
		return errors.New("")
	}
	if nam == nil {
		return errors.New("Year is not entered")
	}
	if thang == nil {
		return errors.New("Month is not entered")
	}
	if ngay == nil {
		return errors.New("Day is not entered")
	}
	return nil
}
func convertDatetimeString(ngay, thang, nam *int) (dateTime string, err error) {
	err = validateData(ngay, thang, nam)
	if err != nil {
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
	ngay, nam := 29, 1200

	dateTime, err := convertDatetimeString(&ngay, nil, &nam)

	if err == nil {
		fmt.Println(dateTime)
	} else {
		fmt.Println(err)
	}

}
