package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/nordew/go-errx"
	"github.com/nordew/warpbid/internal/dto"
	"github.com/nordew/warpbid/internal/models"
)

const (
	usersTable = "users"
)

var (
	ErrFailedToBuildQuery = "failed to build query"
	ErrFailedToCreateUser = "failed to create user"
	ErrFailedToUpdateUser = "failed to update user"
)

func (s *Storage) CreateUser(ctx context.Context, user *models.User) error {
	qb := s.sb.Insert(usersTable).
		Columns(
			"id",
			"username",
			"wallet_address",
			"created_at",
		).
		Values(
			user.ID,
			user.Username,
			user.WalletAddress,
			user.CreatedAt,
		)

	query, args, err := qb.ToSql()
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToBuildQuery, err)
	}

	_, err = s.cockroach.ExecContext(ctx, query, args...)
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToCreateUser, err)
	}

	return nil
}

func (s *Storage) GetUserByFilter(ctx context.Context, filter *dto.GetUserFilter) (*models.User, error) {
	qb := s.sb.Select("*").From(usersTable)

	qb = applyUserFiletrs(qb, filter)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause(ErrFailedToBuildQuery, err)
	}

	var user models.User
	err = s.cockroach.GetContext(ctx, &user, query, args...)
	if err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause(ErrFailedToCreateUser, err)
	}

	return &user, nil
}

func applyUserFiletrs(qb sq.SelectBuilder, filter *dto.GetUserFilter) sq.SelectBuilder {
	if filter.ID != "" {
		qb = qb.Where("id = ?", filter.ID)
	}

	if filter.Username != "" {
		qb = qb.Where("username = ?", filter.Username)
	}

	if filter.WalletAddress != "" {
		qb = qb.Where("wallet_address = ?", filter.WalletAddress)
	}

	return qb
}

func (s *Storage) UpdateUser(ctx context.Context, input *dto.UpdateUserDTO) error {
	qb := s.sb.Update(usersTable)

	if input.Username != "" {
		qb = qb.Set("username", input.Username)
	}

	if input.WalletAddress != "" {
		qb = qb.Set("wallet_address", input.WalletAddress)
	}

	qb = qb.Where("id = ?", input.ID)

	query, args, err := qb.ToSql()
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToBuildQuery, err)
	}

	_, err = s.cockroach.ExecContext(ctx, query, args...)
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToUpdateUser, err)
	}

	return nil
}
