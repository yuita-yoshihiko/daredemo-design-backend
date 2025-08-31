package testutils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/yuita-yoshihiko/daredemo-design-backend/config"
)

func rootDir() string {
	currentPath := ""
	func() {
		_, currentPath, _, _ = runtime.Caller(0)
	}()

	return filepath.Join(currentPath, "../../")
}

func EnvLoad() {
	err := godotenv.Overload(fmt.Sprintf("%s/.env-test", rootDir()))
	if err != nil {
		panic(err)
	}

	if err := env.Parse(&config.Conf); err != nil {
		panic(err)
	}
}

func LoadTestDB() *sqlx.DB {
	if err := godotenv.Load(filepath.Join(rootDir(), ".env-test")); err != nil {
		panic(err)
	}
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err.Error())
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Could not connect to database")
		panic(err.Error())
	}
	return db
}

func LoadFixture(t *testing.T, fixtureDir string) *sqlx.DB {
	db := LoadTestDB()
	fixture, err := testfixtures.New(
		testfixtures.Database(db.DB),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(fixtureDir),
	)
	if err != nil {
		panic(err)
	}
	if err := fixture.Load(); err != nil {
		panic(err)
	}
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Log("failed to close database connection: %w", err)
		}
	})
	return db
}
