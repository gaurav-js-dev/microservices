package service

import (
	"github.com/gaurav-js-dev/microservices/banking/domain"
	"github.com/gaurav-js-dev/microservices/banking/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// Get customer by id
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

// Get all customer
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
