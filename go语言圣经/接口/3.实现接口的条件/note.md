一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口


// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
//接口被类型实现=>类型能实现接口的所有方法->接口实例化=类型实例化


var w io.Writer
w = os.Stdout           // OK: *os.File has Write method
w = new(bytes.Buffer)   // OK: *bytes.Buffer has Write method
w = time.Second         // compile error: time.Duration lacks Write method

var rwc io.ReadWriteCloser
rwc = os.Stdout         // OK: *os.File has Read, Write, Close methods
rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method


这个规则甚至适用于等式右边本身也是一个接口类型

w = rwc                 // OK: io.ReadWriteCloser has Write method
rwc = w                 // compile error: io.Writer lacks Close method

