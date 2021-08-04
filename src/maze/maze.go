package main

import (
	"fmt"
	"os"
)

/**
从文件中读取数组
在windows下的换行符是\r\n，所以遇到\r会读一个0进去，在linux下换行符只有\n
 idea 右下角切换line separator
*/
func readMaze(fileName string) [][]int {
	file, e := os.Open(fileName)
	if e != nil {
		panic(e)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row) //当成每个元素是[]int类型的一元数组 长度为row
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

/**
对点的抽象
*/
type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(dir point) point {
	return point{p.i + dir.i, p.j + dir.j}
}

// 判断下一个值是否有效 有效则返回数组值
func (p point) at(grid [][]int) (int, bool) {

	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true

}

/**
走迷宫
*/
func walk(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start} //待探索的点的队列

	for len(Q) > 0 {
		cur := Q[0] //取出一个点 进行探索
		Q = Q[1:]
		//遇到终点停止
		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			//看是否是走过的坐标
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			// 回到原点排除
			if next == start {
				continue
			}

			// 取当前步骤数
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			Q = append(Q, next)

		}

	}
	return steps

}

func main() {
	maze := readMaze("maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d", val)
		}
		fmt.Println()
	}
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, col := range row {
			fmt.Printf("%3d", col)
		}
		fmt.Println()
	}
}
