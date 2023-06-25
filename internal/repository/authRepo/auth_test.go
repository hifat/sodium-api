package authRepo_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/hifat/sodium-api/internal/domain/authDomain"
	"github.com/hifat/sodium-api/internal/repository/authRepo"
	"github.com/hifat/sodium-api/internal/utils"
	"github.com/hifat/sodium-api/internal/utils/ernos"
	"github.com/hifat/sodium-api/internal/utils/gorm/utype"
	"github.com/hifat/sodium-api/internal/utils/utime"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type testAuthRepoSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock

	authRepo authDomain.AuthRepository
}

func TestAuthRepo(t *testing.T) {
	suite.Run(t, &testAuthRepoSuite{})
}

func (s *testAuthRepoSuite) SetupSuite() {
	dbMock, mock, err := sqlmock.New()
	s.Require().NoError(err)
	dialector := postgres.New(postgres.Config{
		Conn:       dbMock,
		DriverName: "postgres",
	})

	gormMock, err := gorm.Open(dialector, &gorm.Config{})
	s.Require().NoError(err)

	s.db = gormMock
	s.mock = mock
	s.authRepo = authRepo.NewAuthRepository(gormMock)
}

func (s *testAuthRepoSuite) SetupTest() {
	utime.Freeze()
}

func (s *testAuthRepoSuite) TearDownTest() {
	utime.UnFreeze()
}

func (s *testAuthRepoSuite) AfterTest(_, _ string) {
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (s *testAuthRepoSuite) TearDownSuite() {
	sql, err := s.db.DB()
	if err != nil {
		sql.Close()
	}
}

func (s *testAuthRepoSuite) TestAuthRepo_Register() {
	s.Run("success - register", func() {
		ctx := context.Background()
		password, err := utils.HashPassword("12345678")
		s.Require().NoError(err)

		req := authDomain.RequestRegister{
			Username: "sodium",
			Password: password,
			Name:     "Sodiumy",
		}

		s.mock.ExpectBegin()
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`INSERT INTO "users" ("username","password","name","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
			WithArgs(
				req.Username,
				req.Password,
				req.Name,
				utime.Now(),
				utime.Now(),
				nil,
			).WillReturnRows(&sqlmock.Rows{})
		s.mock.ExpectCommit()

		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT "users"."username","users"."name" FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"username", "name"}).
				AddRow(req.Username, req.Name))

		res := authDomain.ResponseRegister{}
		err = s.authRepo.Register(ctx, req, &res)
		s.Require().NoError(err)
		s.Require().NotEmpty(res)
	})
}

func (s *testAuthRepoSuite) TestAuthRepo_Login() {
	s.Run("success - login", func() {
		ctx := context.Background()

		hashedPassword, err := utils.HashPassword("12345678")
		s.Require().NoError(err)

		req := authDomain.RequestLogin{
			Username: "sodium",
			Password: "12345678",
		}

		var (
			name = "Sodiumy"
		)

		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT "id","username","password","name" FROM "users" WHERE username = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"username", "password", "name"}).
				AddRow(req.Username, hashedPassword, name))

		res := authDomain.ResponseRefreshTokenRepo{}
		err = s.authRepo.Login(ctx, req, &res)
		s.Require().NoError(err)
		s.Require().NotEmpty(res)

		s.Require().Equal(req.Username, res.Username)
		s.Require().Equal(name, res.Name)
	})

	s.Run("failed - user not found", func() {
		ctx := context.Background()

		req := authDomain.RequestLogin{
			Username: "sodium_invalid",
		}

		expectedError := errors.New(ernos.M.RECORD_NOTFOUND)

		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT "id","username","password","name" FROM "users" WHERE username = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
			WithArgs(sqlmock.AnyArg()).
			WillReturnError(expectedError)

		res := authDomain.ResponseRefreshTokenRepo{}
		err := s.authRepo.Login(ctx, req, &res)
		s.Require().Error(err)
		s.Require().Equal(err.Error(), ernos.M.RECORD_NOTFOUND)
		s.Require().Empty(res)
	})
}

func (s *testAuthRepoSuite) TestAuthRepo_Logout() {
	s.Run("success - logout", func() {
		ctx := context.Background()

		refreshTokenID := uuid.New()

		s.mock.ExpectBegin()
		s.mock.ExpectExec(
			regexp.QuoteMeta(`DELETE FROM "refresh_tokens" WHERE id = $1`)).
			WithArgs(refreshTokenID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		s.mock.ExpectCommit()

		err := s.authRepo.Logout(ctx, refreshTokenID)
		s.Require().NoError(err)
	})
}

func (s *testAuthRepoSuite) TestAuthRepo_CreateRefreshToken() {
	s.Run("success - create refresh token", func() {
		ctx := context.Background()

		req := authDomain.RequestCreateRefreshToken{
			ID:       uuid.New(),
			Token:    "token",
			Agent:    "web",
			ClientIP: utype.IP("192.168.1.1"),
			UserID:   uuid.New(),
		}

		var (
			isActive  = true
			createdAt = utime.Now()
			updatedAt = utime.Now()
		)

		s.mock.ExpectBegin()
		s.mock.ExpectQuery(
			regexp.QuoteMeta(`INSERT INTO "refresh_tokens" ("token","agent","client_ip","is_active","user_id","created_at","updated_at","id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "id"`)).
			WithArgs(
				req.Token,
				req.Agent,
				req.ClientIP,
				isActive,
				req.UserID,
				createdAt,
				updatedAt,
				req.ID,
			).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(req.ID))
		s.mock.ExpectCommit()

		s.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "refresh_tokens" WHERE "refresh_tokens"."id" = $1`)).
			WithArgs(req.ID).
			WillReturnRows(sqlmock.NewRows([]string{"token", "agent", "client_ip", "is_active", "user_id"}).
				AddRow(
					req.Token,
					req.Agent,
					req.ClientIP,
					isActive,
					req.UserID,
				))

		res, err := s.authRepo.CreateRefreshToken(ctx, req)
		s.Require().NoError(err)
		s.Require().NotEmpty(res)
	})
}
