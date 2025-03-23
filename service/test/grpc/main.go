package main

import (
	"app/generated/proto/enumgrpc"
	"app/generated/proto/servicegrpc"
	"app/generated/proto/sharedgrpc"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	action := flag.String("action", "", "Action to perform: create or get")
	flag.Parse()

	if *action == "" {
		fmt.Println("Usage: go run main.go -action=[create|get]")
		os.Exit(1)
	}

	conn, err := grpc.Dial("localhost:10004", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := servicegrpc.NewQuizzServiceClient(conn)

	switch *action {
	case "create":
		create(client)
	case "get-by-id":
		for i := 0; i < 10; i++ {
			loop(client)
		}
	case "get-by-entity-id":
		getList(client)
	default:
		fmt.Println("Invalid action. Use 'create' or 'get'.")
	}
}

func loop(client servicegrpc.QuizzServiceClient) {
	c := 0
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			err := getById(client)
			if err == nil {
				m.Lock()
				c += 1
				m.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	log.Println(c)
}

func create(client servicegrpc.QuizzServiceClient) {
	res, err := client.CreateQuizz(context.Background(), &servicegrpc.CreateQuizzRequest{
		Ask:        "",
		Time:       10,
		Result:     []string{"1", "2"},
		Option:     []string{"1", "2", "3"},
		EntityType: enumgrpc.EntityType_QUIZZ_VIDEO_LESSON,
		EntityId:   1,
	})

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res)
}

func getById(client servicegrpc.QuizzServiceClient) error {
	_, err := client.GetById(context.Background(), &sharedgrpc.ID{Id: 2})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func getList(client servicegrpc.QuizzServiceClient) {
	res, err := client.GetListByEntityId(context.Background(), &servicegrpc.GetListQuizzRequest{})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res.Total)
}
