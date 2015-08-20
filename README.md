# MQTT wrapper for Go

# Description
Wrapper for org.eclipse.paho.mqtt using simple On (subscribe) and Push ( publish) method with json decoding/encoding

# Installation
    go get github.com/MarinX/mqtt


# Example

    // example
    package main
    
    import (
    	"fmt"
    	"github.com/MarinX/mqtt"
    	"time"
    )
    func main() {
    	cl := mqtt.New("go-client", "tcp://localhost:1883")
    
    	if err := cl.Connect(); err != nil {
    		fmt.Println(err)
    		return
    	}
    
    	fmt.Println("Listening on test/topic")
    	cl.On("test/topic", func(c *mqtt.Context) {
    		fmt.Println(string(c.Payload))
    	})
    
    	fmt.Println("Sending payload")
    	time.Sleep(time.Second)
    
    	for i := 0; i < 5; i++ {
    		if err := cl.Push("test/topic", "HelloPayload"); err != nil {
    			fmt.Println(err)
    		}
    	}
    
    	time.Sleep(time.Second)
    
    	cl.Disconnect()
    }




# License
This library is under the MIT License
# Author
Marin Basic 
