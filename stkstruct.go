package main
import "fmt"

type stk struct {
	data int
	nxt *stk
}

var cur,tmp *stk
var top *stk = nil

func newnode() {
	cur = &stk {}
	cur.nxt = nil
}

func push(x int) {
	newnode()
	cur.data = x
	if top == nil {
		top = cur
	} else {
		cur.nxt = top
		top = cur
	}
}

func pop() {
	if top == nil {
		fmt.Println("stack is empty")
	} else {
		top = top.nxt
	}
}

func display() {
	if top == nil {
		fmt.Println("stack is empty")
	}
	for tmp = top; tmp != nil; tmp = tmp.nxt {
		fmt.Println(tmp.data)
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