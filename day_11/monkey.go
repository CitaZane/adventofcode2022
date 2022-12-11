package main

import (
	"sort"
)

type Monkey struct {
	Items         []int
	Operation     func(int) int
	Test          int
	PositiveThrow int
	NegativeThrow int

	ActivityLevel int
}

type Inspector struct {
	Monkeys []*Monkey
}

func (i *Inspector) InspectRounds(rounds int) {
	for round := 0; round < rounds; round++ {
		for _, monkey := range i.Monkeys {
			for j := 0; j < len(monkey.Items); j++ {
				item, next := monkey.Inspect(j)
				if next == -1 {
					break
				}
				i.Monkeys[next].addItem(item)
				if j == len(monkey.Items)-1 {
					monkey.Items = monkey.Items[j+1:]
				}
			}
		}
	}
}

func (i *Inspector) ActivityCount() int {
	bussiness := []int{}
	for _, m := range i.Monkeys {
		bussiness = append(bussiness, m.ActivityLevel)
	}
	sort.Ints(bussiness)
	num1 := bussiness[len(bussiness)-1]
	num2 := bussiness[len(bussiness)-2]
	return num1 * num2
}

func (m *Monkey) addItem(item int) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) Inspect(i int) (int, int) {
	if len(m.Items) == 0 {
		return -1, -1
	}
	item := m.Items[i]
	item = m.Operation(item)
	item = item / 3
	m.ActivityLevel += 1

	if item%m.Test == 0 {
		return item, m.PositiveThrow
	} else {
		return item, m.NegativeThrow
	}
}
