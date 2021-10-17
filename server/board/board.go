package board

type Board struct {
	Grid [3][3]int
}

// Initizlie the board
func (board *Board) Init() [3][3]int {
	board = &Board{
		[3][3]int{{-1, -1, -1},
			{-1, -1, -1},
			{-1, -1, -1}}}
	return board.Grid
}

// places the players mark choices at the row, column position
func (board *Board) MakeMove(row int, col int, player int) {

	board.Grid[row][col] = player
}

// Return the wining combination
func (board *Board) getHands() [8][3]int {
	return [8][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6}}
}

// Minimax algorithm used to calculate computer's move
func (board *Board) Minimax(depth int, player bool) [3]int {

	var best [3]int
	if player {
		best = [3]int{-1, -1, 1000}
	} else {
		best = [3]int{-1, -1, -1000}
	}

	result, playerWin := board.IsGameOver()
	var reward = 0
	if depth == 0 || result {
		if result {
			if playerWin == 1 {
				reward = -1
			} else if playerWin == 0 {
				reward = 1
			}
		}

		return [3]int{-1, -1, reward}
	}

	for i := 0; i < 9; i++ {
		row := i / 3
		col := i % 3
		if board.Grid[row][col] != -1 {
			continue
		}
		var playerMark int

		if player {
			playerMark = 1
		} else {
			playerMark = 0
		}
		board.Grid[row][col] = playerMark

		score := board.Minimax(depth-1, !player)

		board.Grid[row][col] = -1

		score[0], score[1] = row, col

		if player && score[2] < best[2] {
			best = score
		}
		if !player && score[2] > best[2] {
			best = score
		}
	}
	return best
}

// Check any player wins
func (board *Board) IsGameOver() (bool, int) {

	for _, element := range board.getHands() {

		if board.Grid[element[0]/3][element[0]%3] == board.Grid[element[1]/3][element[1]%3] && board.Grid[element[0]/3][element[0]%3] == board.Grid[element[2]/3][element[2]%3] && board.Grid[element[0]/3][element[0]%3] != -1 {
			return true, board.Grid[element[0]/3][element[0]%3]
		}
	}
	return false, -1
}

// Calculates how many spaces are available in the board
func (board *Board) CalculateDepth() int {
	depth := 0
	for i := 0; i < 9; i++ {
		if board.Grid[i/3][i%3] == -1 {
			depth++
		}
	}
	return depth
}