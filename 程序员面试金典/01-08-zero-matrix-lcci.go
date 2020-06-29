package main

import "fmt"

/**
面试题 01.08. 零矩阵

输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]

[0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
*/

func main() {
	matrix := [3][4]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	var hang []int
	var lie []int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				hang = append(hang, i)
				lie = append(lie, j)
			}
		}
	}
	// 最先想到的笨办法
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, v := range hang {
				if i == v {
					matrix[i][j] = 0
				}
			}
			for _, v := range lie {
				if j == v {
					matrix[i][j] = 0
				}
			}
		}
	}

	fmt.Println(matrix)
	fmt.Println(hang)
	fmt.Println(lie)

}
func s(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make(map[int]bool), make(map[int]bool)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true // // 第几行清0
				col[j] = true // 第几列清0
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if _, ok := row[i]; ok {
				matrix[i][j] = 0
			}
			if _, ok := col[j]; ok {
				matrix[i][j] = 0
			}
		}
	}

}

/**
map 的方式
*/
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make(map[int]int), make(map[int]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i]++
				col[j]++
			}
		}
	}

	for k, v := range row {
		if v > 0 {
			for j := 0; j < n; j++ {
				matrix[k][j] = 0
			}
		}
	}

	for k, v := range col {
		if v > 0 {
			for i := 0; i < m; i++ {
				matrix[i][k] = 0
			}
		}
	}
}

/**
使用两个数组记录需要清理的行和列 遇到需要清理的行定义新数组替换
*/
func setZeroes1(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row, col := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}
	for i := 0; i < m; i++ {
		if row[i] == 1 {
			matrix[i] = make([]int, n)
		}
		for j := 0; j < len(matrix[i]); j++ {
			if col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
