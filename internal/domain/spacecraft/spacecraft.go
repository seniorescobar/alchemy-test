package spacecraft

import "github.com/google/uuid"

type Spacecraft struct {
	ID       uuid.UUID
	Name     string
	Class    string
	Armament []Armament
	Crew     uint
	Image    string
	Value    float64
	Status   Status
}

type Armament struct {
	Title string
	Qty   string
}

type Status string

const (
	StatusOperational Status = "operational"
	StatusDamaged     Status = "damaged"
)
