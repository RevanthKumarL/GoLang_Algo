package main
import "fmt"

	// MaxHeap struct has a slice that holds the array

	type MaxHeap struct{
		array []int
	}

	// Insert add an element to the heap
	func (h *MaxHeap) insert(key int){
		h.array = append(h.array, key)
		h.MaxHeapifyUp(len(h.array)-1)
	}

	// Extract returns the largest key, and removes it from heap

	func (h *MaxHeap) Extract() int{
		extracted:= h.array[0]
		l:= len(h.array)-1

		if len(h.array) == 0{
			fmt.Println("Cannot extract cuz the array len is 0")
			return -1
		}
		h.array[0] = h.array[l]
		h.array = h.array[:l]

		h.MaxHeapifyDown(0)
		return extracted
	}

	// maxheapify will heapify from bottom top
	func (h *MaxHeap) MaxHeapifyUp(index int){
		for h.array[parent(index)] < h.array[index]{
			h.swap(parent(index), index)
			index = parent(index)
		}

	}

	//MaxHeapifyDown will heapify top to down
	func (h *MaxHeap) MaxHeapifyDown(index int){
		lastIndex:= len(h.array)-1
		l,r:= left(index), right(index)
		childToCompare:= 0

		//loop while index has atleast one child
		for l <= lastIndex{
			if l == lastIndex{ //when left child is only child
				childToCompare = l
			} else if h.array[l] > h.array[r]{ //left child-larger
				childToCompare = l
			} else { //right child is larger
				childToCompare = r
			}
			//compare array value of cur.index to larger child and swap if smaller 
			if h.array[index]<h.array[childToCompare]{
				h.swap(index, childToCompare)
				index = childToCompare
				l,r = left(index), right(index)
			} else {
				return
			}
		}
	}




	//get the parent index
	func parent(i int) int{
		return (i-1)/2  // 2*parent-index +1 = leftchild-index
	}
	//get the left child index
	func left(i int) int{
		return 2*i + 1  
	}
	//get the right child index
	func right(i int) int{
		return 2*i + 2
	}

	//swap keys in the array
	func (h *MaxHeap) swap(i1, i2 int){
		h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
	}



	func main(){

		m := &MaxHeap{}
		fmt.Println(m)
		buildHeap:= []int{10,20,30,40,50,7,4,3,8,9}
		for _, v:=  range buildHeap{
			m.insert(v)

			fmt.Println(m)
		}
		for i:=0;i<7;i++{
			m.Extract()
			fmt.Println(m)
		}



}