package response

import (
	"github.com/EvanKA/alterra-agmc/day6/employee-service/pkg/dto"
)

type Meta struct {
	Success bool                `json:"success" default:"true"`
	Message string              `json:"message" default:"true"`
	Info    *dto.PaginationInfo `json:"info"`
}
