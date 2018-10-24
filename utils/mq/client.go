package mq

import "github.com/streadway/amqp"

type Client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewClient(uri string) (*Client, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	cli := &Client{
		conn:    conn,
		channel: channel,
	}

	return cli, nil
}

func (c *Client) PublishBody(ex, rk, body string) error {
	var err error
	channel, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	err = channel.Publish(ex, rk, false, false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		})

	return err
}

func (c *Client) GetConn() *amqp.Connection {
	return c.conn
}
