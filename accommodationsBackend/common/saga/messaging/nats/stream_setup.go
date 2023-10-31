package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
)

func CreateJStream(natsComponent *NATSComponent, streamName string, streamSubjects string) error {
	fmt.Printf("Nats component u CreateJStream: %v\n", natsComponent)
	var js nats.JetStreamContext
	js, err := natsComponent.JetStreamContext()
	if err != nil {
		log.Println("Greska pri dobavljanju konteksta u CreateJStream")
		log.Fatalln(err)
	}
	// Check if the stream already exists; if not, create it.
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Printf("ERROR ZA POSTOJI LI STREAM:  %v", err)
		log.Println(err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)
		si, err := js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		fmt.Println("Stream Info: %v", si)
		if err != nil {
			fmt.Println("Nije napravljen stream %v", streamName)
			return err
		} else {
			fmt.Println("Napravljen je stream %v", streamName)
		}
	} else {
		log.Println("Stream vec postoji")
	}
	return nil
}
