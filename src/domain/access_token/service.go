package access_token

import (
	"github.com/MrBonzzo/udm-bs-oauth-api/src/utils/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}
