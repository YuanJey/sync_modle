package utils

import "github.com/YuanJey/sync_modle/pkg/base_info"

func NewBuildTree(nodes []*base_info.ThirdDept) map[string]*base_info.ThirdDept {
	temp := make(map[string]*base_info.ThirdDept)
	for i := range nodes {
		temp[nodes[i].ThirdUnionId] = nodes[i]
	}
	NewBuildNodes(nodes, temp)
	//去除重复的节点
	return temp
}

func NewBuildNodes(nodes []*base_info.ThirdDept, temp map[string]*base_info.ThirdDept) {
	for i := range nodes {
		if parent, ok := temp[nodes[i].ParentId]; ok {
			parent.Children = append(parent.Children, nodes[i])
			continue
		}
		for _, dept := range temp {
			if dept.Children == nil {
				NewBuildNodes(dept.Children, temp)
			}
		}
	}
}
