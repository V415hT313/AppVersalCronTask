# gRPC Report Generation Service

## How to Run

1. Clone the repository:
    ```bash
    git clone https://github.com/V415hT313/AppVersalCronTask.git
    cd AppVersalCronTask
    ```

2. Generate Go code from protobuf:
    ```bash
    protoc --go_out=. --go-grpc_out=. proto/report.proto
    ```

3. Run the service:
    ```bash
    go run main.go
    ```

## Testing gRPC Calls

You can use a tool like [grpcurl](https://github.com/fullstorydev/grpcurl) to test the gRPC calls:

```bash
grpcurl -plaintext localhost:50051 list
grpcurl -plaintext -d '{"user_id": "user1"}' localhost:50051 report.ReportService/GenerateReport
grpcurl -plaintext localhost:50051 report.ReportService/HealthCheck

