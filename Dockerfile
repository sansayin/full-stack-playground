FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY html .
COPY bin/main .
COPY test.db .
# Expose port 8080 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
