io.Writer类型是用的最广泛的接口之一，因为它提供了所有类型写入bytes的抽象，包括文件类型，内存缓冲区，网络连接，HTTP客户端，
压缩工具，哈希等等。io包中定义了很多其他有用的接口类型。Reader可以代表任意可以读取bytes的类型，Closer可以是任意可以关闭的值
例如一个文件或者网络连接。


package io
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Closer interface {
    Close() error
}

type ReadWriter interface {
    Reader
    Writer
}
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}