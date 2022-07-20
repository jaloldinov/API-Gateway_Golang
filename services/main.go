package services

import (
	"api_gateway/config"
	"api_gateway/genproto/company_service"
	"api_gateway/genproto/position_service"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceManager interface {
	ProfessionService() position_service.ProfessionServiceClient
	AttributeService() position_service.AttributeServiceClient
	CompanyService() company_service.CompanyServiceClient
	PositionService() position_service.PositionServiceClient
}

type grpcClients struct {
	professionService position_service.ProfessionServiceClient
	attributeService  position_service.AttributeServiceClient
	companyService    company_service.CompanyServiceClient
	positionService   position_service.PositionServiceClient
}

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connProfessionService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PositionServiceHost, conf.PositionServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connAttributeService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PositionServiceHost, conf.PositionServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connCompanyService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PositionServiceHost, conf.PositionServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connPositionService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PositionServiceHost, conf.PositionServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		professionService: position_service.NewProfessionServiceClient(connProfessionService),
		attributeService:  position_service.NewAttributeServiceClient(connAttributeService),
		companyService:    company_service.NewCompanyServiceClient(connCompanyService),
		positionService:   position_service.NewPositionServiceClient(connPositionService),
	}, nil
}

func (g *grpcClients) ProfessionService() position_service.ProfessionServiceClient {
	return g.professionService
}

func (g *grpcClients) AttributeService() position_service.AttributeServiceClient {
	return g.attributeService
}

func (g *grpcClients) CompanyService() company_service.CompanyServiceClient {
	return g.companyService
}

func (g *grpcClients) PositionService() position_service.PositionServiceClient {
	return g.positionService
}
