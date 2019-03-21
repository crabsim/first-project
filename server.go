package main
import "fmt"
import "net/http"

func HelloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"HelloWorld")
}
func main(){

	http.HandleFunc("/",HelloWorld)
	http.ListenAndServe(":8080",nil)
}
