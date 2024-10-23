package moon

/*
o                     __...__     *
              *   .--'    __.=-.             o
     |          ./     .-'
    -O-        /      /
     |        /    '"/               *
             |     (+)
            |        \                         .
            |         \
 *          |   |   ___\                  |
             |  .   /  `                 -O-
              \  `~~\                     |
         o     \     \            *
                `\    `-.__           .
    .             `--._    `--'jgs
                       `---~~`                *
            *                   o
*/

import (
	"math"
	"time"
)

// Phase represents a phase of the moon
type Phase uint8

const (
	PhaseNew Phase = iota
	PhaseWaxingCrescent
	PhaseFirstQuarter
	PhaseWaxingGibbous
	PhaseFull
	PhaseWaningGibbous
	PhaseLastQuarter
	PhaseWaningCrescent
	phaseCount
)

var emoji = map[Phase]string{
	PhaseNew:            "🌑",
	PhaseWaxingCrescent: "🌒",
	PhaseFirstQuarter:   "🌓",
	PhaseWaxingGibbous:  "🌔",
	PhaseFull:           "🌕",
	PhaseWaningGibbous:  "🌖",
	PhaseLastQuarter:    "🌗",
	PhaseWaningCrescent: "🌘",
}

var names = map[Phase]string{
	PhaseNew:            "New Moon",
	PhaseWaxingCrescent: "Waxing Crescent",
	PhaseFirstQuarter:   "First Quarter",
	PhaseWaxingGibbous:  "Waxing Gibbous",
	PhaseFull:           "Full Moon",
	PhaseWaningGibbous:  "Waning Gibbous",
	PhaseLastQuarter:    "Last Quarter",
	PhaseWaningCrescent: "Waning Crescent",
}

const (
	lunationDays     = 29.530588
	knownNewMoonUnix = 592500 // 1970-01-07T20:35:00Z
	secondsPerDay    = 86400
)

// GetPhase returns the current moon phase
func GetPhase() Phase {
	return GetPhaseAt(time.Now())
}

// GetPhaseAt returns the moon phase at the given time
func GetPhaseAt(t time.Time) Phase {
	phase := calculatePhase(t)
	spread := (phase * float64(phaseCount)) + 0.5
	if spread >= float64(phaseCount) {
		return PhaseNew
	}
	for i := phaseCount - 1; i > 0; i-- {
		if spread > float64(i) {
			return Phase(i)
		}
	}
	return PhaseNew
}

// Emoji returns the emoji representation of the moon phase
func (p Phase) Emoji() string {
	return emoji[p]
}

// String returns the name of the moon phase
func (p Phase) String() string {
	return names[p]
}

// returns moon phase expressed in [0, 1], where 0 and 1 both represent new moon, 0.5 represents full moon
func calculatePhase(t time.Time) float64 {
	knownNewMoonTime := time.Unix(knownNewMoonUnix, 0)
	phase := (t.UTC().Sub(knownNewMoonTime).Seconds() / float64(secondsPerDay)) / lunationDays
	return phase - math.Floor(phase)
}
