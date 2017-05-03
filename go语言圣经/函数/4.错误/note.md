#错误处理策略

#1)直接传播错误

resp,err:=http.Get(url)
if err!=nil{
	return nil,err
}

doc,err:=html.Parse(resp.Body)
resp.Body.Close()
if err!=nil{
	return nil,fmt.Errorf("parsing %s as HTML:%v",url,err)
}

#fmt.Errorf函数使用fmt.Sprintf格式化错误信息并返回。

#2)当错误发生是偶然的，或由不可预知的问题导致的，一个明智的选择是重新尝试失败的操作。在重试时，我们需要限制重试的时间间隔或重试的次数，防止无限制的重试。


func WaitForServer(url string) error {
    const timeout = 1 * time.Minute
    deadline := time.Now().Add(timeout)
    for tries := 0; time.Now().Before(deadline); tries++ {
        _, err := http.Head(url)
        if err == nil {
            return nil // success
        }
        log.Printf("server not responding (%s);retrying…", err)
        time.Sleep(time.Second << uint(tries)) // exponential back-off
    }
    return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

#3)如果错误发生后，程序无法继续运行，我们可以输出错误信息并结束程序。
#应该在main函数做这个。

if err:=WaitForServer(url);err!=nil{
	fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
	os.Exit(1)
}

调用log.Fatalf可以更简洁，log中的所有函数，都默认会在错误信息之前输出时间信息。

if err := WaitForServer(url); err != nil {
    log.Fatalf("Site is down: %v\n", err)
}

2006/01/02 15:04:05 Site is down: no such domain:
bad.gopl.io

4)
有时，我们只需要输出错误信息就足够了，不需要中断程序的运行。我们可以通过log包提供函数

if err:=Ping();err!=nil{
	log.Printf("ping failed: %v; networking disabled",err)
}
或者标准错误输出错误信息
if err := Ping(); err != nil {
    fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
}

文件结尾错误EOF
io包会产生EOF错误

in := bufio.NewReader(os.Stdin)
for {
    r, _, err := in.ReadRune()
    if err == io.EOF {
        break // finished reading
    }
    if err != nil {
        return fmt.Errorf("read failed:%v", err)
    }
    // ...use r…
}


