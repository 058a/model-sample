package collection

import (
	"errors"
	"model-sample/internal/domain/sell/value"
	"sort"
)

type Durations struct {
	durations []value.Duration
}

var (
	ErrDuplicatedDuration = errors.New("Durations: duplicated duration")
)

func NewDurations(durations []value.Duration) (Durations, error) {
	// durationsをstartAtでソート
	sort.Slice(durations, func(i, j int) bool {
		return durations[i].StartAt().Before(durations[j].StartAt())
	})

	// durationsのstartAtとendAtの期間の重複をチェック
	for i := 0; i < len(durations)-1; i++ {
		endAt := durations[i].EndAt()
		nextStartAt := durations[i+1].StartAt()
		nextEndAt := durations[i+1].EndAt()
		if !nextStartAt.After(endAt) {
			return Durations{}, ErrDuplicatedDuration
		}
		if !nextEndAt.After(endAt) {
			return Durations{}, ErrDuplicatedDuration
		}
	}

	return Durations{
		durations: durations,
	}, nil
}
