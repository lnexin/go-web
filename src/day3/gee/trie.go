package gee

import "strings"

type Node struct {
	Pattern  string  `json:"Pattern"`
	Children []*Node `json:"Children"`
	Part     string  `json:"Part"`
	IsWild   bool    `json:"IsWild"`
}

// 第一个匹配的子节点
func (n *Node) matchFirstChild(part string) *Node {
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			return child
		}

	}
	return nil
}

// 所有匹配的子节点
func (n *Node) matchChildren(part string) []*Node {
	rlts := make([]*Node, 0)
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			rlts = append(rlts, child)
		}
	}
	return rlts
}

func (n *Node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.Pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchFirstChild(part)
	if child == nil {
		child = &Node{Part: part, IsWild: part[0] == '*' || part[0] == ':'}
		n.Children = append(n.Children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *Node) search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(n.Part, "*") {
		if n.Pattern == "" {
			return nil
		}
		return n
	}

	p := parts[height]
	chs := n.matchChildren(p)

	for _, ch := range chs {
		rlt := ch.search(parts, height+1)
		if rlt != nil {
			return rlt
		}
	}
	return nil
}
