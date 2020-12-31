# go-protos
Protobufs in go

## Build
```shell
# Generate go files
$ docker build -t dopltechnologies/api-protos-go -f build.Dockerfile .

# Copy generated files to local dir
$ docker cp $(docker create --rm dopltechnologies/api-protos-go:latest):/output/. ./
```