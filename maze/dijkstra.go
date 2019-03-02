package main

import (
	"fmt"
)

func dijkstra() {
	r := 2
	c := 2
	visited[r][c] = 0
	if parallel {
		fmt.Println("parallel")
		wg.Add(1)
		go searchParallel(1, r, c)
		wg.Wait()
	} else {
		search(1, r, c)
	}
	//time.Sleep(1*time.Second)
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
	defer wg.Done()
	//defer wg.Wait()
	if up(maze, visited, r, c, counter) {
		if visited[r-1][c] > counter {
			visited[r-1][c] = counter
		}
		fmt.Println(wg)
		fmt.Println("dasdasd")
		wg.Add(1)
		go search(counter+1, r-1, c)
		wg.Done()
	}
	if right(maze, visited, r, c, counter) {
		if visited[r][c+1] > counter {
			visited[r][c+1] = counter
		}
		wg.Add(1)
		go search(counter+1, r, c+1)
		wg.Done()
	}
	if down(maze, visited, r, c, counter) {
		if visited[r+1][c] > counter {
			visited[r+1][c] = counter
		}
		wg.Add(1)
		go search(counter+1, r+1, c)
		wg.Done()
	}
	if left(maze, visited, r, c, counter) {
		if visited[r][c-1] > counter {
			visited[r][c-1] = counter
		}
		wg.Add(1)
		go search(counter+1, r, c-1)
		wg.Done()
	}
}
