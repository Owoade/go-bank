package service

import (
	"github.com/Owoade/go-bank/config"
	"github.com/Owoade/go-bank/sql"
)

type Service struct {
	Store      *sql.SQLStore
	ConfigVars *config.Variables
}
