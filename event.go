package debeziummodels

import (
	"time"
)

type Op string

const (
	OpCreate   Op = "c"
	OpUpdate   Op = "u"
	OpDelete   Op = "d"
	OpRead     Op = "r"
	OpTruncate Op = "t"
	OpMessage  Op = "m"
)

type TimeMS int64

func (t TimeMS) Time() time.Time {
	return time.Unix(0, int64(t)*int64(time.Millisecond))
}

type PostgresSource struct {
	Version   string  `json:"version"`
	Connector string  `json:"connector"`
	Name      string  `json:"name"`
	TsMs      TimeMS  `json:"ts_ms"`
	Snapshot  string  `json:"snapshot"`
	DB        string  `json:"db"`
	Sequence  *string `json:"sequence"`
	Schema    string  `json:"schema"`
	Table     string  `json:"table"`
	TxID      int64   `json:"txId"`
	Lsn       int64   `json:"lsn"`
	Xmin      *int64  `json:"xmin"`
}

type Payload[T any, Source any] struct {
	Op     Op     `json:"op"`
	Source Source `json:"source"`
	Before *T     `json:"before,omitempty"`
	After  *T     `json:"after,omitempty"`
	TsMs   TimeMS `json:"ts_ms"`
}

type ChangeEvent[T any, Source any] struct {
	Payload Payload[T, Source] `json:"payload"`
}
