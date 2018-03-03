package v2ex

import "strconv"

type Node struct {
	ID               int    `json:"id"`                //节点 ID
	Name             string `json:"name"`              //节点缩略名
	URL              string `json:"url"`               //节点地址
	Title            string `json:"title"`             //节点名称
	TitleAlternative string `json:"title_alternative"` //备选节点名称
	Topics           int    `json:"topics"`            //节点主题总数
	Header           string `json:"header"`            //节点头部信息
	Footer           string `json:"footer"`            //节点脚部信息
	Created          int    `json:"created"`           //节点创建时间
	Avatar
}

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
