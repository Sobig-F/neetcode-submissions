func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if matrix[0][0] > target || matrix[len(matrix) - 1][len(matrix[0]) - 1] < target {
		return false
	}

	if len(matrix) == 1 {
		if len(matrix[0]) == 1 && matrix[0][0] != target {
			return false
		}

		left := 0
		right := len(matrix[0])
		center := (right - left) / 2 - 1
		
		for {
			if right - left == 1 {
				return matrix[0][left] == target
			}
			if matrix[0][left + center] == target {
				return true
			} else if matrix[0][left + center] > target {
				right = left + center + 1
				center = (right - left) / 2 - 1
			} else if matrix[0][left + center] < target {
				left += center + 1
				center = (right - left) / 2 - 1
			}
		}
		return false
	}

	target_matrix := matrix[len(matrix) / 2 - 1]

	if target_matrix[0] > target {
		return searchMatrix(matrix[:len(matrix) / 2 - 1], target)
	} else if target_matrix[len(target_matrix) - 1] < target {
		return searchMatrix(matrix[len(matrix) / 2:], target)
	} else {
		return searchMatrix([][]int{target_matrix}, target)
	}
	return false
}
