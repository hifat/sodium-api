package userRepo_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/hifat/sodium-api/internal/domain/userDomain"
	"github.com/hifat/sodium-api/internal/repository/userRepo"
	"github.com/hifat/sodium-api/internal/utils/utime"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type testUserRepoSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock

	userRepo userDomain.UserRepository
}

func TestUserRepo(t *testing.T) {
	suite.Run(t, &testUserRepoSuite{})
}

func (s *testUserRepoSuite) SetupSuite() {
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
	s.userRepo = userRepo.NewUserRepository(gormMock)
}

func (s *testUserRepoSuite) SetupTest() {
	utime.Freeze()
}

func (s *testUserRepoSuite) TearDownTest() {
	utime.UnFreeze()
}

func (s *testUserRepoSuite) AfterTest(_, _ string) {
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (s *testUserRepoSuite) TearDownSuite() {
	sql, err := s.db.DB()
	if err != nil {
		sql.Close()
	}
}

func (u *testUserRepoSuite) TestUserRepo_CheckExists() {
	u.Run("success - check user exists", func() {
		ctx := context.Background()

		u.mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT COUNT(*) > 0 FROM "users" WHERE username = $1 AND "users"."deleted_at" IS NULL`)).
			WithArgs("sodium").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(true))

		exists, err := u.userRepo.CheckExists(ctx, "username", "sodium")
		u.Require().NoError(err)
		u.Require().True(exists)
	})

	u.Run("failed - col is not in columns", func() {
		ctx := context.Background()

		exists, err := u.userRepo.CheckExists(ctx, "no_col", "sodium")
		hasErrMessage := regexp.MustCompile("^col must be includes")
		u.Require().True(hasErrMessage.MatchString(err.Error()))
		u.Require().False(exists)
	})

	u.Run("failed - expectValue is empty", func() {
		ctx := context.Background()

		exists, err := u.userRepo.CheckExists(ctx, "username", "")
		u.Require().Error(err)
		u.Require().Equal(err.Error(), "expectValue must be required")
		u.Require().False(exists)
	})
}

func (u *testUserRepoSuite) TestUserRepo_GetFieldsByID() {
	u.Run("success - get filed by ID", func() {
		ctx := context.Background()
		username := "sodium"

		u.mock.ExpectQuery(regexp.QuoteMeta(`SELECT "username" FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL`)).
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"username"}).
				AddRow(username))

		value, err := u.userRepo.GetFieldsByID(ctx, uuid.New(), "username")
		u.Require().NoError(err)
		u.Require().Equal(value, username)
	})
}
