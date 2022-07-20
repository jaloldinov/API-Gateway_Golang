# workspace (GOPATH) configured at /go
FROM golang:1.16 as builder


#
RUN mkdir -p $GOPATH/src/gitlab.udevs.io/urecruit/api_gateway
WORKDIR $GOPATH/src/gitlab.udevs.io/urecruit/api_gateway

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make build && \
    mv ./bin/api_gateway /



FROM alpine
COPY --from=builder api_gateway .
RUN mkdir config
#COPY ./config/rbac_model.conf ./config/rbac_model.conf
#COPY ./config/AuthKey_5RAX23V6QP.p8 ./config/AuthKey_5RAX23V6QP.p8
ENTRYPOINT ["/api_gateway"]
