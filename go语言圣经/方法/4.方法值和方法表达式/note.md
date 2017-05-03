p.Distance()
-》
a:=p.Distance  
a()



eg:
p := Point{1, 2}
q := Point{4, 6}

distanceFromP := p.Distance        // method value
fmt.Println(distanceFromP(q))      // "5"
var origin Point                   // {0, 0}
fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)

scaleP := p.ScaleBy // method value
scaleP(2)           // p becomes (2, 4)
scaleP(3)           //      then (6, 12)
scaleP(10)          //      then (60, 120)


使用场景:
#1 
#一个包的API需要函数值
#+
#调用方希望操作的是某一个绑定了对象的方法的话

func AfterFunc
func AfterFunc(d Duration, f func()) *Timer

type Rocket struct { /* ... */ }
func (r *Rocket) Launch() { /* ... */ }
r := new(Rocket)

time.AfterFunc(10 * time.Second, func() { r.Launch() })
或者
time.AfterFunc(10 * time.Second, r.Launch)


#2.

#当你根据一个变量来决定调用同一个类型的哪个函数时，方法表达式就显得很有用了

#x=类型.方法名 
#x(类型,参数)

p:=Point{1,2}
q:=Point{4,6}

distance:=Point.Distance 
fmt.Printfln(distance(q,p))
fmt.Printf("%T\n",distance)

scale:=(*.Point).ScaleBy
scale(&p,2)
fmt.Println(p)
fmt.Printf("%T\n",scale)




type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
    var op func(p, q Point) Point
    if add {
        op = Point.Add
    } else {
        op = Point.Sub
    }
    for i := range path {
        // Call either path[i].Add(offset) or path[i].Sub(offset).
        path[i] = op(path[i], offset)
    }
}

