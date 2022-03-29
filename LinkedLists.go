package main
import "fmt"

type node struct{
	data int
	next *node
}

type linkedlist struct{
	head *node
	length int
}

func (l *linkedlist) prepend(n *node){

	//we want to insert n as the new head
	//so we push the existing head to second
	//so second node becomes next node to the new node 'n'
	second:= l.head
	l.head = n 
	l.head.next = second
	//we added to the list so...
	l.length++
}
//to print out data in the receiver, we use value receiver 
//not using a pointer receiver
func (l linkedlist) printlistdata() {
	toPrint:= l.head
	for l.length != 0 {
		fmt.Printf("%d ",toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}
//need to change the receiver, so using pointer receiver
func (l *linkedlist) deleteWithValue(value int){
	if l.length == 0{
		return
	}

	if l.head.data == value{
		l.head = l.head.next
		l.length--
		return
	}

	PrevioustoDelete := l.head
	for PrevioustoDelete.next.data != value {
		if PrevioustoDelete.next.next==nil{
			return
		}
		PrevioustoDelete = PrevioustoDelete.next
	}
	PrevioustoDelete = PrevioustoDelete.next.next
	l.length--
}




func main(){
	mylist:= linkedlist{}
	node1:= &node{data: 48}
	node2:= &node{data: 18}
	node3:= &node{data: 16}
	node4:= &node{data: 11}
	node5:= &node{data: 7}
	node6:= &node{data: 2}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)

	mylist.deleteWithValue(7)
	mylist.printlistdata()

	

	
	
	
}