package lab01

import (
	"errors"
	"slices"
)

// Registry manages named counters.
// Its map is unexported: callers use the methods below.
type Registry struct {
	counters map[string]*Counter
}

// NewRegistry returns a ready-to-use registry.
func NewRegistry() *Registry {
	return &Registry{counters: make(map[string]*Counter)} // Starter; keep or revise
}

// Register adds a new counter. Invalid or duplicate names must return an error.
func (r *Registry) Register(name string, initial int) error {
	//if else condition checks for endge cases like empty and repeated names
	if name == "" {
		return errors.New("Empty Name")
	} else if _, exists := r.counters[name]; exists {
		return errors.New("Already Exists")
	} else if initial < 0 {
		return errors.New("Negative initial")
	}

	r.counters[name] = &Counter{Name: name, Value: initial}

	return nil

}

// Increment increments an existing counter; unknown names cause an error.
func (r *Registry) Increment(name string) error {
	//iff name belong to counters then increments value by 1
	c, exists := r.counters[name]
	if !exists {

		return errors.New("Name Not Found")
	}

	c.Value++

	return nil

}

// Get returns a COPY of the counter and whether it was found.
func (r *Registry) Get(name string) (Counter, bool) {
	// gets map of counter type
	// checks if name exists in map
	c, exists := r.counters[name]

	if exists {
		return *c, exists
	}

	return Counter{}, exists

}

// Names returns all names in ascending lexicographic order.
// It must return a fresh slice, independent of registry internals.
func (r *Registry) Names() []string {
	//new slice of names
	names := []string{}

	for name := range r.counters {
		names = append(names, name)
	}

	slices.Sort(names)

	return names
}

// Total returns the sum of all registered values.
func (r *Registry) Total() int {
	//accumulator to register the summation
	total := 0

	for _, c := range r.counters {
		total += c.Value
	}
	return total
}
