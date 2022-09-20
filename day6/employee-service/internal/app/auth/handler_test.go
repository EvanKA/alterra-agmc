package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/EvanKA/alterra-agmc/day6/employee-service/database"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/database/seeder"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/dto"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/factory"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/mocks"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandlerLoginByEmailAndPasswordInvalidPayload(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", nil)
	c.SetPath("/api/v1/auth/login")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.JSONEq(`{"meta": {"success": false,"message": "Invalid parameters or payload","info": null},"error": "bad_request"}`, body)
	}
}

func TestAuthHandlerLoginByEmailAndPasswordUnmatchedEmailAndPassword(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	emailAndPassword := dto.ByEmailAndPasswordRequest{
		Email:    "e0337884@u.nus.edu",
		Password: "1234567890",
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/api/v1/auth/login")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.JSONEq(`{"meta": {"success": false,"message": "Email or password is incorrect","info": null},"error": "bad_request"}`, body)
	}
}

func TestAuthHandlerLoginByEmailAndPasswordSuccess(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	emailAndPassword := dto.ByEmailAndPasswordRequest{
		Email:    "e0337884@u.nus.edu",
		Password: "123abcABC!",
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/api/v1/auth/login")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.LoginByEmailAndPassword(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "jwt")
	}
}

func TestAuthHandlerRegisterByEmailAndPasswordUserAlreadyExist(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	var (
		fullname   = "Evan Kurnia Alim"
		email      = "e0337884@u.nus.edu"
		password   = "123abcABC!"
		divisionID = uint(1)
	)
	emailAndPassword := dto.RegisterEmployeeRequestBody{
		Fullname:   fullname,
		Email:      email,
		Password:   password,
		DivisionID: &divisionID,
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/api/v1/auth/signup")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.RegisterByEmailAndPassword(c)) {
		asserts.Equal(409, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Created value already exists")
	}
}

func TestAuthHandlerRegisterByEmailAndPasswordInvalidPayload(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	var (
		fullname   = "Evan Kurnia Alim"
		password   = "123abcABC!"
		divisionID = uint(3)
	)
	emailAndPassword := dto.RegisterEmployeeRequestBody{
		Fullname:   fullname,
		Password:   password,
		DivisionID: &divisionID,
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/api/v1/auth/signup")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.RegisterByEmailAndPassword(c)) {
		asserts.Equal(400, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "Invalid parameters or payload")
	}
}

func TestAuthHandlerRegisterByEmailAndPasswordSuccess(t *testing.T) {
	// setup database
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()

	// setup context
	var (
		fullname   = "Lippo Karawaci"
		email      = "lk@gmail.com"
		password   = "123abcABC!"
		divisionID = uint(2)
	)
	emailAndPassword := dto.RegisterEmployeeRequestBody{
		Fullname:   fullname,
		Email:      email,
		Password:   password,
		DivisionID: &divisionID,
	}
	e := echo.New()
	echoMock := mocks.EchoMock{E: e}
	payload, err := json.Marshal(emailAndPassword)
	if err != nil {
		t.Fatal(err)
	}
	c, rec := echoMock.RequestMock(http.MethodPost, "/", bytes.NewBuffer(payload))
	c.Request().Header.Set("Content-Type", "application/json")
	c.SetPath("/api/v1/auth/signup")

	// setup handler
	asserts := assert.New(t)
	db := database.GetConnection()
	factory := factory.Factory{EmployeeRepository: repository.NewEmployeeRepository(db)}
	authHandler := NewHandler(&factory)

	// testing
	if asserts.NoError(authHandler.RegisterByEmailAndPassword(c)) {
		asserts.Equal(200, rec.Code)

		body := rec.Body.String()
		asserts.Contains(body, "id")
		asserts.Contains(body, "fullname")
		asserts.Contains(body, "email")
		asserts.Contains(body, "jwt")
	}
}
