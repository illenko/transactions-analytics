package integration

import (
	"context"
	"github.com/illenko/analytics-service/internal/database"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

// BaseIntegrationTestSuite represents the note repository test suite.
type BaseIntegrationTestSuite struct {
	suite.Suite
	Log                *slog.Logger
	Ctx                context.Context
	DB                 *gorm.DB
	pgContainer        *postgres.PostgresContainer
	pgConnectionString string
}

func (suite *BaseIntegrationTestSuite) SetupSuite() {
	suite.Ctx = context.Background()
	pgContainer, err := postgres.RunContainer(
		suite.Ctx,
		testcontainers.WithImage("postgres:16-alpine"),
		postgres.WithDatabase("analytics-service"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	suite.NoError(err)

	connStr, err := pgContainer.ConnectionString(suite.Ctx, "sslmode=disable")
	suite.NoError(err)

	db, err := gorm.Open(pg.Open(connStr), &gorm.Config{})
	suite.NoError(err)

	suite.pgContainer = pgContainer
	suite.pgConnectionString = connStr
	suite.DB = db

	sqlDB, err := suite.DB.DB()
	suite.NoError(err)

	err = sqlDB.Ping()
	suite.NoError(err)

	suite.Log = slog.Default()

	migration := database.NewMigration(suite.Log, db)

	err = migration.Execute("../../../migrations")
	suite.NoError(err)
}

func (suite *BaseIntegrationTestSuite) TearDownSuite() {
	err := suite.pgContainer.Terminate(suite.Ctx)
	suite.NoError(err)
}
