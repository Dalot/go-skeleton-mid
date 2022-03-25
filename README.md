# Go Skeleton mid size project
A suggestion of the initial structure of a mid size project

## Run

### Local
```bash
go run main.go
```

#### Docker
```bash
docker build -t go-skeleton-mid .
docker run --rm -e ENV=dev -p8000:8000 go-skeleton-mid
```

## Build

### Local
```bash
go build cmd/app/main.go
```

## Test

```bash
make test
```
 
