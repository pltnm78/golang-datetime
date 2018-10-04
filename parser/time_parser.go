package parser

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type TimeParams struct {
	Hour, Min, Sec, Nsec int
}

// regexp Time
var regTime01 = regexp.MustCompile(`(0?[1-9]|1[0-2])\.([0-5][0-9])\.([0-5][0-9])[ \t]?([ap]\.?m\.?)`) // hh.MM.II[ \t]?meridian
var regTime02 = regexp.MustCompile(`(0?[1-9]|1[0-2]):([0-5][0-9]):([0-5][0-9])[ \t]?([ap]\.?m\.?)`)   // hh:MM:II[ \t]?meridian
var regTime03 = regexp.MustCompile(`(0?[1-9]|1[0-2])\.([0-5][0-9])[ \t]?([ap]\.?m\.?)`)               // hh.MM[ \t]?meridian
var regTime04 = regexp.MustCompile(`(0?[1-9]|1[0-2]):([0-5][0-9])[ \t]?([ap]\.?m\.?)`)                // hh:MM[ \t]?meridian
var regTime05 = regexp.MustCompile(`(0?[1-9]|1[0-2])[ \t]?([ap]\.?m\.?)`)                             // hh[ \t]?meridian
var regTime06 = regexp.MustCompile(`[tT]?([01][0-9]|2[0-4])\.([0-5][0-9])\.([0-5][0-9])\.([0-9]+)`)   // [tT]?HH.MM.II.frac
var regTime07 = regexp.MustCompile(`[tT]?([01][0-9]|2[0-4]):([0-5][0-9]):([0-5][0-9])\.([0-9]+)`)     // [tT]?HH:MM:II.frac
var regTime08 = regexp.MustCompile(`[tT]?([01][0-9]|2[0-4])\.([0-5][0-9])\.([0-5][0-9])`)             // [tT]?HH.MM.II
var regTime09 = regexp.MustCompile(`[tT]?([01][0-9]|2[0-4]):([0-5][0-9]):([0-5][0-9])`)               // [tT]?HH:MM:II
var regTime10 = regexp.MustCompile(`[tT]?([01][0-9]|2[0-4])([0-5][0-9])([0-5][0-9])`)                 // [tT]?HHMMII
var regTime11 = regexp.MustCompile(`[tT]?([01][0-9]|2[0-4])\.([0-5][0-9])`)                           // [tT]?HH.MM
var regTime12 = regexp.MustCompile(`[tT]?([01][0-9]|2[0-4]):([0-5][0-9])`)                            // [tT]?HH:MM
var regTime13 = regexp.MustCompile(`[tT]?([01][0-9]|2[0-4])([0-5][0-9])`)                             // [tT]?HHMM

func GetTime(dt string) (params TimeParams, eliminated string, err error) {
	params, eliminated, err = parseTime(strings.ToLower(dt))

	if err != nil {
		return
	}

	return
}

// This parses "time" factors by datetime format string.
func parseTime(str string) (prm TimeParams, eliminated string, err error) {
	var hour, min, sec, nsec int

	if !isPrepared {
		prepare()
		isPrepared = true
	}

	// check regexp Time
	matchTime, matchId := checkMatchRegexp(str, timeRegexpList)
	if matchId > 0 {
		switch matchId {
		case 1, 2:
			hour, min, sec, nsec, err = validateTime(matchTime[1], matchTime[2], matchTime[3], "", matchTime[4])
		case 3, 4:
			hour, min, sec, nsec, err = validateTime(matchTime[1], matchTime[2], "", "", matchTime[3])
		case 5:
			hour, min, sec, nsec, err = validateTime(matchTime[1], "", "", "", matchTime[2])
		case 6, 7:
			hour, min, sec, nsec, err = validateTime(matchTime[1], matchTime[2], matchTime[3], matchTime[4], "")
		case 8, 9, 10:
			hour, min, sec, nsec, err = validateTime(matchTime[1], matchTime[2], matchTime[3], "", "")
		case 11, 12, 13:
			hour, min, sec, nsec, err = validateTime(matchTime[1], matchTime[2], "", "", "")
		}

		if err == nil {
			prm = TimeParams{hour, min, sec, nsec}
			eliminated = strings.Replace(str, matchTime[0], "", -1)
			return
		} else {
			return TimeParams{0, 0, 0, 0}, str, nil
		}
	} else {
		return TimeParams{0, 0, 0, 0}, str, nil
	}

	return
}

func validateTime(hour string, min string, sec string, nsec string, meridian string) (h int, m int, s int, ns int, e error) {
	// parse to int from input string
	// if the arguments is empty, apply 0

	if h, e = strconv.Atoi(hour); e != nil {
		h = 0
	} else {
		if meridian != "" {
			tmpRune := []rune(meridian)
			firstChar := string(tmpRune[0])

			// P.M.
			if (firstChar == "P") || (firstChar == "p") {
				h = h + 12
			}
		}
	}
	if min == "" {
		m = 0
	} else {
		if m, e = strconv.Atoi(min); e != nil {
			return
		}
	}
	if sec == "" {
		s = 0
	} else {
		if s, e = strconv.Atoi(sec); e != nil {
			return
		}
	}
	if nsec == "" {
		ns = 0
	} else {
		// fit 9 digit
		for len(nsec) < 9 {
			nsec += "0"
		}
		nsec = nsec[0:9]

		if ns, e = strconv.Atoi(nsec); e != nil {
			return
		}
	}

	// validate format
	if (h < 0 || h > 24) || (m < 0 || m > 59) || (s < 0 || s > 59) || (ns < 0 || ns > 999999999) {
		e = errors.New("Validate Error: invalid time format.")
		return 0, 0, 0, 0, e
	}

	return
}
