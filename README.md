# WEB_TERMINAL
**It takes input from the user and returns correponding outputs.**

It is a service which takes commands from the user as the input and return the output of the corresponding command. 

It provides 2 service.

1) /health: this endpoint ensures that app is running and up.
Example: `http://localhost:9090/health`

2) /home: this endpoint provide user interface to enter their commands.
Internally "/cmd" endpoint is called. All the output of the data are sent to "/cmd" endpoint.

# There are 2 ways to run it:
## 1) By cloning repository

1) Clone the repo:
    
    `git clone https://github.com/viveksahu26/virtual_terminal.git`

2) Jump to the directory.

    `cd virtual_terminal`

3) Execute main program. 

    `go run main.go`
    
    You can provide your own custom port
    
     Replace PORT number by your own number in .env file.

*NOTE:*: Make sure that Port 9090 is free. By defaul Port is 9090. But you can customize accordingly by passing port number after command(go run main.go).

4) Check the health of program to ensure that app is running.

    http://localhost:9090/health

5) Now, go to /home endpoint

    http://localhost:9090/home

    Enter only non-interactive commands and press enter or Send button.

## 2) By docker image(recommended)
Docker image: https://hub.docker.com/r/viveksahu26/web_terminal

1) docker pull viveksahu26/web_terminal:v1.1

2) docker run -it --name web -p 9090:9090 viveksahu26/web_terminal:v1.1

3)  And repeat step 4 and 5 of above. 
