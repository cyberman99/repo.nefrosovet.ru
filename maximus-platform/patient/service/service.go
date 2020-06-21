package service

type Service interface {
	Appointment() // Appointment().Param() or Appointment().Program()
	Clinic()
	ConfirmCode()
	Employee()
	Implementation() // Implementation().Param()
	Invite()
	Patient()
}
