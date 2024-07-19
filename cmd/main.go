package main

import (
	"fmt"
	"log"
	"service/api"
	"service/config"
	"service/pkg"
	"service/producer"
)

func main() {

	cfg, err := config.Load("..")
	if err != nil {
		log.Fatal(err)
	}

	mongosh, err := pkg.NewConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	pro, err := producer.DialProducer()
	if err != nil {
		log.Fatal(err)
	}

	r := api.NewGin(mongosh, *pro)
	
	addr := fmt.Sprintf(":%s", cfg.ServicePort)
	r.Run(addr)
}