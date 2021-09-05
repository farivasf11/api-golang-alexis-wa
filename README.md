# api-golang-alexis-wa
A basic api made with golang for the Go Bootcamp Q3 2021 of the Wizeline Academy

#Requeriments
1. Need to have installed Go (1.16.7 or superior)

#How to run the server
1. Open a terminal on the path of the directory
2. The server runs on the localhost and requires and port to open it, set the port by a environment variable with the name of "API_PORT" followed by the command go run . 
    eg: "API_PORT=:3030 go run ."
3. The server can handle only two routes: the home route "/" and "/anime-quotes". 
    Go to a browser and enter the localhost:<API_PORT> and you can made the request to any of the two routes availables.