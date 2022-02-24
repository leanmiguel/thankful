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

// type Entry struct {
// 	UserId      string   `json:"user_id"`
// 	CreatedTime string   `json:"created_time"`
// 	Entries     []string `json:"entries"`
// }

// type Entries []Entry

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
		Region:   aws.String("us-west-2"),
		LogLevel: aws.LogLevel(aws.LogDebugWithEventStreamBody),
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

// result, err := db.GetItem(&dynamodb.GetItemInput{
// 	TableName: aws.String("thankful_entries"),
// 	Key: map[string]*dynamodb.AttributeValue{
// 		"user_id":      {S: aws.String("lean")},
// 		"created_time": {S: aws.String("2022-02-22T11:53:28Z")},
// 	},
// })

// if err != nil {
// 	log.Fatalf("Got error calling GetItem: %s", err)
// }

// if result.Item == nil {
// 	fmt.Println("bad news")
// }

// item := Entry{}

// fmt.Println(result.Item)
// err = dynamodbattribute.UnmarshalMap(result.Item, &item)

// if err != nil {
// 	panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
// }
