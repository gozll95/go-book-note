允许多个只读操作并行执行，但写操作会完全互斥-->"多读单写"锁   sync.RWMutex 

读 和 读  无锁
读 和 写  会上锁 可以默认是串行了


场景:
 	在100刀的存款消失时不做记录多少还是会让我们有一些恐慌，Bob写了一个程序，每秒运行几百次来检查他的银行余额。他会在家，在工作中，甚至会在他的手机上来运行这个程序。银行注意到这些陡增的流量使得存款和取款有了延时，因为所有的余额查询请求是顺序执行的，这样会互斥地获得锁，并且会暂时阻止其它的goroutine运行。

	由于Balance函数只需要读取变量的状态，所以我们同时让多个Balance调用并发运行事实上是安全的，只要在运行的时候没有存款或者取款操作就行。


var mu sync.RWMutex 
var balance int 
func Balance()int{
	mu.RLock()  // readers lock
	defer mu.RUnlock()
	return balance 
}