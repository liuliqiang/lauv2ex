package http

import (
	"github.com/yetship/lauv2ex"
)

var _ v2ex.SiteService = &RawV2exService{}

type RawV2exService struct {
	//获取社区统计信息
	ApiUrl string
}

func NewRawV2exService(apiUrl string) *RawV2exService {
	return &RawV2exService{
		ApiUrl: apiUrl,
	}
}

func (s *RawV2exService) GetStats() (*v2ex.Stats, error) {
	url := s.ApiUrl + "/site/stats.json"

	stat := &v2ex.Stats{}
	err := v2ex.GetAPIData(url, stat)
	if err != nil {
		return nil, err
	}
	return stat, err
}

func (s *RawV2exService) GetInfo() (*v2ex.Info, error) {
	url := s.ApiUrl + "/site/info.json"

	info := &v2ex.Info{}
	err := v2ex.GetAPIData(url, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
