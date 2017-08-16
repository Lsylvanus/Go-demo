package main

import "fmt"

type person struct {
	id      int
	name    string
	country string
}

type interface_person interface {
	introduction()
}

func (p *person) introduction() {
	fmt.Println("My name is : ", p.name)
}

func main() {

	var zhu person
	zhu.id = 1001
	zhu.name = "Lsylvanus"
	zhu.country = "China"

	fmt.Println("Lsylvanus=", zhu)

	zhu.introduction()

	var o valueable
	o = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	var t f
	o = t
	showValue(o)
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice *s.count
}

type f float32

func (s f) getValue() float32 {
	return 0.1
}

type valueable interface {
	getValue() float32
}

func showValue(asset valueable) {
	fmt.Printf("value of the asset is %f\n", asset.getValue())
}
