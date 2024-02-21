package collection

import (
	"errors"
	"model-sample/internal/domain/sell/value"
	"sort"
	"time"
)

type Durations struct {
	durations []value.Duration
}

var (
	ErrDurationDuplicated = errors.New("Durations: duration duplicated")
	ErrDurationNotFound   = errors.New("Durations: duration not found")
)

func NewDurations(durations []value.Duration) (Durations, error) {
	// durationsをstartAtでソート
	sort.Slice(durations, func(i, j int) bool {
		return durations[i].StartAt().Before(durations[j].StartAt())
	})

	// 該当するendAt < 次のstartAtの場合はエラー
	for i := 0; i < len(durations)-1; i++ {
		endAt := durations[i].EndAt()
		nextStartAt := durations[i+1].StartAt()
		if !nextStartAt.After(endAt) {
			return Durations{}, ErrDurationDuplicated
		}
	}

	return Durations{
		durations: durations,
	}, nil
}

func (d *Durations) Remove(duration value.Duration) error {
	//見つからない場合はエラー
	for i := 0; i < len(d.durations); i++ {
		if d.durations[i].Equals(duration) {
			d.durations = append(d.durations[:i], d.durations[i+1:]...)
			return nil
		}
	}

	return ErrDurationNotFound
}

func (d *Durations) Add(target value.Duration) {
	buffer := make([]value.Duration, 0)

	for i := 0; i < len(d.durations); i++ {
		endAt := d.durations[i].EndAt()
		startAt := d.durations[i].StartAt()

		if target.StartAt().Before(startAt) && target.EndAt().After(endAt) {
			continue
		}

		if !endAt.Before(target.StartAt()) {
			add, err := value.NewDuration(target.StartAt().Add(1*time.Second), endAt)
			if err != nil {
				return
			}
			buffer = append(buffer, add)
			continue
		}

		if !startAt.After(target.EndAt()) {
			add, err := value.NewDuration(startAt, target.EndAt().Add(-1*time.Second))
			if err != nil {
				return
			}
			buffer = append(buffer, add)
			continue
		}
	}

	buffer = append(buffer, target)
	sort.Slice(buffer, func(i, j int) bool {
		return buffer[i].StartAt().Before(buffer[j].StartAt())
	})

	d.durations = buffer
}
