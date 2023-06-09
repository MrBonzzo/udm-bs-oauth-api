package access_token

import (
	"strings"

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
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

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
