package moon

import (
	"testing"
	"time"
)

func TestKnownMoonPhases(t *testing.T) {
	tests := []struct {
		date  time.Time
		phase Phase
	}{
		{
			date:  mustParseTime("1900-01-01T12:00:00Z"),
			phase: PhaseNew,
		},
		{
			date:  mustParseTime("2000-01-06T19:24:01Z"),
			phase: PhaseNew,
		},
		{
			date:  mustParseTime("1980-07-20T05:51:00Z"),
			phase: PhaseFirstQuarter,
		},
		{
			date:  mustParseTime("1990-09-05T01:46:00Z"),
			phase: PhaseFull,
		},
		{
			date:  mustParseTime("2000-07-24T11:03:00Z"),
			phase: PhaseLastQuarter,
		},
	}

	for _, test := range tests {
		t.Run(test.date.String(), func(t *testing.T) {
			actual := GetPhaseAt(test.date)
			if test.phase != actual {
				t.Errorf("expected %s, got %s", test.phase, actual)
			}
		})
	}
}

func mustParseTime(t string) time.Time {
	parsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		panic(err)
	}
	return parsed
}
