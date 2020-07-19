package main
import "fmt"

var a [20]int
var top int = -1

func push(x int) {
	if top == 19 {
		fmt.Println("stack is full")
	} else {
		top++
		a[top] = x
	}
}

func pop() {
	if top == -1 {
		fmt.Println("stack is empty")
	} else {
		top--
	}
}

func display() {
	if top == -1 {
		fmt.Println("Stack is Empty")
	}
	for i := top; i >= 0; i-- {
		fmt.Println(a[i],"\t")		
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