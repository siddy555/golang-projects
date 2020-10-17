package main
import(
  "fmt"
  "net/http"
)

func main(){
  http.HandleFunc("/", GoServer)
  http.ListenAndServe(":9090", nil)
}

func GoServer(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello, this is a sample web server build using GO")
}

