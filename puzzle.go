package main

import (
	"fmt"
)

type move struct {
	from int
	to   int
	over int
}

func (m move) String() string {
	return fmt.Sprintf("Move from %v, to %v, over %v", m.from, m.to, m.over)
}

//   0
//  1 2
// 3 4 5
//6 7 8 9
type board [10]bool

func (b board) String() (ans string) {
	ans += "\n   "
	if b[0] {
		ans += "X\n"
	} else {
		ans += " \n"
	}

	ans += "  "
	if b[1] {
		ans += "X "
	} else {
		ans += "  "
	}
	if b[2] {
		ans += "X\n"
	} else {
		ans += " \n"
	}

	ans += " "
	if b[3] {
		ans += "X "
	} else {
		ans += "  "
	}
	if b[4] {
		ans += "X "
	} else {
		ans += "  "
	}
	if b[5] {
		ans += "X\n"
	} else {
		ans += " \n"
	}

	if b[6] {
		ans += "X "
	} else {
		ans += "  "
	}
	if b[7] {
		ans += "X "
	} else {
		ans += "  "
	}
	if b[8] {
		ans += "X "
	} else {
		ans += "  "
	}
	if b[9] {
		ans += "X\n"
	} else {
		ans += " \n"
	}

	return
}

//   0
//  1 2
// 3 4 5
//6 7 8 9
var moveList = [...]move{
	move{3, 0, 1},
	move{5, 0, 2},
	move{6, 1, 3},
	move{8, 1, 4},
	move{7, 2, 4},
	move{9, 2, 5},
	move{0, 3, 1},
	move{5, 3, 4},
	move{0, 5, 2},
	move{3, 5, 4},
	move{1, 6, 3},
	move{8, 6, 7},
	move{2, 7, 4},
	move{9, 7, 8},
	move{6, 8, 7},
	move{1, 8, 4},
	move{7, 9, 8},
	move{2, 9, 5},
}

func init() {
}

func getMoves(b board) []board {
	list := make([]board, 0)
	for _, m := range moveList {

		if b[m.from] && b[m.over] && !b[m.to] {
			fmt.Println(m)
			temp := b
			temp[m.over] = false
			temp[m.from] = false
			temp[m.to] = true
			list = append(list, temp)
		}
	}

	return list
}

func main() {
	println("Hello world")
	fmt.Printf("Boo\n")
	aMove := move{0, 0, 0}
	fmt.Printf("move: %v\n", aMove)

	aBoard := board{false, true, true, true, true, true, true, true, true, true}
	fmt.Printf("%v\n", aBoard)
	fmt.Printf("%v\n", getMoves(aBoard))
}
