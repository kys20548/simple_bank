package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kys20548/simple_bank/db/sqlc"
	"os"
	"testing"
)

func newTestServer(t *testing.T, store db.Store) *Server {

	server := NewServer(store)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
