# build stage
FROM golang:alpine AS builder
RUN apk add make \
    gcc \ 
    g++
ADD . /src
RUN cd /src && make build

# final stage
FROM alpine
WORKDIR /app
COPY --from=builder /src/bara /app/
RUN chmod +x ./bara
CMD ./bara