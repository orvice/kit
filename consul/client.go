package consul

import "github.com/hashicorp/consul/api"

type Client struct {
	Client *api.Client
}

func NewClient(cfg *api.Config) (*Client, error) {
	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: client,
	}, nil
}

func (c *Client) KVGet(key string) ([]byte, error) {
	p, _, err := c.Client.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}
	return p.Value, nil
}
