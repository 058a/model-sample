package value_test

import (
	"errors"
	"model-sample/internal/domain/sell/value"
	"testing"
	"time"
)

func TestNewDuration(t *testing.T) {
	// Setup
	t.Parallel()

	type args struct {
		startAt time.Time
		endAt   time.Time
	}
	type want struct {
		startAt time.Time
		endAt   time.Time
	}

	passedStartAt := time.Now()
	passedEndAt := time.Now().Add(1 * time.Hour)
	failedStartAt := time.Now().Add(1 * time.Hour)
	failedEndAt := time.Now()
	tests := []struct {
		name        string
		args        args
		want        want
		wantErr     bool
		wantErrType error
	}{
		{
			name: "success",
			args: args{
				startAt: passedStartAt,
				endAt:   passedEndAt,
			},
			want: want{
				startAt: passedStartAt,
				endAt:   passedEndAt,
			},
			wantErr:     false,
			wantErrType: nil,
		},
		{
			name: "fail/startAt must be before endAt",
			args: args{
				startAt: failedStartAt,
				endAt:   failedEndAt,
			},
			want:        want{},
			wantErr:     true,
			wantErrType: value.ErrInvalidDuration,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := value.NewDuration(tt.args.startAt, tt.args.endAt)
			if !tt.wantErr {
				if err != nil {
					t.Errorf("NewDuration() error = %v, wantErr %v", err, tt.wantErr)
				}
				if got.StartAt() != tt.want.startAt {
					t.Errorf("NewDuration() = %v, want %v", got.StartAt(), tt.want.startAt)
				}
				if got.EndAt() != tt.want.endAt {
					t.Errorf("NewDuration() = %v, want %v", got.EndAt(), tt.want.endAt)
				}
				return
			}

			if !errors.Is(err, tt.wantErrType) {
				t.Errorf("NewDuration() error = %v, wantErr %v", err, tt.wantErrType)
			}
		})
	}
}
