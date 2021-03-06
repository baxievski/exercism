// Package matrix implements a solution for the matrix exercism problem
package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents a matrix of numbers
type Matrix [][]int

// New constructs a matrix of numbers
func New(i string) (Matrix, error) {
	rows := strings.Split(i, "\n")
	w := 0
	matrix := make([][]int, len(rows))
	for i, r := range rows {
		rEls := strings.Split(strings.TrimSpace(r), " ")
		if w == 0 {
			w = len(rEls)
		} else if len(rEls) != w {
			return nil, fmt.Errorf("matrix has different number of elements in rows")
		}
		row := make([]int, len(rEls))
		for j, e := range rEls {
			elem, err := strconv.Atoi(e)
			if err != nil {
				return nil, err
			}
			row[j] = elem
		}
		matrix[i] = row
	}

	return matrix, nil
}

// Rows returns the rows of the matrix
func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i, r := range m {
		rows[i] = make([]int, len(r))
		for j, e := range r {
			rows[i][j] = e
		}
	}
	return rows
}

// Cols returns the columns of the matrix
func (m Matrix) Cols() [][]int {
	columns := make([][]int, len(m[0]))
	for i := range m[0] {
		columns[i] = make([]int, len(m))
	}
	for i, r := range m {
		for j, e := range r {
			columns[j][i] = e
		}
	}
	return columns
}

// Set sets the value of an element in the matrix
func (m Matrix) Set(i, j, v int) bool {
	if i < 0 {
		return false
	}
	if j < 0 {
		return false
	}
	if i >= len(m) {
		return false
	}
	if j >= len(m[0]) {
		return false
	}
	m[i][j] = v
	return true
}
