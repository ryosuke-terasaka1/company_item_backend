package guest

import "item_management/usecases/interactors"

type GuestController struct {
	Interactor interactors.CompanyInteractor
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
