FROM alpine:edge AS base 
RUN apk update
RUN apk upgrade
RUN apk add --update go=1.20.1-r0 gcc=12.2.1_git20220924-r9 g++=12.2.1_git20220924-r9

WORKDIR /src
ENV CGO_ENABLED=1
COPY go.* .
COPY docs .
COPY api .
COPY cmd .
COPY db .
COPY graph .
COPY server .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/main -a -ldflags '-linkmode external -extldflags "-static"' ./cmd
#RUN CGO_ENABLED=1 GOOS=linux go build -o /app -a -ldflags '-linkmode external -extldflags "-static"' .
FROM base AS unit-test
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    mkdir /out && go test -v -coverprofile=/out/cover.out ./...

FROM golangci/golangci-lint:latest-alpine AS lint-base

FROM base AS lint
RUN --mount=target=. \
    --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
    golangci-lint run --timeout 10m0s ./...

FROM scratch AS unit-test-coverage
COPY --from=unit-test /out/cover.out /cover.out

FROM scratch AS bin-unix
COPY --from=build /out/main /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/main /main.exe

FROM bin-${TARGETOS} as bin
