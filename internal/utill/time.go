package utill

import "time"

func Parse(dmy, hm string) (time.Time, error) {
	return time.Parse("02/01/2006 15:04", dmy+" "+hm)
}

func ParseDate(dmy string) (time.Time, error) {
	return time.Parse("02/01/2006", dmy)
}

func Format(tm time.Time) (dmy, hm string) {
	dmy = tm.Format("02/01/2006")
	hm = tm.Format("15:04")
	return dmy, hm
}

func FormatHHmmTZ0(tm time.Time) string {
	return tm.In(time.UTC).Format("15:04")
}
