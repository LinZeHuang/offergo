package just_to_offer

import "sync"

//在go实现里面，有一个很方便的sync.Once，在单例模式里面时最好用的
type SingleStruct struct{}

var singleInfo *SingleStruct
var onceMu sync.Once

func GetSingleInfo()*SingleStruct{
	onceMu.Do(func() {
		singleInfo = &SingleStruct{}
	})
	return singleInfo
}

//这里也实现其他几种简单的实现方式
//1.单协程安全
func GetSingleInfoOnlyOneCoroutine()*SingleStruct{
	if singleInfo == nil{
		singleInfo= &SingleStruct{}
	}
	return singleInfo
}

//2.加锁实现
var mu sync.Mutex
func GetSingleInfoWithMu()*SingleStruct{
	mu.Lock()
	defer mu.Unlock()
	if singleInfo == nil{
		singleInfo= &SingleStruct{}
	}
	return singleInfo
}

//3.加锁优化
func GetSingleInfoWithMuBetter()*SingleStruct{
	if singleInfo == nil{
		mu.Lock()
		defer mu.Unlock()
		if singleInfo == nil{
			singleInfo= &SingleStruct{}
		}
	}
	return singleInfo
}

