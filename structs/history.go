package structs

type Stats struct {
	Count int
	Value float64
}

type History struct {
	ArmToStats map[Arm]Stats
}
