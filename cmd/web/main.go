package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/douglasmg7/snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// username:password@protocol(address)/dbname?param=value
	// dsn := flag.String("dsn", "web:pass@tcp(localhost:3306)/snippetbox?parseTime=true", "MySQL data source name")
	// dsn := flag.String("dsn", "web:pass@tcp(host.docker.internal:3306)/snippetbox?parseTime=true", "MySQL data source name")
	// dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")	// Working out of docker.
	dsn := flag.String("dsn", "web:pass@tcp(mysql:3306)/snippetbox?parseTime=true", "MySQL data source name") // Using docker.
	flag.Parse()

	// Use the slog.New() function to initialize a new structured logger, which
	// writes to the standard out stream and uses the default settings.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create data base pool.
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Initialize a new instance of our application struct, containing the
	// dependencies (for now, just the structured logger).
	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	app.logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
