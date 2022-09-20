package auth

import (
	"context"
	"strings"
	"testing"

	"github.com/EvanKA/alterra-agmc/day6/employee-service/database"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/database/seeder"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/dto"
	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/factory"
	"github.com/stretchr/testify/assert"
)

func TestAuthServiceLoginByEmailAndPasswordSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		payload     = dto.ByEmailAndPasswordRequest{
			Email:    "e0337884@u.nus.edu",
			Password: "123abcABC!",
		}
	)
	res, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.Equal(payload.Email, res.Email)
	asserts.Len(strings.Split(res.JWT, "."), 3)
}

func TestAuthServiceLoginByEmailAndPasswordRecordNotFound(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	var (
		asserts     = assert.New(t)
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		payload     = dto.ByEmailAndPasswordRequest{
			Email:    "lpkr@gmail.com",
			Password: "123abcABC!",
		}
	)
	_, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 404")
	}
}

func TestAuthServiceLoginByEmailAndPasswordunmatchedEmailAndPassword(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		payload     = dto.ByEmailAndPasswordRequest{
			Email:    "e0337884@u.nus.edu",
			Password: "1234567890",
		}
	)
	_, err := authService.LoginByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 400")
	}
}

func TestAuthServiceRegisterByEmailAndPasswordSuccess(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		divisionID  = uint(1)
		payload     = dto.RegisterEmployeeRequestBody{
			Fullname:   "Lippo Village",
			Email:      "lv@gmail.com",
			Password:   "123abcABC!",
			DivisionID: &divisionID,
		}
	)
	payload.FillDefaults()
	res, err := authService.RegisterByEmailAndPassword(ctx, &payload)
	if err != nil {
		t.Fatal(err)
	}
	asserts.NotEmpty(res.ID)
	asserts.Equal(payload.Fullname, res.Fullname)
	asserts.Equal(payload.Email, res.Email)
	asserts.Len(strings.Split(res.JWT, "."), 3)
}

func TestAuthServiceRegisterByEmailAndPasswordUserExist(t *testing.T) {
	database.GetConnection()
	seeder.NewSeeder().DeleteAll()
	seeder.NewSeeder().SeedAll()
	asserts := assert.New(t)
	var (
		authService = NewService(factory.NewFactory())
		ctx         = context.Background()
		divisionID  = uint(1)
		payload     = dto.RegisterEmployeeRequestBody{
			Fullname:   "Evan Kurnia Alim",
			Email:      "e0337884@u.nus.edu",
			Password:   "123abcABC!",
			DivisionID: &divisionID,
		}
	)
	_, err := authService.RegisterByEmailAndPassword(ctx, &payload)
	if asserts.Error(err) {
		asserts.Equal(err.Error(), "error code 409")
	}
}
