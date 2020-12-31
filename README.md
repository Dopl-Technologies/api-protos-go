# go-protos
Protobufs in go

## Build
```shell
# Get the date of the last commit. This will force a pull from the repo if there has been a commit
$ CACHEBUST=$(curl https://api.github.com/repos/Dopl-Technologies/api-protos/commits/main 2>&1 | grep '"date"' | tail -n 1)

# Generate go files
$ docker build -t dopltechnologies/api-protos-go --build-arg CACHEBUST="$CACHEBUST" -f build.Dockerfile .

# Copy generated files to local dir
$ docker cp $(docker create --rm dopltechnologies/api-protos-go:latest):/output/. ./
```