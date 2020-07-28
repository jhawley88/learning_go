package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (current *Node) Insert(newValue int) {
	if newValue > current.Val {
		//move right
		if current.Right == nil {
			current.Right = &Node{Val: newValue}
		} else {
			current.Right.Insert(newValue)
		}
	} else if newValue <= current.Val {
		//move left
		if current.Left == nil {
			current.Left = &Node{Val: newValue}
		} else {
			current.Left.Insert(newValue)
		}
	}
}

func (current *Node) FindMin() *Node {
	currentNode := current

	// While the node has a child that is less than it, move to it
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode
}

func (current *Node) FindMax() *Node {
	currentNode := current
	// While the node has a child that is less than it, move to it
	for currentNode.Right != nil {
		currentNode = currentNode.Right
	}
	return currentNode
}

func (current *Node) Search(key int) bool {

	if current == nil {
		return false
	}

	if key > current.Val {
		//move right
		return current.Right.Search(key)
	} else if key < current.Val {
		return current.Left.Search(key)
		//move left
	}

	return true
}

// Efficient in order traversal algo
func (root *Node) MorrisTraversal() []int {
	ordered := []int{}
	tourist := root

	for tourist != nil {
		guide := tourist.Left

		//If the guide is on a node
		if guide != nil {
			// fmt.Println("guide", guide)
			// fmt.Println("guide is here")
			for guide.Right != nil && guide.Right != tourist {
				guide = guide.Right
			}

			if guide.Right == nil {
				guide.Right = tourist
				tourist = tourist.Left
			} else {
				guide.Right = nil
				ordered = append(ordered, tourist.Val)
				tourist = tourist.Right
			}
		} else {
			ordered = append(ordered, tourist.Val)
			tourist = tourist.Right
		}
		// tourist = nil
	}
	return ordered
}

func (root *Node) postorderTraversal() []int {
	ans := []int{}
	var cur, prev *Node
	cur = root
	for cur != nil {
		if cur.Right == nil {
			ans = append(ans, cur.Val)
			cur = cur.Left
		} else {
			prev = cur.Right
			for prev.Left != nil && prev.Left != cur {
				prev = prev.Left
			}
			if prev.Left == nil {
				prev.Left = cur
				ans = append(ans, cur.Val)
				cur = cur.Right
			} else {
				prev.Left = nil
				cur = cur.Left
			}
		}

	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return ans
}

func (root *Node) bfs() []int {
	ordered := []int{}
	queue := []Node{}
	queue = append(queue, *root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Left != nil {
			queue = append(queue, *current.Left)
		}
		if current.Right != nil {
			queue = append(queue, *current.Right)
		}
		ordered = append(ordered, current.Val)
	}
	return ordered
}

func main() {
	// tree := &Node{Val: 100}
	// fmt.Println(tree)
	// tree.Insert(80)
	// tree.Insert(60)
	// tree.Insert(120)
	// tree.Insert(150)
	// fmt.Println(tree)
	// fmt.Println(tree.Right)
	// fmt.Println(tree.Search(60))
	// tree.FindMin()
	// tree.FindMax()
	// dfs := tree.postorderTraversal()
	// fmt.Println("dfs", dfs)
	// MorrisTraversal := tree.MorrisTraversal()
	// fmt.Println("MorrisTraversal", MorrisTraversal)
	// bfs := tree.bfs()
	// fmt.Println("bfs", bfs)
}
