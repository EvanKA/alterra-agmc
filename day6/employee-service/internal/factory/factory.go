package factory

import (
	"github.com/EvanKA/alterra-agmc/day6/employee-service/database"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/repository"
)

type Factory struct {
	EmployeeRepository repository.Employee
	DivisionRepository repository.Division
	RoleRepository     repository.Role
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		repository.NewEmployeeRepository(db),
		repository.NewDivisionRepository(db),
		repository.NewRoleRepository(db),
	}
}
