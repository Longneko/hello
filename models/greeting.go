package models

import (
	"errors"
	"sort"
	"sync"
	"time"
)

var DefaultGreetingRepo *GreetingRepository

type Greeting struct {
	Name string `form:"name"`
	Time time.Time
}

type GreetingRepository struct {
	storage map[time.Time]Greeting
	lock    *sync.RWMutex
}

func GetDefaultGreetingRepo() (repo *GreetingRepository, err error) {
	if DefaultGreetingRepo == nil {
		err = errors.New("GreetingRepository is not initialized!")
		return
	}
	repo = DefaultGreetingRepo
	return
}

func NewGreetingRepository() *GreetingRepository {
	return &GreetingRepository{
		storage: make(map[time.Time]Greeting),
		lock:    new(sync.RWMutex),
	}
}

func (r *GreetingRepository) Store(g Greeting) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.storage == nil {
		r.storage = make(map[time.Time]Greeting)
	}

	r.storage[g.Time.UTC()] = g
}

func (r *GreetingRepository) GetAll(ascending bool) map[time.Time]Greeting {
	return r.storage
}

func (r *GreetingRepository) GetSorted(ascending bool) []Greeting {
	r.lock.RLock()
	greetingsSorted := make([]Greeting, 0, len(r.storage))
	for _, g := range r.storage {
		greetingsSorted = append(greetingsSorted, g)
	}
	r.lock.RUnlock()

	var compareFunc func(int, int) bool
	if ascending {
		compareFunc = func(i, j int) bool {
			return greetingsSorted[i].Time.Before(greetingsSorted[j].Time)
		}
	} else {
		compareFunc = func(i, j int) bool {
			return greetingsSorted[i].Time.After(greetingsSorted[j].Time)
		}
	}
	sort.Slice(greetingsSorted, compareFunc)

	return greetingsSorted
}
