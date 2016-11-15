package tree

import (
    "os"
    "io/ioutil"
    fp "path/filepath"
    "crypto/sha256"
    "fmt"
    "strings"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func InitGimMasterTree(root string) *GimMasterTree {
    gt := GimMasterTree{root, nil, ""}
    path := fp.Join(gt.MasterDir, ".gim")
    if !Fexists(path) {
        os.Mkdir(path, 0755)
    }
    gt.MasterRecordDir = path
    gt.MasterRoot = gt.GenerateTree(gt.MasterDir)
    return &gt
}

func (gt GimMasterTree) GenerateTree(path string) *TNode {
    if path == ".gim" || path == ".git" {
        return nil
    }
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
        check(err)
        tn.Children = make([]*TNode, 0)
        for _, file := range files {
            child := NewTNode()
            nodeMode := file.Mode()
            if nodeMode.IsDir() {
                child = gt.GenerateTree(fp.Join(path, file.Name()))
            } else {
                child = gt.GenerateTNode(fp.Join(path, file.Name()))
            }
            if child != nil {
                child.Parent = tn
                tn.Children = append(tn.Children, child)
            }
        }
        tn.Type = "tree"
        tn.Path = path
        return tn
    }
    return gt.GenerateTNode(path)
}

func (gt GimMasterTree) GenerateTNode(fpath string) *TNode {
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
    b64sha256 := B64Enc(<-ch)
    lastModified := stat.ModTime().UnixNano()
    fn := NewTNode() 
    fn.Data = fmt.Sprintf("%s %d %s", b64sha256, lastModified, fpath)
    fn.Type = "file"
    fn.Path = fpath
    return fn
}

func WalkNode(node *TNode, level int) {
    if node != nil {
        fmt.Printf("%s node.data: %s node.type %s\n",
            strings.Repeat("----", level), node.Data, node.Type)
        for _, child := range node.Children {
            WalkNode(child, level+1)
        }
    }
}

func Fexists(path string) bool {
    _, err := os.Stat(path)
    if err != nil {
        return false
    }
    return true
}

func (gt GimMasterTree) WriteBlob(blob Blob) string {
    _, err := os.Stat(gt.MasterRecordDir)
    if err != nil {
        os.Mkdir(gt.MasterRecordDir, 0755)
    }
    var f *os.File
    path := fp.Join(gt.MasterRecordDir,
            fmt.Sprintf("%s.blob", Base64FSCompat(blob.ShaHash)))
    if !Fexists(path) {
        f, err = os.Create(path)
        check(err)
        defer f.Close()
        data := fmt.Sprintf("%s %d %s", blob.ShaHash, blob.ModTime, blob.ActualPath)
        f.Write([]byte(data))
        return Base64FSCompat(blob.ShaHash)
    }
    return "preexisting"
}

/*
 * Writes a tree to a file. Within each .tree file contains the hashes of its children
 * and the filename of the .tree is the sha hash of the children.
 * So here is the issue that we are facing. Let's say two files have the same contents
 * or two directories have the same contents. How do we make sure that there are no
 * duplicates in the records while also being able to properly record said duplicates.
 */

func (gt GimMasterTree) WriteTree(root *TNode) string {
    if root != nil {
        println("MEEP", root.Path, root.Type)
        switch (root.Type) {
            case "file":
                tokens := strings.Split(root.Data, " ")
                modTime,_ := strconv.ParseInt(tokens[1], 10, 64)
                blob := Blob{tokens[0], modTime, tokens[2]}
                return gt.WriteBlob(blob)
            case "tree":
                childs := make([][]string, len(root.Children))
                for i, c := range root.Children {
                    childs[i] = make([]string, 2)
                    childs[i][0] = gt.WriteTree(c)
                    childs[i][1] = c.Type
                }
                // collecting together hashes
                hashes := make([]string, len(childs))
                for i, ht := range childs {
                    hashes[i] = ht[0]
                }
                joined := strings.Join(hashes, "")
                sha := sha256.Sum256([]byte(joined))
                ch := make(chan []byte, 1)
                ch <- sha[:]
                treeHash := B64Enc(<-ch)
                close(ch)
                path := fp.Join(gt.MasterRecordDir,
                        fmt.Sprintf("%s.tree", Base64FSCompat(treeHash)))
                if Fexists(path) {
                    // cycle through numeric suffixes until one is found
                    // (i.e. <hash>.1.tree)
                }
                f, err := os.Create(path)
                check(err)
                defer f.Close()
                f.Write([]byte(root.Path + "\n"))
                for _, ht := range childs {
                    f.Write([]byte(fmt.Sprintf("%s %s\n", ht[0], ht[1])))
                }
                return treeHash
        }
    }
    return ""
}
