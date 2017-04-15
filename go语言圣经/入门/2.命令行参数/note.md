  输入来自于程序外部:文件、网络连接、其他程序的输出、敲键盘的用户、命令行参数或者其他类似输入员

  1）命令行参数
  os包的Args变量获取
  os.Args变量是一个[]string
  os.Args[0]:命令本身的名字
  os.Args[1]:传的第一个参数
  os.Args[1:len(os.Args)]<=>os.Args[1:]