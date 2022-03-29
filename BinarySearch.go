
//Binary search tree

package main
import "fmt"


var count int
//define a node
type Node struct {
	key int
	left *Node
	right *Node
}

//make an insert method
//key to add should not be in the tree
func (n *Node) insert(k int) {
	if n.key < k {
		//move right
		if n.right == nil{
			n.right=&Node{key: k}
		} else {
			n.right.insert(k)
		}
	} else if n.key > k {
		//move left
		if n.left == nil {
			n.left =&Node{key: k}
		} else {
			n.left.insert(k)
		}
	}
}

//search
//return true if there is a node with that value
func(n *Node) search(k int) bool{

	count++

	if n ==nil{
		return false
	}

	if n.key < k{
		//move right
	return	n.right.search(k)
	} else if n.key > k{
		//move left
	return n.left.search(k)

	}
	return true
}

func main() {
	tree:= &Node{key: 100}
	tree.insert(200)
	tree.insert(300)

	tree.insert(52)
	tree.insert(203)
	tree.insert(19)
	tree.insert(76)
	tree.insert(150)
	tree.insert(310)
	tree.insert(7)
	tree.insert(24)
	tree.insert(88)
	tree.insert(276)

	tree.search(76)
	fmt.Println(tree)
	fmt.Println(tree.search(400))

	fmt.Println(count)



}