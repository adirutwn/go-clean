#Docker multi-stage builds

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------

#Builder stageresources
FROM golang:1.13-alpine as builder

# Force the go compiler to use modules
ENV GO111MODULE=on

# Update OS package and install Git
RUN apk update && apk add git openssh && apk add build-base

# Set working directory
WORKDIR /go/src/github.com/adirutwn/go-clean

# Install wait-for
RUN wget https://raw.githubusercontent.com/eficode/wait-for/master/wait-for -O /usr/local/bin/wait-for &&\
    chmod +x /usr/local/bin/wait-for

# Copy Go dependency file
ADD go.mod go.mod
ADD go.sum go.sum
ADD app app
ADD Makefile Makefile

RUN go mod download

# Install fresh for hot-reload local development
RUN go get github.com/pilu/fresh
#
# Install go tool for convert go test output to junit xml
RUN go get -u github.com/jstemmer/go-junit-report
RUN go get github.com/axw/gocov/gocov
RUN go get github.com/AlekSi/gocov-xml

# Set Docker's entry point commands
RUN cd app/ && go build -o /go/bin/app.bin


# ------------------------------------------------------------------------------
# Deployment image
# ------------------------------------------------------------------------------

#App stage
FROM golang:1.13-alpine

RUN apk add --no-cache tini tzdata
RUN addgroup -g 211000 -S appgroup && adduser -u 211000 -S appuser -G appgroup

# Set working directory
WORKDIR /app

#Get artifact from buiber stage
RUN mkdir -p migrations
RUN mkdir -p recovery_migrations

#COPY --from=builder /out/ /out/
COPY --from=builder /go/bin/app.bin /app/app.bin
COPY --from=builder /go/src/github.com/adirutwn/app/migrations/ migrations/
COPY --from=builder /usr/local/bin/wait-for /usr/local/bin/wait-for
COPY --from=builder /go/pkg/mod/github.com/adirutwn/ /go/pkg/mod/github.com/adirutwn/

# Set Docker's entry point commands
RUN chown -R appuser:appgroup /go/pkg/mod/github.com/adirutwn/ /app
USER appuser
ENTRYPOINT ["/sbin/tini","-sg","--","/app/app.bin"]
