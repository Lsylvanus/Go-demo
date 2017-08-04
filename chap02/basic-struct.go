package main

import "fmt"

func main() {

	type Person struct {
		id      int
		name    string
		country string
	}

	type Employee struct {
		Person          //匿名字段
		salary int
		int             //用内置类型作为匿名字段
		country string  //类似于重载
	}

	/*var em1 Employee = Employee{}
	em1.Person = Person{1000, "Lsylvanus", "foshan"}
	em1.salary = 5000
	em1.int = 100 //使用时注意其意义，此处无
	em1.addr = "luoding"*/
	//初始化方式不一样，但是结果一样
	em1 := Employee{Person:Person{1000, "Lsylvanus", "foshan"}, salary:5000, int:100, country:"luoding"}
	fmt.Println(em1)

	fmt.Println("live addr(em1.country) = ", em1.country)
	fmt.Println("work addr(em1.Person.country) = ", em1.Person.country)
	em1.int = 200  //修改匿名字段的值

	//def and init 1
	var zhu Person
	zhu.id = 1001
	zhu.name = "Lsylvanus"
	zhu.country = "China"

	fmt.Println("Lsylvanus=", zhu)

	//def and init 2
	michael := Person{1002, "michael", "PRC"}
	fmt.Println("michael=", michael)

}
