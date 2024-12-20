package global

import (
	"database/sql"

	"github.com/newit-hieutm/go-backend/configs"
	"github.com/newit-hieutm/go-backend/pkg/sqlc/godev"
	"go.uber.org/zap"
)

var (
	Logs         *zap.Logger
	ConfigGlobal configs.Config
	Dbtx         *sql.DB
	Queries      *godev.Queries
)
