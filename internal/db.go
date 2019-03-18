package internal

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/utrack/clay/v2/transport/httptransport"
)

type DB interface {
	DoSmth(ctx context.Context, message proto.Message, opts httptransport.OptionUnaryInterceptor) error
}
