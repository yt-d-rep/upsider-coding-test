package shared

import "time"

type (
	Clock interface {
		Now() time.Time
	}
	clock struct{}

	TimeRange struct {
		from time.Time
		to   time.Time
	}
)

func (c *clock) Now() time.Time {
	return time.Now()
}

func NewTimeRange(from, to time.Time) (*TimeRange, error) {
	if from.After(to) {
		return nil, &ValidationError{Field: "time range", Err: "from must be before to"}
	}
	return &TimeRange{from: from, to: to}, nil
}
func (t *TimeRange) From() time.Time {
	return t.from
}
func (t *TimeRange) To() time.Time {
	return t.to
}
