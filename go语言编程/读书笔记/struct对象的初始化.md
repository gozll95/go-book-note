type Rect struct { 
	x, y float64
	width, height float64 
}



rect1 := new(Rect)
rect2 := &Rect{}
rect3 := &Rect{0, 0, 100, 200}
rect4 := &Rect{width: 100, height: 200}


func NewRect(x, y, width, height float64) *Rect{ 
	return &Rect{x, y, width, height}
}

