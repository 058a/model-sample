package collection

import (
	"fmt"
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
					validDuration2,
					validDuration1,
				},
			},
			want: Durations{
				durations: []value.Duration{
					validDuration2,
					validDuration1,
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

func TestDurations_Remove(t *testing.T) {
	// Setup
	t.Parallel()

	currentAt := time.Now()

	duration1, err := value.NewDuration(currentAt, currentAt.Add(1*time.Hour))
	if err != nil {
		t.Fatal(err)
	}

	duration2, err := value.NewDuration(currentAt.Add(1*time.Hour+1*time.Second), currentAt.Add(2*time.Hour))
	if err != nil {
		t.Fatal(err)
	}

	durations := []value.Duration{
		duration1,
		duration2,
	}

	d, err := NewDurations(durations)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
		{
			name:    "fail/duration not found",
			wantErr: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			// When
			err = d.Remove(duration1)

			if !tt.wantErr {
				if err != nil {
					t.Errorf("Durations.Remove() error = %v", err)
				}
				return
			}

			if err == nil {
				t.Errorf("Durations.Remove() error = %v", err)
			}

		})
	}
}

func TestDurations_Add(t *testing.T) {
	// Setup
	t.Parallel()

	currentAt := time.Now()

	duration1, err := value.NewDuration(currentAt, currentAt.Add(1*time.Hour))
	if err != nil {
		t.Fatal(err)
	}

	duration2, err := value.NewDuration(currentAt.Add(1*time.Hour+1*time.Second), currentAt.Add(2*time.Hour))
	if err != nil {
		t.Fatal(err)
	}

	duration3, err := value.NewDuration(currentAt.Add(2*time.Hour+1*time.Second), currentAt.Add(3*time.Hour))
	if err != nil {
		t.Fatal(err)
	}

	duration4, err := value.NewDuration(currentAt.Add(3*time.Hour+1*time.Second), currentAt.Add(4*time.Hour))
	if err != nil {
		t.Fatal(err)
	}

	target, err := value.NewDuration(currentAt.Add(2*time.Hour+2*time.Second), currentAt.Add(2*time.Hour+3*time.Second))
	if err != nil {
		t.Fatal(err)
	}

	durations := []value.Duration{
		duration1,
		duration2,
		duration3,
		duration4,
	}

	d, err := NewDurations(durations)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			for _, v := range d.durations {
				fmt.Printf("durations: %v %v\n", v.StartAt().Format(time.RFC3339), v.EndAt().Format(time.RFC3339))
			}
			fmt.Println("-----")

			// When
			d.Add(target)

			if !tt.wantErr {
				if err != nil {
					t.Errorf("Durations.Add() error = %v", err)
				}
				for _, v := range d.durations {
					fmt.Printf("durations: %v %v\n", v.StartAt().Format(time.RFC3339), v.EndAt().Format(time.RFC3339))
				}
				return
			}

			if err == nil {
				t.Errorf("Durations.Add() error = %v", err)
			}

		})
	}
}
