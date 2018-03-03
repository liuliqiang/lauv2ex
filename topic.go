package v2ex

// =======================================
// 主题 API
type Topic struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	URL             string `json:"url"`
	Content         string `json:"content"`
	ContentRendered string `json:"content_rendered"`
	Replies         int    `json:"replies"`
	Member          Member `json:"member"`
	Node            Node   `json:"node"`
	Created         int    `json:"created"`
	LastModified    int    `json:"last_modified"`
	LastTouched     int    `json:"last_touched"`
}

type Reply struct {
	ID              int    `json:"id"`     //Reply ID
	Thanks          int    `json:"thanks"` //感谢数量
	Content         string `json:"content"`
	ContentRendered string `json:"content_rendered"`
	Member          Member `json:"member"`
	Created         int    `json:"created"`
	LastModified    int    `json:"last_modified"`
}

type TopicService interface {
	// id: topic ID
	GetRepliesByTopicID(id int) ([]Reply, error)
	// 获取节点ID下所有主题
	GetTopicsByNodeID(nodeid int) ([]Topic, error)
	//获取节点下所有主题
	GetTopicsByNodename(name string) ([]Topic, error)
	// 相当于首页右侧的 10 大每天的内容
	GetHot() ([]Topic, error)
	// 最新的主题
	GetLatest() ([]Topic, error)
	GetTopicByID(id int) (*Topic, error)
	//获取指定用户的所有主题
	GetTopicsByUsername(name string) ([]Topic, error)
}
