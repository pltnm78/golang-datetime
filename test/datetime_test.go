package gotest

import (
	".."
	"fmt"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	var errmsg string

	if errmsg = checkDate("19990131", 1999, 1, 31); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("1999/01/31", 1999, 1, 31); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("2001/2/3", 2001, 2, 3); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("12/22/78", 1978, 12, 22); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("1/17/2006", 2006, 1, 17); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("3/4/6", 2006, 3, 4); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("5/12", time.Now().Year(), 5, 12); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("10/27", time.Now().Year(), 10, 27); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("2008-6", 2008, 6, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("1978-12", 1978, 12, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("2008-6-30", 2008, 6, 30); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("78-12-22", 1978, 12, 22); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("8-6-21", 2008, 6, 21); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("30-6-2008", 2008, 6, 30); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("22.12.1978", 1978, 12, 22); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("30.6.75", 1975, 6, 30); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("22\t12.78", 1978, 12, 22); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("May-09-80", 1980, 5, 9); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("Apr-17-1790", 1790, 4, 17); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("78-Dec-22", 1978, 12, 22); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("1814-MAY-17", 1814, 5, 17); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("30-June 2008", 2008, 6, 30); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("22DEC78", 1978, 12, 22); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("14 III 1879", 1879, 3, 14); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("1.September.2010", 2010, 9, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("July 1st, 2008", 2008, 7, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("April 17, 1790", 1790, 4, 17); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("May.9,78", 1978, 5, 9); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("June 2008", 2008, 6, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("DEC1978", 1978, 12, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("March 1879", 1879, 3, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("July 1st,", time.Now().Year(), 7, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("Apr 17", time.Now().Year(), 4, 17); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("May.9", time.Now().Year(), 5, 9); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("2008 June", 2008, 6, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("1978-XII", 1978, 12, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("1879.MArCH", 1879, 3, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("1 July", time.Now().Year(), 7, 1); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("17 Apr", time.Now().Year(), 4, 17); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDate("9.May", time.Now().Year(), 5, 9); errmsg != "" {
		t.Fatal(errmsg)
	}
}

func TestParseTime(t *testing.T) {
	var errmsg string

	if errmsg = checkTime("04:08", 4, 8, 0, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("19.19", 19, 19, 0, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("0102", 1, 2, 0, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("2113", 21, 13, 0, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("04.08.37", 4, 8, 37, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("21:37:08", 21, 37, 8, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("065709", 6, 57, 9, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("171819", 17, 18, 19, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("04.08.37.81412", 4, 8, 37, 814120000); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("19:21:03.032453", 19, 21, 3, 32453000); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("23:59:05.1230", 23, 59, 5, 123000000); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("4 am", 4, 0, 0, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("5PM", 17, 0, 0, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("10:08 am", 10, 8, 0, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("7:19A.M", 7, 19, 0, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("2:05:37 am", 2, 5, 37, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("11:02:59P.M.", 23, 2, 59, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("10.05.33 P.m.", 22, 5, 33, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkTime("01.59pM.", 13, 59, 0, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
}

func TestCompositePattern(t *testing.T) {
	var errmsg string

	if errmsg = checkDateTime("2008-08-07 18:11:31", 2008, 8, 7, 18, 11, 31, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
	if errmsg = checkDateTime("2008-07-01T22:35:17.02", 2008, 7, 1, 22, 35, 17, 20000000); errmsg != "" {
		t.Fatal(errmsg)
	}
}

func TestRandomPattern(t *testing.T) {
	var errmsg string

	if errmsg = checkDateTime("11110213 T230113", 1111, 2, 13, 23, 1, 13, 0); errmsg != "" {
		t.Fatal(errmsg)
	}
}

func checkDate(testData string, year int, month int, day int) (msg string) {
	return checkDateTime(testData, year, month, day, 0, 0, 0, 0)
}

func checkTime(testData string, hour int, min int, sec int, nsec int) (msg string) {
	return checkDateTime(testData, time.Now().Year(), int(time.Now().Month()), time.Now().Day(), hour, min, sec, nsec)
}

func checkDateTime(testData string, year int, month int, day int, hour int, min int, sec int, nsec int) (msg string) {
	result, err := datetime.GetDatetime(testData, time.Local)

	if err != nil {
		msg = "Fatal error!"
		fmt.Println(err)
		return
	}

	if (result.Year() != year) || (result.Month() != time.Month(month)) || (result.Day() != day) || (result.Hour() != hour) || (result.Minute() != min) || (result.Second() != sec) || (result.Nanosecond() != nsec) {
		msg = "Error: Not match!"
		fmt.Println(testData, result, nsec, result.Nanosecond())
		return
	}

	return ""
}
