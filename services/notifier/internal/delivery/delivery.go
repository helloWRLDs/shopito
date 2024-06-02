package delivery

import (
	"context"
	"shopito/services/notifier/internal/service"
	"shopito/services/notifier/protobuf"
)

type Delivery struct {
	protobuf.UnimplementedNotifierServiceServer
	serv service.Service
}

func New(serv *service.NotifierService) *Delivery {
	return &Delivery{
		serv: serv,
	}
}

func (d *Delivery) SendEmail(ctx context.Context, request *protobuf.SendEmailRequest) (*protobuf.SendEmailResponse, error) {
	response := &protobuf.SendEmailResponse{Success: false}

	err := d.serv.SendEmailService(request.GetTo(), request.GetSubject(), request.GetBody())
	if err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (d *Delivery) SendAllEmail(ctx context.Context, request *protobuf.SendAllEmailRequest) (*protobuf.SendAllEmailResponse, error) {
	response := &protobuf.SendAllEmailResponse{Success: false}
	
	err := d.serv.SendAllEmailService(request.GetSubject(), request.GetBody())
	if err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}
