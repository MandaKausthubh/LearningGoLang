package main

import (
    "fmt"
)


func BinarySearch( arr *[]int, l, r, val int ) (int) {
    l--; r++;
    for r - l > 1 {
        mid := l + (r-l)/2
        if (*arr)[mid] <= val {
            l = mid
        } else {
            r = mid
        }
    }
    return l
}

type Node struct {
    val int
    left_child *Node
    right_child *Node
}

func InsertNode ( root *Node, new_node *Node ) *Node {

    if (new_node == nil) {
        fmt.Println("New node is nil")
        return root
    }
    if (root == nil) {
        fmt.Println("Root is nil")
        return new_node
    }

    if (root.val >= new_node.val) {
        if (root.left_child == nil) {
            root.left_child = new_node
        } else {
            InsertNode(root.left_child, new_node)
        }
    } else {
        if (root.right_child == nil) {
            root.right_child = new_node
        } else {
            InsertNode(root.right_child, new_node)
        }
    }
    return root
}

func PrintTreeInOrder( root *Node ) {
    if root == nil {
        return
    }

    PrintTreeInOrder(root.left_child)
    fmt.Println(root.val)
    PrintTreeInOrder(root.right_child)
}


func main() {
    arr := []int{1,3,4,5,7,11, -3, -5, -2, -2};
    var root *Node = nil

    for _, ele := range arr {
        fmt.Println("Inserting element", ele)
        var new_node Node = Node{val: ele}
        root = InsertNode(root, &new_node)
    }

    PrintTreeInOrder(root)
}
