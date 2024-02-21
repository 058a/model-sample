package value

import (
	"time"

	"errors"
)

type Duration struct {
	startAt time.Time
	endAt   time.Time
}

var (
	ErrInvalidDuration = errors.New("Duration: startAt must be before endAt")
)

func NewDuration(startAt, endAt time.Time) (Duration, error) {
	if startAt.After(endAt) {
		return Duration{}, ErrInvalidDuration
	}

	return Duration{
		startAt: startAt,
		endAt:   endAt,
	}, nil
}

func (d Duration) StartAt() time.Time {
	return d.startAt
}

func (d Duration) EndAt() time.Time {
	return d.endAt
}

func (d Duration) Equals(other Duration) bool {
	return d.StartAt() == other.StartAt() && d.EndAt() == other.EndAt()
}

func (d *Duration) ChnageStartAt(v time.Time) {
	d.startAt = v
}

func (d *Duration) ChnageEndAt(v time.Time) {
	d.endAt = v
}
