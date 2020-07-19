package main
import "fmt"

type node struct {
	data int
	prev, nxt *node
}

var tmp, cur *node
var head *node

func newnode() {
	cur = &node {}
	cur.nxt = nil
	cur.prev = nil
}

func insnode(head *node) {
	if head.data < cur.data {
		if head.nxt == nil {
			head.nxt = cur
		} else {
			insnode(head.nxt)
			
		}
	} else {
		if head.prev == nil {
			head.prev = cur
		} else {
			insnode(head.prev)
		}
	}
}

func insert(x int) {
	newnode()
	cur.data = x
	if head == nil {
		head = cur
	} else {
		insnode(head)
	}
}

func displayin(head *node) {
	if head == nil {
		return
	}
	displayin(head.prev)
	fmt.Print(head.data," ")
	displayin(head.nxt)
}

func displaypre(head *node) {
	if head == nil {
		return
	}
	fmt.Print(head.data," ")
	displayin(head.prev)
	displayin(head.nxt)
}

func displaypost(head *node) {
	if head == nil {
		return
	}
	displayin(head.prev)
	displayin(head.nxt)
	fmt.Print(head.data," ")
}

func display() {
	if head == nil {
		fmt.Println("Empty")
	} else {
		var choice int
		fmt.Print("1. Preorder\n2. Inorder\n3. Postorder: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			displaypre(head)
		case 2:
			displayin(head)
		case 3:
			displaypost(head)
		}
		fmt.Println()
	}
}

var flag int

func sch(head *node, x int) {
	if head != nil {
		sch(head.prev, x)	
		if head.data == x {
			flag = 1
		}
		sch(head.nxt, x)
	}
}

func search(x int) {
	flag = 0
	if head == nil {
		fmt.Println("Empty")
		return
	}
	if head.data == x {
		fmt.Println("Found")
	} else {
		sch(head, x)
		if flag == 1 {
			fmt.Println("Found")
		} else {
			fmt.Println("Not Found")
		}
	} 
}


func scannum() int {
	var x int
	fmt.Print("Enetr Number: ")
	fmt.Scan(&x)
	return x
}

func main() {
	var choice, x, exit int
	for exit != -1 {
		fmt.Println("1. Insert\n2. Display\n3. Search")
		fmt.Print("4. Exit\nEnter Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			x = scannum()
			insert(x)
		case 2:
			display()
		case 3:
			x = scannum()
			search(x)
		case 4:
			exit = -1
		default:
			fmt.Println("Enetr Correct Option")
		}
	}
}