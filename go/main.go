package main
import "fmt"
import "os/exec"

func main() {
	var (
		cmdOut []byte
		err    error
	)
	binary, lookErr := exec.LookPath("mvn")
    if lookErr != nil {
        panic(lookErr)
    }
    if cmdOut,err = exec.Command(binary, "-version").Output(); err != nil { 
        fmt.Println("Error: ", err)
    }
	fmt.Printf(string(cmdOut))
	
    if cmdOut,err = exec.Command("cmd", "/C", "dir", "d:").Output(); err != nil { 
        fmt.Println("Error: ", err)
    }
	fmt.Printf(string(cmdOut))
}