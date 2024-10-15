package utils

import (
	"database/sql"
	"fmt"
	"os"
	"path"
  
  "portal/sql_bindings"

	"github.com/tursodatabase/go-libsql"
)

func ConnectToTurso(cfg *Config) (*sql_bindings.Queries, *libsql.Connector, *sql.DB) {
  connector, err := libsql.NewEmbeddedReplicaConnector(path.Join("./", cfg.DatabaseFileName), cfg.TursoAddress, libsql.WithAuthToken(cfg.TursoToken))
  if err != nil {
    fmt.Println("Error creating connector:", err.Error())
    os.Exit(1)
  }
  db := sql.OpenDB(connector)
  executor := sql_bindings.New(db)
  return executor, connector, db
}
