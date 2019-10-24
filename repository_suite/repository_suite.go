package repository_suite

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-pg/pg/v9"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	DB *pg.DB
}

// SetupSuite setup at the beginning of test
func (s *RepositoryTestSuite) SetupSuite() {
	var err error

	s.DB = pg.Connect(
		&pg.Options{
			User:     "postgres",
			Password: "postgres",
			Network:  "tcp",
			Addr:     "0.0.0.0:5555",
			Database: "bara",
		},
	)
	_, err = s.DB.Exec("SELECT 1")
	require.NoError(s.T(), err)

	dir, err := os.Getwd()
	require.NoError(s.T(), err)

	file, err := ioutil.ReadFile(fmt.Sprintf(`%s/../../db/v1.sql`, dir))

	require.NoError(s.T(), err)

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		if request != "" {
			_, err = s.DB.Exec(request)
			require.NoError(s.T(), err)
		}
	}

}

// ClearDatabase ...
func (s *RepositoryTestSuite) ClearDatabase() {
	var tableNames []string
	_, err := s.DB.Query(&tableNames, "SELECT table_name FROM information_schema.tables WHERE table_schema='public'")

	require.NoError(s.T(), err)
	for _, table := range tableNames {
		s.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table))
	}
}
