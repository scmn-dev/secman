FROM --platform=${BUILDPLATFORM} golang:1.14.3-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o secman.go .

FROM scratch AS bin-unix
COPY --from=build secman.go /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build secman.go /example.exe

FROM bin-${TARGETOS} AS bin