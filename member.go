package v2ex

// =======================================
// 成员 API
type Avatar struct {
	AvatarMini   string `json:"avatar_mini"`
	AvatarNormal string `json:"avatar_normal"`
	AvatarLarge  string `json:"avatar_large"`
}

//用户的自我介绍，及其登记的社交网络信息
type Member struct {
	Avatar

	Status   string `json:"status"`
	ID       int    `json:"id"`
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
	Created  int    `json:"created"`
}

type MemberService interface {
	GetMemberByID(int) (*Member, error)
	GetMemberByUsername(string) (*Member, error)
}
