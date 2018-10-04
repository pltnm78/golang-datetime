package parser

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type DateParams struct {
	Year, Month, Day int
}

// regexp Date
var regDate01 = regexp.MustCompile(`([0-9]{4})([0-9]{2})([0-9]{2})`)                             // YYMMDD
var regDate02 = regexp.MustCompile(`([0-9]{4})/([0-9]{1,2})/([0-9]{1,2})`)                       // YY/mm/dd
var regDate03 = regexp.MustCompile(`([0-9]{1,2})[\t-]([0-9]{1,2})-([0-9]{4})`)                   // dd[\t-]mm-YY
var regDate04 = regexp.MustCompile(`([0-9]{4})-([0-9]{1,2})-([0-9]{1,2})`)                       // YY-mm-dd
var regDate05 = regexp.MustCompile(`([0-9]{1,2})/([0-9]{1,2})/([0-9]{1,4})`)                     // mm/dd/y
var regDate06 = regexp.MustCompile(`([0-9]{1,2})/([0-9]{1,2})`)                                  // mm/dd
var regDate07 = regexp.MustCompile(`([0-9]{1,4})-([0-9]{1,2})-([0-9]{1,2})`)                     // y-mm-dd
var regDate08 = regexp.MustCompile(`([0-9]{4})-([0-9]{1,2})`)                                    // YY-mm
var regDate09 = regexp.MustCompile(`([0-9]{1,2})\t([0-9]{1,2})\.([0-9]{2})`)                     // dd\tmm.yy
var regDate10 = regexp.MustCompile(`([0-9]{1,2})\.([0-9]{1,2})\.([0-9]{4})`)                     // dd.mm.YY
var regDate11 = regexp.MustCompile(`([0-9]{1,2})\.([0-9]{1,2})\.([0-9]{2})`)                     // dd.mm.yy
var regDate12 = regexp.MustCompile(`([a-z]{3})-([0-9]{2})-([0-9]{1,4})`)                         // M-DD-y
var regDate13 = regexp.MustCompile(`([0-9]{1,4})-([a-z]{3})-([0-9]{2})`)                         // y-M-DD
var regDate14 = regexp.MustCompile(`([0-9]{1,2})[ \t.-]*([a-z]{1,9})[ \t.-]*([0-9]{1,4})`)       // dd([ \t.-])*m([ \t.-])*y
var regDate15 = regexp.MustCompile(`([a-z]{1,9})[ .\t-]*([0-9]{1,2})[,.stndrh\t ]+([0-9]{1,4})`) // m[ .\t-])*dd[,.stndrh\t ]+y
var regDate16 = regexp.MustCompile(`([a-z]{1,9})[ \t.-]*([0-9]{4})`)                             // m([ \t.-])*YY
var regDate17 = regexp.MustCompile(`([a-z]{1,9})[ .\t-]*([0-9]{1,2})[,.stndrh\t ]*`)             // m[ .\t-])*dd[,.stndrh\t ]*
var regDate18 = regexp.MustCompile(`([0-9]{4})[ \t.-]*([a-z]{1,9})`)                             // YY([ \t.-])*m
var regDate19 = regexp.MustCompile(`([0-9]{1,2})[ .\t-]*([a-z]{1,9})`)                           // dd([ .\t-])*m

func GetDate(dt string) (params DateParams, eliminated string, err error) {
	params, eliminated, err = parseDate(strings.ToLower(dt))

	if err != nil {
		return
	}

	return
}

// This parses "date" factors by datetime format string.
func parseDate(str string) (prm DateParams, eliminated string, err error) {
	var year, month, day int

	if !isPrepared {
		prepare()
		isPrepared = true
	}

	// check regexp Date
	matchDate, matchId := checkMatchRegexp(str, dateRegexpList)
	if matchId > 0 {
		switch matchId {
		case 1, 2, 4, 7, 13:
			year, month, day, err = validateDate(matchDate[1], matchDate[2], matchDate[3])
		case 5, 12, 15:
			year, month, day, err = validateDate(matchDate[3], matchDate[1], matchDate[2])
		case 6, 17:
			year, month, day, err = validateDate("", matchDate[1], matchDate[2])
		case 8:
			year, month, day, err = validateDate(matchDate[1], matchDate[2], "")
		case 3, 9, 10, 11, 14:
			year, month, day, err = validateDate(matchDate[3], matchDate[2], matchDate[1])

			// Additional decision for "dd.mm.yy"
			if (err == nil) && (matchId == 11) {
				tmpYear, _ := strconv.Atoi(matchDate[3])
				if tmpYear < 61 {
					// Skip date format dd.mm.yy to prioritize time format
					return DateParams{time.Now().Year(), int(time.Now().Month()), time.Now().Day()}, str, nil
				}
			}
		case 16:
			year, month, day, err = validateDate(matchDate[2], matchDate[1], "1")
		case 18:
			year, month, day, err = validateDate(matchDate[1], matchDate[2], "1")
		case 19:
			year, month, day, err = validateDate("", matchDate[2], matchDate[1])
		}

		if err == nil {
			prm = DateParams{year, month, day}
			eliminated = strings.Replace(str, matchDate[0], "", -1)
			return
		} else {
			return DateParams{time.Now().Year(), int(time.Now().Month()), time.Now().Day()}, str, nil
		}
	} else {
		return DateParams{time.Now().Year(), int(time.Now().Month()), time.Now().Day()}, str, nil
	}

	return
}


func validateDate(year string, month string, day string) (y int, m int, d int, e error) {
	// parse to int from input string
	// if the arguments is empty, apply now

	if year == "" {
		y = time.Now().Year()
	} else {
		if y, e = strconv.Atoi(year); e != nil {
			return
		}

		// the rule of two-digit "year" is the same as PHP
		if y < 70 {
			y = y + 2000
		} else if y < 100 {
			y = y + 1900
		}
	}
	if month == "" {
		m = 1
	} else {
		if m, e = strconv.Atoi(month); e != nil {
			if m = monthStringToInt(month); m == 0 {
				e = errors.New("Validate Error: invalid date format.")
				return 0, 0, 0, e
			}
		}
	}
	if day == "" {
		d = 1
	} else {
		if d, e = strconv.Atoi(day); e != nil {
			return
		}
	}

	// validate format
	if (y < 0 || y > 9999) || (m < 0 || m > 12) || (d < 0 || d > 31) {
		e = errors.New("Validate Error: invalid date format.")
		return 0, 0, 0, e
	}

	return
}

func monthStringToInt(str string) int {
	switch str {
	case "jan", "january", "i":
		return 1
	case "feb", "february", "ii":
		return 2
	case "mar", "march", "iii":
		return 3
	case "apr", "april", "iv":
		return 4
	case "may", "v":
		return 5
	case "jun", "june", "vi":
		return 6
	case "jul", "july", "vii":
		return 7
	case "aug", "august", "viii":
		return 8
	case "sep", "sept", "september", "ix":
		return 9
	case "oct", "october", "x":
		return 10
	case "nov", "november", "xi":
		return 11
	case "dec", "december", "xii":
		return 12
	}

	return 0
}
