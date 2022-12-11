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

func (i *Inspector) InspectRounds(rounds int, postProcess func(int) int) {
	for round := 0; round < rounds; round++ {
		for _, monkey := range i.Monkeys {
			monkey.ActivityLevel += len(monkey.Items)
			for j := 0; j < len(monkey.Items); j++ {
				item := monkey.Inspect(j)
				item = postProcess(item)

				if item%monkey.Test == 0 {
					i.Monkeys[monkey.PositiveThrow].addItem(item)
				} else {
					i.Monkeys[monkey.NegativeThrow].addItem(item)
				}
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

func (m *Monkey) Inspect(i int) int {
	if len(m.Items) == 0 {
		return -1
	}
	item := m.Items[i]
	return m.Operation(item)
}
