package refresh

import (
	"context"
	"fmt"

	"yirv2/auth/internal/repository"
	"yirv2/auth/internal/repository/entity"
	"yirv2/auth/internal/services/tokenizer"

	"github.com/google/uuid"
)

type Service interface {
	Refresh(ctx context.Context, token string) (string, string, error)
}

type service struct {
	dao           repository.DAO
	tokenaizerSrv tokenizer.Service
}

func New(
	dao repository.DAO,
	tokenaizerSrv tokenizer.Service,
) Service {
	return &service{
		dao:           dao,
		tokenaizerSrv: tokenaizerSrv,
	}
}

func (s *service) Refresh(ctx context.Context, token string) (string, string, error) {
	claims, err := s.tokenaizerSrv.ParseClaims(ctx, token)
	if err != nil {
		return "", "", fmt.Errorf("parse token: %w", err)
	}

	userIDStr, ok := claims["x-user_id"].(string)
	if !ok {
		return "", "", fmt.Errorf("fake token (w/o x-user_id)")
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return "", "", fmt.Errorf("parse token: %w", err)
	}

	userQuery := s.dao.NewUserQuery(ctx)

	userDB, err := userQuery.GetUserByPK(userID)
	if err != nil {
		return "", "", fmt.Errorf("get user: %w", err)
	}
	user := userDB.ToDomain()

	if token != *user.Token {
		return "", "", fmt.Errorf("fake refresh token") // TODO: починить 500тки, возвращать норм ошибки
	}

	// TODO: подумать над паттерном композит
	access, refresh, err := s.tokenaizerSrv.GeneratePair(
		ctx,
		map[string]any{"x-user_id": user.Id},
	)
	if err != nil {
		return "", "", fmt.Errorf("get tokens: %w", err)
	}

	user.Token = &refresh
	if _, err := userQuery.UpdateUser(entity.User{}.FromDomain(user)); err != nil {
		return "", "", fmt.Errorf("update user: %w", err)
	}

	return access, refresh, nil
}
