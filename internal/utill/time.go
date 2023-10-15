package utill

import "time"

func Parse(dmy, hm string) (time.Time, error) {
	return time.Parse("02/01/2006 15:04", dmy+" "+hm)
}

func Format(tm time.Time) (dmy, hm string) {
	dmy = tm.Format("02/01/2006")
	hm = tm.Format("15:04")
	return dmy, hm
}
