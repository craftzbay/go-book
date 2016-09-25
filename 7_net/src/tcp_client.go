package main 
import ( "fmt";"io/ioutil";"net";"os" ) 

func main() {
  tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:8088") 
  checkError(err) 

  // холболт тогтоох
  conn, err := net.DialTCP("tcp", nil, tcpAddr) 
  checkError(err) 

  // серверээс мэдээлэл унших
  result, err := ioutil.ReadAll(conn) 
  checkError(err) 

  // авсан хариуг дэлгэцэнд хэвлэж харуулах
  fmt.Println(string(result))
} 

func checkError(err error) { 
  if err != nil { 
	fmt.Fprintf(os.Stderr, "Алдаа: %s", err.Error()) 
	os.Exit(1) 
  } 
} 