package http

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/yetship/lauv2ex"
)

var _ v2ex.TopicService = &TopicService{}

type TopicService struct {
	ApiUrl string
}

func NewTopicService(apiUrl string) *TopicService {
	return &TopicService{
		ApiUrl: apiUrl,
	}
}

func (s *TopicService) GetTopicsByNodeID(nodeid int) ([]v2ex.Topic, error) {
	url := s.ApiUrl + "/topics/show.json?node_id=" + strconv.Itoa(int(nodeid))
	var topics []v2ex.Topic

	err := v2ex.GetAPIData(url, &topics)
	return topics, err
}

func (s *TopicService) GetTopicsByNodename(name string) ([]v2ex.Topic, error) {
	if name == "" {
		return nil, errors.New("Invalid node name")
	}

	url := s.ApiUrl + "/topics/show.json?node_name=" + name
	var topics []v2ex.Topic

	err := v2ex.GetAPIData(url, &topics)

	return topics, err
}

func (s *TopicService) GetTopicByID(id int) (*v2ex.Topic, error) {
	var topics []*v2ex.Topic
	url := s.ApiUrl + "/topics/show.json?id=" + strconv.Itoa(int(id))

	err := v2ex.GetAPIData(url, &topics)

	return topics[0], err
}

func (s *TopicService) GetLatest() ([]v2ex.Topic, error) {
	url := s.ApiUrl + "/topics/latest.json"
	var topics []v2ex.Topic

	err := v2ex.GetAPIData(url, &topics)

	return topics, err
}

func (s *TopicService) GetHot() ([]v2ex.Topic, error) {
	url := s.ApiUrl + "/topics/hot.json"

	var topics []v2ex.Topic
	err := v2ex.GetAPIData(url, &topics)

	return topics, err
}

func (s *TopicService) GetTopicsByUsername(name string) ([]v2ex.Topic, error) {
	if name == "" {
		return nil, errors.New("Invalid username ")
	}

	url := s.ApiUrl + "/topics/show.json?username=" + name
	var topics []v2ex.Topic

	err := v2ex.GetAPIData(url, &topics)
	return topics, err
}

func (s *TopicService) GetRepliesByTopicID(id int) ([]v2ex.Reply, error) {
	url := s.ApiUrl + "/replies/show.json?topic_id=" + strconv.Itoa(int(id))
	var replies []v2ex.Reply

	err := v2ex.GetAPIData(url, &replies)

	return replies, err
}
