package channels

import (
	"github.com/labstack/gommon/log"
	"io"
	"net"
	"os"
)

//发送与接受 ： <-
func main(){
	conn,err := net.Dial("tcp","localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func (){
		io.Copy(os.Stdout,conn)
		log.Printf("done")
		done <- struct{}{}
	}()
	
	mustCopy(conn,os.Stdin)
	conn.Close()
	<-done
}

func mustCopy (dst io.Writer,src io.Reader) {
	if _,err := io.Copy(dst,src);err!=nil {
		log.Fatal(err)
	}
}