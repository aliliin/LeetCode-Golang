package main

import "fmt"

/**
旋转矩阵
*/
func main() {
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)
	if len(matrix) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > i {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
			if j >= n/2 {
				matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
			}
		}
	}
	fmt.Println(matrix)
}
