package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/nordew/go-errx"
	"github.com/nordew/warpbid/internal/dto"
	"github.com/nordew/warpbid/internal/models"
)

const (
	auctionsTable = "auctions"
)

var (
	ErrFailedToCreateAuction = "failed to create auction"
	ErrFailedToGetAuction    = "failed to get auction"
	ErrFailedToUpdateAuction = "failed to update auction"
)

func (s *Storage) CreateAuction(ctx context.Context, auction *models.Auction) error {
	qb := s.sb.Insert(auctionsTable).
		Columns(
			"id",
			"title",
			"description",
			"start_time",
			"end_time",
			"start_price",
			"current_price",
			"status",
			"created_at",
			"updated_at",
		).
		Values(
			auction.ID,
			auction.Title,
			auction.Description,
			auction.StartTime,
			auction.EndTime,
			auction.StartPrice,
			auction.CurrentPrice,
			auction.Status,
			auction.CreatedAt,
			auction.UpdatedAt,
		)

	query, args, err := qb.ToSql()
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToBuildQuery, err)
	}

	_, err = s.cockroach.ExecContext(ctx, query, args...)
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToCreateAuction, err)
	}

	return nil
}

func (s *Storage) GetAuctionByFilter(ctx context.Context, filter *dto.GetAuctionFilter) (*models.Auction, error) {
	qb := s.sb.Select("*").From(auctionsTable)
	qb = applyAuctionFilters(qb, filter)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause(ErrFailedToBuildQuery, err)
	}

	var auction models.Auction
	err = s.cockroach.GetContext(ctx, &auction, query, args...)
	if err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause(ErrFailedToGetAuction, err)
	}

	return &auction, nil
}

func applyAuctionFilters(qb sq.SelectBuilder, filter *dto.GetAuctionFilter) sq.SelectBuilder {
	if filter.ID != "" {
		qb = qb.Where("id = ?", filter.ID)
	}

	if filter.Title != "" {
		qb = qb.Where("title = ?", filter.Title)
	}

	if filter.Description != "" {
		qb = qb.Where("description = ?", filter.Description)
	}

	if !filter.StartTime.IsZero() {
		qb = qb.Where("start_time = ?", filter.StartTime)
	}

	if !filter.EndTime.IsZero() {
		qb = qb.Where("end_time = ?", filter.EndTime)
	}

	if filter.StartPrice != 0 {
		qb = qb.Where("start_price = ?", filter.StartPrice)
	}

	if filter.CurrentPrice != 0 {
		qb = qb.Where("current_price = ?", filter.CurrentPrice)
	}

	if filter.Status != "" {
		qb = qb.Where("status = ?", filter.Status)
	}

	return qb
}

func (s *Storage) UpdateAuction(ctx context.Context, input *dto.UpdateAuctionDTO) error {
	qb := s.sb.Update(auctionsTable)

	if input.Title != "" {
		qb = qb.Set("title", input.Title)
	}

	if input.Description != "" {
		qb = qb.Set("description", input.Description)
	}

	if !input.EndTime.IsZero() {
		qb = qb.Set("end_time", input.EndTime)
	}

	qb = qb.Where("id = ?", input.ID)

	query, args, err := qb.ToSql()
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToBuildQuery, err)
	}

	_, err = s.cockroach.ExecContext(ctx, query, args...)
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToUpdateAuction, err)
	}

	return nil
}
