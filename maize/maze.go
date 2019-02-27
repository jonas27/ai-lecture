package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	maze     map[int][]int
	visited  map[int][]int
	parallel bool
	wg       = sync.WaitGroup{}
)

func main() {
	maze = initMap()
	parallel = true
	//wg = sync.WaitGroup{}
	//wg.Add(1)
	//for _, i := range maze {
	//	fmt.Println(i)
	//}
	bfs()
}

func initMap() map[int][]int {
	add := false
	length := 7
	width := 10
	maze := make(map[int][]int)
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
	return maze
}

func bfs() {
	r := 2
	c := 2
	visited = make(map[int][]int)
	for i := 0; i < 7; i++ {
		max := 20
		visited[i] = []int{max, max, max, max, max, max, max, max, max, max}
	}
	visited[r][c] = 0
	if parallel {
		defer func() {

			wg.Wait()
		}()
		fmt.Println("para")
		//wg.Add(1)
		searchParallel(1, r, c)
	} else {
		search(1, r, c)
	}
	//wg.Wait()
	time.Sleep(2 * time.Second)
	for l := 0; l < 7; l++ {
		fmt.Println(visited[l])
	}
}

func search(counter int, r int, c int) {
	if up(maze, visited, r, c, counter) {
		if visited[r-1][c] > counter {
			visited[r-1][c] = counter
		}
		search(counter+1, r-1, c)
	}
	if right(maze, visited, r, c, counter) {
		if visited[r][c+1] > counter {
			visited[r][c+1] = counter
		}
		search(counter+1, r, c+1)
	}
	if down(maze, visited, r, c, counter) {
		if visited[r+1][c] > counter {
			visited[r+1][c] = counter
		}
		search(counter+1, r+1, c)
	}
	if left(maze, visited, r, c, counter) {
		if visited[r][c-1] > counter {
			visited[r][c-1] = counter
		}
		search(counter+1, r, c-1)
	}
}

func searchParallel(counter int, r int, c int) {
	if up(maze, visited, r, c, counter) {
		defer wg.Done()
		if visited[r-1][c] > counter {
			visited[r-1][c] = counter
		}
		wg.Add(1)
		go search(counter+1, r-1, c)
	}
	if right(maze, visited, r, c, counter) {
		defer wg.Done()
		if visited[r][c+1] > counter {
			visited[r][c+1] = counter
		}
		wg.Add(1)
		go search(counter+1, r, c+1)
	}
	if down(maze, visited, r, c, counter) {
		defer wg.Done()
		if visited[r+1][c] > counter {
			visited[r+1][c] = counter
		}
		wg.Add(1)
		go search(counter+1, r+1, c)
	}
	if left(maze, visited, r, c, counter) {
		defer wg.Done()
		if visited[r][c-1] > counter {
			visited[r][c-1] = counter
		}
		wg.Add(1)
		go search(counter+1, r, c-1)
	}
	return
}

func up(maze map[int][]int, visited map[int][]int, r int, c int, counter int) bool {
	if r > 0 && maze[r-1][c] == 0 && visited[r-1][c] > counter {
		return true
	}
	return false
}
func down(maze map[int][]int, visited map[int][]int, r int, c int, counter int) bool {
	if r < 6 && maze[r+1][c] == 0 && visited[r+1][c] > counter {
		return true
	}
	return false
}
func right(maze map[int][]int, visited map[int][]int, r int, c int, counter int) bool {
	if c < 9 && maze[r][c+1] == 0 && visited[r][c+1] > counter {
		return true
	}
	return false
}
func left(maze map[int][]int, visited map[int][]int, r int, c int, counter int) bool {
	if c > 0 && maze[r][c-1] == 0 && visited[r][c-1] > counter {
		return true
	}
	return false
}
