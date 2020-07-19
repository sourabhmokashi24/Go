package main
import "fmt"

type node struct {
	data int
	nxt *node
}

var cur, tmp *node
var head *node = nil

func newnode() {
	cur = &node {}
	cur.nxt = nil
}

func insert(x int) {
	newnode()
	if head == nil {
		cur.data = x
		head = cur
	} else {
		tmp = head
		for tmp.nxt != nil {
			tmp = tmp.nxt
		}
		cur.data = x
		tmp.nxt = cur
	}
}

func diplay() {
	fmt.Println("List:")
	for tmp = head; tmp != nil; tmp = tmp.nxt {
		fmt.Print(tmp.data)
		if tmp.nxt != nil {
			fmt.Print("-->")
		}
	}
	fmt.Println()
}

func search(x int) int {
	for tmp = head; tmp != nil; tmp = tmp.nxt {
		if tmp.data == x {
			return 1
		} 
	}
	return 0
}

func insertatfist(x int) {
	newnode()
	if head == nil {
		head = cur
		cur.data = x
	} else {
		cur.nxt = head
		head = cur
		cur.data = x
	}
}

func reverse(tmp *node) {
	if tmp.nxt == nil {
		head = tmp
		return
	}
	reverse(tmp.nxt)
	var q *node = tmp.nxt
	q.nxt = tmp
	tmp.nxt = nil
}

func reversedisplay(tmp *node) {
	if tmp == nil {
		return
	}
	reversedisplay(tmp.nxt)
	fmt.Print(tmp.data)
	if tmp != head {
		fmt.Print("-->")
	} else {
		fmt.Println()
	}
}

func insertatmid(x, y int) {
	newnode()
	tmp = head
	for tmp.data != y {
		tmp = tmp.nxt
	}
	cur.nxt = tmp.nxt
	tmp.nxt = cur
	cur.data = x
}

func deletelast() {
	var prev *node
	prev = head
	tmp = head.nxt
	for tmp.nxt != nil {
		tmp = tmp.nxt
		prev = prev.nxt
	}
	prev.nxt = nil
}

func deleteatmid(x int) {
	if head.data == x {
		tmp = head
		head = head.nxt
		tmp = nil
	} else {
		var prev *node
		prev = head
		tmp = head.nxt
		for tmp.data != x {
			tmp = tmp.nxt
			prev = prev.nxt
		}
		prev.nxt = tmp.nxt
		tmp.nxt = nil
	}
}

func deletefirst() {
	if head.nxt == nil {
		head = nil
	} else {
		head = head.nxt
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
		fmt.Println("1. Insert\n2. InsertAtFirst\n3. InsertAtMid\n4. Display\n5. ReverseDisplay\n6. Search")
		fmt.Println("7. DeleteLast\n8. DeleteFisrt\n9. DeleteMid\n10. Reverse\n11. Exit")
		fmt.Print("Enter Your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			x = scannum()
			insert(x)
		case 2:
			x = scannum()
			insertatfist(x)
		case 3:
			x = scannum()
			var index int
			index = scannum()
			insertatmid(x, index)
		case 4:
			diplay()
		case 5:
			reversedisplay(head)
		case 6:
			x = scannum()
			var flag = search(x)
			if flag == 1 {
				fmt.Println("Found")
			} else {
				fmt.Println("Not Found")
			}
		case 7:
			deletelast()
		case 8:
			deletefirst()
		case 9:
			x = scannum()
			var flag = search(x)
			if flag == 1 {
				deleteatmid(x)
			} else {
				fmt.Println("Not Found")
			}
		case 10:
			reverse(head)
		case 11:
			exit = -1
		default:
			fmt.Println("Enetr Correct Option")
		}
	}
}