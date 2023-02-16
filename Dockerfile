######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY bin/main .

# Expose port 8080 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
