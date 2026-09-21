package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// define an application struct to hold the application-wide dependencies for the web application
type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//creating a structured LOGGER writes the output in the TERMINAL that writes it in default settingbrew install mysql
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//define a new command-line flag for the MySQL DSN string
	dsn := flag.String("dsn", "web:laaq2003@/snippetbox?parseTime=true", "MySQL data source name")
	db, err := openDB(*dsn) //calling the sql.open()method from the function openDB()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	//we also defer a call to db.close(), so that the connection pool is closed before the main() function exits.
	defer db.Close()

	app := &application{
		logger: logger,
	}

	//use the INFO() method to log the starting server message
	logger.Info("starting server", "addr", *addr)

	//listening for changes on the SERVER
	err = http.ListenAndServe(*addr, app.routes())

	//use the ERROR() method to log any error message returned by the server and use the OS.EXIT() to terminate the application
	logger.Error(err.Error())
	os.Exit(1)
}

// the openDB() function wraps sql.open() and return a sql.DB connection pool for a given DSN
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

	return nil, err
}
