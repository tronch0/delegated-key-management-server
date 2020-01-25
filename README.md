
###### - **DON'T USE THIS CODE IN PRODUCTION**


# Delegated Key-Management Server



**Key management solution** for storing keys at designated servers to remove the need to store passwords on Client's end.
This approach enforces the separation of concerns between the Client and the Server/s; Client can't access the keys, Server can't access the data.

Program steps:
1. The Client hides the data from Server by blinding the data (use hash-to-group technique then exponentiate to random secret by the Client) and sends it to Server.
2. The Server takes the given input from the Client and exponentiate it to its secret, then sends it back to the Client.
3. Client exponentiate to the inverse of its original random secret, which will unhide his original   
Data and keys are separated, 
4. The Client can now send the commitment with the Server's secret embedded into it.

## Running the program
```
go run main.go
```
##### Tests
```
go test -v ./crypto/.
```
