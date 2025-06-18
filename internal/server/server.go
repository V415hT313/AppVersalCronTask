package server

import (
    "context"
    "net"
    "sync"
    "github.com/V415hT313/AppVersalCronTask/proto"
    "go.uber.org/zap"
    "google.golang.org/grpc"
)

type ReportServer struct {
    proto.UnimplementedReportServiceServer
    reports map[string]string
    mu      sync.Mutex
}

func (s *ReportServer) GenerateReport(ctx context.Context, req *proto.ReportRequest) (*proto.ReportResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    reportID := "report-" + req.UserId
    s.reports[reportID] = "Report for " + req.UserId
    zap.L().Info("Generated report", zap.String("report_id", reportID), zap.String("user_id", req.UserId))

    return &proto.ReportResponse{ReportId: reportID}, nil
}

func (s *ReportServer) HealthCheck(ctx context.Context, req *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
    return &proto.HealthCheckResponse{Status: "OK"}, nil
}

func StartGRPCServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        zap.L().Fatal("Failed to listen", zap.Error(err))
    }
    s := grpc.NewServer()
    proto.RegisterReportServiceServer(s, &ReportServer{reports: make(map[string]string)})
    zap.L().Info("Server listening", zap.String("address", lis.Addr().String()))
    if err := s.Serve(lis); err != nil {
        zap.L().Fatal("Failed to serve", zap.Error(err))
    }
}
