# build stage
FROM golang:alpine AS builder
RUN apk add make \
    gcc \ 
    g++
ADD . /src
RUN cd /src && make build
RUN cd /src && make build-sandbox

# final stage
FROM node:12.12.0-alpine
WORKDIR /app
COPY --from=builder /src/bara /app/
COPY --from=builder /src/sandbox-cli /app/
RUN adduser -D execUser
RUN chmod 700 ./bara
RUN chmod +x ./sandbox-cli
CMD ./bara