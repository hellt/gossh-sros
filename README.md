A basic example that demonstrates how to use Go [golang.org/x/crypto/ssh](https://pkg.go.dev/golang.org/x/crypto@v0.0.0-20201112155050-0c6587e931a9/ssh) package to remotely execute commands on Nokia SR OS routers.

Nokia SR OS requires a login shell to be spawned and bytes written to Stdin. 