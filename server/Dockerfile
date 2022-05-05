FROM golang:1.18.0-stretch as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest 
RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Riyadh /etc/localtime
RUN echo "Asia/Riyadh" >  /etc/timezone
WORKDIR /app
COPY --from=builder /app .
RUN ["chmod", "+x", "./app"]
CMD ./app 