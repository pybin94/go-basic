package main

import (
	"fmt"
	"strings"
)

func multiple(a, b int) int {
	// multiple(a int, b int) a와 b의 타입이 같을 때 a의 타입을 생략할 수 있다.
	return a + b
}

// multiple value function
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// variadic parameter - 가변 매개변수
func animals(words ...string) {
	fmt.Println(words)
}

// naked return
func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// defer - 함수가 끝난 후 실행된다.
func deferReturn(words string) {
	defer fmt.Println("I'm done")
	fmt.Println(words)
}

// loop
func superAdd(numbers ...int) int {

	for index, number := range numbers {
		fmt.Println("range", index, number)
	}

	total := 0
	for _, number := range numbers { // index를 사용하고 싶지 않을 때 _(ignore)한다
		total += number
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println("for", numbers[i])
	}
	return total
}

// map
func maps() {
	// map[key type]value type
	opps := map[string]string{"name": "opps", "age": "12"}

	for key, value := range opps {
		fmt.Println("maps", key, value)
	}
}

// if else
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

// switch
func learnSwitch1(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func learnSwitch2(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return true
	}
	return false
}

// pointer &: 주소(address), *: 주소에 저장된 값
func memoryPointer() {
	a := 2
	b := 5
	c := &a // c는 a의 주소를 저장한다
	*c = 20 // *c는 a의 값을 보여준다

	// c는 a(&a)의 주소와 연결이 되어있고 *c(&a 주소의 값)를 20으로 바꿈으로 a = 20이 된다.
	fmt.Println("memoryPointer", a, &b, c)
}

// array and slice
func arrayAndSlice() {
	// array
	names := [5]string{"opps", "viin", "aah"}
	names[3] = "baa"
	names[4] = "cba"

	// slice
	animals := []string{"dog", "cat", "monkey"}

	// ERROR!!!
	// animals[3] = "rabbit"
	// append(animals, "rabbit")

	// SUCCESS!!!
	animals = append(animals, "rabbit")

	fmt.Println("arrayAndSlice", names, animals)
}

// struct
type person struct {
	name    string
	age     int
	favFood []string
}

func learnStruct() {
	favFood := []string{"kimchi", "ramen"}
	// opps := person{"opps", 18, favFood}
	opps := person{name: "opps", age: 18, favFood: favFood}
	fmt.Println(opps.name, opps)
}

func main() {
	fmt.Println(multiple(2, 3))

	totalLength, upperName := lenAndUpper("opps")
	// totalLength, _ := lenAndUpper("opps")
	// 함수의 return을 제한하고 싶으면 '_' 로 표기한다.
	fmt.Println(totalLength, upperName)

	// 배열로 반환된다.
	animals("tiger", "rabbit", "cat", "dog")

	fmt.Println(lenAndUpper2("opps"))

	deferReturn("I'm the First")

	fmt.Println("total", superAdd(1, 2, 3, 4, 5, 6))

	maps()

	fmt.Println("canIDrink", canIDrink(16))

	fmt.Println("learnSwitch1", learnSwitch1(16))
	fmt.Println("learnSwitch2", learnSwitch2(16))

	memoryPointer()

	arrayAndSlice()

	learnStruct()
}
