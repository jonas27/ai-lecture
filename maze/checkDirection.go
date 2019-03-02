package main

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
