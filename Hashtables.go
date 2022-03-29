//Hash - fast with insert delete or search
// downside - collisions. two names might have same hash code
//collision handling 1) open addressing 2) separate chaining using linked lists

package main

import (
	"fmt"
	
)

const arraySize = 7

//hash table part(array): 
//hashTable struct
type hashTable struct{
	array [arraySize]*bucket
} 
//bucket is a linked list in each slot of 

//insert() method
//search() method
//delete() method


//bucket part (linked list):
//bucket struct
type bucket struct{
	head *bucketNode
}
//bucketNode struct
type bucketNode struct{
	key string
	next *bucketNode
}
///insert() method
func (h *hashTable) insert(key string){
	index := hash(key)
	h.array[index].insert(key)
}
/*
//search() method
func (h *hashTable) search(key string){
	index := hash(key)
}
//delete() method
func (h *hashTable) delete(key string){
	index := hash(key)
}
*/
//insert will take in a key and delete it from the hash table
func (b *bucket) insert(k string){
	newNode:= &bucketNode{key: k}
	newNode.next = b.head
}
//search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool{
	currentNode:= b.head
	for currentNode != nil{
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	} 
	return false
} 


//hash
func hash(key string) int{
	sum:= 0
	for _, v := range key{
		sum +=int(v)
	}
	return sum % arraySize
}

//init will create a bucket in each slot of the hash table
func Init() *hashTable {
	result:= &hashTable{}
	for i:= range result.array {
		result.array[i]= &bucket{}
	}
	return result
}

func main(){
	testhashTable:= Init()
	fmt.Println(testhashTable)
	fmt.Println(hash("Randy"))

	testBucket:= &bucket{}
	testBucket.insert("Randy")
	fmt.Println(testBucket.search("Randy"))
	fmt.Println(testBucket.search("Eric"))
}
