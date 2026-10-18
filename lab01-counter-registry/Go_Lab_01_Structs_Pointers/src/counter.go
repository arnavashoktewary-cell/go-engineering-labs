package lab01

import "errors"

// Counter represents a named nonnegative count.
type Counter struct {
	Name  string
	Value int
}

// NewCounter constructs a Counter or returns an error for invalid input.
// TODO: reject empty names and negative initial values.
func NewCounter(name string, initial int) (*Counter, error) {
	if name == "" || initial < 0 {
		return nil, errors.New("Empty Name")
	}

	c := Counter{name, initial}

	return &c, nil
}

// Increment adds one to the original Counter.
func (c *Counter) Increment() {

	c.Value = c.Value + 1

}

// Add adds a nonnegative delta. Negative values cause an error and no mutation.
func (c *Counter) Add(delta int) error {
	if delta < 0 {
		return errors.New("Negative Delta")
	}
	c.Value += delta

	return nil

}

// Snapshot returns a value copy of the current Counter.
func (c Counter) Snapshot() Counter {

	return c

}
