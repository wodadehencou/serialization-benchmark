# Benchmark some serialization and RPC framework

## Benchmark environment

| framework  | tools | version  |
| --- | --- | --- |
| golang | go | 1.14.4 |
| protobuffer | protoc | 3.13.0 |
| protobuffer go | protoc-gen-go | 1.25.0 |
| gogo protobuf | protoc-gen-gogoslick | 1.3.1 |


## Test Case

### Simple Message

| Field | Type | Length |
| --- | ---- | ---- |
| Str | string | 32 |
| Byte1 | bytes | 1024 |
| Byte2 | bytes | 64 |
| Timestamp | int64 |  |

### SimpleArr Message

| Field | Type | Length |
| --- | ---- | ---- |
| Simples | Simple Array | 4 |


