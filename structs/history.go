package structs

import "fmt"

type Stats struct {
	Count            int
	RunningAvgReward float64
	TimeSinceLast    int
}

func (s Stats) String() string {
	return fmt.Sprintf("Count: %d, RunningAvgReward: %f, TimeSinceLast: %d", s.Count, s.RunningAvgReward, s.TimeSinceLast)
}

type History struct {
	ArmToStats map[*Arm]*Stats
}

func NewHistory(arms []*Arm) *History {
	armToStats := make(map[*Arm]*Stats)
	for _, arm := range arms {
		armToStats[arm] = &Stats{}
	}
	return &History{
		ArmToStats: armToStats,
	}
}

func (h *History) Update(chosenArm *Arm, reward float64) {
	armStats := h.ArmToStats[chosenArm]
	armStats.Count++
	count := float64(armStats.Count)
	armStats.RunningAvgReward = ((count-1.0)/count)*armStats.RunningAvgReward + (1.0/count)*reward
	for arm, stats := range h.ArmToStats {
		if arm == chosenArm {
			stats.TimeSinceLast = 0
		} else {
			stats.TimeSinceLast++
		}
	}
}
