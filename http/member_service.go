package http

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/yetship/lauv2ex"
)

var _ v2ex.MemberService = &MemberService{}

type MemberService struct {
	ApiUrl string
}

func NewMemberService(apiUrl string) *MemberService {
	return &MemberService{
		ApiUrl: apiUrl,
	}
}

func (s *MemberService) GetMemberByID(id int) (*v2ex.Member, error) {
	url := s.ApiUrl + "/members/show.json?id=" + strconv.Itoa(int(id))
	var member v2ex.Member

	err := v2ex.GetAPIData(url, &member)

	return &member, err
}

func (s *MemberService) GetMemberByUsername(name string) (*v2ex.Member, error) {
	if name == "" {
		return nil, errors.New("Invalid username")
	}

	url := s.ApiUrl + "/members/show.json?username=" + name
	var member v2ex.Member

	err := v2ex.GetAPIData(url, &member)

	return &member, err
}
