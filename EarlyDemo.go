package main

import (
	"fmt"
	"reflect"
)

type BankAccount struct {
	accountNumber int
	accountHolder string
	balance       float64
}

type Student struct {
	studentNumber int
	name          string
	gpa           float64
}

func main() {
	////	demo2 := []string{"Comp510", "comp490", "Comp151", "Comp426"}
	////demo3 := []int{1, 2, 435, 456, 567, 67}
	////demo4 := make(map[string]int)
	//demo5 := BankAccount{
	//	accountNumber: 23425123541345,
	//	accountHolder: "Comp510",
	//	balance:       1234.56,
	//}
	//Stu1 := Student{
	//	studentNumber: 5675367343,
	//	name:          "Stu Dent",
	//	gpa:           3.000001,
	//}
	//
	//type Prof struct {
	//	classlist []string
	//	name      string
	//	id        int
	//	rank      string
	//}
	//
	//type myData struct {
	//	data int
	//}
	//
	//demo6 := Prof{
	//	classlist: []string{"Comp510", "comp490", "Comp151", "Comp426"},
	//	name:      "J Santore",
	//	id:        234234,
	//	rank:      "Professor",
	//}
	//basicReflectionDemo(demo5)
	//basicReflectionDemo(Stu1)
	//mystery := reflect.ValueOf(demo6)
	//fmt.Println("+++++++++++++++++++++++++++++++")
	//fmt.Println("Reflect.Value:", mystery)
	//for fieldNum := 0; fieldNum < mystery.NumField(); fieldNum++ {
	//	fmt.Println(mystery.Field(fieldNum))
	//}

	demo7 := []int{12323, 5345, 47567, 64567}
	demo8 := []float64{1.2234, 4.646, 55656.445}
	buildMap(demo7)
	buildMap(demo8)
}

func basicReflectionDemo(mystery interface{}) {
	mysteryType := reflect.TypeOf(mystery)
	kind := mysteryType.Kind()
	fmt.Println("The type of the parameter is: ", mysteryType)
	fmt.Println("The kind of that parameter is: ", kind)
}

func buildMap(data interface{}) {
	mystery := reflect.TypeOf(data)
	dataStructure := reflect.MakeMap(reflect.MapOf(
		reflect.TypeOf("d"), mystery))
	actualData := reflect.ValueOf(data)
	for loc := 0; loc < actualData.Len(); loc++ {
		fmt.Println(actualData.Index(loc))
	}
	fmt.Println(dataStructure)
}
