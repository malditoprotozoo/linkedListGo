package main

import "fmt"

// Node : Defining a type Node
type Node struct {
	body int
	pos  int
	prev *Node
	next *Node
}

// List : Conjunct of Nodes
type List struct {
	length int
	start  *Node
	end    *Node
}

// Append : Append new Node to List
func (list *List) Append(newNode *Node) {
	if list.length == 0 {
		list.start = newNode
		list.end = newNode
	} else {
		lastNode := list.end
		lastNode.next = newNode
		list.end = newNode
	}
	newNode.pos = list.length
	list.length++
}

// Insert : Insert new Node to List in a specific position
func (list *List) Insert(node, pos int) {
	newNode := new(Node)
	newNode.body = node
	newNode.pos = pos
	newNode.next = nil
	if pos == 0 {
		newNode.next = list.start
		list.start = newNode
		return
	}
	newNode2 := new(Node)
	newNode2 = list.start
	for i := 0; i < pos-1; i++ {
		newNode2 = newNode2.next
	}
	newNode.next = newNode2.next
	newNode2.next = newNode
}

// Show : Show all Nodes on List
func (list *List) Show() {
	length := 0
	start := list.start
	fmt.Println("The nodes on the list are:")
	if start != nil {
		for start != nil {
			fmt.Printf("Data = %d \n", start.body)
			length++
			start = start.next
		}
	} else {
		fmt.Println("The list is empty")
	}
	list.length = length
}

// Find : Find item on the list
func (list *List) Find(item int) int {
	index := 0
	start := list.start
	for start != nil {
		if start.body != item {
			start = start.next
		} else {
			fmt.Printf("The item %d is in the position %d \n", item, start.pos)
			start = start.next
			index++
		}
	}
	if index != 0 {
		return index
	}
	return -1
}

// FindWithPos : Find what item is in a given position
func (list *List) FindWithPos(pos int) int {
	start := list.start
	for start != nil {
		if start.pos == pos {
			return start.body
		}
		start = start.next
	}
	return -1
}

// RemoveItem : Remove item in given position
func (list *List) RemoveItem(pos int) {
	if list.start == nil {
		return
	}
	newNode := new(Node)
	newNode = list.start
	if pos == 0 {
		list.start = list.start.next
		return
	}
	for i := 0; newNode != nil && i < pos-1; i++ {
		newNode = newNode.next
	}
	if newNode == nil || newNode.next == nil {
		return
	}
	newNode2 := new(Node)
	newNode2 = newNode.next.next
	newNode.next = newNode2
}

// CleanList : Removes all items
func (list *List) CleanList() {
	for list.start != nil {
		list.start = nil
	}
}

// ShowNext : Show the item that comes next to given position
func (list *List) ShowNext(pos int) int {
	start := list.start
	for start != nil {
		if start.pos == pos+1 {
			return start.body
		}
		start = start.next
	}
	return -1
}

// ShowPrev : Show the item that comes before togiven position
func (list *List) ShowPrev(pos int) int {
	start := list.start
	for start != nil {
		if start.pos == pos-1 {
			return start.body
		}
		start = start.next
	}
	return -1
}

// ShowFirst : Show the first iteration of a Node
func (list *List) ShowFirst(item int) int {
	start := list.start
	for start != nil {
		if start.body == item {
			return start.pos
		}
		start = start.next
	}
	return -1
}

// First : Show the first Node
func (list *List) First() {
	fmt.Printf("The first node on the list is %d at the position %d \n", list.start.body, list.start.pos)
}

// Last : Show the position after the last node
func (list *List) Last() {
	start := list.start
	for start != nil {
		if start.pos == list.length-1 {
			fmt.Printf("The position after the last node is %d \n", start.pos+1)
		}
		start = start.next
	}
}

func main() {
	MyList := &List{}
	fmt.Println("How many nodes do you want?")
	var input int
	fmt.Scanln(&input)
	for i := 0; i < input; i++ {
		fmt.Println("Add a Node")
		var node int
		fmt.Scanln(&node)
		addNode := new(Node)
		addNode.body = node
		MyList.Append(addNode)
	}
	for {
		fmt.Println(`
			Write the number of what you want to do now: 
			1. Show My List 
			2. Search An Item 
			3. Search What Item is in Given Position 
			4. Show the item that comes next a position 
			5. Show the item that comes before a position 
			6. Show the first iterarion of a Node 
			7. Insert node in given position
			8. Remove a node
			9. Remove all items
			10. Show the first node on the list
			11. Show the position after the last node
			12. Show the list backwards`)
		var command int
		fmt.Scanln(&command)
		if command == 1 {
			MyList.Show()
		} else if command == 2 {
			fmt.Println("Search your Data")
			var search int
			fmt.Scanln(&search)
			if MyList.Find(search) < 0 {
				fmt.Println("Item not Found")
			}
		} else if command == 3 {
			fmt.Println("Write the position")
			var pos int
			fmt.Scanln(&pos)
			if MyList.FindWithPos(pos) < 0 {
				fmt.Println("Position Unavailable")
			} else {
				fmt.Printf("In position %d there is the data %d \n", pos, MyList.FindWithPos(pos))
			}
		} else if command == 4 {
			fmt.Println("Write a Position")
			var pos int
			fmt.Scanln(&pos)
			if MyList.ShowNext(pos) < 0 {
				fmt.Println("Position Unavailable")
			} else {
				fmt.Printf("In position %d there is the data %d \n", pos+1, MyList.ShowNext(pos))
			}
		} else if command == 5 {
			fmt.Println("Write a Position")
			var pos int
			fmt.Scanln(&pos)
			if MyList.ShowPrev(pos) < 0 {
				fmt.Println("Position Unavailable")
			} else {
				fmt.Printf("In position %d there is the item %d \n", pos-1, MyList.ShowPrev(pos))
			}
		} else if command == 6 {
			fmt.Println("What Node are you Looking For?")
			var node int
			fmt.Scanln(&node)
			if MyList.ShowFirst(node) < 0 {
				fmt.Println("That Node is not in the List")
			} else {
				fmt.Printf("The first iteration of node %d is in the position %d \n", node, MyList.ShowFirst(node))
			}
		} else if command == 7 {
			fmt.Println("Add a Node")
			var node int
			fmt.Scanln(&node)
			fmt.Println("In what position?")
			var pos int
			fmt.Scanln(&pos)
			MyList.Insert(node, pos)
			MyList.Show()
		} else if command == 8 {
			fmt.Println("Select a position")
			var pos int
			fmt.Scanln(&pos)
			MyList.RemoveItem(pos)
			MyList.Show()
		} else if command == 9 {
			MyList.CleanList()
			MyList.Show()
		} else if command == 10 {
			MyList.First()
		} else if command == 11 {
			MyList.Last()
		} else {
			fmt.Println("Command not Found")
		}
	}
}
