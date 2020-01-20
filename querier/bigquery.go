package querier

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// BigQueryService expl
type BigQueryService struct {
	Client *bigquery.Client
}

// NewBQ expl
func NewBQ(ctx context.Context, credPath string) (*BigQueryService, error) {
	creds, err := getCreds(ctx, credPath)
	if err != nil {
		return nil, err
	}

	client, err := bigquery.NewClient(ctx, creds.ProjectID, option.WithCredentials(creds))
	if err != nil {
		return nil, err
	}

	return &BigQueryService{client}, nil
}

// DoQuery expl
func (*BigQueryService) DoQuery(ctx context.Context, query bigquery.Query, outChan chan<- []bigquery.Value, doneChan chan<- bool) error {
	it, err := query.Read(ctx)
	if err != nil {
		return errors.Wrap(err, "While querying BigQuery")
	}

	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			doneChan <- true
			break
		}
		if err != nil {
			return errors.Wrap(err, "While reading query response")
		}
		outChan <- values
	}

	return nil
}

// GetQuota expl
func (*BigQueryService) GetQuota(ctx context.Context, query bigquery.Query) (bytes int64, err error) {
	query.DryRun = true
	query.Location = "EU"

	job, err := query.Run(ctx)
	if err != nil {
		return
	}

	status := job.LastStatus()
	if err != nil {
		return
	}

	return status.Statistics.TotalBytesProcessed, nil
}

// MakeSQLquery expl
func (bq *BigQueryService) MakeSQLquery(query string) *bigquery.Query {
	return bq.Client.Query(query)
}
