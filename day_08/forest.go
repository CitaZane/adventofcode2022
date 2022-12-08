package main

type Forest struct {
	Map          [][]int
	Width        int
	Height       int
	VisibleTrees int
	ScenicScore  int
}

func newForest() Forest {
	return Forest{
		Map:         [][]int{},
		ScenicScore: 0,
	}
}

func (forest *Forest) findTopScenicScore() {
	for col := range forest.Map {
		for row := range forest.Map[col] {
			score := forest.CountScernicScore(col, row)

			if score > forest.ScenicScore {
				forest.ScenicScore = score
			}
		}
	}
}
func (forest *Forest) CountScernicScore(col, row int) int {
	if forest.treeIsOnEdge(col, row) {
		return 0
	}
	treeHeight := forest.Map[col][row]
	leftScore := forest.leftScenicScore(col, row, treeHeight)
	rightScore := forest.rightScenicScore(col, row, treeHeight)
	topScore := forest.topScenicScore(col, row, treeHeight)
	bottomScore := forest.bottomScenicScore(col, row, treeHeight)

	return leftScore * rightScore * topScore * bottomScore
}

func (forest *Forest) leftScenicScore(col, row, treeHeight int) int {
	score := 0
	for i := row - 1; i >= 0; i-- {
		if forest.Map[col][i] < treeHeight {
			score += 1
		} else {
			return score + 1
		}
	}
	return score
}

func (forest *Forest) rightScenicScore(col, row, treeHeight int) int {
	score := 0
	for i := row + 1; i < forest.Width; i++ {
		if forest.Map[col][i] < treeHeight {
			score += 1
		} else {
			return score + 1
		}
	}
	return score
}

func (forest *Forest) topScenicScore(col, row, treeHeight int) int {
	score := 0
	for i := col - 1; i >= 0; i-- {
		if forest.Map[i][row] < treeHeight {
			score += 1
		} else {
			return score + 1
		}
	}
	return score
}

func (forest *Forest) bottomScenicScore(col, row, treeHeight int) int {
	score := 0
	for i := col + 1; i < forest.Height; i++ {
		if forest.Map[i][row] < treeHeight {
			score += 1
		} else {
			return score + 1
		}
	}
	return score
}

func (forest *Forest) CountVisibleTrees() {
	for col := range forest.Map {
		for row := range forest.Map[col] {
			forest.CheckIfVisible(col, row)
		}
	}
}

func (forest *Forest) CheckIfVisible(col, row int) {
	if forest.treeIsOnEdge(col, row) {
		forest.VisibleTrees += 1
		return
	}
	forest.innerTreeVisibility(row, col)

}

func (forest *Forest) treeIsOnEdge(row, col int) bool {
	return col == 0 || row == 0 || col == forest.Height-1 || row == forest.Width-1
}

func (forest *Forest) innerTreeVisibility(row, col int) {
	treeHeight := forest.Map[col][row]

	if visible := forest.checkLeft(row, col, treeHeight); visible {
		forest.VisibleTrees += 1
		return
	}
	if visible := forest.checkRight(row, col, treeHeight); visible {
		forest.VisibleTrees += 1
		return
	}

	if visible := forest.checkTop(row, col, treeHeight); visible {
		forest.VisibleTrees += 1
		return
	}
	if visible := forest.checkBottom(row, col, treeHeight); visible {
		forest.VisibleTrees += 1
		return
	}
}

func (forest *Forest) checkLeft(row, col, treeHeight int) bool {
	for i := 0; i < row; i++ {

		if forest.Map[col][i] >= treeHeight {
			return false
		}
	}
	return true
}

func (forest *Forest) checkRight(row, col, treeHeight int) bool {
	for i := forest.Width - 1; i > row; i-- {

		if forest.Map[col][i] >= treeHeight {
			return false
		}
	}
	return true
}
func (forest *Forest) checkTop(row, col, treeHeight int) bool {
	for i := 0; i < col; i++ {
		if forest.Map[i][row] >= treeHeight {
			return false
		}
	}
	return true
}
func (forest *Forest) checkBottom(row, col, treeHeight int) bool {
	for i := forest.Height - 1; i > col; i-- {

		if forest.Map[i][row] >= treeHeight {
			return false
		}
	}
	return true
}
