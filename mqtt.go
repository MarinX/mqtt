// mqtt
package mqtt

import (
	"encoding/json"
	"errors"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

type Conn struct {
	opts   *MQTT.ClientOptions
	client *MQTT.Client
}

type Context struct {
	Payload []byte
	Error   error
}

type Handler func(c *Context)

func New(clientid string, conn string) *Conn {
	opts := MQTT.NewClientOptions().AddBroker(conn).SetClientID(clientid)
	return &Conn{
		opts: opts,
	}
}

func (t *Conn) Connect() error {
	t.client = MQTT.NewClient(t.opts)

	if token := t.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (t *Conn) Disconnect() {
	t.client.Disconnect(500)
}

func (t *Conn) On(topic string, h Handler) error {

	if token := t.client.Subscribe(topic, 0, func(client *MQTT.Client, m MQTT.Message) {
		ctx := new(Context)

		if !client.IsConnected() {
			ctx.Error = errors.New("Client not connected")
			h(ctx)
			return
		}
		client.Lock()
		ctx.Payload = m.Payload()
		h(ctx)

		client.Unlock()

	}); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (t *Conn) Push(topic string, payload interface{}) error {
	buff, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	token := t.client.Publish(topic, 0, false, buff)
	token.Wait()
	return token.Error()
}

func (t *Context) JSON(rcv interface{}) error {
	return json.Unmarshal(t.Payload, rcv)
}

