package common

import (
	"bytes"
	"context"
	"github.com/jinzhu/copier"
	"github.com/trancer-nature/galaxy-common/util/format"
	"github.com/trancer-nature/galaxy-common/xerr"
	"go.opentelemetry.io/otel/trace"
)

type Response struct {
	Status    uint32
	Data      interface{}
	Message   string
	RequestID string
}

func getRequestID(ctx context.Context) string {
	span := trace.SpanContextFromContext(ctx)

	spanID := xerr.GetSpanID(span)
	traceID := xerr.GetTraceID(span)

	b := bytes.Buffer{}
	b.WriteString(spanID)

	if traceID != "" {
		b.WriteString("_")
		b.WriteString(traceID)
	}

	return b.String()
}

func NewResponse(ctx context.Context, resp interface{}) {
	requestID := getRequestID(ctx)
	response := &Response{
		Status:    xerr.DefaultCode,
		Message:   format.EmptyString,
		RequestID: requestID,
	}
	_ = copier.Copy(resp, response)
}
