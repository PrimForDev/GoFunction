package basic

// Boolean
var TypeBool bool = true

// String
var TypeString = "Hello World"

// Integer
//int (Range: 2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems)
//int8 (Range: -128 to 127)
//int16 (Range: -32768 to 32767)
//int32 (Range: -2147483648 to 2147483647)
//int64 (Range: -9223372036854775808 to 9223372036854775807)
var TypeInt int = 10

// Floating point number
//float32 (Range: -3.4e+38 to 3.4e+38)
//float64 (Range: -1.7e+308 to +1.7e+308)
var TypeFloat = 19.99

const PI = 3.14

var TypeArray = [3]int{1, 2, 3}
var TypeSlice = []int{4, 5, 6}

// ประกาศ struct นอกฟังก์ชัน เพื่อให้สามารถใช้ได้ทั้งโปรแกรม
type Preson struct {
	name   string
	age    int
	job    string
	salary int
}

// แก้ไขฟังก์ชันให้คืนค่าเป็น struct ที่เราสร้างไว้
// กำหนดชนิดคืนค่าเป็น Person ซึ่งเป็น struct ที่เราได้ประกาศไว้
func GetStruct() Preson {
	var TypeStruct Preson
	TypeStruct.name = "Hege"
	TypeStruct.age = 45
	TypeStruct.job = "Teacher"
	TypeStruct.salary = 6000
	return TypeStruct

}

// แก้ไขฟังก์ชันให้คืนค่าเป็น map
func GetMap() map[string]string {
	var TypeMap = make(map[string]string)
	TypeMap["brand"] = "Ford"
	TypeMap["model"] = "Mustang"
	TypeMap["year"] = "1964"
	return TypeMap
}
