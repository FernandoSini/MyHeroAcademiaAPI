#guardar pra caso dê merda
FROM golang:alpine AS build_base

RUN apk add --no-cache git

RUN mkdir /app
ADD . /app
# Set the Current Working Directory inside the container
WORKDIR /app/my-hero-academia-api
RUN cd /app/my-hero-academia-api && mkdir tmp

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/my-hero-academia-api .

# Start fresh from a smaller image
FROM alpine:3.14.6
RUN apk add ca-certificates

#ENV URL=mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false
#ENV SECRET_KEY=ASKLJHDJASDKJASHDKJASHDAKJSDHKASJDHJKAS
#ENV API_PORT=8080
#ENV DATABASENAME=MyHeroDataBase

COPY --from=build_base /app/my-hero-academia-api/out/my-hero-academia-api /app/my-hero-academia-api

COPY --from=build_base /app/my-hero-academia-api/.env .
# This container exposes port 8080 to the outside world
#EXPOSE 8080


# Run the binary program produced by `go install`
CMD ["/app/my-hero-academia-api"]