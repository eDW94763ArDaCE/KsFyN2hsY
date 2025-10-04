// 代码生成时间: 2025-10-04 21:48:53
package main

import (
	"context"
	"fmt"
	"log"
	"math"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	"pb "option_pricing_service/v1" // Assuming the proto definition is in option_pricing_service/v1
)

// OptionPricingServiceServer is the server API for OptionPricingService service.
type OptionPricingServiceServer struct {
	pb.UnimplementedOptionPricingServiceServer

	// Add your service logic here
}

// CalculateOptionPrice calculates the price of an option based on the Black-Scholes model.
func (s *OptionPricingServiceServer) CalculateOptionPrice(ctx context.Context, req *pb.CalculateOptionPriceRequest) (*pb.CalculateOptionPriceResponse, error) {
	// Validate the request parameters
	if req.GetS() <= 0 || req.GetK() <= 0 || req.GetT() <= 0 || req.GetR() <= 0 || req.GetSigma() <= 0 {
		return nil, fmt.Errorf("invalid parameter values")
	}

	// Calculate d1 and d2 using the Black-Scholes formula
	d1 := (math.Log(req.GetS()/req.GetK()) + (req.GetR() + req.GetSigma()*req.GetSigma()/2)*req.GetT()) / (req.GetSigma() * math.Sqrt(req.GetT()))
	d2 := d1 - req.GetSigma()*math.Sqrt(req.GetT())

	// Calculate the call and put option prices
	callOptionPrice := req.GetS() * math.Exp(-1*req.GetR()*req.GetT()) * math.NormCDF(d1) - req.GetK()*math.Exp(-1*req.GetR()*req.GetT())*math.NormCDF(d2)
	putOptionPrice := req.GetK()*math.Exp(-1*req.GetR()*req.GetT())*math.NormCDF(-d2) - req.GetS()*math.Exp(-1*req.GetR()*req.GetT())*math.NormCDF(-d1)

	// Return the calculated prices
	return &pb.CalculateOptionPriceResponse{
		CallOptionPrice:  callOptionPrice,
		PutOptionPrice:  putOptionPrice,
	}, nil
}

// Serve starts the gRPC server.
func Serve() error {
	lis, err := grpc.Listen(":50051", grpc.Insecure())
	if err != nil {
		return err
	}
	defer lis.Close()
	s := grpc.NewServer()
	pb.RegisterOptionPricingServiceServer(s, &OptionPricingServiceServer{})
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Serve(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
