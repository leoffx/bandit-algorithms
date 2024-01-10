package arm

type Arm interface {
	Mean() float64
	Pull() float64
	String() string
}

func InitializeArms(numArms int) []Arm {
	arms := make([]Arm, numArms)
	for i := 0; i < numArms; i++ {
		arms[i] = NewBernouliArm()
	}
	return arms
}
