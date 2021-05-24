package util

import (
	"fmt"
	"strconv"
	"time"
)

// DateFormat01 ...
const DateFormat01 = "2006-01-02 15:04:05"

// DateFormat02 ...
const DateFormat02 = "2006-01-02T15:04:05+09:00"

// DateFormat03 ...
const DateFormat03 = "20060102150405"

// DateFormat04 ...
const DateFormat04 = "060102150405"

// FromDateToDateStringV1 ...
func FromDateToDateStringV1(from string) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromDateToDateStringV1 : func FromDateToDateStringV1(from string) (string, error) start.")))

	loc, _ := time.LoadLocation("Asia/Tokyo")

	// dtTs, err := time.Parse("2006-01-02", from)
	dtTs, err := time.ParseInLocation("2006-01-02T15:04:05+09:00", from, loc)
	if err != nil {
		return "", err
	}

	convertedDtTs := dtTs.Format("2006-01-02")

	return convertedDtTs, nil
}

// FromDateTimeToIS8601DZwStringV1 ...
func FromDateTimeToIS8601DZwStringV1(from string) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromDateTimeToIS8601DZwStringV1 : func FromDateTimeToIS8601DZwStringV1(from string) (string, error) start.")))

	loc, _ := time.LoadLocation("Asia/Tokyo")

	// dtTs, err := time.Parse("2006-01-02T15:04:05+09:00", from)
	dtTs, err := time.ParseInLocation("2006-01-02T15:04:05+09:00", from, loc)
	if err != nil {
		// dtTs, err := time.Parse("2006-01-02 15:04:05", from)
		dtTs, err = time.ParseInLocation("2006-01-02 15:04:05", from, loc)
		if err != nil {
			// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromDateTimeToIS8601DZwStringV1 : time.ParseInLocation(\"2006-01-02 15:04:05\", from, loc) err != nil")))
			return "", err
		}
	}

	// // dtTs, err := time.Parse("2006-01-02 15:04:05", from)
	// dtTs, err := time.ParseInLocation("2006-01-02 15:04:05", from, loc)
	// if err != nil {
	// 	apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromDateTimeToIS8601DZwStringV1 : time.ParseInLocation(\"2006-01-02 15:04:05\", from, loc) err != nil")))
	// 	return "", err
	// }

	convertedDtTs := dtTs.Format("2006-01-02T15:04:05+09:00")

	return convertedDtTs, nil
}

// FromIS8601DZwToDateTimeStringV1 ...
func FromIS8601DZwToDateTimeStringV1(from string) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromIS8601DZwToDateTimeStringV1 : func FromIS8601DZwToDateTimeStringV1(from string) (string, error) start.")))

	loc, _ := time.LoadLocation("Asia/Tokyo")

	// dtTs, err := time.Parse("2006-01-02T15:04:05+09:00", from)
	dtTs, err := time.ParseInLocation("2006-01-02T15:04:05+09:00", from, loc)
	if err != nil {
		return "", err
	}

	convertedDtTs := dtTs.Format("2006-01-02 15:04:05")

	return convertedDtTs, nil
}

// FromIS8601DZwToDateTimeStringV2 ...
func FromIS8601DZwToDateTimeStringV2(from string) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromIS8601DZwToDateTimeStringV2 : func FromIS8601DZwToDateTimeStringV2(from string) (string, error) start.")))

	loc, _ := time.LoadLocation("Asia/Tokyo")

	// dtTs, err := time.Parse("2006-01-02T15:04:05.000+09:00", from)
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromIS8601DZwToDateTimeStringV2 : time.ParseInLocation layout=2006-01-02T15:04:05.000+09:00")))
	dtTs, err := time.ParseInLocation("2006-01-02T15:04:05.000+09:00", from, loc)
	if err != nil {
		// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromIS8601DZwToDateTimeStringV2 : time.ParseInLocation layout=2006-01-02T15:04:05.00+09:00")))
		dtTs, err = time.ParseInLocation("2006-01-02T15:04:05.00+09:00", from, loc)
	}
	if err != nil {
		// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromIS8601DZwToDateTimeStringV2 : time.ParseInLocation layout=2006-01-02T15:04:05.0+09:00")))
		dtTs, err = time.ParseInLocation("2006-01-02T15:04:05.0+09:00", from, loc)
	}
	if err != nil {
		// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromIS8601DZwToDateTimeStringV2 : time.ParseInLocation layout=2006-01-02T15:04:05+09:00")))
		dtTs, err = time.ParseInLocation("2006-01-02T15:04:05+09:00", from, loc)
	}
	if err != nil {
		return "", err
	}

	convertedDtTs := dtTs.Format("2006-01-02 15:04:05.000")

	return convertedDtTs, nil
}

// FromIS8601DZwToDateTimeStringV3 ...
func FromIS8601DZwToDateTimeStringV3(from string) (string, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromIS8601DZwToDateTimeStringV3 : func FromIS8601DZwToDateTimeStringV3(from string) (string, error) start.")))

	loc, _ := time.LoadLocation("Asia/Tokyo")

	// dtTs, err := time.Parse("2006-01-02T15:04:05.000+09:00", from)
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromIS8601DZwToDateTimeStringV2 : time.ParseInLocation layout=2006-01-02T15:04:05.000+09:00")))
	dtTs, err := time.ParseInLocation("2006-01-02T15:04:05.000+09:00", from, loc)
	if err != nil {
		// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromIS8601DZwToDateTimeStringV2 : time.ParseInLocation layout=2006-01-02T15:04:05.00+09:00")))
		dtTs, err = time.ParseInLocation("2006-01-02T15:04:05.00+09:00", from, loc)
	}
	if err != nil {
		// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromIS8601DZwToDateTimeStringV2 : time.ParseInLocation layout=2006-01-02T15:04:05.0+09:00")))
		dtTs, err = time.ParseInLocation("2006-01-02T15:04:05.0+09:00", from, loc)
	}
	if err != nil {
		// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromIS8601DZwToDateTimeStringV2 : time.ParseInLocation layout=2006-01-02T15:04:05+09:00")))
		dtTs, err = time.ParseInLocation("2006-01-02T15:04:05+09:00", from, loc)
	}
	if err != nil {
		return "", err
	}

	convertedDtTs := dtTs.Format("060102150405")

	return convertedDtTs, nil
}

// FromDateStringToTimeObjV1 ...
func FromDateStringToTimeObjV1(from string) (time.Time, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromDateStringToTimeObjV1 : func FromDateStringToTimeObjV1(from string) (time.Time, error) start.")))
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromDateStringToTimeObjV1 : from=%s", from)))

	loc, _ := time.LoadLocation("Asia/Tokyo")

	// t, err := time.Parse("2006-01-02T15:04:05+09:00", from)
	t, err := time.ParseInLocation("2006-01-02T15:04:05+09:00", from, loc)
	if err == nil {
		// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromDateStringToTimeObjV1 : 2006-01-02T15:04:05+09:00 format")))
		return t, nil
	}

	// t, err = time.Parse("2006-01-02 15:04:05", from)
	t, err = time.ParseInLocation("2006-01-02 15:04:05", from, loc)
	if err == nil {
		// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromDateStringToTimeObjV1 : 2006-01-02 15:04:05 format")))
		return t, nil
	}

	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : FromDateStringToTimeObjV1 : fail")))
	return NowTimeV1(), err
}

// NowTimeV1 ...
func NowTimeV1() time.Time {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : NowTimeV1 : func NowTimeV1() start.")))

	return time.Now()
}

// TimeFormatMillisecondV1 ...
func TimeFormatMillisecondV1(t time.Time) string {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : TimeFormatMillisecondV1 : func TimeFormatMillisecondV1(t time.Time) string start.")))

	return fmt.Sprintf("%s.%03d", t.Format("2006-01-02 15:04:05"), t.Nanosecond()/int(time.Millisecond))
}

// NowDateTimeStringAddSecondsV1 ...
func NowDateTimeStringAddSecondsV1(seconds int) string {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : NowDateTimeStringAddSecondsV1 : func NowDateTimeStringAddSecondsV1(seconds int) string start.")))

	return NowTimeV1().Add(time.Duration(seconds) * time.Second).Format("2006-01-02 15:04:05")
}

// TimeAddSecondsV1 ...
func TimeAddSecondsV1(now time.Time, seconds int) string {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : TimeAddSecondsV1 : func TimeAddSecondsV1(now time.Time, seconds int) string start.")))

	return now.Add(time.Duration(seconds) * time.Second).Format("2006-01-02 15:04:05")
}

// BeginningMonthV1 ...
func BeginningMonthV1(t time.Time, addMonths int) time.Time {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : BeginningMonthV1 : func BeginningMonthV1(t time.Time, addMonths int) time.Time start.")))

	loc, _ := time.LoadLocation("Asia/Tokyo")
	at := t.AddDate(0, addMonths, 0)
	bmt := time.Date(at.Year(), at.Month(), 1, 0, 0, 0, 0, loc)

	return bmt
}

// TimeComparisonNowAndIS8601DZwDateTimeStringV1 ...
func TimeComparisonNowAndIS8601DZwDateTimeStringV1(timeString string) (int, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : TimeComparisonNowAndIS8601DZwDateTimeStringV1 : func TimeComparisonNowAndIS8601DZwDateTimeStringV1(timeString string) (int, error) start.")))

	loc, _ := time.LoadLocation("Asia/Tokyo")

	nt := NowTimeV1()

	// ct, err := time.Parse("2006-01-02 15:04:05", timeString)
	ct, err := time.ParseInLocation("2006-01-02 15:04:05", timeString, loc)
	if err != nil {
		return 0, err
	}

	if nt.Before(ct) {
		// -> true  :nt < ct
		return 1, nil
	} else if nt.After(ct) {
		// -> true  :nt > ct
		return -1, nil
	}
	// nt == ct
	return 0, nil
}

// TimeComparisonNowAndDateTimeStringV1 ...
func TimeComparisonNowAndDateTimeStringV1(timeString string) (int, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : TimeComparisonNowAndDateTimeStringV1 : func TimeComparisonNowAndDateTimeStringV1(timeString string) (int, error) start.")))

	nt := NowTimeV1()
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : TimeComparisonNowAndDateTimeStringV1 : nt=%v", nt)))

	ct, err := FromDateStringToTimeObjV1(timeString)
	if err != nil {
		return 0, err
	}
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : TimeComparisonNowAndDateTimeStringV1 : ct=%v", ct)))

	if nt.Before(ct) {
		// -> true  :nt < ct
		return 1, nil
	} else if nt.After(ct) {
		// -> true  :nt > ct
		return -1, nil
	}
	// nt == ct
	return 0, nil
}

// TimeComparisonDateTimeStringV1 ...
func TimeComparisonDateTimeStringV1(timeString1 string, timeString2 string) (int, error) {
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : TimeComparisonDateTimeStringV1 : func TimeComparisonDateTimeStringV1(timeString1 string, timeString2 string) (int, error) start.")))

	nt, err := FromDateStringToTimeObjV1(timeString1)
	if err != nil {
		return 0, err
	}
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : TimeComparisonDateTimeStringV1 : nt=%v", nt)))

	ct, err := FromDateStringToTimeObjV1(timeString2)
	if err != nil {
		return 0, err
	}
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : TimeComparisonDateTimeStringV1 : ct=%v", ct)))

	if nt.Before(ct) {
		// -> true  :nt < ct
		return 1, nil
	} else if nt.After(ct) {
		// -> true  :nt > ct
		return -1, nil
	}
	// nt == ct
	return 0, nil
}

// CalcAgeV1 ...
func CalcAgeV1(birthdate string) (int, error) {

	nt := NowTimeV1()
	// apllog.Debugf(apllogdefaultutil.LogMessageV1(fmt.Sprintf("util/date_util.go : CalcAgeV1 : nt=%v", nt)))

	nowString := nt.Format("20060102")

	loc, _ := time.LoadLocation("Asia/Tokyo")

	// dtTs, err := time.Parse("2006-01-02", birthdate)
	birthdateDtTs, err := time.ParseInLocation("2006-01-02T15:04:05+09:00", birthdate, loc)
	if err != nil {
		return 0, err
	}

	birthdateString := birthdateDtTs.Format("20060102")

	nowInt, err := strconv.Atoi(nowString)
	if err != nil {
		return 0, err
	}
	birthdateInt, err := strconv.Atoi(birthdateString)
	if err != nil {
		return 0, err
	}

	// (今日の日付 - 誕生日) / 10000 = 年齢
	age := (nowInt - birthdateInt) / 10000

	return age, nil
}
