package just_to_offer

//go语言数组跟C++数据有些不同的地方就是，go更多的是用切片去实现并且有默认值,只能修改len(s)内的值才能作用到函数外面，无法确认默认值是否也是函数本来
// 的值，出现替换后的数组覆盖原数组的情况很难检查，看了bytes.Replace 实现，确实新建一个byte数组更加可靠
func ReplactBlack(s []byte) []byte {
	if s == nil || len(s) <= 0 {
		//避免对新值修改时修改到原s
		return append([]byte(nil), s...)
	}
	count := getByteCount(s, ' ')
	if count <= 0 {
		return append([]byte(nil), s...)
	}
	changeByte := []byte{'0', '2', '%'}
	newByte := make([]byte, len(s)+count*2)
	oldIndex := 0
	newIndex := 0
	for blackIndex, checkByte := range s {
		if checkByte != ' ' {
			continue
		}
		newIndex += copy(newByte[newIndex:], s[oldIndex:blackIndex])
		newIndex += copy(newByte[newIndex:], changeByte)
		oldIndex = blackIndex + 1
	}
	copy(newByte[newIndex:], s[oldIndex:])
	return newByte
}

func getByteCount(s []byte, old byte) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] == old {
			n++
		}
	}
	return n
}
