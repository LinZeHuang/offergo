package just_to_offer

import "errors"

//对应剑指offer面试题3

var NotFindErr = errors.New("找不到相同的数字")
//首先是能而外开辟一个map的内存的方法
func FindTheSameNumberInTheArray(target []int) (int, error) {
	if target == nil || len(target) <= 0 {
		return 0, NotFindErr
	}
	checkMap := make(map[int]bool, len(target))
	for _, checkNumber := range target {
		if checkMap[checkNumber] {
			return checkNumber, nil
		}
		checkMap[checkNumber] = true
	}
	return 0, NotFindErr
}

//不开辟而外内存，但是会修改到target的值，已经不是纯函数了
func FindTheSameNubmerInTheArrayNotWithMap(target []int) (int, error) {
	if target == nil || len(target) <= 0 {
		return 0, NotFindErr
	}
	for checkKey, checkNumber := range target {
		if checkKey != checkNumber {
			if checkNumber == target[checkNumber] {
				return checkNumber, nil
			}
			target[checkKey], target[checkNumber] = target[checkNumber], target[checkKey]
		}
	}
	return 0, NotFindErr
}
