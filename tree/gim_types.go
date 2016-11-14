package tree

type TNode struct {
    Parent *TNode
    Children []*TNode
    Data string
    NType string
}

func NewTNode() *TNode {
    var n TNode
    n.Parent = nil
    n.Children = nil
    n.Data = ""
    n.NType = ""
    return &n
}

type GimTree struct {
    MasterDir string
    MasterRoot *TNode
    MasterRecordDir string
}

func NewGimTree() GimTree {
    var t GimTree
    t.MasterDir = ""
    t.MasterRoot = nil
    return t
}
