func Parse(input string)(s *Syntax,err error){
	defer func(){
		if p:=recover();p!=nil{
			err=fmt.Errorf("internal error:%v",[])
		}
	}()

	//...parser...


}

//我们也可以通过调用runtime.Stack往错误信息中添加完整的堆栈调用信息。