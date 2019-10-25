FROM golang:1-alpine AS build
RUN apk add --no-cache git
WORKDIR /src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

FROM vimagick/youtube-dl
COPY --from=build /server /server
ENTRYPOINT ["/server"]
