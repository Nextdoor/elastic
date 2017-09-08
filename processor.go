package elastic

import "context"

// RequestProcessor specifies the interface for processing requests.
type RequestProcessor interface {
	Process(ctx context.Context, req *Request)
}
