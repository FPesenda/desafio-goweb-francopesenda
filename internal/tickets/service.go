package tickets

import (
	"context"
	"desafio-goweb-francopesenda/internal/domain"
)

type Service interface {
	GetTicketsByCountry(context.Context, string) ([]domain.Ticket, error)
	AverageDestination(context.Context, string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (serv *service) GetTicketsByCountry(ctx context.Context, country string) ([]domain.Ticket, error) {
	listTickets, err := serv.repository.GetTicketByDestination(ctx, country)
	if err != nil {
		return nil, err
	}
	return listTickets, nil
}

func (serv *service) AverageDestination(ctx context.Context, country string) (float64, error) {
	listAllTickets, errGetAll := serv.repository.GetAll(ctx)
	if errGetAll != nil {
		return 0.0, errGetAll
	}
	listTickets, errGetByCountry := serv.repository.GetTicketByDestination(ctx, country)
	if errGetByCountry != nil {
		return 0.0, errGetByCountry
	}
	return (float64(len(listTickets)) / float64(len(listAllTickets))) * 100, nil
}
