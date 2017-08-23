package main

import (
	"fmt"
	"os"
	"math/rand"
	"strconv"
)

const defaultName string = "Arthur" //Or without string...

//Declaration of some variable.
//We can use "const" in place of "var".

//log.Fatal() || panic("") - когда нельзя дальше продолжать

var (
	someFruit string = "Apple"
	someNumber uint = 0 //uint (unsigned int)
)

func readFile()  {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Error in file")
		//err := errors.New("error message")
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func declareVariables() {
	//two ways to declare variable.
	var hello string = "Hello" //we can use var x = "Hello".
	fmt.Println(hello)
	world := "world"
	fmt.Println(world)
	fmt.Println(hello + " " + world, 1 + 1, "\n")

	var floatNumber float64 = 14.432
	var intNumber int = 23;
	//type conversion
	amount := floatNumber + float64(intNumber)
	fmt.Println(amount, "\n")

	//Array's
	var array [5]int
	for i := 0; i < len(array); i++ {
		array[i] = i
	}
	fmt.Println(array, "\n")

	//Another way to see array values
	someArray := [5]int{0, 4, 6, 2, 7}
	fmt.Print("SomeArray values: ")
	for value := range someArray {
		//total += value
		fmt.Print(someArray[value], " ")
	}

	fmt.Println("\n")

	//Slices
	var slice1 = []int {1, 4, 6, 8} //this is slice
	fmt.Println(slice1[0:2]) //С какого по какой символ


	intStr := [3]int{1, 2, 3} //array
	for _, r := range intStr { //range  - цикл по масиву
		fmt.Print(r, " ")
	}

	slice2 := append(slice1, 10, 2)
	//slice3 := make([]int, 5)
	//use "copy(slice1, slice2)" to copy slice1 in slice1.
	fmt.Println(slice2, "\n")

}

func input() {
	fmt.Println("What's your name? ") //Use "fmt.Print", if you want to print in one line.
	var inputname string
	fmt.Scanf("%v", &inputname) //"%v" - the value in a default format's.
	fmt.Println("Hello, ", inputname, "!")
}

func forIfFunc()  {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, " - even")
		} else { //If u need another condition -  "else if --condition-- {".
			fmt.Println(i, " - odd")
		}
	}
}

func switchFunc() {
	var i uint = 3
	//You can use switch without expression. In case you should add condition.
	switch i {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	default: fmt.Println("Unknown Number")
	}
}

func containerMap()  {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	//declare map, like arrays.
	//elements := map[string]string{
	//	"H": "Hydrogen",
	//...
	//}

	//Check. did we have this item in container. Return "true" if yes and vice versa.
	name, ok := elements["Un"]
	fmt.Println(name, ok)

	fmt.Println(elements["Li"])
}

func containerWithMoreInfo()  {
	elements := map[string]map[string]string{
		"H": {"name":  "Hydrogen", "state": "gas"},
		"He": {"name":  "Helium", "state": "gas"},
		"Li": {"name":  "Lithium", "state": "solid"},
		"Be": {"name":  "Beryllium", "state": "solid"},
		"B": {"name":  "Boron", "state": "solid"},
		"C": {"name":  "Carbon", "state": "solid"},
		"N": {"name":  "Nitrogen", "state": "gas"},
		"O": {"name":  "Oxygen", "state": "gas"},
		"F": {"name":  "Fluorine", "state": "gas"},
		"Ne": {"name":  "Neon", "state": "gas"},
	}

	if element, ok := elements["Ne"]; ok {
		fmt.Println(element["name"], element["state"])
	} else {
		fmt.Println("About this element, nothing is known")
	}
}

func pointers() {
	i := 1
	iPtr := &i
	fmt.Println(i, *iPtr)
	*iPtr = 2
	fmt.Println(i, *iPtr)
}

func genRandomArray(n int) []int {
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}

	return nums
}

func transformStringtoInt()  {
	sNum := "1"
	num, err := strconv.Atoi(sNum) //second argument - it's err
	if err != nil {panic("Error")}
	fmt.Println("\nNum in string", num)
}

func main() {
	//argv := os.Args[1:] //var argv []string = os.Args[1:]


	//var cc chan int // chane;

	//goroutine
	//const nrutines
	//ch := make(chan int, 10) //chanel too
	//Посмотри про интерфейсы
}
