package main

import (
    "./tree"
)

func main() {
    mtree := tree.NewGimTree()
    mtree.MasterRecordDir = "./gim"
    mtree.MasterDir = "."
    mtree.GenerateMasterTree()
    tree.WalkNode(mtree.MasterRoot, 0)
}

