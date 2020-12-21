package slicing

import (
	"errors"
)

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up

	initialDirection = Right
)

type Matrix struct {
	size   int
	slices [][]int
	used   [][]bool
}

//NewMatrix returns new Matrix entity
func NewMatrix(slices [][]int) (Matrix, error) {
	var m Matrix

	m.size = len(slices)

	for _, slice := range slices {
		if len(slice) != m.size {
			return *new(Matrix), errors.New("proposed slice is not of n*n size")
		}
	}

	m.slices = slices

	return m, nil
}

//Transform returns []int slice, that are calculated according to proposed algorithm
func (m *Matrix) Transform() []int {
	var row, col, usedElements int
	direction := initialDirection
	elementsNumber := m.size * m.size
	oneDimSlice := make([]int, 0, elementsNumber)

	m.used = make([][]bool, m.size)
	for i, _ := range m.slices {
		m.used[i] = make([]bool, m.size)
	}

	for usedElements < elementsNumber {
		oneDimSlice = append(oneDimSlice, m.slices[row][col])
		m.used[row][col] = true
		usedElements++

		switch direction {
		case Right:
			if col < m.size-1 && !m.used[row][col+1] {
				col++
			} else {
				row++
				direction = Down
			}

		case Down:
			if row < m.size-1 && !m.used[row+1][col] {
				row++
			} else {
				col--
				direction = Left
			}

		case Left:
			if col > 0 && !m.used[row][col-1] {
				col--
			} else {
				row--
				direction = Up
			}

		case Up:
			if row > 0 && !m.used[row-1][col] {
				row--
			} else {
				col++
				direction = Right
			}
		}

	}

	return oneDimSlice
}
