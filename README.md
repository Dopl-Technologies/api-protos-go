# go-protos
Protobufs in go

## Build
```shell
# Generate go files
$ docker build -t dopltechnologies/api-protos-go --build-arg CACHEBUST="$(curl https://api.github.com/repos/Dopl-Technologies/api-protos/commits/main 2>&1 | grep '"date"' | tail -n 1)" -f build.Dockerfile .

# Copy generated files to local dir
$ docker cp $(docker create --rm dopltechnologies/api-protos-go:latest):/output/. ./
```