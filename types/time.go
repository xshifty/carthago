package types

import (
	"time"
)

type UnixTimestamp int64

func NewUnixTimestamp() UnixTimestamp {
	return UnixTimestamp(time.Now().Unix())
}

func (timestamp UnixTimestamp) DatetimeString() string {
	return time.Unix(int64(timestamp), 0).String()
}
