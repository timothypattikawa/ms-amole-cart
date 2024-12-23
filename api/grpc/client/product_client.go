package client

import (
	"context"
	"time"

	"github.com/spf13/viper"
	pb "github.com/timothypattikawa/amole-services/cart-service/api/grpc/protos/product"
	"google.golang.org/grpc"
)

type ProductgClientgRPC struct {
	v      *viper.Viper
	client pb.ProductStockClient
}

func NewProductClientgRPC(v *viper.Viper, conn *grpc.ClientConn) *ProductgClientgRPC {

	client := pb.NewProductStockClient(conn)

	return &ProductgClientgRPC{
		v:      v,
		client: client,
	}
}

func (p *ProductgClientgRPC) GetProductInfo(ctx context.Context, productId int64) (*pb.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	arg := pb.ProductRequest{
		TbapID: int64(productId),
	}
	resp, err := p.client.ProductInfo(ctx, &arg)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ProductgClientgRPC) TakeStockForATC(ctx context.Context, in *pb.TakeStockForATCkRequest) (*pb.TakeStockForATCResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := p.client.TakeStockForATC(ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
