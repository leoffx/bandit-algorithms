package structs

import "fmt"

type Stats struct {
	Count            int
	RunningAvgReward float64
}

func (s Stats) String() string {
	return fmt.Sprintf("Count: %d, RunningAvgReward: %f", s.Count, s.RunningAvgReward)
}

type History struct {
	ArmToStats map[*Arm]Stats
}

// Updates Arm's Running Average Reward and Count
func (h *History) Update(arm *Arm, reward float64) {
	armStats := h.ArmToStats[arm]
	armStats.Count += 1
	count := float64(armStats.Count)
	armStats.RunningAvgReward = ((count-1.0)/count)*armStats.RunningAvgReward + (1.0/count)*reward
	h.ArmToStats[arm] = armStats
}
