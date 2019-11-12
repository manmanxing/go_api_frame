package util

import "time"

const DateFormat = "2006-01-02"          //格式化日期
const TimeFormat = "2006-01-02 15:04:05" //格式化时间

func JudgeDate(a ...string) bool {
	if len(a) == 0 {
		return false
	}
	for i := range a {
		_, err := time.Parse(DateFormat, a[i])
		if err != nil {
			return false
		}
	}
	return true
}
