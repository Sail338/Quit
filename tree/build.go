package tree

import (
    "os"
    "io/ioutil"
    fp "path/filepath"
    "crypto/sha256"
    b64 "encoding/base64"
    "fmt"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func (gt GimTree) GenerateMasterTree() {
    gt.MasterRoot = gt.GenerateTree(gt.MasterDir)
}

func (gt GimTree) GenerateTree(path string) *TNode {
    f, err := os.Open(path)
    check(err)
    var stat os.FileInfo
    stat, err = f.Stat()
    check(err)
    mode := stat.Mode()
    if mode.IsDir() {
        tn := NewTNode()
        var files []os.FileInfo
        files, err = ioutil.ReadDir(path)
        tn.Children = make([]*TNode, len(files))
        check(err)
        for i, file := range files {
            var child *TNode
            nodeMode := file.Mode()
            if nodeMode.IsDir() {
                child = gt.GenerateTree(fp.Join(path, file.Name()))
            } else {
                child = gt.GenerateTNode(fp.Join(path, file.Name()))
            }
            child.Parent = tn
            tn.Children[i] = child
        }
        tn.NType = "tree"
        return tn
    }
    return gt.GenerateTNode(path)
}

func (gt GimTree) GenerateTNode(fpath string) *TNode {
    f, err := os.Open(fpath)
    check(err)
    var stat os.FileInfo
    stat, err = f.Stat()
    check(err)
    size := int(stat.Size())
    n := 0
    filedata := make([]byte, 0)
    buffer := make([]byte, 4096)
    for size > 0 {
        n, err = f.Read(buffer)
        check(err)
        filedata = append(filedata, buffer...)
        size -= n
    }
    ch := make(chan []byte, 1)
    sha := sha256.Sum256(filedata)
    ch <- sha[:]
    close(ch)
    b64sha256 := b64.StdEncoding.EncodeToString(<-ch)
    lastModified := stat.ModTime().UnixNano()
    gt.WriteBlob(Base64FSCompat(b64sha256) + ".blob", b64sha256, lastModified)
    fn := NewTNode() 
    fn.Data = Base64FSCompat(b64sha256) + ".blob"
    fn.NType = "file"
    return fn
}

func WalkNode(node *TNode, level int) {
    fmt.Printf("%s node.data: %s node.type %s\n",
        strings.Repeat("----", level), node.Data, node.NType)
    for _, child := range node.Children {
        WalkNode(child, level+1)
    }
}

func (gt GimTree) WriteBlob(blobname, shaHash string, lastMod int64) {
    _, err := os.Stat(gt.MasterRecordDir)
    if err != nil {
        os.Mkdir(gt.MasterRecordDir, 0666)
    }
    var f *os.File
    f, err = os.Create(fp.Join(gt.MasterRecordDir, blobname + ".blob"))
    check(err)
    data := fmt.Sprintf("%s %s", shaHash, lastMod)
    f.Write([]byte(data))
}
