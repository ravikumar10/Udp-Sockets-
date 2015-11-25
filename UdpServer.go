package main
import (
    "fmt" 
    "net"  
)


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
    _,err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
    if err != nil {
        fmt.Printf("Couldn't send response %v", err)
    }
}


func main() {
    
    addr := net.UDPAddr{
        Port: 2000,
        IP: net.ParseIP("127.0.0.1"),
    }
    ser, err := net.ListenUDP("udp", &addr)
    if err != nil {
        fmt.Printf("Some error %v\n", err)
        return
    }
	num:=0
    for {
			p := make([]byte, 2048)
        _,remoteaddr,err := ser.ReadFromUDP(p)
        fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		
		num++
        if err !=  nil {
            fmt.Printf("Some error  %v", err)
			num++
            continue
        }
		fmt.Println("total packet recieved : ",num)
    }
	
}
