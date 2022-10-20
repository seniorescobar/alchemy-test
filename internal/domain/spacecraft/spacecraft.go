package spacecraft

type Spacecraft struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Class     string     `json:"class"`
	Armaments []Armament `json:"armaments"`
	Crew      uint       `json:"crew"`
	Image     string     `json:"image"`
	Val       float64    `json:"value"`
	Status    Status     `json:"status"`
}

type Armament struct {
	Title string `json:"title"`
	Qty   string `json:"qty"`
}

type Status string

const (
	StatusOperational Status = "operational"
	StatusDamaged     Status = "damaged"
)
