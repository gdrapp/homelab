FROM --platform=amd64 golang:1.19 as build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

ENV GOARCH=amd64

RUN CGO_ENABLED=0 go build -o /hass_homekit_cam_k8s

FROM --platform=amd64 gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=build /hass_homekit_cam_k8s /app/hass_homekit_cam_k8s

USER nonroot:nonroot

ENTRYPOINT [ "/app/hass_homekit_cam_k8s" ]