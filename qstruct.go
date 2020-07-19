package main
import "fmt"

type queue struct {
	data int
	nxt *queue
}

var cur, tmp *queue
var front, end *queue = nil, nil

func newnode() {
	cur = &queue {}
	cur.nxt = nil
}

func insert(x int) {
	newnode()
	cur.data = x
	if front == nil && end == nil {
		front, end = cur, cur
	} else {
		end.nxt = cur
		end = cur
	}
	fmt.Println(x, "inserted")
}

func delete() {
	if front == nil && end == nil {
		fmt.Println("queue is empty")
	} else if front == end {
		fmt.Println(front.data, "deleted")
		front, end = nil, nil
	} else {
		fmt.Println(front.data, "deleted")
		front = front.nxt
	}
}

func display() {
	if front == nil && end == nil {
		fmt.Println("queue is empty")
		return
	}
	for tmp = front; tmp != nil; tmp = tmp.nxt {
		fmt.Println(tmp.data)
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