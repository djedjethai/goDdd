package memory

import (
	"time"
)

type Beer struct {
	ID        string
	Name      string
	Brewery   string
	Abv       float32
	ShortDesc string
	Created   time.Time
}
