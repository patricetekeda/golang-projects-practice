package services


import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)


func SendToTopic(topicName string, messageBody []byte) error {
	// create a new Service Bus client
	client, err := azservicebus.NewClientFromConnectionString("your_connection_string", nil)
	if err != nil {
		return fmt.Errorf("failed to create Service Bus client: %w", err)
	}

	defer client.Close(context.Background())

	sender, err := client.NewSender(topicName, nil)
	if err != nil {
		return fmt.Errorf("failed to create sender: %w", err)
	}
	defer sender.Close(context.Background())

	// create a new message
	msg := &azservicebus.Message{
		Body: messageBody,
	}

	// send the message
	if err := sender.SendMessage(context.Background(), msg, nil); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	
	fmt.Println("Message sent successfully to topic:", topicName)
	return nil
}