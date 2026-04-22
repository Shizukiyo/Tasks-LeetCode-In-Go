// https://leetcode.com/problems/rotate-image/description/
package main

func rotate(matrix [][]int)  {
    n := make([][]int, len(matrix))
    for i := range matrix {
        n[i] = make([]int, len(matrix[i]))
        copy(n[i], matrix[i])
    }

    for i := range matrix {
        for j := range matrix {
            matrix[i][len(matrix)-j-1] = n[j][i]
        }
    }
}
