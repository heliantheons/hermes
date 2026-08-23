FROM alpine:3.22

RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --chmod=0555 build/hermes /app/hermes

USER 65532:65532
ENTRYPOINT ["/app/hermes"]
