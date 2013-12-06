package timeutil

import "time"

func  ParseCn( timeStr string) (time.Time, error){
	return Parse( "2006-01-02 15:04:05", timeStr)
}

func  Parse( format string, timeStr string) (time.Time, error){
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation( format, timeStr, loc)
}
