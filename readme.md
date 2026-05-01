## parameter and arguments

- parameter-> variable in function definition
- argument-> value passed to function when calling

```code
func sum(a, b int) { // parameter or formal parameter
	z := a + b
	fmt.Println(z)
}
sum(10, 20) // argument or actual parameter
```

## function

- standard function
- init function
- anonymous function
- immidiately invoked function expression(IIFE)
- first order function
- higer order function/first class function
- callback function

##### init function-> you can not call this , computer calls this function automatically, it call by system before execute main function

#### global->init->main->other function

```code
func init() {
	fmt.Println("init function called")

	// fmt.Println(a)
}
```

#### anonymous function-> function without name, you can not call this function, you can assign this function to a variable and call that variable

```code
sum := func(a int, b int) {
		fmt.Println("this is anonymous function", a, b)
	}
```

#### first order function-> function that can be assigned to a variable, passed as an argument to another function, or returned from another function example-> standard function, anonymous function, immidiately invoked function expression(IIFE) all are first order function

#### simply A function can be: assigned to a variable, passed as an argument to another function, or returned from another function taht is first order function

```code

```

#### higer order function/first class function-> a function that can take another function as an argument or return a function as a result

#### it would be higer order function if any one of the following three

- take parameter as a function
- return a function
- both take parameter as a function and return a function

```code
func processOperation(a int, b int, op func(p, q int)) {
	op(a, b)

}
func sum(a, b int) {
	z := a + b
	fmt.Println(z)
}

func main() {

	processOperation(2 , 5 , sum)

}

returning a function example:

func call() func(a, b int) {
	return sum

}
func sum(a, b int) {
	z := a + b
	fmt.Println(z)
}

func main() {


	add := call()
	add(10, 20)
}

```

#### callback function-> a function that is passed as an argument to another function and is executed after some operation is completed

```code
func processOperation(a int, b int, op func(p, q int)) { // op func(p, q int) is callback function
	op(a, b)
}

func main() {
	processOperation(2, 5, func(p, q int) {
		fmt.Println(p + q)
	})
}
```

#### why it is first class function-> because it can be assigned to a variable, passed as an argument to another function, or returned from another function
