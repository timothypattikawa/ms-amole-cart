package service

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	rpc "github.com/timothypattikawa/amole-services/cart-service/api/grpc/client"
	pb "github.com/timothypattikawa/amole-services/cart-service/api/grpc/protos/product"
	"github.com/timothypattikawa/amole-services/cart-service/internal/dto"
	"github.com/timothypattikawa/amole-services/cart-service/internal/repository"
	"github.com/timothypattikawa/amole-services/cart-service/internal/repository/postgres"
	exception "github.com/timothypattikawa/amole-services/cart-service/pkg/errors"
	"github.com/timothypattikawa/amole-services/cart-service/pkg/utils"
)

type CartService interface {
	AddToCart(ctx context.Context, request dto.AddToCartRequest) (dto.AddToCartResponse, error)
}

type CartServiceImpl struct {
	cr    repository.CartRepository
	v     *viper.Viper
	pgrpc *rpc.ProductgClientgRPC
}

func NewCartService(cr repository.CartRepository,
	v *viper.Viper, db *pgxpool.Pool,
	prpc *rpc.ProductgClientgRPC) CartService {
	return &CartServiceImpl{
		cr:    cr,
		v:     v,
		pgrpc: prpc,
	}
}

func (cs CartServiceImpl) AddToCart(ctx context.Context, request dto.AddToCartRequest) (dto.AddToCartResponse, error) {
	var activeCart *postgres.TbAmoleCart
	var err error
	activeCart, err = cs.cr.GetCartByMemberId(ctx, int32(request.UserId))
	if err != nil {
		if err == sql.ErrNoRows {
			activeCart, err = cs.createNewCartForMember(ctx, int32(request.UserId))
			if err != nil {
				return dto.AddToCartResponse{}, exception.NewBusinessProcessError("Somtehing wen't wrong!", http.StatusInternalServerError)
			}
		} else {
			return dto.AddToCartResponse{}, exception.NewBusinessProcessError("Somtehing wen't wrong!", http.StatusInternalServerError)
		}
	}

	var response dto.AddToCartResponse

	err = cs.cr.ExecTx(ctx, func(q *postgres.Queries) error {
		var cartItems postgres.TbAmoleCartItem

		productInfo, err := cs.pgrpc.GetProductInfo(ctx, int64(request.ProductId))
		if err != nil {
			log.Printf("err while hit product grpc err{%v}", err.Error())
			return err
		}

		cartItems, err = q.GetCarItemsByCartIdAmdProductid(ctx, postgres.GetCarItemsByCartIdAmdProductidParams{
			TaciCartID:    int32(activeCart.TacID),
			TaciProductID: int32(request.ProductId),
		})

		if err != nil {
			if err == sql.ErrNoRows {
				cartItems, err = q.InsertCartItemsByCartId(ctx, postgres.InsertCartItemsByCartIdParams{
					TaciCartID:    int32(activeCart.TacID),
					TaciProductID: int32(request.ProductId),
					TaciQty:       int32(request.Qty),
					TaciPrice:     int32(request.Qty) * productInfo.TbapPrice,
				})
				if err != nil {
					log.Printf("error while save cart item err{%v}", err)
					return exception.NewBusinessProcessError("Somtehing wen't wrong!", http.StatusInternalServerError)
				}
				log.Printf("sucessful for save cart items to cart %v", cartItems)
				return nil
			} else {
				log.Printf("error while get cart item err{%v}", err)
				return exception.NewBusinessProcessError("Somtehing wen't wrong!", http.StatusInternalServerError)
			}
		}

		err = q.UpdateCartItemByCartId(ctx, postgres.UpdateCartItemByCartIdParams{
			TaciID:        cartItems.TaciID,
			TaciProductID: cartItems.TaciProductID,
			TaciQty:       int32(request.Qty),
			TaciPrice:     int32(request.Qty) * productInfo.TbapPrice,
		})

		if err != nil {
			log.Printf("error when updated cart item err{%v}", err)
			return exception.NewBusinessProcessError("Something wen't wrong!", http.StatusInternalServerError)
		}

		// Hit product service for add to cart stock
		resRPC, err := cs.pgrpc.TakeStockForATC(ctx, &pb.TakeStockForATCkRequest{
			Id:       int64(request.ProductId),
			QtyStock: int64(request.Qty),
		})

		if err != nil {
			log.Printf("fail to hit product for atc err{%v} data{%v}", err, request)
			return exception.NewBusinessProcessError("out of stock!!", http.StatusBadRequest)
		}

		response = dto.AddToCartResponse{
			SuccessTakeStock: true,
			Id:               resRPC.Id,
			ProductName:      resRPC.ProductName,
			Price:            resRPC.Price,
		}

		return nil
	})

	if err != nil {
		return dto.AddToCartResponse{}, nil
	}

	return response, nil
}

func (cs CartServiceImpl) createNewCartForMember(ctx context.Context, memberId int32) (*postgres.TbAmoleCart, error) {
	newCart, err := cs.cr.InsertCart(ctx, postgres.CreateCartParams{
		TacMemberID:   memberId,
		TacTotalPrice: 0,
		TacStatus:     utils.ActiveStatus,
	})

	if err != nil {
		return nil, err
	}

	return newCart, err
}
