package grpc

import (
	"context"

	"github.com/bxcodec/go-clean-arch/bmi"
	"github.com/bxcodec/go-clean-arch/proto/proto_models"
	"google.golang.org/grpc"
)

type GrpcRoute struct {
	server *grpc.Server
}

func NewGRPCBmiRoute(server *grpc.Server) *GrpcRoute {
	return &GrpcRoute{
		server: server,
	}
}

func (g GrpcRoute) RegisterBmiHandler(handler proto_models.BMICalculatorServer) {
	proto_models.RegisterBMICalculatorServer(g.server, handler)
}

type BmiGrpcHandler struct {
	BmiService *bmi.Service
}

func NewBmiGrpcHandler(svc *bmi.Service) proto_models.BMICalculatorServer {
	return &BmiGrpcHandler{
		BmiService: svc,
	}
}

func (g *BmiGrpcHandler) CalculateBMI(ctx context.Context, req *proto_models.BMICalculateRequest) (*proto_models.BMICalculateResponse, error) {
	weight := req.GetWeight()
	height := req.GetHeight()

	bmi, err := g.BmiService.Calculate(context.Background(), weight, height)
	if err != nil {
		return &proto_models.BMICalculateResponse{}, err
	}
	return &proto_models.BMICalculateResponse{Bmi: bmi}, nil
}
