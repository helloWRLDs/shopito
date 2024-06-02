package delivery

import (
	"context"
	"shopito/services/notifier/internal/service"
	"shopito/pkg/protobuf/notifier"
)

type Delivery struct {
	notifierproto.UnimplementedNotifierServiceServer
	serv service.Service
}

func New(serv *service.NotifierService) *Delivery {
	return &Delivery{
		serv: serv,
	}
}

func (d *Delivery) SendEmail(ctx context.Context, request *notifierproto.SendEmailRequest) (*notifierproto.SendEmailResponse, error) {
	response := &notifierproto.SendEmailResponse{Success: false}

	err := d.serv.SendEmailService(request.GetTo(), request.GetSubject(), request.GetBody())
	if err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (d *Delivery) SendAllEmail(ctx context.Context, request *notifierproto.SendAllEmailRequest) (*notifierproto.SendAllEmailResponse, error) {
	response := &notifierproto.SendAllEmailResponse{Success: false}
	
	err := d.serv.SendAllEmailService(request.GetSubject(), request.GetBody())
	if err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}
