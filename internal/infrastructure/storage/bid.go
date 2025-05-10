package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/nordew/go-errx"
	"github.com/nordew/warpbid/internal/dto"
	"github.com/nordew/warpbid/internal/models"
)

const (
	bidsTable = "bids"
)

var (
	ErrFailedToCreateBid       = "failed to create bid"
	ErrFailedToGetBid          = "failed to get bid"
	ErrFailedToCreateBidsBatch = "failed to create bids batch"
)

func (s *Storage) CreateBid(ctx context.Context, bid *models.Bid) error {
	qb := s.sb.Insert(bidsTable).
		Columns(
			"id",
			"auction_id",
			"user_id",
			"amount",
			"created_at",
		).
		Values(
			bid.ID,
			bid.AuctionID,
			bid.UserID,
			bid.Amount,
			bid.CreatedAt,
		)

	query, args, err := qb.ToSql()
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToBuildQuery, err)
	}

	_, err = s.cockroach.ExecContext(ctx, query, args...)
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToCreateBid, err)
	}

	return nil
}

func (s *Storage) CreateBidsBatch(ctx context.Context, bids []*models.Bid) error {
	if len(bids) == 0 {
		return nil
	}

	tx, err := s.cockroach.BeginTxx(ctx, nil)
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToCreateBidsBatch, err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	qb := sq.Insert(bidsTable).
		Columns(
			"id",
			"auction_id",
			"user_id",
			"amount",
			"created_at",
		)

	for _, bid := range bids {
		qb = qb.Values(
			bid.ID,
			bid.AuctionID,
			bid.UserID,
			bid.Amount,
			bid.CreatedAt,
		)
	}

	query, args, err := qb.RunWith(tx).ToSql()
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToBuildQuery, err)
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToCreateBidsBatch, err)
	}

	return nil
}

func (s *Storage) GetBidByFilter(ctx context.Context, filter *dto.GetBidFilter) (*models.Bid, error) {
	qb := s.sb.Select("*").From(bidsTable)
	qb = applyBidFilters(qb, filter)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause(ErrFailedToBuildQuery, err)
	}

	var bid models.Bid
	err = s.cockroach.GetContext(ctx, &bid, query, args...)
	if err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause(ErrFailedToGetBid, err)
	}

	return &bid, nil
}

func applyBidFilters(qb sq.SelectBuilder, filter *dto.GetBidFilter) sq.SelectBuilder {
	if filter.ID != "" {
		qb = qb.Where("id = ?", filter.ID)
	}
	if filter.AuctionID != "" {
		qb = qb.Where("auction_id = ?", filter.AuctionID)
	}
	if filter.UserID != "" {
		qb = qb.Where("user_id = ?", filter.UserID)
	}
	if filter.Amount != 0 {
		qb = qb.Where("amount = ?", filter.Amount)
	}

	return qb
}
