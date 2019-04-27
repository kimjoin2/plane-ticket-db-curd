package utils

import (
	"strconv"
	"time"
)

func TimeToStringUntilMin(t time.Time) string {
	res := strconv.Itoa(t.Year()) + "/" + strconv.Itoa(int(t.Month())) + "/"
	res += strconv.Itoa(t.Day()) + " " + strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute())
	return res
}
