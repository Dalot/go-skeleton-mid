# Prepare
FROM golang:1.17-alpine as baseimg

RUN apk --no-cache upgrade && apk --no-cache add git make

WORKDIR /go/src/app/
COPY ./ /go/src/app

# Build
FROM baseimg as builder

COPY . ./
RUN make build

# Run
FROM scratch

COPY --from=builder /go/src/app/server /
COPY --from=builder /go/src/app/internal/database/messages.csv  /internal/database/messages.csv
EXPOSE 8000
CMD ["./server"]
