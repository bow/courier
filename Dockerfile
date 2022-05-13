FROM golang:1.18-alpine AS builder

WORKDIR /src

RUN apk add --update --no-cache build-base~=0 make~=4 git~=2

COPY .git /src/.git

RUN git checkout -- . && make bin

# -- #

FROM golang:1.18-alpine

ARG REVISION
ARG BUILD_TIME

LABEL org.opencontainers.image.title="courier"
LABEL org.opencontainers.image.revision="${REVISION}"
LABEL org.opencontainers.image.created="${BUILD_TIME}"

COPY --from=builder /src/bin/courier /bin/courier

WORKDIR /runtime
ENTRYPOINT ["courier"]
