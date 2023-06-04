package structs

import "fmt"

type Stats struct {
	Count                 int
	AvgRewardWhenUsed     float64
	AvgRewardWhenEligible float64
	TimeSinceLast         int
}

func (s Stats) String() string {
	return fmt.Sprintf("Count: %d, AvgRewardWhenUsed: %f, TimeSinceLast: %d", s.Count, s.AvgRewardWhenUsed, s.TimeSinceLast)
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

func (h *History) Update(chosenArm *Arm, eligibleArms []*Arm, reward float64) {
	h.updateAvgRewardWhenUsed(chosenArm, reward)
	h.updateAvgRewardWhenEligible(eligibleArms, reward)
	h.updateTimeSinceLast(chosenArm)
}

func (h *History) updateAvgRewardWhenUsed(chosenArm *Arm, reward float64) {
	armStats := h.ArmToStats[chosenArm]
	armStats.Count++
	count := float64(armStats.Count)
	armStats.AvgRewardWhenUsed = ((count-1.0)/count)*armStats.AvgRewardWhenUsed + (1.0/count)*reward
}

func (h *History) updateAvgRewardWhenEligible(eligibleArms []*Arm, reward float64) {
	for _, arm := range eligibleArms {
		armStats := h.ArmToStats[arm]
		count := float64(armStats.Count)
		armStats.AvgRewardWhenEligible = ((count-1.0)/count)*armStats.AvgRewardWhenEligible + (1.0/count)*reward
	}
}

func (h *History) updateTimeSinceLast(chosenArm *Arm) {
	for arm, stats := range h.ArmToStats {
		if arm == chosenArm {
			stats.TimeSinceLast = 0
		} else {
			stats.TimeSinceLast++
		}
	}
}
