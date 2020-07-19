package main
import "fmt"

var a [10]int
var front, end int = -1, -1

func insert(x int) {
	if front == -1 && end == -1 {
		front -= -1
		end -= -1
		a[end] = x
	} else if end == 9 {
		fmt.Println("queue is full")
	} else {
		end++
		a[end] = x
	}
}

func delete() {
	if front == -1 && end == -1 {
		fmt.Println("queue is empty")
	} else if front == end {
		front, end = -1, -1
	} else {
		front++
	}
}

func display() {
	if front == -1 && end == -1 {
		fmt.Println("queue is empty")
		return
	}
	for i := front; i <= end; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	var choice, x int
	for choice != 4 {
		fmt.Print("1. Insert\n2. Delete\n3. Display\n4. Exit: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("Enter number: ")
			fmt.Scan(&x)
			insert(x)
		case 2:
			delete()
		case 3:
			display()
		case 4:
		default:
			fmt.Println("Enter Correct Choice")
		}
	}
}