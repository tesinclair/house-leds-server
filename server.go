package main 

import(
	"fmt"
	"net"
	"bufio"
	"./client"
	"strings"
)

const (
	SERVER = ""
	PORT = "3333"
	TYPE = "tcp"

)

var clientList []client.Client

func handleConnection(c *client.Client){
	netData, err := bufio.NewReader(c.Connection).ReadString('\n')
	if err != nil{
		fmt.Println("Error reading: ", err.Error())
		return
	}
	message := strings.TrimSpace(string(netData))
	fmt.Println("[Message]: ", message)
	c.Write([]byte("Message recieved."))
	c.Close()
}

func handleComand(command string){
	commandParts := strings.Split(command, "@")
	for i=0; i<len(clientList); i++{
		if clientList[i].isReciever{
			switch commandParts[0]{
			case "SWITCH":
				if commandParts[1] == "on"{
					clientList[i].sendMessage("Lights@On")
				}
				else if commandParts[1] == "off"{
					clientList[i].sendMessage("Lights@Off")
				}
			}
		}
	}
}
func main(){
	l, err := net.Listen(TYPE, ":" + PORT)
	if err != nil{
		fmt.Println("Error listening: ", err.Error())
		return
	}
	
	defer l.Close()
	fmt.Println("Listening on " +  PORT)
	for{
		conn, err := l.Accept()
		fmt.Println("test")
		if err != nil{
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil{
			fmt.Println("Error reading: ", err.Error())
			return
		}
		reciever := strings.TrimSpace(string(netData))
		newClient := *client.NewClient(conn, reciever)
		clientList = append(clientList, newClient)
		go handleRequest(&newClient)
	}
}