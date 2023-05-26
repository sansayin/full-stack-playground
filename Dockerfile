FROM golang:alpine AS base 

WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
COPY api .
COPY rpc .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

ARG TARGETOS
ARG TARGETARCH
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/rest  ./api

RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/rpc  ./rpc

EXPOSE 8888 

CMD /out/rest   
