package main

import (
	"context"
	microclient "github.com/micro/go-micro/client"
	proto "go-rpc-demo-client/proto"
	"log"
)

// go run main.go
func main() {

	animeService := proto.NewAnimeServiceClient("AnimeService", microclient.DefaultClient)

	res, err := animeService.QueryAnimeInfo(context.TODO(), &proto.QueryAnimeInfoRequest{AnimeId: 1})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(res.EnName)
}
