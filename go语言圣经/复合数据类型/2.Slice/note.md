slice 包含 指针、长度、容量

复制slice,其实就是创建一个别名，还是对本身的操作

因为slice值包含指向第一个slice元素的指针，因此向函数传递slice将允许在函数内部修改底层数组的元素。换句话说，复制一个slice只是对底层的数组创建了一个新的slice别名（§2.3.2）。下面的reverse函数在原内存空间将[]int类型的slice反转，而且它可以用于任意长度的slice。


func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}


\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
一个slice可以用来模拟一个stack。最初给定的空slice对应一个空的stack，然后可以使用append函数将新的值压入stack：

stack = append(stack, v) // push v
stack的顶部位置对应slice的最后一个元素：

top := stack[len(stack)-1] // top of stack
通过收缩stack可以弹出栈顶的元素

stack = stack[:len(stack)-1] // pop