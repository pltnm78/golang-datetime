package datetime

import "regexp"

var dateRegexpList []*regexp.Regexp
var timeRegexpList []*regexp.Regexp

var isPrepared = false

func prepare() {
	// register regexp Date
	dateRegexpList = append(dateRegexpList, regDate01)
	dateRegexpList = append(dateRegexpList, regDate02)
	dateRegexpList = append(dateRegexpList, regDate03)
	dateRegexpList = append(dateRegexpList, regDate04)
	dateRegexpList = append(dateRegexpList, regDate05)
	dateRegexpList = append(dateRegexpList, regDate06)
	dateRegexpList = append(dateRegexpList, regDate07)
	dateRegexpList = append(dateRegexpList, regDate08)
	dateRegexpList = append(dateRegexpList, regDate09)
	dateRegexpList = append(dateRegexpList, regDate10)
	dateRegexpList = append(dateRegexpList, regDate11)
	dateRegexpList = append(dateRegexpList, regDate12)
	dateRegexpList = append(dateRegexpList, regDate13)
	dateRegexpList = append(dateRegexpList, regDate14)
	dateRegexpList = append(dateRegexpList, regDate15)
	dateRegexpList = append(dateRegexpList, regDate16)
	dateRegexpList = append(dateRegexpList, regDate17)
	dateRegexpList = append(dateRegexpList, regDate18)
	dateRegexpList = append(dateRegexpList, regDate19)

	// register regexp Time
	timeRegexpList = append(timeRegexpList, regTime01)
	timeRegexpList = append(timeRegexpList, regTime02)
	timeRegexpList = append(timeRegexpList, regTime03)
	timeRegexpList = append(timeRegexpList, regTime04)
	timeRegexpList = append(timeRegexpList, regTime05)
	timeRegexpList = append(timeRegexpList, regTime06)
	timeRegexpList = append(timeRegexpList, regTime07)
	timeRegexpList = append(timeRegexpList, regTime08)
	timeRegexpList = append(timeRegexpList, regTime09)
	timeRegexpList = append(timeRegexpList, regTime10)
	timeRegexpList = append(timeRegexpList, regTime11)
	timeRegexpList = append(timeRegexpList, regTime12)
	timeRegexpList = append(timeRegexpList, regTime13)
}

func checkMatchRegexp(str string, list []*regexp.Regexp) (result []string, matchId int) {
	matchId = 0

	for i := range list {
		m := list[i].FindAllStringSubmatch(str, -1)
		if len(m) > 0 {
			result = m[0]
			matchId = i + 1
			return
		}
	}

	return
}
