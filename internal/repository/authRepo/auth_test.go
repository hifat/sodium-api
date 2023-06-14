package authRepo_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hifat/sodium-api/internal/domain/authDomain"
	"github.com/hifat/sodium-api/internal/repository/authRepo"
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

func (u *testAuthRepoSuite) TestAuthRepo_Register() {
	u.Run("success - register", func() {
		var ctx = context.Background()

		req := authDomain.RequestRegister{
			Username: "sodium",
			Password: "12345678",
			Name:     "Sodiumy",
		}

		u.mock.ExpectBegin()
		u.mock.ExpectQuery(
			regexp.QuoteMeta(`INSERT INTO "users" ("username","password","name","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
			WithArgs(
				req.Username,
				req.Password,
				req.Name,
				utime.Now(),
				utime.Now(),
				nil,
			).WillReturnRows(&sqlmock.Rows{})
		u.mock.ExpectCommit()

		u.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT "users"."username","users"."name" FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT 1`)).
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"username", "name"}).AddRow("sodium", "Sodiumy"))

		res := authDomain.ResponseRegister{}
		err := u.authRepo.Register(ctx, req, &res)
		u.Require().NoError(err)
		u.Require().NotEmpty(res)
	})
}
