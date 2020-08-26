# build stage
FROM golang:alpine AS builder
RUN apk add make \
    gcc \ 
    g++
ADD . /src
RUN cd /src && make build

# final stage
FROM node:12.12.0-alpine
WORKDIR /app
COPY --from=builder /src/bara /app/
COPY --from=builder /src/sandbox-cli /app/
RUN addgroup -g 2000 execuser && \
    adduser -D -u 2001 -G execuser execuser

EXPOSE 8080
ENV PORT=8080 
RUN chmod 700 ./bara
RUN chmod +x ./sandbox-cli
CMD ./bara