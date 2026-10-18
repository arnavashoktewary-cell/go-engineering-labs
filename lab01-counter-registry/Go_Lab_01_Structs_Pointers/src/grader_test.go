package lab01

import (
	"reflect"
	"testing"
)

// Public autograder: 10 tests × 10 points. Please don't modify these tests.
// Implement the API in counter.go and registry.go; write your own tests in student_test.go.

func TestCounter_ConstructorValid(t *testing.T) {
	c, err := NewCounter("requests", 7)
	if err != nil {
		t.Fatalf("valid constructor returned unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("valid constructor returned nil Counter")
	}
	if c.Name != "requests" || c.Value != 7 {
		t.Errorf("NewCounter(requests, 7) = %+v; want name=requests and value=7", *c)
	}
	zero, err := NewCounter("zero", 0)
	if err != nil || zero == nil || zero.Name != "zero" || zero.Value != 0 {
		t.Errorf("initial value zero must be accepted: counter=%v error=%v", zero, err)
	}
}

func TestCounter_ConstructorInvalid(t *testing.T) {
	for _, tc := range []struct {
		name    string
		initial int
	}{
		{"", 0},
		{"", 5},
		{"bad", -1},
	} {
		got, err := NewCounter(tc.name, tc.initial)
		if err == nil || got != nil {
			t.Errorf("NewCounter(%q, %d) = (%v,%v); want (nil, non-nil error)", tc.name, tc.initial, got, err)
		}
	}
}

func TestCounter_Increment(t *testing.T) {
	c := Counter{Name: "events", Value: 3}
	c.Increment()
	c.Increment()
	if c.Value != 5 || c.Name != "events" {
		t.Errorf("after two increments: %+v; want Value=5, Name=events", c)
	}
}

func TestCounter_Add(t *testing.T) {
	c := Counter{Name: "events", Value: 8}
	if err := c.Add(3); err != nil || c.Value != 11 {
		t.Errorf("Add(3): err=%v value=%d; want nil and 11", err, c.Value)
	}
	if err := c.Add(0); err != nil || c.Value != 11 {
		t.Errorf("Add(0): err=%v value=%d; want nil and 11", err, c.Value)
	}
	if err := c.Add(-9); err == nil || c.Value != 11 {
		t.Errorf("Add(-9): err=%v value=%d; want non-nil error and unchanged value 11", err, c.Value)
	}
}

func TestCounter_Snapshot(t *testing.T) {
	original := Counter{Name: "events", Value: 12}
	copy := original.Snapshot()
	if copy != original {
		t.Errorf("Snapshot() = %+v; want %+v", copy, original)
	}
	copy.Name = "other"
	copy.Value = 999
	if original.Name != "events" || original.Value != 12 {
		t.Errorf("editing snapshot unexpectedly changed original: %+v", original)
	}
}

func TestRegistry_Register(t *testing.T) {
	r := NewRegistry()
	if r == nil {
		t.Fatal("NewRegistry returned nil")
	}
	if err := r.Register("alpha", 2); err != nil {
		t.Fatalf("Register valid name: %v", err)
	}
	c, ok := r.Get("alpha")
	if !ok || c.Name != "alpha" || c.Value != 2 {
		t.Errorf("Get after Register: counter=%+v found=%v", c, ok)
	}
	if err := r.Register("alpha", 55); err == nil {
		t.Error("duplicate registration must return an error")
	}
	if err := r.Register("", 10); err == nil {
		t.Error("empty name registration must return an error")
	}
	if err := r.Register("invalid", -3); err == nil {
		t.Error("negative initial value registration must return an error")
	}
	if len(r.Names()) != 1 {
		t.Errorf("failed registrations must not add counters; Names()=%v", r.Names())
	}
	c, ok = r.Get("alpha")
	if !ok || c.Value != 2 {
		t.Errorf("duplicate registration changed initial value: counter=%+v found=%v", c, ok)
	}
}

func TestRegistry_Increment(t *testing.T) {
	r := NewRegistry()
	if r == nil {
		t.Fatal("NewRegistry returned nil")
	}
	if err := r.Register("requests", 4); err != nil {
		t.Fatalf("Register: %v", err)
	}
	if err := r.Increment("requests"); err != nil {
		t.Errorf("Increment(existing) error: %v", err)
	}
	c, found := r.Get("requests")
	if !found || c.Value != 5 {
		t.Errorf("after increment, Get = %+v found=%v; want 5", c, found)
	}
	if err := r.Increment("unknown"); err == nil {
		t.Error("Increment(unknown) should return an error")
	}
	if len(r.Names()) != 1 {
		t.Errorf("Increment(unknown) should not create an entry: %v", r.Names())
	}
}

func TestRegistry_GetSnapshot(t *testing.T) {
	r := NewRegistry()
	if r == nil {
		t.Fatal("NewRegistry returned nil")
	}
	if err := r.Register("abc", 8); err != nil {
		t.Fatalf("Register: %v", err)
	}
	c, found := r.Get("abc")
	if !found || c.Name != "abc" || c.Value != 8 {
		t.Fatalf("Get(existing) = %+v, %v; want {abc 8},true", c, found)
	}
	c.Value = 900
	c.Name = "changed"
	other, found := r.Get("abc")
	if !found || other.Name != "abc" || other.Value != 8 {
		t.Errorf("Get must return an independent copy; got %+v found=%v", other, found)
	}
	missing, found := r.Get("not-here")
	if found || missing != (Counter{}) {
		t.Errorf("Get(missing) = %+v, %v; want Counter{}, false", missing, found)
	}
}

func TestRegistry_Names(t *testing.T) {
	r := NewRegistry()
	if r == nil {
		t.Fatal("NewRegistry returned nil")
	}
	for _, name := range []string{"zebra", "alpha", "middle", "Beta"} {
		if err := r.Register(name, 1); err != nil {
			t.Fatalf("Register(%q): %v", name, err)
		}
	}
	want := []string{"Beta", "alpha", "middle", "zebra"} // Go string ordering
	got := r.Names()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Names()=%v; want sorted %v", got, want)
	}
	if len(got) != 4 {
		t.Fatalf("Names length=%d; want 4", len(got))
	}
	got[0] = "tampered"
	if !reflect.DeepEqual(r.Names(), want) {
		t.Errorf("caller modifying Names result must not affect registry")
	}
}

func TestRegistry_Total(t *testing.T) {
	r := NewRegistry()
	if r == nil {
		t.Fatal("NewRegistry returned nil")
	}
	if got := r.Total(); got != 0 {
		t.Errorf("empty Total=%d; want 0", got)
	}
	for _, tc := range []struct {
		name string
		n    int
	}{
		{"a", 2}, {"b", 8}, {"c", 0},
	} {
		if err := r.Register(tc.name, tc.n); err != nil {
			t.Fatalf("Register(%q): %v", tc.name, err)
		}
	}
	if got := r.Total(); got != 10 {
		t.Errorf("Total=%d; want 10", got)
	}
	if err := r.Increment("a"); err != nil {
		t.Errorf("Increment(a): %v", err)
	}
	if got := r.Total(); got != 11 {
		t.Errorf("Total after increment=%d; want 11", got)
	}
}
