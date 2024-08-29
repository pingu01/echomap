# EchoMap
EchoMap is a tool made with Golang for port scanning and network mapping. 


# Install 
## Pre-requisites
- Golang ([Download](https://go.dev/dl/))

## Installation
```bash
git clone https://github.com/pingu01/echomap.git
cd echomap
go install
go build -o echomap main.go
sudo mv echomap /usr/local/bin
```

# Usage
```bash
echomap -t <target> -p <port1,port2>
```
