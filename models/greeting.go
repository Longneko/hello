package models

import "time"

type Greeting struct {
	Name string `form:"name"`
	Time time.Time
}
