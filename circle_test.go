package main

import "testing"

func TestSimulateKillsStepPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected error. But not found any.")
		}
	}()
	circle := new(Circle)
	circle.Init(1, 1)
	t.Fatalf("Failed")
}

func TestSimulateKillsNumberPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected error. But not found any.")
		}
	}()
	circle := new(Circle)
	circle.Init(0, 1)
	t.Fatalf("Failed")
}

func TestSimulateKillsSmallCircle(t *testing.T) {
	circle := new(Circle)
	circle.Init(2, 1)
	alivePerson := circle.SimulateKills()
	if alivePerson.pos != 0 {
		t.Fatalf("Failed, expected: %v result: %v are not equal.", 0, alivePerson.pos)
	}
}

func TestSimulateKillsLargeCircle(t *testing.T) {
	circle := new(Circle)
	circle.Init(120, 1)
	alivePerson := circle.SimulateKills()
	if alivePerson.pos+1 != 113 {
		t.Fatalf("Failed, expected: %v result: %v are not equal.", 113, alivePerson.pos)
	}
}
