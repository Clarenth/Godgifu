package handlers

import (
	"godgifu/modules/account/services"
)

type handler struct {
	AccountService services.AccountService
}

func NewAccountHandlers(service services.AccountService) AccountHandlers {
	return &handler{
		AccountService: service,
	}
}
