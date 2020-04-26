package service

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fsantiag/track-progress/src/database"
	"github.com/fsantiag/track-progress/src/repository"
)

func ProcessTaskMessage(channel <-chan *sqs.Message) {
	for {
		message := <-channel

		fmt.Println("Got this message", message)
		session, err := database.NewSession()
		if err != nil {
			fmt.Println(err.Error())
		}
		process(channel, session)
	}
}

func process(channel <-chan *sqs.Message, session repository.SessionInterface) {
	//TODO save task to database
}
