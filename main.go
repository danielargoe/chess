package main

import "fmt"

func main() {

	// declare variable(s)
	board := [][]string{
		{"r", "n", "b", "q", "k", "b", "n", "r"},
		{"p", "p", "p", "p", "p", "p", "p", "p"},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{"P", "P", "P", "P", "P", "P", "P", "P"},
		{"R", "N", "B", "K", "Q", "B", "N", "R"},
	}

	// display the menu
	display_menu()

	// display the board
	display_board(board)

}

// function to display menu
func display_menu() {

	// display menu
	fmt.Println(
		"\n\t\t\tWelcome to Chess!" +
			"\n\t2-Players black is lowercase, white is uppercase" +
			"\n\t  Enter commands such as i.e.: white -> a2->a4")

}

// function to display the chess board
func display_board(board [][]string) {

	// calculate length of rows
	row := len(board)

	// display the row
	fmt.Println("\n\t\t         a b c d e f g h")

	// display spacing
	fmt.Println("\t\t         ---------------")

	// iterate through each row on board
	for i := 0; i < len(board); i++ {

		// display current row number
		fmt.Printf("\t\t     %d | ", row)

		// iterate through each column on board
		for j := 0; j < len(board); j++ {

			// display current piece on board
			fmt.Printf("%s ", board[i][j])

		}

		// display current row number
		fmt.Printf("| %d\n", row)

		// decrement row
		row--

	}

	// display spacing
	fmt.Println("\t\t         ---------------")

	// display the row
	fmt.Println("\t\t         a b c d e f g h")

}
