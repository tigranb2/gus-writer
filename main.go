package main

import (
	"context"
	"fmt"
	"gus-writer/db"
	"log"
	"os"
	"paxos/randstring"
	"strconv"
)

type data struct {
	name string
}

func main() {
	arguments := os.Args
	if len(arguments) < 4 {
		fmt.Println("Please specify redis address, number of acknowledgements needed, and number of writes")
		return
	}
	redisAddr := arguments[1] //"localhost:6379"
	ackNeeded, err := strconv.Atoi(arguments[2])
	totalWrites, err := strconv.Atoi(arguments[3])
	if err != nil {
		log.Fatalf("input formatted incorrectly: %v", err)
	}

	d, err := db.NewDatabase(redisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	for i := 0; i < totalWrites; i++ {
		name := randstring.FixedLengthString(10) //randomly generate value to write, length of 10.
		d.Client.Set(context.TODO(), "name", data{name: name}, 0)
		d.Client.Wait(context.TODO(), ackNeeded, 0) //timeout = 0

		val, err := d.Client.Get(context.TODO(), "name").Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(val)
	}
}
