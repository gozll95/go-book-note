\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
如果考虑效率的话，较大的结构体通常会用指针的方式传入和返回，

func Bonus(e *Employee, percent int) int {
    return e.Salary * percent / 100
}

返回结构体的地址:
pp := &Point{1, 2}



\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
可比较的结构体类型和其他可比较的类型一样，可以用于map的key类型。

type address struct {
    hostname string
    port     int
}

hits := make(map[address]int)
hits[address{"golang.org", 443}]++



\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
type Point struct{
	X,Y int
}

type Circle struct{
	Center Point 
	Radius int 
}

type Wheel struct{
	Circle Circle
	Spokes int
}


var w Wheel
w.Circle.Center.X = 8
w.Circle.Center.Y = 8
w.Circle.Radius = 5
w.Spokes = 20




\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
#匿名成员
#匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。
type Point struct{
	X,Y int
}


type Circle struct {
    Point
    Radius int
}

type Wheel struct {
    Circle
    Spokes int
}

var w Wheel
w.X = 8            // equivalent to w.Circle.Point.X = 8
w.Y = 8            // equivalent to w.Circle.Point.Y = 8
w.Radius = 5       // equivalent to w.Circle.Radius = 5
w.Spokes = 20