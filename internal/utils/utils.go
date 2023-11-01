package utils

import (
	"image/color"
	"math"
)

type Vec2 struct {
	X, Y int
}

type Cell struct {
	Alive    bool
	Position Vec2
	Color    color.RGBA
}

func NewVec2(x, y int) Vec2 {
	return Vec2{x, y}
}

func InitGrid(Width int, Height int, CellSize int) [][]*Cell {
	grid := make([][]*Cell, Width/CellSize)

	for x := range grid {
		grid[x] = make([]*Cell, Height/CellSize)
		for y := range grid[x] {
			grid[x][y] = &Cell{
				Alive:    math.Mod(float64(x), 2) == 0,
				Position: NewVec2(x, y),
			}
		}
	}

	return grid
}
