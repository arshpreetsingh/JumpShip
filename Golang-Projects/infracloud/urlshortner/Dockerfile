FROM golang:1.18.3 as builder

# Set the Current Working Directory inside the container

WORKDIR $GOPATH/src/github.com/arshpreetsingh/urlshortner

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd/urlshortner cmd/main.go

ENTRYPOINT ["cmd/urlshortner"]

# Now run Second build for smaller Image!

############################
#STEP 2 build a small image
# ############################
# FROM alpine

# WORKDIR /

# RUN apk update

# #Copying Go Binary
# COPY --from=builder /go/src/github.com/arshpreetsingh/urlshortner .
#EXPOSE 8080
#CMD cmd/urlshortner