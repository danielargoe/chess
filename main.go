package main

import (
	"fmt"
)

func main() {

	// declare variable(s)
	var board [][]byte
	var white byte
	var black byte
	var winner byte

	// initialize variable(s)
	white = 'w'
	black = 'b'

	// display the menu
	display_menu()

	// setup board
	board = setup_board()

	// display the board
	display_board(board)

	// iterate until someone wins
	play_chess(board, white, black, &winner)

	// display winner
	display_winner(winner)

}

// function to display the winner
func display_winner(winner byte) {

	// determine who won
	if winner == 'w' {

		// display that white won
		fmt.Printf("\nWhite wins!")

	} else {

		// display that black won
		fmt.Printf("\nBlack wins!")

	}

}

// function to get the current player's move
func get_move(current_player byte) string {

	// declare variable(s)
	var move string

	// display output
	fmt.Printf("\n%c pick a move: > ", current_player)

	// get move from input
	fmt.Scan(&move)

	// return move
	return move

}

// function to valid the current player's move
func is_valid(current_player byte, current_move string) bool {

	// determine if move has length of 6 for now
	return len(current_move) == 6

}

// function to perform moves on the board
func perform_move(current_move string, board [][]byte) {

	// declare variable(s)
	col_lookup := map[byte]int{
		'a': 0,
		'b': 1,
		'c': 2,
		'd': 3,
		'e': 4,
		'f': 5,
		'g': 6,
		'h': 7,
	}
	row_lookup := map[byte]int{
		'8': 0,
		'7': 1,
		'6': 2,
		'5': 3,
		'4': 4,
		'3': 5,
		'2': 6,
		'1': 7,
	}

	// get source column
	source_col := col_lookup[current_move[0]]

	// get source row
	source_row := row_lookup[current_move[1]]

	// get destination column
	dest_col := col_lookup[current_move[4]]

	// get destination row
	dest_row := row_lookup[current_move[5]]

	// move piece
	board[dest_row][dest_col] = board[source_row][source_col]

	// replace source piece
	board[source_row][source_col] = '.'

}

// function to play chess
func play_chess(board [][]byte, white byte, black byte, winner *byte) {

	// declare variable(s)
	var moves_white []string
	var moves_black []string
	var moves_count int
	var current_player byte
	var current_move string
	var valid_move bool

	// initialize variable(s)
	current_player = white
	moves_count = 1

	// display board
	display_board(board)

	// infinite loop until game completes
	for {

		// display current move number
		fmt.Printf("\nMove: #%d\n", moves_count)

		// get current move
		current_move = get_move(current_player)

		// validate move
		valid_move = is_valid(current_player, current_move)

		// invalid move
		if !valid_move {

			// display invalid input
			fmt.Println("Invalid move: Pick a different move")
			continue

		}

		// perform move
		perform_move(current_move, board)

		// increment moves_count
		moves_count++

		// store piece in corresponding slice
		if current_player == 'w' {

			// store in moves_white slice
			moves_white = append(moves_white, current_move)

		} else {

			// store in moves_black slice
			moves_black = append(moves_black, current_move)

		}

		// display board
		display_board(board)

		// validate checkmate
		// is_checkmate()

		// update current player
		if current_player == white {

			// update players for next turn
			current_player = black

		} else {

			// update players for next turn
			current_player = white

		}

	}

}

// function to setup the board
func setup_board() [][]byte {

	// return the template for the board
	return [][]byte{
		{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'},
		{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
		{'R', 'N', 'B', 'K', 'Q', 'B', 'N', 'R'},
	}

}

// function to display menu
func display_menu() {

	// display menu
	fmt.Println(
		"\nWelcome to Chess!" +
			"\n2-Players black is lowercase, white is uppercase" +
			"\nEnter commands such as i.e.: white -> a2->a4" +
			"\nWhite always starts first.")

}

// function to display the chess board
func display_board(board [][]byte) {

	// calculate length of rows
	row := len(board)

	// display the columns
	fmt.Println("\n    a b c d e f g h")

	// display spacing
	fmt.Println("    ---------------")

	// iterate through each column on board
	for i := 0; i < len(board); i++ {

		// display current row number
		fmt.Printf("%d | ", row)

		// iterate through each row on board
		for j := 0; j < len(board); j++ {

			// display current piece on board
			fmt.Printf("%c ", board[i][j])

		}

		// display current row number
		fmt.Printf("| %d\n", row)

		// decrement row
		row--

	}

	// display spacing
	fmt.Println("    ---------------")

	// display the columns
	fmt.Println("    a b c d e f g h")

}
