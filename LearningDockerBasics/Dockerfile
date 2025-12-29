# Set base image to build the rest of the image off of 
FROM golang:1.24.3

# Set working directory for inside container (essentially like CD ./app inside container)
WORKDIR /app

#Copy source code to container
COPY . .

#Compile go code into a binary
RUN go build main.go

#Expose 8080 port
EXPOSE 8080

#Execute binary
CMD [ "./main" ]
