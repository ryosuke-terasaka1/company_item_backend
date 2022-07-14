package guest

import (
	"item_management/server/usecases/companies"
)

type GuestController struct {
	Interactor companies.CompanyInteractor
}

func InitGuestController(sqlHandler gateway.SQLHandler) *GuestController {
	return *GuestController{
		Interactor: usecase.GuestINteractor{
			GuestPort: &gateway.GuestRepository{
				sqlHandler: sqlHandler,
			},
		},
	}
}
