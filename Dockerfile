# stage 1
FROM golang:1.18-alpine as stage

WORKDIR /state-machine-app
COPY go.mod go.sum ./
RUN go mod download
# copy the source from the current directory to the Working Directory inside the container
COPY . .

ENV GO111MODULE=auto
RUN CGO_ENABLED=0 GOOS=linux go build github.com/isgo-golgo13/state-machine-go/svc/app

# stage 2
FROM golang:1.18-alpine 

WORKDIR /root/
COPY --from=stage /state-machine-app .
CMD ["./app"]