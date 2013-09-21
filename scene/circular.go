// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// +build !fast

package scene

func (n *Node) checkForCircular() {
	top := n.Top()
	visited := make(map[*Node]bool)
	top.Traverse(func(i int, current *Node) bool {
		_, visitedAlready := visited[current]
		if visitedAlready {
			panic(CircularErr)
		}
		visited[current] = true
		return true
	})
}
