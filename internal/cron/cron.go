package cron

import (
    "context"
    "log"
    "time"

    "github.com/jasonlvhit/gocron"
    "github.com/V415hT313/AppVersalCronTask/internal/logger"
    "github.com/V415hT313/AppVersalCronTask/proto"
    "google.golang.org/grpc"
)

func StartCronJob() {
    s := gocron.NewScheduler()
    s.Every(10).Second().Do(func() {
        conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
        if err != nil {
            logger.Log.Error("Failed to dial server", zap.Error(err))
            return
        }
        defer conn.Close()

        client := proto.NewReportServiceClient(conn)
        userIDs := []string{"user1", "user2", "user3"}
        for _, userID := range userIDs {
            _, err := client.GenerateReport(context.Background(), &proto.ReportRequest{UserId: userID})
            if err != nil {
                logger.Log.Error("Failed to generate report", zap.String("user_id", userID), zap.Error(err))
            }
        }
    })

    <-s.Start()
    logger.Log.Info("Cron job started")
}
