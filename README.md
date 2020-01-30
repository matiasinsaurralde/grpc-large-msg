# Instructions

Build the project:

```
$ go build
```

Run the project:

```
$ ./grpc-large-msg
```

To test with different message sizes, set the `GRPC_MAX_SIZE` environment variable, which represents the number of bytes to use. It will also generate enough test data to use it:

```
$ GRPC_MAX_SIZE=20000000 ./grpc-large-msg
```

## To deploy:

Generate a bundle:

```
$ tyk bundle build -y
```

or:

```
$ tyk-cli bundle build -y
```

Attach the bundle to an API (no auth is required, the hooks provided here are just `Pre` hooks, so `use_keyless` can be set to `true`):

```json
{
    "custom_middleware_bundle": "bundle.zip"
}
```

Deploy the bundle to a local HTTP server and set the bundle settings in `tyk.conf`:

```json
{
    "bundle_base_url": "http://localhost/",
    "enable_bundle_downloader": true
}
```

Configure the appropriate gRPC settings in `tyk.conf`. `grpc_recv_max_size` and `grpc_send_max_size` set the gRPC message size limits on the Tyk side:

```json
    "coprocess_options": {
        "enable_coprocess": true,
        "coprocess_grpc_server": "tcp://127.0.0.1:9111",
        "grpc_recv_max_size": 10000000,
        "grpc_send_max_size": 10000000
    },
```

When the gRPC client receives a message larger than the specified size, it will throw the following error, where the first value is the actual message size and the second one is the max message size:

```
[Jan 30 17:16:31] ERROR Dispatch error api_id=3 api_name=Tyk Test API error=rpc error: code = ResourceExhausted desc = grpc: received message larger than max (9999226 vs. 4194304) mw=CoProcessMiddleware org_id=default origin=::1 path=/quickstart/headers
```