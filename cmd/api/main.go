package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Duane-Arzu/test1/internal/data"
	_ "github.com/lib/pq"
)

// func main() {
// 	var settings serverConfig
// 	flag.IntVar(&settings.port, "port", 4000, "Server Port")
// 	flag.StringVar(&settings.environment, "env", "development", "Environment(development|staging|production)")
// 	//read the dsn
// 	flag.StringVar(&settings.db.dsn, "db-dsn", "postgres://comments:comments@localhost/comments?sslmode=disable", "PostgreSQL DSN")
// 	flag.Parse()

// 	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

// 	//the call to openDB() sets up our connection pool
// 	db, err := openDB(settings)
// 	if err != nil {
// 		logger.Error(err.Error())
// 		os.Exit(1)
// 	}

// 	//release the database connection before exiting
// 	defer db.Close()

// 	logger.Info("Database Connection Pool Established")

// 	appInstance := &applicationDependencies{
// 		config:       settings,
// 		logger:       logger,
// 		commentModel: data.CommentModel{DB: db},
// 	}

// 	apiServer := &http.Server{
// 		Addr:         fmt.Sprintf(":%d", settings.port),
// 		Handler:      appInstance.routes(),
// 		IdleTimeout:  time.Minute,
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 10 * time.Second,
// 		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
// 	}

// 	logger.Info("Starting Server", "address", apiServer.Addr, "environment", settings.environment)
// 	err = apiServer.ListenAndServe()
// 	if err != nil {
// 		logger.Error(err.Error())
// 		os.Exit(1)
// 	}
// }

const appVersion = "7.0.0"

type serverConfig struct {
	port        int
	environment string
	db          struct {
		dsn string
	}
}

type applicationDependencies struct {
	config       serverConfig
	logger       *slog.Logger
	productModel data.ProductModel
	reviewModel  data.ReviewModel
}

func main() {
	var setting serverConfig

	flag.IntVar(&setting.port, "port", 4000, "Server port")
	flag.StringVar(&setting.environment, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&setting.db.dsn, "db-dsn", "postgres://products:darzu12@localhost/products?sslmode=disable", "PostgreSQL DSN")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// the call to openDB() sets up our connection pool
	db, err := openDB(setting)
	if err != nil {
		logger.Error("Database connection failed")
		os.Exit(1)
	}
	// release the database resources before exiting
	defer db.Close()

	logger.Info("Database connection pool established")

	appInstance := &applicationDependencies{
		config:       setting,
		logger:       logger,
		productModel: data.ProductModel{DB: db},
		reviewModel:  data.ReviewModel{DB: db},
	}

	apiServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.port),
		Handler:      appInstance.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("Starting server", "address", apiServer.Addr, "environment", setting.environment)
	err = apiServer.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

}

func openDB(settings serverConfig) (*sql.DB, error) {
	// open a connection pool
	db, err := sql.Open("postgres", settings.db.dsn)
	if err != nil {
		return nil, err
	}

	// set a context to ensure DB operations don't take too long
	ctx, cancel := context.WithTimeout(context.Background(),
		5*time.Second)
	defer cancel()

	// let's test if the connection pool was created
	// we trying pinging it with a 5-second timeout
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	// return the connection pool (sql.DB)
	return db, nil

}

//This is the command to push an existing repository from the command line
// git remote add origin https://github.com/Duane-Arzu/flower.git
// git branch -M main
// git push -u origin main

//hello
