# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.16.3
# Adding labels
LABEL project="ASCII-ART-WEB" \
      authors="Dias1c, nrblzn" \
      description="That's ascii web project" \
      link="https://git.01.alem.school/Dias1c/ascii-art-web"
# Copy the local package files to the container's workspace.
RUN mkdir /ascii-art-web
ADD . /ascii-art-web
WORKDIR /ascii-art-web
# Build the program
RUN go build -o main .
# Document that the service listens on port 8080.
EXPOSE 8080
# Run the main command by default when the container starts.
# ENTRYPOINT /main
CMD ["./main"]