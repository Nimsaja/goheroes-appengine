package db

import (
	"context"
	"testing"

	"github.com/lima1909/goheroes-appengine/service"
)

func TestList(t *testing.T) {
	m := NewMemService()

	fh, err := m.List(context.TODO(), "")
	if err != nil {
		t.Errorf("no err expected: %v", err)
	}

	if len(m.heroes) != len(fh) {
		t.Errorf("%v != %v", len(m.heroes), len(fh))
	}
}

func TestListFilter(t *testing.T) {
	m := NewMemService()

	fh, err := m.List(context.TODO(), "Alex M")
	if err != nil {
		t.Errorf("no err expected: %v", err)
	}
	if 1 != len(fh) {
		t.Errorf("%v != %v", 1, len(fh))
	}
}

func TestAdd(t *testing.T) {
	m := NewMemService()
	size := len(m.heroes)

	m.Add(context.TODO(), service.Hero{ID: 99, Name: "Test"})
	size = size + 1
	if len(m.heroes) != size {
		t.Errorf("%v != %v", len(m.heroes), size)
	}
}

func TestAddAndList(t *testing.T) {
	m := NewMemService()
	m.Add(context.TODO(), service.Hero{ID: 99, Name: "Alex M"})

	fh, err := m.List(context.TODO(), "Alex M")
	if err != nil {
		t.Errorf("no err expected: %v", err)
	}
	if 2 != len(fh) {
		t.Errorf("%v != %v", 2, len(fh))
	}
}

func TestDelete(t *testing.T) {
	m := NewMemService()
	size := len(m.heroes)

	m.Delete(context.TODO(), 1)
	size = size - 1
	if len(m.heroes) != size {
		t.Errorf("%v != %v", len(m.heroes), size)
	}
}

func TestAddAndGetAndDelete(t *testing.T) {
	m := NewMemService()
	size := len(m.heroes)

	// Add
	m.Add(context.TODO(), service.Hero{ID: 99, Name: "Test"})
	size = size + 1
	if len(m.heroes) != size {
		t.Errorf("%v != %v", len(m.heroes), size)
	}

	// Get
	h, err := m.GetByID(context.TODO(), 99)
	if err != nil {
		t.Errorf("no err expected: %v", err)
	}
	if h.ID != 99 {
		t.Errorf("99 != %v", h.ID)
	}

	// Del
	m.Delete(context.TODO(), 99)
	size = size - 1
	if len(m.heroes) != size {
		t.Errorf("%v != %v", len(m.heroes), size)
	}
}
