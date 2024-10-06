FROM golang:1.23-alpine AS build
# For gcc compilers
RUN apk add --update --no-cache build-base
RUN apk add --update --no-cache upx

WORKDIR /go/src/app

COPY . /go/src/app/
RUN go mod download
# Add options for building go with static library
RUN GOOS=linux CGO_ENABLED=1 go build -ldflags "-s -w -extldflags '-static'" -o ./backend
# Compress the binary file using Ultimate Packer for eXecutables
RUN upx ./backend

# Stage 2 final
FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/app/backend /backend

ENTRYPOINT ["/backend"]

EXPOSE 8080