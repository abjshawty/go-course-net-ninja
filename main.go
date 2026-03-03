// File explaining the basics of Go.
// Based on the Go Tutorial by NetNinja (https://youtube.com/playlist?list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM&si=KelHCoiAI_jtTlRC)

/**
* Packages are essentially a group of related files
* Using the "main" informs the Go compiler that this package will be built into a standalone app
* Similar to Java
 */
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var strongest string = "Satoru Gojo"

// String formatter + Console interface
/**
* The main function is the one that will run when we launch this program
* There can only be one per project
 */
func main() {
	// basics_lesson()
	// strings_lesson()
	// slices_lesson()
	// loops_lesson()
	// booleans_lesson()
	// functions_lesson()
	// fmt.Println(getInitials("Timmy Kouadio-Bhegnin"))
	// package_lesson()
	// maps_lesson()
	// pass_by_value_lesson()
	// pointers_lesson()
	// structs_lesson()
	// receiver_functions_lesson()
	// receiver_functions_with_pointers_lesson()
	// user_input_lesson()
	// switch_lesson()
	interface_lesson()
}

func basics_lesson() {
	/*
	* As you already should know if you're here, variables are boxes containg values, essentially
	* Variables can be declared in a few ways:
	* - var name type = value
	* - var name = value
	* - var name
	* - name := value
	 */

	// Strings
	var first_name string = "Timmy"
	var last_name string = "Kouadio-Bhegnin"

	// Numbers (Integers). See https://go.dev/ref/spec#Numeric_types
	var birth_year int = 2002
	var number_of_holes int8 = 7
	var favorite_number = 17

	var absolute_value uint = 8

	// Numbers (Floats). See https://go.dev/ref/spec#Numeric_types
	var net_worth float32 = 35.0
	var pi_from_memory float64 = 3.14

	// Arrays (Unchangeable length)
	var ages [3]int = [3]int{20, 25, 30}
	var songs = [5]string{"yeet", "who", "coffee", "nod", "outside"}

	// Slices (Use Arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	// Slice ranges
	rangeOne := songs[1:3]
	rangeTwo := songs[2:]
	rangeThree := songs[:3]
	/**
	* Printing and formatting Strings
	*
	* We use the fmt package for both of these tasks.
	* "Print", does just that
	* "Println" adds a new line at the end of the string
	* "Printf formats the string Using
	* - %d for Integers
	* - %v for all values
	* - %s for strings
	* - %q for strings in quotes
	* - %f (%.2f) for floats (can control rounding)
	* - %T for types
	* "Sprintf" formats a string using the same arguments, then returns the string without printing. (See https://pkg.go.dev/fmt)
	 */

	fmt.Print("Same ")
	fmt.Print("line\n")
	fmt.Println("Different")
	fmt.Println("lines...")

	fmt.Print("Morning. My name is ", first_name, last_name, ".\n")
	fmt.Println("I was born in", birth_year)

	var test_string = fmt.Sprintf("My body has %d holes, though my %vth favorite number is %d", number_of_holes, absolute_value, favorite_number)
	fmt.Printf("%s has $%f, and around %.3f friends\n", first_name, net_worth, pi_from_memory)
	fmt.Println(test_string)
	fmt.Println(ages[0])
	fmt.Println(songs[2])
	fmt.Println(scores, len(scores))
	fmt.Println(rangeOne, rangeTwo, rangeThree)
}

func strings_lesson() {
	greeting := "Hi! I am not Claude. But I can still code..."
	fmt.Println(strings.Contains(greeting, "hi"))
	fmt.Println(strings.Count(greeting, "i"))
	fmt.Println(strings.Replace(greeting, "code", "program", 1))
	fmt.Println(strings.ReplaceAll(greeting, "code", "program"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.ToLower(greeting))
	fmt.Println(strings.Split(greeting, " "))
	fmt.Println(strings.Repeat(greeting, 3))
}

func slices_lesson() {
	ages := []int{21, 54, 30}
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages, 21)
	fmt.Println(index)
	var names []string = []string{"Timmy", "Claude", "John", "Jane", "Romaric", "Gemini", "Sam", "Elon"}
	fmt.Println(names)
}

func loops_lesson() {
	/**
	* Loops are pretty much the same as in Java
	 */
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	x := 0
	for x < 5 {
		fmt.Println("x is ", x)
		x++
	}

	var names []string = []string{"Timmy", "Claude", "John", "Jane", "Romaric", "Gemini", "Sam", "Elon"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("Value at index %d is %s\n", index, value)
	}

	for _, value := range names {
		fmt.Printf("Value is %s\n", value)
	}

	fmt.Println(names)
	for _, value := range names {
		value = "Replacement Test 1"
		fmt.Println(value)
	}
	fmt.Println(names)
	for index, value := range names {
		fmt.Printf("Value at index %d is %s\n", index, value)
		names[index] = "Replacement Test 2"
	}
	fmt.Println(names)
}

func booleans_lesson() {
	/**
	* Bool
	* Booleans are true or false values
	 */
	var true_value bool = true
	var false_value bool = false
	fmt.Printf("%T %v\n", true_value, true_value)
	fmt.Printf("%T %v\n", false_value, false_value)

	number := 80

	fmt.Println("Number is less than 100: ", number < 100)
	fmt.Println("Number is greater than 100: ", number > 100)
	fmt.Println("Number is less than or equal to 100: ", number <= 100)
	fmt.Println("Number is greater than or equal to 100: ", number >= 100)
	fmt.Println("Number is equal to 100: ", number == 100)
	fmt.Println("Number is not equal to 100: ", number != 100)

	if number > 100 {
		fmt.Println("Condition check: Number is greater than 100")
	} else if number == 100 {
		fmt.Println("Condition check: Number is equal to 100")
	} else {
		fmt.Println("Condition check: Number is less than 100")
	}

	var names []string = []string{"Timmy", "Claude", "John", "Jane", "Romaric", "Gemini", "Sam", "Elon"}

	for index, value := range names {
		if index == 2 {
			fmt.Println("Skipping index 2")
			continue
		}
		if index == 5 {
			fmt.Println("Breaking at index 5")
			break
		}
		fmt.Printf("Value at index %d is %s\n", index, value)
	}
}

func functions_lesson() {
	/**
	* Functions are a repeatable block of code, that takes arguments and returns values
	* The syntax is func name(args) return_type {}
	 */
	greet("abjshawty")
	cycle_names([]string{"Timmy", "Claude", "John", "Jane", "Romaric", "Gemini", "Sam", "Elon"}, greet)
	fmt.Println(circle_area(10))
}

func maps_lesson() {
	/**
	* Maps allow us to store key - value pairs
	* The keys might be of any primitive tupe, and the values can be of "any" type, but all keys must have the same type, as well as all values
	 */
	menu := map[string]float64{
		"whole chicken": 3999.99,
		"half chicken":  2499.99,
		"rice":          499.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["rice"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// Integers as key type
	var phonebook map[int]string = map[int]string{
		2250711115686: "Timmy",
		2250101993388: "Christiane",
		2250505201515: "Yesso",
	}

	fmt.Println(phonebook)
	phonebook[2250101993388] = "Mom"
	fmt.Println(phonebook)
}

func pass_by_value_lesson() {
	/**
	* Go (GoLang) is a "Pass by value" language
	* It means when we pass variables around, what is transferred is not the memory adress of the data, but its value
	* Go makes copies of values to pass to functions
	*
	* Go variable types can be separated into two distinct groups;
	* Group A (Value / Non Pointer Values): Strings, Ints, Floats, Booleans, Arrays, Structs
	* Group B (Reference / Pointer Wrapper Values): Slices, Maps, Functions
	 */

	var name string = "Dabura"
	update_name_fail(name)
	fmt.Println(name)

	// This update failed because what was actually passed to it is as *copy* of name

	name = update_name_by_returning(name)
	fmt.Println(name)

	// Here, it worked because we *assigned* the returned *value* to our variable
	// This essentially drops the current value and adopts the new one

	alter := map[string]string{
		"Deku":    "One For All",
		"AFO":     "All For One",
		"Star":    "Command",
		"Tomura":  "Decay",
		"Chisaki": "Overhaul",
	}
	fmt.Println(alter)
	update_map(alter)
	fmt.Println(alter)

	/**
	* This works because of how Go manages memory...
	* Types in group B also pass a copy of their content.
	* However, the content of those is not the expected value, but a memory adress referencing the actual content we want.
	* As such, when we make a modification to those, the modification is applied to what lies in that copied adress, which is just the same object.
	 */
}

func pointers_lesson() {
	/**
	 * As established previously, we saw that some data types serve as memory adress containers for values
	 * We call those, "Pointers".
	 * However, it is worth noting that we can also manually create pointers for non pointer values.
	 */

	var score int = 97
	scorePointer := &score
	fmt.Println("Memory adress of score is: ", scorePointer)
	fmt.Println("Value at adress: ", *scorePointer)

	update_score(&score)

	fmt.Println("Memory adress of score is: ", scorePointer)
	fmt.Println("Value at adress: ", *scorePointer)
}

func structs_lesson() {
	/**
	* A struct can be defined as a blueprint to define a data model or data type
	 */
	mybill := newBill("Muchacho")
	fmt.Println(mybill)
}

func receiver_functions_lesson() {
	/**
	* Probably the closest thing to a method.
	* Takes in a copy of the object
	 */
	myBill := newBill("AI")
	fmt.Println(myBill.format())
}

func receiver_functions_with_pointers_lesson() {
	var b bill = newBill("Doofenschmirz")
	fmt.Println(b.format())
	b.updateTip(50)
	fmt.Println(b.format())
	b.addItem("Morgan", 100)
	fmt.Println(b.format())
}

func user_input_lesson() {
	var b bill = createBill()
	b.selfPrint()
}

func switch_lesson() {
	var b bill = newBill("Mustang")
	promptOptions(b)
	b.selfPrint()
}

func interface_lesson() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, shape := range shapes {
		printShapeInfo(shape)
	}
}

// Secondary Functions

func greet(name string) {
	fmt.Printf("Hello %s, doing well?\n", name)
}

func cycle_names(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

func circle_area(radius float64) float64 {
	return math.Pi * radius * radius
}

func getInitials(name string) (string, string) {
	words := strings.Split(name, " ")
	initials := ""
	for _, word := range words {
		initials += string(word[0])
	}
	return string(initials[0]), string(initials[1])
}

func update_name_fail(name string) {
	name = "Satoru Gojo"
}

func update_name_by_returning(name string) string {
	name = "Ryomen Sukuna"
	return name
}

func update_map(mapObj map[string]string) {
	mapObj["Satoru Gojo"] = "Limitless"
}

func update_score(p *int) {
	*p = 60
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"Sam": 1.9, "El": 700.4},
		tip:   0,
	}
	return b
}

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func (b *bill) format() string {
	var fs string = "Bill breakdown: \n"
	var total float64 = 0

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Total: ", total)
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Tip: ", b.tip)
	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b bill) selfPrint() {
	fmt.Println(b.format())
}

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	b.selfPrint()
	return b
}

func (b bill) saveBill() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a -- Add Item, s -- Save bill, t -- Add tip)", reader)

	switch opt {
	case "a":
		fmt.Println("You chose A.\nAdd Item: ")
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Invalid price value")
			return
		}
		b.addItem(name, p)
		b.selfPrint()
	case "s":
		fmt.Println("You chose S.\nSave Bill: ")
		b.saveBill()
	case "t":
		fmt.Println("You chose T.\nAdd Tip: ")
		tip, _ := getInput("Tip amount: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Invalid tip value")
			return
		}
		b.updateTip(t)
		b.selfPrint()
	default:
		fmt.Println("Invalid option")
	}
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is; %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n", s, s.circumf())
}
