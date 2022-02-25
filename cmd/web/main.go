package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"leanmiguel/thankful/pkg/models/dynamo"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Config struct {
	Addr      string
	StaticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	entries  *dynamo.EntryModel
}

func main() {

	cfg := new(Config)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	awsCfg := &aws.Config{
		Region: aws.String("us-west-2"),
		// LogLevel: aws.LogLevel(aws.LogDebugWithHTTPBody),
	}
	mySession := session.Must(session.NewSession())

	db := dynamodb.New(mySession, awsCfg)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		entries:  &dynamo.EntryModel{DB: db},
	}

	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	infoLog.Printf("Starting server on %s", cfg.Addr)

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
