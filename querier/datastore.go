package querier

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// DatastoreService expl
type DatastoreService struct {
	Client *datastore.Client
}

// NewDS expl
func NewDS(ctx context.Context, credPath string) (*DatastoreService, error) {
	creds, err := getCreds(ctx, credPath)
	if err != nil {
		return nil, err
	}

	client, err := datastore.NewClient(ctx, creds.ProjectID, option.WithCredentials(creds))
	if err != nil {
		return nil, err
	}

	return &DatastoreService{client}, nil
}

// DoQuery expl
func (ds *DatastoreService) DoQuery(ctx context.Context, query datastore.Query, outChan chan<- map[datastore.Key]interface{}, doneChan chan<- bool) error {
	it := ds.Client.Run(ctx, &query)

	for {
		var entry []interface{}
		key, err := it.Next(&entry)
		if err == iterator.Done {
			doneChan <- true
			break
		}
		if err != nil {
			return errors.Wrap(err, "While reading query response")
		}
		outChan <- map[datastore.Key]interface{}{*key: entry}
	}

	return nil
}

// MakeGQLquery expl
func (*DatastoreService) MakeGQLquery(kind string, args map[string]interface{}) *datastore.Query {
	query := datastore.NewQuery(kind)
	for arg, value := range args {
		if arg == "Limit" {
			query = query.Limit(value.(int))
		} else {
			query = query.Filter(arg+"=", value)
		}
	}
	return query
}
