type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
    Position  string
    Salary    int
    ManagerID int
}


var dilbert Employee

1.通过点
dilbert.Salary -= 5000

2.通过取地址
position := &dilbert.Position
*position = "Senior " + *position

3.
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"




#后面的语句通过EmployeeByID返回的结构体指针更新了Employee结构体的成员
func EmployeeByID(id int) *Employee { /* ... */ }

fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

id := dilbert.ID
EmployeeByID(id).Salary = 0 // fired for... no real reason