package main
import "fmt"

//Queue represents a queue that holds a slice

type Queue struct{
	items []int
}

//Enqueue adds value at the end
func (q *Queue) Enqueue(i int){
	q.items = append(q.items, i)
}



//Dequeue removes a value at the front
//returns the Removed value
func (q *Queue) Dequeue() int{
	toRemove:= q.items[0]
	q.items = q.items[1:]
	return toRemove
}


func main(){
	myqueue:= Queue{}
	fmt.Println(myqueue)
	myqueue.Enqueue(100)
	myqueue.Enqueue(200)
	myqueue.Enqueue(300)
	fmt.Println(myqueue.items)
	myqueue.Dequeue()
	fmt.Println(myqueue)
}