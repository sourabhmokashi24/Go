package main
import "fmt"

type node struct {
	data int
	nxt *node
}

var tmp, cur *node
var head *node = nil

func newnode() {
	cur = &node {}
	cur.nxt = cur
}

func insert() {
	newnode()
	var x int
	fmt.Print("\nEnter number to Insert: ")
	fmt.Scan(&x)
	cur.data = x
	if head == nil {
		head = cur
	} else {
		tmp = head
		for tmp.nxt != head {
			tmp = tmp.nxt
		}
		tmp.nxt = cur
		cur.nxt = head
	}
}

func search(x int) int {
	if head == nil {
		return -1
	}
	tmp = head
	var flag int
	for {
		if tmp.data == x {
			flag =1
			break
		}
		tmp = tmp.nxt
		if tmp == head {
			break
		}
	}
	if flag == 1 {
		return 1
	} else {
		return 0
	}
}

func insatfirst() {
	newnode()
	var x int
	fmt.Print("\nEnter number to Insert at First: ")
	fmt.Scan(&x)
	cur.data = x
	if head == nil {
		head = cur
	} else {
		cur.nxt = head
		tmp = head
		for tmp.nxt != head {
			tmp = tmp.nxt
		}
		tmp.nxt = cur
		head = cur
	}
}

func insatmid() {
	var x int
	fmt.Print("\nEnetr Number to Insert at Mid: ")
	fmt.Scan(&x)
	newnode()
	cur.data = x
	if head == nil {
		head = cur
	} else {
		fmt.Print("\nEnter number to Insert After: ")
		var y int
		fmt.Scan(&y)
		var flag int = search(y)
		if flag == 1 {
			tmp = head
			for tmp.data != y {
				tmp = tmp.nxt
			}
			cur.nxt = tmp.nxt
			tmp.nxt = cur
		} else {
			fmt.Println("\nNumber not Found")
		}
	}
}

func delete() {
	var x int
	var prev *node
	if head == nil {
		fmt.Println("\nList is Empty")
		return
	}
	fmt.Print("\nEnter number to Delete: ")
	fmt.Scan(&x)
	if head.data == x {
		tmp = head
		for tmp.nxt != head {
			tmp = tmp.nxt
		}
		if tmp == head{
			head = nil
			return
		}
		tmp.nxt = head.nxt
		head = nil
		head = tmp.nxt
	} else {
		var flag int
		flag = search(x)
		if flag == 1 {
			prev = head
			tmp = head.nxt
			for tmp.data != x {
				tmp = tmp.nxt
				prev = prev.nxt
			}
			prev.nxt = tmp.nxt
			tmp.nxt = nil
		} else {
			fmt.Println("\nNumber not Found")
		}
	}
}

func display() {
	if head == nil {
		fmt.Println("\nList is Empty")
	} else {
		fmt.Print("\nList: ")
		tmp = head
		for {
			fmt.Print(tmp.data)
			tmp = tmp.nxt
			if tmp == head {
				break
			}
			fmt.Print("->")
		}
		fmt.Println()
	}
}

func main() {
	var choice, x, exit int
	for exit != -1 {
		fmt.Println("1. Insert\n2. InsertAtFirst\n3. InsertAtMid\n4. Display\n5. Delete\n6. Search")
		fmt.Print("7.Exit\nEnter Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			insert()
		case 2:
			insatfirst()
		case 3:
			insatmid()
		case 4:
			display()
		case 5:
			delete()
		case 6:
			fmt.Print("Enter number to search: ")
			fmt.Scan(&x)
			var flag = search(x)
			if flag == 1 {
				fmt.Println("Found")
			} else if flag == -1{
				fmt.Println("\nList is Empty")
			} else {
				fmt.Println("Not Found")
			}
		case 7:
			exit = -1
		default:
			fmt.Println("Enetr Correct Option")
		}
	}
}