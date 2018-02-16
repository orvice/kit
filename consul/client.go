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
	if p == nil {
		return []byte{}, nil
	}
	return p.Value, nil
}

func (c *Client) KVSet(key string, value []byte) error {
	p := &api.KVPair{Key: key, Value: value}
	_, err = kv.Put(p, nil)
	return err
}
