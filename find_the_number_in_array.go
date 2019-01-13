package just_to_offer

//在有规律的二维数组中查找指定的值是否存在
func Find(target [][]int, wantNumber int) bool {
	if target == nil || len(target) <= 0 {
		return false
	}
	//检查这个二维切片是否满足二维数组
	tmpLen := 0
	for _, rowInfo := range target {
		if rowInfo == nil || len(rowInfo) <= 0 || (tmpLen != 0 && tmpLen != len(rowInfo)) {
			return false
		}
		tmpLen = len(rowInfo)
	}
	//从右上角开始检查
	col := len(target)
	row := len(target[0])
	for colIndex, rowIndex := 0, row-1; colIndex <= col-1 && rowIndex >= 0; {
		if target[colIndex][rowIndex] == wantNumber {
			return true
		} else if target[colIndex][rowIndex] < wantNumber {
			colIndex ++
		} else {
			rowIndex--
		}
	}
	return false
}
