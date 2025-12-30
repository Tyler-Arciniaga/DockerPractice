# Set base image to build the rest of the image off of 
# This is the first build stage and is pretty heavy
FROM golang:1.24.3 AS build

# Set working directory for inside container (essentially like CD /app inside container)
WORKDIR /app

#Copy source code to container
COPY . .

#Compile go code into a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Set new base image to light weight Linux image
FROM alpine:latest

#Needed to run curl command as healthcheck for app service
RUN apk add --no-cache curl

WORKDIR /app

#Copy executable from first stage to this stage
COPY --from=build /app/server .

#Expose 8080 port
EXPOSE 8080

#Execute binary
CMD [ "./server" ]
