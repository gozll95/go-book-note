方法可以被声明到任意类型，只要不是一个指针或者一个interface

只有类型(Point)和指向他们的指针(*Point),才是可能会出现在接收器声明里的两种接收器。此外，为了避免歧义，在声明方法时，
如果一个类型名本身是一个指针的话，是不允许其出现在接收器中的。

type P *int 
func (P)f(){/*...*/}  // compile error: invalid receiver type

这样可以
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}


r:=&Point{1,2}
r.ScaleBy(2)
fmt.Println(*r)


或者这样：

p := Point{1, 2}
pptr := &p
pptr.ScaleBy(2)
fmt.Println(p) // "{2, 4}"
或者这样:

p := Point{1, 2}
(&p).ScaleBy(2)
fmt.Println(p) // "{2, 4}"


不过后面两种方法有些笨拙。幸运的是，go语言本身在这种地方会帮到我们。如果接收器p是一个Point类型的变量，并且其方法需要一个Point指针作为接收器，我们可以用下面这种简短的写法：

p.ScaleBy(2)
编译器会隐式地帮我们用&p去调用ScaleBy这个方法。这种简写方法只适用于“变量”，包括struct里的字段比如p.X，以及array和slice内的元素比如perim[0]。我们不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取得到：

Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal
但是我们可以用一个*Point这样的接收器来调用Point的方法，因为我们可以通过地址来找到这个变量，只要用解引用符号*来取到该变量即可。编译器在这里也会给我们隐式地插入*这个操作符，所以下面这两种写法等价的：

pptr.Distance(q)
(*pptr).Distance(q)


------------------------------------------------------------------------------------------>
不论是接收器的实际参数和其接收器的形式参数相同，比如两者都是类型T或者都是类型*T：

Point{1, 2}.Distance(q) //  Point
pptr.ScaleBy(2)         // *Point
或者接收器形参是类型T，但接收器实参是类型*T，这种情况下编译器会隐式地为我们取变量的地址：

p.ScaleBy(2) // implicit (&p)
或者接收器形参是类型*T，实参是类型T。编译器会隐式地为我们解引用，取到指针指向的实际变量：

pptr.Distance(q) // implicit (*pptr)