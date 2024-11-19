package storage

import (
	"fmt"
	"time"
)

func dateToString(date uint64) string {
	if date == 0 {
		return "1970-01-01"
	}
	t := time.Unix(int64(date*24*3600), 0).UTC()
	return t.Format("2006-01-02")
}

func timestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp/1e9, timestamp%1e9).UTC()
}

func timestampFromTime(t time.Time) int64 {
	return t.UnixNano()
}

type TimeRange struct {
	MinTimestamp int64
	MaxTimestamp int64
}

func (tr *TimeRange) String() string {
	start := TimestampToHumanReadableFormat(tr.MinTimestamp)
	end := TimestampToHumanReadableFormat(tr.MaxTimestamp)
	return fmt.Sprintf("[%s..%s]", start, end)
}

func TimestampToHumanReadableFormat(timestamp int64) string {
	t := timestampToTime(timestamp).UTC()
	return t.Format("2006-01-02T15:04:05.999Z")
}

func timestampToPartitionName(timestamp int64) string {
	t := timestampToTime(timestamp)
	return t.Format("2006_01")
}

func (tr *TimeRange) fromPartitionName(name string) error {
	t, err := time.Parse("2006_01", name)
	if err != nil {
		return fmt.Errorf("cannot parse partition name %q: %w", name, err)
	}
	tr.fromPartitionTime(t)
	return nil
}

func (tr *TimeRange) fromPartitionTimestamp(timestamp int64) {
	t := timestampToTime(timestamp)
	tr.fromPartitionTime(t)
}

func (tr *TimeRange) fromPartitionTime(t time.Time) {
	y, m, _ := t.UTC().Date()
	minTime := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
	maxTime := time.Date(y, m+1, 1, 0, 0, 0, 0, time.UTC)
	tr.MinTimestamp = minTime.Unix() * 1e3
	tr.MaxTimestamp = maxTime.Unix()*1e3 - 1
}

const msecPerDay = 24 * 3600 * 1000

const msecPerHour = 3600 * 1000
