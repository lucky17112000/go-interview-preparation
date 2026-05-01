## function->

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
