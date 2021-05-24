package dto

type AccessTokenDTO struct {
	ErrCode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type UserInfoDTO struct {
	ErrCode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	UserId  string `json:"UserId"`
}

type Reuslter interface {
	GetErrCode() int
	GetErrmsg() string
}

func (s *AccessTokenDTO) GetErrCode() int {
	return s.ErrCode
}

func (s *AccessTokenDTO) GetErrmsg() string {
	return s.Errmsg
}

func (s *UserInfoDTO) GetErrCode() int {
	return s.ErrCode
}

func (s *UserInfoDTO) GetErrmsg() string {
	return s.Errmsg
}
