package ticTacToe

import (
	"fmt"
	"github.com/jonas27/a-star/astar"
	"testing"
)

func TestIni(t *testing.T) {
	ini(m, "me", "")
	m.PrintValues()

}

func TestSolveWithInput(t *testing.T) {
	//var r,c int
	m := astar.Init(0, 0, 3, 3)
	ini(m, "me", "")
	fmt.Println("Make an input. Should be two integers between 0 and 2 and space separated and press enter.")
	for i := 0; i < 5; i++ {
		for {
			//fmt.Scanf("%d%d/n",&r,&c)
			if mark(1, 1) {
				break
			}

		}
		solve()
	}

	m.PrintValues()
	//solve(m,r,c)
}
