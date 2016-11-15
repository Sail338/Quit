package main

import (
    "./tree"
)

func main() {
    mtree := tree.InitGimMasterTree(".")
    println(mtree.MasterRoot == nil)
    mtree.WriteTree(mtree.MasterRoot)
}

