FROM golang:1.18.3 as builder

# Set the Current Working Directory inside the container

WORKDIR $GOPATH/src/github.com/arshpreetsingh/hellofresh

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o bin/main -ldflags="-w -s"

#RUN go build -o bin/main

#RUN bin/main hf_test_calculation_fixtures.json 9AM 9PM 10115 pizza banana chicken

# Now run Second build for smaller Image!

############################
#STEP 2 build a small image
############################
FROM alpine

WORKDIR /

RUN apk update

#Copying Go Binary
RUN tail -f /dev/null

COPY --from=builder $GOPATH/src/github.com/arshpreetsingh/hellofresh/bin .
