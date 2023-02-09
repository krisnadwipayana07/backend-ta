package timeconvert

import (
	"fmt"
	"strconv"
	"time"
)

func UnixTimestampConvert(UnixTime string) time.Time {
	fmt.Println(UnixTime)
	i, err := strconv.ParseInt(UnixTime, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	// fmt.Println(tm)
	return tm
}
