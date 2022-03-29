package main
import "fmt"

//stack represents stack that holds a slice

type stack struct{
	items []int
}

//push will add a value at the end of a stack
func (s *stack) push(i int){
	s.items = append(s.items, i)
}


//pop will remove a value at the end of a stack
// returns the removed value

func (s *stack) pop() int{

	l:= len(s.items)-1
	toRemove:= s.items[l]  //LIFO
	s.items = s.items[:l] //starts from begin. 
	//and stops at the last before element
	//leaving or popping the last one out of the stack!
	return toRemove
}

func main(){
	mystack:= stack{}
	fmt.Println(mystack)

	mystack.push(100)
	mystack.push(200)
	mystack.push(300)
	fmt.Println(mystack.items)

	mystack.pop()
	fmt.Println(mystack)
}