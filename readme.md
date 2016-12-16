#收集的常用 有用的Golang工具函数


##案例

```shell
//判断两个浮点数是否相等
func EqualFloaf64(floatOne, floatTwo, precision float64) bool {
	if precision < 0.0 {
		precision = math.SmallestNonzeroFloat64
	}
	return math.Abs(floatOne-floatTwo) <= (precision * math.Min(math.Abs(floatOne), math.Abs(floatTwo)))
}

//高效的Bytes 转 String
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//高效的String 转  Bytes
func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
```