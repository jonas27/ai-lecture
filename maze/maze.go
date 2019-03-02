package main

import (
	"sync"
)

var (
	maze     = make(map[int][]int)
	visited  = make(map[int][]int)
	parallel bool
	wg       = sync.WaitGroup{}
)

func main() {
	initMap()
	initVisitedMap()
	parallel = true
	//wg = sync.WaitGroup{}
	//wg.Add(1)
	//for _, i := range maze {
	//	fmt.Println(i)
	//}
	//dijkstra()
	aStar()
}

func initMap() {
	add := false
	length := 7
	width := 10
	walls := []int{1, 11, 12, 14, 16, 17, 18, 24, 28, 32, 33, 34, 36, 37, 38, 51, 52, 54, 55, 56, 58, 62}
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			for _, a := range walls {
				if a == i*10+j {
					add = true
					break
				}
			}
			if add {
				maze[i] = append(maze[i], 1)
				add = false
			} else {
				maze[i] = append(maze[i], 0)
			}
		}
	}
}

func initVisitedMap() {
	for i := 0; i < 7; i++ {
		max := 20
		visited[i] = []int{max, max, max, max, max, max, max, max, max, max}
	}
}
