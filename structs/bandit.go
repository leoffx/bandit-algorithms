package structs

type Arm struct{}

func (a Arm) Draw() float64 {
	return 0.0
}

type Bandit struct {
	Arms []Arm
}

func (b Bandit) GetEligibleArms() []Arm {
	return b.Arms
}
