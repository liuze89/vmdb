package storage

import (
	"fmt"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/encoding"
)

type TSID struct {
	MetricGroupID uint64
	MetricID      uint64
}

var marshaledTSIDSize = func() int {
	var t TSID
	dst := t.Marshal(nil)
	return len(dst)
}()

func (t *TSID) Marshal(dst []byte) []byte {
	dst = encoding.MarshalUint64(dst, t.MetricGroupID)
	dst = encoding.MarshalUint64(dst, t.MetricID)
	return dst
}

func (t *TSID) Unmarshal(src []byte) ([]byte, error) {
	if len(src) < marshaledTSIDSize {
		return nil, fmt.Errorf("too short src; got %d bytes; want %d bytes", len(src), marshaledTSIDSize)
	}

	t.MetricGroupID = encoding.UnmarshalUint64(src)
	src = src[8:]
	t.MetricID = encoding.UnmarshalUint64(src)
	src = src[8:]

	return src, nil
}

func (t *TSID) Less(b *TSID) bool {
	if t.MetricGroupID != b.MetricGroupID {
		return t.MetricGroupID < b.MetricGroupID
	}
	return t.MetricID < b.MetricID
}
