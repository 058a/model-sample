package collection

import (
	"model-sample/internal/domain/sell/value"
	"testing"
	"time"
)

func TestNewDurations(t *testing.T) {
	// Setup
	t.Parallel()

	currentAt := time.Now()
	validDuration1, err := value.NewDuration(currentAt, currentAt.Add(1*time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	validDuration2, err := value.NewDuration(currentAt.Add(1*time.Hour+1*time.Second), currentAt.Add(2*time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	validDuration3, err := value.NewDuration(currentAt.Add(1*time.Hour), currentAt.Add(2*time.Hour))
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		durations []value.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    Durations
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				durations: []value.Duration{
					validDuration1,
					validDuration2,
				},
			},
			want: Durations{
				durations: []value.Duration{
					validDuration1,
					validDuration2,
				},
			},
			wantErr: false,
		},
		{
			name: "fail/durations are duplicated",
			args: args{
				durations: []value.Duration{
					validDuration3,
					validDuration1,
				},
			},
			want:    Durations{},
			wantErr: true,
		},
		{
			name: "fail/durations are duplicated",
			args: args{
				durations: []value.Duration{
					validDuration3,
					validDuration1,
					validDuration2,
				},
			},
			want:    Durations{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// When
			_, err := NewDurations(tt.args.durations)

			// Then
			if !tt.wantErr {
				if err != nil {
					t.Errorf("NewDurations() error = %v", err)
				}
				return
			}

			if err == nil {
				t.Errorf("NewDurations() error = %v", err)
			}
		})
	}
}
