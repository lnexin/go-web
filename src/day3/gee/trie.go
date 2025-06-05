package gee

import "strings"

type node struct {
	pattern  string
	children []*node
	part     string
	isWild   bool
}

// 第一个匹配的子节点
func (n *node) matchFirstChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}

	}
	return nil
}

// 所有匹配的子节点
func (n *node) matchChildren(part string) []*node {
	rlts := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			rlts = append(rlts, child)
		}
	}
	return rlts
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchFirstChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == '*' || part[0] == ':'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
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
