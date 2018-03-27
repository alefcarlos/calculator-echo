##URL DE EXEMPLO: https://blog.golang.org/docker

# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang AS builder

# Go dep!
RUN go get -u github.com/golang/dep/...

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/alefcarlos/calculator-echo

WORKDIR /go/src/github.com/alefcarlos/calculator-echo

RUN dep ensure

# Build the calculator-service command inside the container.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/calculator-echo -v github.com/alefcarlos/calculator-echo/src/cmd/webapi

FROM scratch

# Copiar arquivo .env específico para o docker
#COPY docker.env .env

# Copiar arquivo binário
COPY --from=builder /bin/calculator-echo calculator-echo

CMD [ "./calculator-echo" ]