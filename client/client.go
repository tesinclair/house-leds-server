package client

import (
	"net"
	"log"
	"github.com/nu7hatch/gouuid"
)

type Client struct {
	Id string
	Connection net.Conn
	isReciever bool
}

func NewClient(c net.Conn, reciever bool) *Client{
	u, err := uuid.NewV4()
	if err != nil{
		log.Fatal("unable to assign uuid to client")
	}
	return &Client{
		Id: u.String(),
		Connection: c,
		isReciever: reciever,
	}
}
func (c *Client) SendMessage(message string){
	sendMessage := message + "\n"
	c.Connection.Write([]byte(string(sendMessage)))
}
func (c *Client) CloseConnection(){
	c.Connection.Close()
}