package main

import (
	"fmt"
)

func aStar() {
	r := 2
	c := 2
	visited[r][c] = 0
	searchA(1, c, r)
	//time.Sleep(1*time.Second)
	for l := 0; l < 7; l++ {
		fmt.Println(visited[l])
	}
}

func searchA(counter int, r int, c int) {
	bd := bestDirection(r, c)
	for i := 0; i < 4; i++ {
		if up(maze, visited, r, c, counter) && bd[i] == i {
			goUp(counter, r, c)
		} else if right(maze, visited, r, c, counter) && bd[i] == i {
			goRight(counter, r, c)
		} else if down(maze, visited, r, c, counter) && bd[i] == i {
			goDown(counter, r, c)
		} else if left(maze, visited, r, c, counter) && bd[i] == i {
			goLeft(counter, r, c)
		}
	}
}

func bestDirection(r int, c int) [4]int {
	rEnd := 4
	cEnd := 6
	arr := [4]int{}
	if r > rEnd {
		if c < cEnd {
			arr[0] = 0
			arr[1] = 1
			arr[2] = 2
			arr[3] = 3
		} else {
			arr[0] = 0
			arr[1] = 2
			arr[2] = 1
			arr[3] = 3
		}
	} else {
		if c < cEnd {
			arr[0] = 3
			arr[1] = 1
			arr[2] = 2
			arr[3] = 0
		} else {
			arr[0] = 3
			arr[1] = 2
			arr[2] = 1
			arr[3] = 0
		}
	}
	return arr
}

func goUp(counter int, r int, c int) {
	if visited[r-1][c] > counter {
		visited[r-1][c] = counter
	}
	searchA(counter+1, r-1, c)
}

func goRight(counter int, r int, c int) {
	if visited[r][c+1] > counter {
		visited[r][c+1] = counter
	}
	searchA(counter+1, r, c+1)
}

func goDown(counter int, r int, c int) {
	if visited[r+1][c] > counter {
		visited[r+1][c] = counter
	}
	searchA(counter+1, r+1, c)
}

func goLeft(counter int, r int, c int) {
	if visited[r][c-1] > counter {
		visited[r][c-1] = counter
	}
	searchA(counter+1, r, c-1)
}
