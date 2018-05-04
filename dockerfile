FROM golang:1.9.2 as builder
WORKDIR /go/src/github.com/MarcvanMelle/face-tome
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /home/scratchuser
COPY --from=builder /etc/ssl /etc/ssl
COPY --from=builder /go/src/github.com/MarcvanMelle/face-tome/app .
USER 10001
ENTRYPOINT [ "./app" ]
EXPOSE 8080 8081
