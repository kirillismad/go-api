package gin_test

import (
	"bytes"
	"encoding/json"
	"go-api/internal/gin"
	"go-api/pkg"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	g "github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type GinTestSuite struct {
	suite.Suite
	config gin.Config
	db     *gorm.DB
	router *g.Engine
}

func (s *GinTestSuite) SetupSuite() {
	driver := pkg.Must(sqlite.WithInstance(pkg.Must(s.db.DB()), &sqlite.Config{}))
	migrator := pkg.Must(migrate.NewWithDatabaseInstance("file://../../migrations", "main", driver))
	if err := migrator.Up(); err != nil {
		panic(err)
	}

}

func (s *GinTestSuite) TearDownSuite() {
	os.Remove("../../gorm_test.db")
}

func (s *GinTestSuite) TearDownTest() {
	s.db.Exec("DELETE FROM entities;")
}

func (s *GinTestSuite) TearDownSubTest() {
	s.db.Exec("DELETE FROM entities;")
}

// tests

func (s *GinTestSuite) TestPingOk() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	s.router.ServeHTTP(w, req)

	s.Equal(200, w.Code)
}

func (s *GinTestSuite) TestCreateEntityOk() {
	body := map[string]interface{}{
		"uuid_field":     "6ac99ab2-76fe-43fc-8c6a-13a06e0609b6",
		"int_field":      1,
		"float_field":    1.1,
		"datetime_field": "2023-12-01T09:59:48.839Z",
		"string_field":   "string",
		"bool_field":     true,
	}

	req := pkg.Must(http.NewRequest("POST", "/entities", bytes.NewReader(pkg.Must(json.Marshal(body)))))
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	s.Equal(201, w.Code)
}

func (s *GinTestSuite) TestCreateEntityBadRequestInvalidInput() {

}

func (s *GinTestSuite) TestCreateEntityBadRequestUnique() {

}

func TestGin(t *testing.T) {
	config := pkg.Must(gin.GetConfig())
	db := pkg.Must(gin.BuildDb(config))
	router := pkg.Must(gin.Build(config, pkg.Must(gin.BuildLogger(config)), db))

	suite.Run(t, &GinTestSuite{config: config, db: db, router: router})
}
