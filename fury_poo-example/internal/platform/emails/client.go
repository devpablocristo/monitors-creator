package emails

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c Client) Send(payload []byte, destiny string) error {
	return nil
}
