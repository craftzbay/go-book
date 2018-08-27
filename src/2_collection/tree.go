package main

import "fmt"

// модны нэг зангилаа
type Node struct {
  right, left *Node // зүүн, баруун зангилаа
  value string	 // зангилаан дээрх утга
}

// мод бүтэц
type BinaryTree struct {
  root *Node // оройн элементийг заана
}

func addNode(val string) *Node {
  return &Node{nil, nil, val}
}

func (b *BinaryTree) Insert(val string) (n *Node) {
  if b.root == nil {
    n = addNode(val)
    b.root = n
  } else {
    n = b.insert(b.root, val)
  }
  return
}

func (b *BinaryTree) insert(root *Node, val string) *Node {
  switch {
  case root == nil:
    return addNode(val)
  case val <= root.value:
    root.left = b.insert(root.left, val)
  case val > root.value:
    root.right = b.insert(root.right, val)
  }
  return root
}

func (b *BinaryTree) Print() {
  printTree(b.root)
}

func printTree(n *Node) { 
    if n == nil {
        return;
    }
    printTree(n.left); 
    fmt.Printf("%s\n", n.value)
    printTree(n.right); 
}

func find(n *Node, val string) {
  if (n == nil) {
    return
  }
  if n.value == val {
    fmt.Printf("%s үг модонд олдлоо!\n", val)
  }
  if val <= n.value {
    find(n.left, val)   
  } else {
    find(n.right, val)   
  }
}

func (b *BinaryTree) Find(val string) {
  find(b.root, val)
}

func main() {
  b := new(BinaryTree)
  b.Insert("lemon")
  b.Insert("apple")
  b.Insert("grape")
  b.Insert("orange")
  b.Insert("plum")
  b.Insert("pear")
  
  b.Print()
  
  b.Find("pear")
}