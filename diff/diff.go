package diff

import "fmt"

type diffTool struct{}

func NewDiffTool() *diffTool {
	return &diffTool{}
}

func (d *diffTool) LCS(X string, Y string) string {
	m := len(X) + 1
	n := len(Y) + 1
	C := createMatrix(m, n)
	for i := range X {
		nI := i + 1
		for j := range Y {
			nJ := j + 1
			if X[i] == Y[j] {
				diagonal := C[nI-1][nJ-1]
				C[nI][nJ] = diagonal + 1
			} else {
				xPrev := C[nI-1][nJ]
				yPrev := C[nI][nJ-1]
				C[nI][nJ] = max(xPrev, yPrev)
			}
		}
	}
	return backtrack(C, X, Y, m-1, n-1)
}

func (d *diffTool) LCSLine(X []string, Y []string) []string {
	m := len(X) + 1
	n := len(Y) + 1
	C := createMatrix(m, n)
	for i := range X {
		nI := i + 1
		for j := range Y {
			nJ := j + 1
			if X[i] == Y[j] {
				diagonal := C[nI-1][nJ-1]
				C[nI][nJ] = diagonal + 1
			} else {
				xPrev := C[nI-1][nJ]
				yPrev := C[nI][nJ-1]
				C[nI][nJ] = max(xPrev, yPrev)
			}
		}
	}
	return backtrackLine(C, X, Y, m-1, n-1)
}

func backtrack(C [][]int, X string, Y string, i int, j int) string {
	if i == 0 || j == 0 {
		return ""
	}
	nI := i - 1
	nJ := j - 1
	if X[nI] == Y[nJ] {
		return backtrack(C, X, Y, i-1, j-1) + string(X[nI])
	}
	if C[i][j-1] > C[i-1][j] {
		return backtrack(C, X, Y, i, j-1)
	}
	return backtrack(C, X, Y, i-1, j)
}

func backtrackLine(C [][]int, X []string, Y []string, i int, j int) []string {
	if i == 0 || j == 0 {
		return []string{}
	}
	nI := i - 1
	nJ := j - 1
	if X[nI] == Y[nJ] {
		return append(backtrackLine(C, X, Y, i-1, j-1), string(X[nI]))
	}
	if C[i][j-1] > C[i-1][j] {
		return backtrackLine(C, X, Y, i, j-1)
	}
	return backtrackLine(C, X, Y, i-1, j)
}

func createMatrix(m int, n int) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func printMatrix(matrix [][]int) string {
	str := ""
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			str = str + fmt.Sprint(matrix[i][j]) + " "
		}
		str = str + "\n"
	}
	return str
}
