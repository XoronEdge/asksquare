#### Run the Applications
Here is the steps to run it with `docker-compose`
#move to directory
$ cd workspace

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/XoronEdge/asksquare.git

#move to project
$ cd asksquare

# Create .env file and put content in it
$ touch .env

# Run the application
$ docker-compose up --build


# Execute the call
$ curl localhost:3000



