package main

func aggregate(num1 int, num2 int, op func(int, int) int) int {
	return op(num1, num2)
}
func add_(a, b int) int {
	return a + b
}
func multi_(a, b int) int {
	return a * b
}
func div_(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
func sub_(a, b int) int {
	return a + b
}

func main() {
	println(
		aggregate(12, 12, add_),
		aggregate(12, 12, multi_),
		aggregate(12, 12, div_),
		aggregate(12, 12, sub_),
	)
}

// function as data is called function as 1st class citizen
// higher order function is function that takes function as a argument and use that function inside the body
