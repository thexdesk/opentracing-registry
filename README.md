opentracing-registry
---

This package registers middlware for `github.com/docker/distribution` that instruments registry with distributed tracing using `github.com/opentracing/opentracing-go`.

# Building

This project is built using `github.com/moby/buildkit`. So we will get a few dependencies first.

Install `buildctl` the buildkit client.
```
$ go get -u github.com/moby/buildkit/cmd/buildctl 
```

Run a buildkit daemon and set an environment variable to point your `buildctl` towards the containerized daemon.
```
$ docker run -d --name buildkit --privileged -p 1234:1234 tonistiigi/buildkit:standalone --addr tcp://0.0.0.0:1234
$ export BUILDKIT_HOST=tcp://0.0.0.0:1234
```

Buildkit depends on a specific `runc` version, so we need to build and install that.
```
$ go get -d github.com/opencontainers/runc
$ cd $GOPATH/src/github.com/opencontainers/runc
$ git checkout 74a17296470088de3805e138d3d87c62e613dfc4
$ make
$ sudo make install
```

Make!
```
$ make
```
