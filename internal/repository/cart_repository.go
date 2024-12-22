package repository

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/timothypattikawa/amole-services/cart-service/internal/repository/postgres"
)

type CartRepository interface {
	ExecTx(ctx context.Context, fn func(q *sqlc.Queries) error) error
	GetCartByActiveStatus(ctx context.Context, arg int32) (*sqlc.GetCartAndCartItemsByMemberIdAndActiveStatusRow, error)
	GetCartByMemberId(ctx context.Context, arg int32) (*sqlc.TbAmoleCart, error)
	InsertCart(ctx context.Context, arg sqlc.CreateCartParams) (*sqlc.TbAmoleCart, error)
}

type CartRepositoryImpl struct {
	db *pgxpool.Pool
	q  *sqlc.Queries
}

func NewCartRepository(db *pgxpool.Pool) CartRepository {
	return &CartRepositoryImpl{
		db: db,
		q:  sqlc.New(db),
	}
}

func (cr CartRepositoryImpl) ExecTx(ctx context.Context, fn func(q *sqlc.Queries) error) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tx, err := cr.db.Begin(ctx)
	if err != nil {
		return err
	}
	q := sqlc.New(cr.db).WithTx(tx)
	err = fn(q)

	if err != nil {
		log.Printf("error while query database %v", err)
		tx.Rollback(ctx)
	}

	tx.Commit(ctx)
	return nil
}

func (cr CartRepositoryImpl) GetCartByActiveStatus(ctx context.Context, arg int32) (*sqlc.GetCartAndCartItemsByMemberIdAndActiveStatusRow, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	result, err := cr.q.GetCartAndCartItemsByMemberIdAndActiveStatus(ctx, arg)
	if err != nil {
		log.Printf("error GetCartByActiveStatus err{%v}", err)
		return nil, err
	}

	return &result, err
}

func (cr CartRepositoryImpl) InsertCart(ctx context.Context, arg sqlc.CreateCartParams) (*sqlc.TbAmoleCart, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	result, err := cr.q.CreateCart(ctx, arg)
	if err != nil {
		log.Printf("error InsertCart err{%v}", err)
		return nil, err
	}

	return &result, err
}

func (cr CartRepositoryImpl) GetCartByMemberId(ctx context.Context, arg int32) (*sqlc.TbAmoleCart, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	result, err := cr.q.GetCartByMemberId(ctx, arg)
	if err != nil {
		log.Printf("error GetCartByMemberId err{%v}", err)
		return nil, err
	}

	return &result, err
}
