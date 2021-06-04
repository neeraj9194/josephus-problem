package main

import (
	"fmt"
)

type Person struct {
	pos    int
	isDead bool
}

type Circle struct {
	person []Person
	step   int // no of skips before killing
}

// Initilize circle with peoples and set the step
func (c *Circle) Init(num int, step int) {
	if num < 1 {
		panic(fmt.Errorf("number of people should be greater than or equal to 1"))
	}
	if step >= num || step < 1 {
		panic(fmt.Errorf("step per kill should be less then number of people and greater than 0"))
	}

	c.step = step
	c.person = make([]Person, num)
	for index := range c.person {
		c.person[index].pos = index
		c.person[index].isDead = false
	}
}

func (c *Circle) SimulateKills() Person {
	i := 0
	for {
		currentPerson := &c.person[i]
		nextPos := i + c.step

		if nextPos >= len(c.person) {
			// if array is finished circle back
			nextPos -= len(c.person)
		}
		if !currentPerson.isDead {
			targetPerson := &c.person[nextPos]
			// current persion kills target.
			if targetPerson.isDead {
				nextPos = nextAlive(c.person, nextPos)
				targetPerson = &c.person[nextPos]
			}
			targetPerson.isDead = true
			fmt.Printf("%v kills % v\n", currentPerson.pos+1, targetPerson.pos+1)
		}
		i = nextAlive(c.person, nextPos)
		if i == currentPerson.pos {
			// If only one erson is alive then return
			return c.person[i]
		}
	}
}

func nextAlive(p []Person, currentPos int) int {
	pos := currentPos + 1
	if pos == len(p) {
		pos = 0
	}
	for {
		if p[pos].isDead {
			pos += 1
			if pos == len(p) {
				pos = 0
			}
		} else {
			return pos
		}
	}
}
