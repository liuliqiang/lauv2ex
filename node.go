package v2ex

import "strconv"

type NodeService interface {
	//通过节点ID获取单个节点信息
	GetNodeByID(id uint32) (node Node, err error)
	//通过节点名字获取单个节点信息
	GetNodeByName(name string) (node Node, err error)
	//获取全部节点
	GetNodes() (nodes Nodes, err error)
}

//通过节点ID获取单个节点信息
func GetNodeByID(id uint32) (node Node, err error) {
	url := v2exAPI + "/nodes/show.json?id=" + strconv.Itoa(int(id))
	err = getAPIData(url, &node)
	return
}

//通过节点名字获取单个节点信息
func GetNodeByName(name string) (node Node, err error) {
	if name == "" {
		return node, ErrInvalidUsername
	}
	url := v2exAPI + "/nodes/show.json?name=" + name
	err = getAPIData(url, &node)
	return
}

type Nodes []Node

//获取全部节点
func GetNodes() (nodes Nodes, err error) {
	url := v2exAPI + "/nodes/all.json"
	err = getAPIData(url, &nodes)
	return
}
