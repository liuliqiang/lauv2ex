package v2ex

import "strconv"

// =======================================
// 成员 API
type Avatar struct {
	AvatarMini   string `json:"avatar_mini"`
	AvatarNormal string `json:"avatar_normal"`
	AvatarLarge  string `json:"avatar_large"`
}

type Node struct {
	ID               uint32 `json:"id"`                //节点 ID
	Name             string `json:"name"`              //节点缩略名
	URL              string `json:"url"`               //节点地址
	Title            string `json:"title"`             //节点名称
	TitleAlternative string `json:"title_alternative"` //备选节点名称
	Topics           uint32 `json:"topics"`            //节点主题总数
	Header           string `json:"header"`            //节点头部信息
	Footer           string `json:"footer"`            //节点脚部信息
	Created          int64  `json:"created"`           //节点创建时间
	Avatar
}

//用户的自我介绍，及其登记的社交网络信息
type Member struct {
	Status   string `json:"status"`
	ID       uint32 `json:"id"`
	URL      string `json:"url"`
	Username string `json:"username"`
	Website  string `json:"website"`
	Twitter  string `json:"twitter"`
	Psn      string `json:"psn"`
	Github   string `json:"github"`
	Btc      string `json:"btc"`
	Location string `json:"location"`
	TagLine  string `json:"tagline"`
	Bio      string `json:"bio"`
	Avatar
	Created int64 `json:"created"`
}

//通过用户的ID获取
func GetMemberByID(id uint32) (member Member, err error) {
	url := v2exAPI + "/members/show.json?id=" + strconv.Itoa(int(id))
	err = getAPIData(url, &member)
	return
}

//通过用户的名字获取
func GetMemberByUsername(name string) (member Member, err error) {
	if name == "" {
		return member, ErrInvalidUsername
	}
	url := v2exAPI + "/members/show.json?username=" + name
	err = getAPIData(url, &member)
	return
}

type UserService interface {
	//通过用户的名字获取
	GetMemberByUsername(name string) (member Member, err error)
}
