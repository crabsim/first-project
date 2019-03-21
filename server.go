package main
import "fmt"
import "net/http"

func HelloWorld(w http.ResposeWriter, r *http.Request)
{	fmt.Fprintf(w,"HelloWorld")}
func main()
{
	http.HandleFunc("/",helloWorld)
	http.ListenandServe(":8080",nil)
}