package guest

type GuestController struct {
	Interactor usecases.CompanyInteractor
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
