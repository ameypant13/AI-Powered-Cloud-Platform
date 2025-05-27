package cloudwatch

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"time"
)

type CloudWatchFetcher struct {
	Client *cloudwatch.Client
}

func NewCloudWatchFetcher() (*CloudWatchFetcher, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	return &CloudWatchFetcher{Client: cloudwatch.NewFromConfig(cfg)}, nil
}

func (cw *CloudWatchFetcher) FetchCPUUtilization(instanceID string, start, end time.Time) ([]types.Datapoint, error) {
	out, err := cw.Client.GetMetricStatistics(context.TODO(), &cloudwatch.GetMetricStatisticsInput{
		StartTime:  &start,
		EndTime:    &end,
		MetricName: aws.String("CPUUtilization"),
		Namespace:  aws.String("AWS/EC2"),
		Period:     aws.Int32(300), // 5 minutes
		Statistics: []types.Statistic{types.StatisticAverage},
		Dimensions: []types.Dimension{
			{
				Name:  aws.String("InstanceId"),
				Value: aws.String(instanceID),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return out.Datapoints, nil
}
