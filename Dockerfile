FROM alpine:edge

# Install packages
RUN apk update && \
    apk add go make

# Build project
WORKDIR /build
COPY . .
RUN go get && \
    make build-linux && \
    mv /build/target /app

# Clean up
RUN rm -rf /build
RUN apk del go make && \
    apk cache clean

# Start container
WORKDIR /app
EXPOSE 5002
ENTRYPOINT [ "./chatthing_linux" ]
