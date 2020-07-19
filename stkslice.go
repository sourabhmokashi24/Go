package main
import "fmt"

var a []int
var top = -1

func push(x int) {
	a = append(a, x)
	top++
	fmt.Println(a[top:top+1], "push")
}

func pop() {
	if top == -1 {
		fmt.Println("Stack is Empty")
		return
	}
	fmt.Println(a[top:top+1], "pop")
	a = a[:top]
	top--
}

func display() {
	if top == -1 {
		fmt.Println("Stack is Empty")
	}
	for i := top; i >= 0; i-- {
		fmt.Println(a[i])
	}
}

func main() {
	var choice, x int
	for choice != 4 {
		fmt.Print("1. Push\n2. Pop\n3. Display\n4. Exit: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("Enter number: ")
			fmt.Scan(&x)
			push(x)
		case 2:
			pop()
		case 3:
			display()
		case 4:
		default:
			fmt.Println("Enter Correct Choice")
		}
	}
}