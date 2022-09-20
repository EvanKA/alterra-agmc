package division

import (
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/dto"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/middleware"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/pkg/util"
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.Use(middleware.JWTMiddleware(dto.JWTClaims{}, util.JWT_SECRET))
	g.GET("", h.Get)
	g.GET("/:id", h.GetById)
	g.PUT("/:id", h.UpdateById)
	g.DELETE("/:id", h.DeleteById)
	g.POST("", h.Create)
}
