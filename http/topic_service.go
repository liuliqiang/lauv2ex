package http

import (
	"strconv"

	"fmt"

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

func (s *TopicService) TopicSummary(t *v2ex.Topic) string {
	return fmt.Sprintf("%6d %3d  %s", t.ID, t.Replies, t.Title)
}

func (s *TopicService) TopicDetail(topic *v2ex.Topic) string {
	str := fmt.Sprintf("Title: %s\n", topic.Title)
	str += fmt.Sprintf("Url: %s\n", topic.URL)
	str += fmt.Sprintf("Author: %s\n", topic.Member.Username)
	str += fmt.Sprintf("Content: %s\n", topic.Content)
	str += fmt.Sprintf("Replics: %d\n", topic.Replies)
	str += fmt.Sprintf("-----------\n")

	replies, err := s.GetRepliesByTopicID(topic.ID)
	if err != nil {
		return "error occur"
	}
	for idx, reply := range replies {
		str += fmt.Sprintf("\tId: %d,  ", idx)
		str += fmt.Sprintf("%s\n", s.ReplyDetail(&reply))
		str += fmt.Sprintln("\t-------------")
	}

	return str
}

func (s *TopicService) ReplyDetail(r *v2ex.Reply) string {
	str := fmt.Sprintf("\tUser: %s,  ", r.Member.Username)
	str += fmt.Sprintf("\tTks: %d\n", r.Thanks)
	str += fmt.Sprintf("\tContent: %s", r.Content)

	return str
}
