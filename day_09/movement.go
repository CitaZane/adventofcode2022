package main

type Movement struct {
	Direction string
	Count     int
}

type Position struct {
	X int
	Y int
}

type Simulator struct {
	Movements []Movement

	Rope []Position

	// Tail Position
	// Head Position

	PlacesVisited map[Position]bool
}

func NewSimulator(movements []Movement, ropeLength int) Simulator {
	startPosition := Position{X: 0, Y: 0}
	rope := make([]Position, ropeLength)
	for i := 0; i < ropeLength; i++ {
		rope[i] = startPosition
	}

	return Simulator{
		Movements: movements,
		Rope:      rope,
		// Tail: startPosition,
		// Head: startPosition,

		PlacesVisited: map[Position]bool{startPosition: true},
	}
}

func (s *Simulator) ExecuteMovements() {
	for _, move := range s.Movements {
		s.Movement(move)
	}
}

func (s *Simulator) Movement(move Movement) {
	for i := 0; i < move.Count; i++ {
		s.MoveHead(move.Direction, 0)
		for j := 0; j < len(s.Rope)-1; j++ {
			s.MoveTail(move.Direction, j+1, j)
		}
	}
}

func (s *Simulator) MoveHead(direction string, head int) {
	switch direction {
	case "U":
		s.Rope[head].Y += 1
	case "D":
		s.Rope[head].Y -= 1
	case "L":
		s.Rope[head].X -= 1
	case "R":
		s.Rope[head].X += 1
	}
}

func (s *Simulator) MoveTail(direction string, tail int, head int) {
	if move, direction := s.TailNeedsToFollow(head, tail); move {
		switch direction {
		case "U":
			s.Rope[tail].Y += 1
		case "D":
			s.Rope[tail].Y -= 1
		case "L":
			s.Rope[tail].X -= 1
		case "R":
			s.Rope[tail].X += 1
		case "UR":
			s.Rope[tail].Y += 1
			s.Rope[tail].X += 1
		case "DR":
			s.Rope[tail].Y -= 1
			s.Rope[tail].X += 1
		case "UL":
			s.Rope[tail].Y += 1
			s.Rope[tail].X -= 1
		case "DL":
			s.Rope[tail].Y -= 1
			s.Rope[tail].X -= 1
		}
		if tail == len(s.Rope)-1 {
			s.PlacesVisited[s.Rope[tail]] = true
		}
	}
}

func (s *Simulator) TailNeedsToFollow(head, tail int) (bool, string) {
	if s.Rope[tail].X == s.Rope[head].X {
		if s.Rope[tail].Y == s.Rope[head].Y { //the same position
			return false, ""
		}
		if s.Rope[tail].Y > s.Rope[head].Y { //tail upwards
			if s.Rope[head].Y+1 == s.Rope[tail].Y {
				return false, ""
			} else {
				return true, "D"
			}
		}
		if s.Rope[tail].Y < s.Rope[head].Y {
			if s.Rope[head].Y-1 == s.Rope[tail].Y {
				return false, ""
			} else {
				return true, "U"
			}
		}
	}
	if s.Rope[tail].Y == s.Rope[head].Y {
		if s.Rope[tail].X > s.Rope[head].X { //tail upwards
			if s.Rope[head].X+1 == s.Rope[tail].X {
				return false, ""
			} else {
				return true, "L"
			}
		}
		if s.Rope[tail].X < s.Rope[head].X { //tail upwards
			if s.Rope[head].X-1 == s.Rope[tail].X {
				return false, ""
			} else {
				return true, "R"
			}
		}
	}
	if s.Rope[head].Y > s.Rope[tail].Y { //upper diognals
		if s.Rope[tail].Y+1 == s.Rope[head].Y { //only one
			if s.Rope[head].X+1 < s.Rope[tail].X {
				return true, "UL"
			} else if s.Rope[head].X-1 > s.Rope[tail].X {
				return true, "UR"
			}
			return false, ""
		} else {
			if s.Rope[head].X > s.Rope[tail].X {
				return true, "UR"
			} else {
				return true, "UL"
			}
		}
	}
	if s.Rope[head].Y < s.Rope[tail].Y { //lower diognals
		if s.Rope[tail].Y-1 == s.Rope[head].Y {
			if s.Rope[head].X+1 < s.Rope[tail].X {
				return true, "DL"
			} else if s.Rope[head].X-1 > s.Rope[tail].X {
				return true, "DR"
			}
			return false, ""
		} else if s.Rope[tail].X > s.Rope[head].X {
			return true, "DL"
		} else {
			return true, "DR"
		}
	}
	return false, ""
}
