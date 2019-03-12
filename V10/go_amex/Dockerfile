############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
COPY . $GOPATH/src/dock/v10/go_amex
WORKDIR $GOPATH/src/dock/v10/go_amex
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
#RUN go build -o /go/bin/v01
#RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /app .

RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /go/bin/v10/go_amex .


############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/v10/go_amex /go/bin/v10/go_amex
# Run the hello binary.
ENTRYPOINT ["/go/bin/v10/go_amex"]

# docker restart todo_01
#  `docker run -p 80:80 <container>`

