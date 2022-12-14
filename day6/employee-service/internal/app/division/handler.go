package division

import (
	"net/http"

	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/dto"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/factory"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/pkg/enum"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/pkg/util"
	pkgdto "github.com/EvanKA/alterra-agmc/day6/employee-service/pkg/dto"
	res "github.com/EvanKA/alterra-agmc/day6/employee-service/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Get(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	_, err := util.ParseJWTToken(authHeader)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}

	payload := new(pkgdto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Find(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(http.StatusOK, result.Data, "Get divisions success", &result.PaginationInfo).Send(c)
}

func (h *handler) GetById(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	_, err := util.ParseJWTToken(authHeader)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}

	payload := new(pkgdto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.FindByID(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) UpdateById(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := util.ParseJWTToken(authHeader)
	if (err != nil) || (jwtClaims.RoleID != uint(enum.Admin)) {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}

	payload := new(dto.UpdateDivisionRequestBody)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	result, err := h.service.UpdateById(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) DeleteById(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := util.ParseJWTToken(authHeader)
	if (err != nil) || (jwtClaims.RoleID != uint(enum.Admin)) {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}

	payload := new(pkgdto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	result, err := h.service.DeleteById(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) Create(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := util.ParseJWTToken(authHeader)
	if (err != nil) || (jwtClaims.RoleID != uint(enum.Admin)) {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}

	payload := new(dto.CreateDivisionRequestBody)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	division, err := h.service.Store(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(division).Send(c)
}
