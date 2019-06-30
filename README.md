# CloudStorage [![Godoc](https://godoc.org/github.com/dcbCIn/CloudStorage?status.svg)](https://godoc.org/github.com/dcbCIn/CloudStorage)
Example system for [MidCloud](https://github.com/dcbCIn/MidCloud) usage, an Adaptive middleware for cloud services, with transparent choice of best cloud.
The basic conceptions are that the message will be sent to the server that has the lower cost to do the task and is available.

## Installation

Install [MidCloud](https://github.com/dcbCIn/MidCloud), and then run this project.

For MidCloud installation:

Standard `go get`:

```
$ go get -u github.com/dcbCIn/MidCloud/...
$ go get -u cloud.google.com/go/storage/...
$ go get -u google.golang.org/api/option/...
$ go get -u github.com/minio/minio-go/...
```

See [MidCloud](https://github.com/dcbCIn/MidCloud) for more information.

## Usage & Example

Run first the nameServer, then the googleStorage and/or awsStorage, then the clientStorate.

For and examples see the [CloudStorage Godoc](http://godoc.org/github.com/dcbCIn/CloudStorage).

## But Why?!

There exists a long list of middlewares, but MidCloud comes with the proposal of reduce the costs of multi-cloud 
applications, and to eradicate any downtime possible. 

Can be from a multi-cloud storage (see [CloudStorage](https://github.com/dcbCIn/CloudStorage) project), multi-cloud faas, 
or any other you want to create).