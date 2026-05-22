package main

import "fmt"

func main() {
	var hello string
	hello = "Hello world"
	var hellov2 string = "hello world!"
	fmt.Println(hello + " " + hellov2)

	hello = "change change"
	fmt.Println(hello)

	var (
		name string = "test"
		age  int    = 22
	)

	fmt.Print("Name: ", name)
	fmt.Print("Age: ", age)

	namev2 := "SS"
	fmt.Println(namev2)

	const (
		a = 1
		b
		c
		d
		e = 3
	)

	fmt.Println(a, b, c, d, e)

	const (
		C0 = iota + 1
		C1
		C2 = iota
	)

	fmt.Println(C0, C1, C2)

	const ( // iota drops to 0
		C3 = iota
	)
	fmt.Println(C3)

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(numbers[2])
	numbers[0] = 87
	fmt.Println(numbers[0])
	fmt.Println("Array len: ", len(numbers))

	var numbersv2 [5]int = [5]int{1, 2}
	fmt.Println(numbersv2)

	var numbersv3 = [...]int{1, 2, 3, 4, 5}
	var numbersv4 = [...]int{1, 2, 3}
	fmt.Println(numbersv3)
	fmt.Println(numbersv4)

	colors := [3]string{2: "red", 0: "blue", 1: "green"} // we can put each custom index
	fmt.Println(colors)

	numbersRowsxCol := [3][2]int{
		{1, 2},
		{4, 5},
		{3, 6},
	}
	fmt.Println(numbersRowsxCol)

	var numbersRowsxColV2 [3][2]int
	numbersRowsxColV2[0] = [2]int{1, 2}
	numbersRowsxColV2[1] = [2]int{3, 4}
	numbersRowsxColV2[2] = [2]int{5, 6}

	fmt.Println(numbersRowsxColV2)

	arr := [3]int{1, 2, 3}
	newArr := arr

	newArr[0] = 11

	fmt.Println("Old arr: ", arr)
	fmt.Println("New arr: ", newArr)

	x := 2
	switch x {
	case 1:
		fmt.Println("x=1")
	case 2:
		fmt.Println("x=2")
		fallthrough //make go to the next case despite condition
	case 3:
		fmt.Println("x=3")
	case 4:
		fmt.Println("x=4")
	}

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	var i2 = 1
	for i2 < 10 {
		fmt.Println(i2 * i2)
		i2++
	}

	str := "Hello"
	for index, value := range str {
		fmt.Printf("Index: %d, Value %c\n", index, value)
	}

	for _, value := range str {
		fmt.Printf("%c\n", value)
	}

	var users = [3]string{"Tom", "John", "Amorim"}
	for index, value := range users {
		fmt.Println(index, value)
	}

	for i := 0; i < len(users); i++ {
		fmt.Print(users[i] + " ")
	}
	fmt.Print("\n")

	// mark for loop
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i = %d, j=%d\n", i, j)

			if i == 2 && j == 2 {
				fmt.Println("Exiting from outer loop...")

				break OuterLoop
			}
		}
	}

	fmt.Println("Loop exited")

	add(4, 5)
	add(7, 8)

	addV2(1, 2, 3.4, 5.6, 1.2)

	var variable = 8
	fmt.Println("var before: ", variable)
	increment(variable)
	fmt.Println("Var after: ", variable)

	adds(1, 2, 3)
	adds(1, 2, 3, 4)
	adds(4, 5, 6, 7, 8)

	var addsNums = []int{1, 2, 3, 4, 5}
	adds(addsNums...)

	///////////

	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "x")
	s = append(s, "s")
	fmt.Println("After append: ", s)

	copyS := make([]string, len(s))
	copy(copyS, s)
	fmt.Println("Coppied: ", copyS)

	l := s[2:5]
	fmt.Println("slice1: ", l)

	l = s[:5]
	fmt.Println("slc2: ", l)

	l = s[2:]
	fmt.Println("slc3: ", l)

	t := []string{"a", "b", "c"}
	fmt.Println("raw: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)

	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, 2)
		for j := 0; j < 2; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("3x2 matrix: ", twoD)

	m := make(map[string]int)
	m["k2"] = 7
	m["k3"] = 17

	fmt.Println("map: ", m)

	v1 := m["k2"]
	fmt.Println(v1)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	_, prs := m["k2"]
	fmt.Println("exists: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s->%s\n", k, v)
	}

	for k := range kvs {
		fmt.Printf("key: %s\n", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func add(x int, y int) {
	fmt.Println(x + y)
}

func addV2(x, y int, a, b, c float32) {
	var z = x + y
	var d = a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}

func increment(x int) {
	fmt.Println("X before: ", x)
	x += 20
	fmt.Println("X after: ", x)
}

func adds(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}
