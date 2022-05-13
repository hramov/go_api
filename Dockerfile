FROM golang:alpine as build
RUN apk --no-cache add -U ca-certificates

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./bin/server /server
COPY logs logs/
COPY uploads/ uploads/
EXPOSE 3000 3001
ENTRYPOINT ["/server"]