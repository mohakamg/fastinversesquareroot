# FISR
Fast Inverse Square Root Algorithm used for Quake III in GoLang

This program is implemented in GoLang Version 1.17

# Package
The Repo consists of a package fisr that implements the algorithm and can
be used as a standalone package.
To fetch the Package:
```bash 

go get github.com/mohakamg/fastinversesquareroot
```

To use the package:
```go

import (
	"fmt"
	"github.com/mohakamg/fastinversesquareroot/fisr"
)

func main() {
	if answer, err := fisr.FastInverseSquereRoot(2); err != nil {
		fmt.Println(answer)
	}
}
```

# Binary
The package can also be directly used as an executable to compute the Inverse Square Root

## Compile
```bash

# Create a Projects Directory
mkdir -p ~/projects
cd ~/project

# Set up the GO Path
export GOPATH=`pwd`
export PATH=$GOPATH/bin:$PATH

# Set up this repo
mkdir -p src/github.com/mohakamg/
cd src/github.com/mohakamg/
git clone https://github.com/mohakamg/fastinversesquareroot.git
cd fastinversesquareroot

# Make sure everything is fetched and formatted
go mod tidy

# Compile the Binary
cd cmd/fisrserver
go buiid . 
```

Note: In the last `go build .` command, extra flags can be specified to optimize performance
and space depending on OS and Architecture. 
Also, this binary is self contained and does not need anything else to run.

## Using the Binary
The binary file server ships with two modes:
1. Standalone: This standalone mode can be used like an bash command with the argument `number` passed to it
```bash
./fisrserver --number 2
```
Output: `2021/08/25 16:54:30 Result:  0.7069296`

On an unix machine this binary can optionally be moved to `/usr/local/bin` to be used as an installed software
```bash
cp ./firserver /usr/local/bin
fisrserver --number 2
2021/08/25 16:54:30 Result:  0.7069296
```

2. Server Mode: An HTTP Server to calculate the Inverse Square Root
To start the Server:

```bash
./fisrserver --server-endpoint="localhost:4000"
```

Make Requests to it:
```bash
wget -q -O - localhost:4000/fisr --post-data='{"instances": [2, 3]}'
```
Response: `{"result":{"2":0.7069296,"3":0.5768461}}`

# Docker
This repo also ships with a multistage Dockerfile that can be used to build in a temporary container
and export another distroless container to execute the server

1. Build and export
```
docker build -t fsir .
```

2. Run
2.1 Standalone
```
docker run fsir --number=2
```

2.2 Server
```
docker run -p 4000:4000 fsir --server-endpoint="0.0.0.0:4000"
```
