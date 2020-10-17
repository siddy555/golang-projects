# WEBSERVER USING GO

### Line by Line explanation of Code.
&nbsp;
```Shell
package main
```
This means that this code with be executed as an executable instead of being treated as a library. Also, this is the starting point to run a program.

&nbsp;
```Shell
import(
  "fmt"
  "net/http"
)
```
Here we are importing packages. You can import packages from standard libraries of GO, third party or your own packages!

&nbsp;
```Shell
func main()
```
A function named "main" has been declared. This is a  special function. Thhe program execution starts here.

&nbsp;
```Shell
http.HandleFunc("/", GoServer)
```
The HandleFunc uses DefaultServeMux, which is basically an HTTP router which routes traffic based on path specified(prefix). In our case "/"

&nbsp;
```Shell
http.ListenAndServe(":9090", nil)
```
This is the webserver. Notice that our webserver is listening on port 9090 for all domains/IPs which lead to this server. If you are running on your PC, it's loaclhost. "nil" means we want to use the default ServerMux

&nbsp;
```Shell
GoServer(w http.ResponseWriter, r *http.Request)
```
This is a handler function which takes two arguments as input parameters.
1. Variable name: "w"  Data type: "http.ResponseWriter". Used to send response
2. Variable name:"r" Data type: "http.Request". Used to read the request.

&nbsp;
```Shell
fmt.Fprintf(w, "Hello, this is a sample web server built using GO")
```
We use the Fprintf function from fmt library to display output. In this case, a webpage. Lets keep things simple as of now :)
