// Sscanf 用于扫描 str 中的数据，并根据 format 指定的格式
// 将扫描出的数据填写到参数列表 a 中
// 当 r 中的数据被全部扫描完毕或者扫描长度超出 format 指定的长度时
// 则停止扫描（换行符会被当作空格处理）
func Sscanf(str string, format string, a ...interface{}) (n int, err error)