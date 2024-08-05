// Package base implements server generated by protoc-gen-go.
package base

import (
	"context"
	"swclabs/swix/internal/mod/proto/base"
)

// NewBaseServer creates a new BaseServer.
func NewBaseServer() base.BaseServer {
	return &_Base{}
}

// Base implements base.BaseServer.
type _Base struct {
	base.UnimplementedBaseServer
}

// HealthCheck implements base.BaseServer.
func (s *_Base) HealthCheck(_ context.Context, _ *base.Empty) (*base.HealthCheckReply, error) {
	return &base.HealthCheckReply{Message: "OK"}, nil
}

// WorkerCheck implements base.BaseServer.
func (s *_Base) WorkerCheck(_ context.Context, _ *base.WorkerMessage) (*base.Empty, error) {
	return &base.Empty{}, nil
}

// WorkerCheckResult implements base.BaseServer.
func (s *_Base) WorkerCheckResult(_ context.Context, _ *base.WorkerMessage) (*base.WorkerReply, error) {
	return &base.WorkerReply{Message: "OK"}, nil
}
