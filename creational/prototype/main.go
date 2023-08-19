package main

import "fmt"

type Cloneable interface {
	clone() Cloneable
}

type Maze struct {
	name string
}

type MazeGame struct {
	mazeProto Maze
}

func (m MazeGame) createMaze() Maze {
	return Maze{m.mazeProto.name + "_clone"}
}

func main() {
	mazeGame := MazeGame{}
	mazeGame.mazeProto.name = "maze1"

	maze := mazeGame.createMaze()

	fmt.Println(maze.name)
}
