package framework

import (
	"strings"
)

type TreeNode struct {
	children []*TreeNode
	handler  func(ctx *Context)
	param    string
	parent   *TreeNode
}

func NewTreeNode() *TreeNode {
	return &TreeNode{
		children: make([]*TreeNode, 0),
		handler:  nil,
		param:    "",
		parent:   nil,
	}
}

func isGeneral(param string) bool {
	return strings.HasPrefix(param, ":")
}

func (tn *TreeNode) Insert(pathname string, handler func(ctx *Context)) {
	node := tn

	params := strings.Split(pathname, "/")

	for _, param := range params {
		child := node.findChild(param)

		if child == nil {
			child = &TreeNode{
				param:    param,
				children: make([]*TreeNode, 0),
				parent:   node,
			}
			node.children = append(node.children, child)
		}

		node = child
	}

	node.handler = handler
}

func (tn *TreeNode) findChild(param string) *TreeNode {
	for _, child := range tn.children {
		if child.param == param {
			return child
		}
	}
	return nil
}

func (tn *TreeNode) Search(pathname string) *TreeNode {
	params := strings.Split(pathname, "/")
	return dfs(tn, params)
}

func dfs(node *TreeNode, params []string) *TreeNode {
	currentParam := params[0]
	isLastParam := len(params) == 1

	for _, child := range node.children {

		if isLastParam {
			if isGeneral(child.param) {
				return child
			}

			if child.param == currentParam {
				return child
			}

			continue
		}

		if !isGeneral(child.param) && child.param != currentParam {
			continue
		}

		result := dfs(child, params[1:])

		if result != nil {
			return result
		}
	}

	return nil
}

func (tn *TreeNode) ParseParams(pathname string) map[string]string {
	node := tn

	pathname = strings.TrimSuffix(pathname, "/")
	paramArr := strings.Split(pathname, "/")

	paramDicts := make(map[string]string)
	for i := len(paramArr) - 1; i >= 0; i-- {
		if isGeneral(node.param) {
			paramDicts[node.param] = paramArr[i]
		}
		node = node.parent
	}

	return paramDicts
}
