文件、目录、网络通讯

I/O的核心是接口io.Reader和io.Writer 

f,_:=os.Open("/etc/passwd")
r:=bufio.NewReader(f)
s,ok:=r.ReadString('\n')重输入中读取一行