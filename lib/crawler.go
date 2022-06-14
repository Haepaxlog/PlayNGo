package lib

import (
	"os/exec"
	"bytes"
	"regexp"
	"io/ioutil"
	"github.com/golang-collections/collections/stack"
	//"path/filepath"
)

type Node struct{
	path string
	visited bool
	children []*Node
}

func getNode(node *Node) (*Node, error){
	// Read Subdirectories
	entries, err := ioutil.ReadDir(node.path)
	if err != nil {
		return nil, err
	}

	// If No Children, Return Node with nil Children
	if len(entries) == 0 {
		if node.visited {
			return &Node{node.path, node.visited, nil}, nil
		}
		return &Node{node.path, node.visited, nil}, nil
	}

	// If Node was already set, Return Node itself
	if len(node.children) > 0{
		return node, nil
	}

	// If Children, Return Node with Children
	for _, entry := range entries {
		if entry.IsDir(){
			if node.visited {
				node.children = append(node.children, &Node{node.path + "/" + entry.Name(), node.visited, nil})
			}
			node.children = append(node.children, &Node{node.path + "/" + entry.Name(), node.visited, nil})
		}
	}

	return node, nil
}


//Call
//	root := &Node{"/home/conni1410/Music/", false, nil}
//	getTree(root)

func getTree(node *Node) {
	// Get Tree on Root Node with Stack
	node, _ = getNode(node)
	stack := stack.New()
	stack.Push(node)
	for stack.Len() > 0 {
		node = stack.Peek().(*Node)
		node, _ = getNode(node)
		if node.visited || len(node.children) == 0 {
			stack.Pop()
			continue
		}
		for i := len(node.children) - 1; i>= 0; i-- {
			stack.Push(node.children[i])
		}
		node.visited = true

	}

}

// Deprecated but works
func getDir(path string) ([]string, error){
	bin := "ls"
	arg1 := path

	cmd := exec.Command(bin, arg1)

	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	out := bytes.Split(stdout,[]byte("\n"))
	strOut := make([]string, len(out))

	for i := 0; i < len(out); i++ {
		strOut[i] = string(out[i])
		}

	strOut = strOut[:(len(strOut) - 1)]
	return strOut, nil
}

func GetAudioPlaylist(path string) ([]string, error){
	files, err := getDir(path)
	if err != nil {
		return nil, err
	}

	count := 0
	r, err := regexp.Compile(`.*\.(wav|mp3)`)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(files); i++ {
		if r.MatchString(files[i]){
			count += 1
		}
	}

	playlist := make([]string, count)
	y := 0
	for i := 0; i < len(files); i++ {
		if r.MatchString(files[i]){
			playlist[y] = files[i]
			y++
		}
	}
	return playlist, nil
}
