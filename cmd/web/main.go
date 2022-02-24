package main

import (
	"log"
	"net/http"
)

type Item struct {
	UserId      string   `json:"user_id"`
	CreatedTime string   `json:"created_time"`
	Entries     []string `json:"entries"`
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/today", serveTodayScreen)
	mux.HandleFunc("/api/days", home)

	// awsCfg := &aws.Config{
	// 	Region:   aws.String("us-west-2"),
	// 	LogLevel: aws.LogLevel(aws.LogDebugWithEventStreamBody),
	// }
	// mySession := session.Must(session.NewSession())
	// svc := dynamodb.New(mySession, awsCfg)

	// result, err := svc.GetItem(&dynamodb.GetItemInput{
	// 	TableName: aws.String("thankful_entries"),
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"user_id":      {S: aws.String("lean")},
	// 		"created_time": {S: aws.String("2022/02/21 19:23:37")},
	// 	},
	// })

	// if err != nil {
	// 	log.Fatalf("Got error calling GetItem: %s", err)
	// }

	// if result.Item == nil {
	// 	fmt.Println("bad news")
	// }

	// item := Item{}

	// fmt.Println(result.Item)
	// err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	// if err != nil {
	// 	panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	// }

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

	// fmt.Println(time.Now())
}
