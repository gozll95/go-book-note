#练习 7.2： 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，返回一个新的Writer类型把原来的Writer封装在里面和一个表示写入新的Writer字节数的int64类型指针

func CountingWriter(w io.Writer) (io.Writer, *int64)
#练习 7.3： 为在gopl.io/ch4/treesort (§4.4)的*tree类型实现一个String方法去展示tree类型的值序列。

