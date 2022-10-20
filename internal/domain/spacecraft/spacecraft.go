package spacecraft

type Spacecraft struct {
	ID       int
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
