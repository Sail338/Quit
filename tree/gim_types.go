package tree

type TNode struct {
    Parent *TNode
    Children []*TNode
    Data string
    Type string
    Path string
}

func NewTNode() *TNode {
    return &TNode{nil, nil, "", "", ""}
}

type GimMasterTree struct {
    MasterDir string
    MasterRoot *TNode
    MasterRecordDir string
}

func NewGimMasterTree() GimMasterTree {
    return GimMasterTree{"", nil, ""}
}

type Blob struct {
    ShaHash string
    ModTime int64
    ActualPath string
}
