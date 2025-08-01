package main

import (
	"context"

	raft "github.com/FrancescoLuzzi/raft/raft"
)

type RaftModuleClient interface {
	AppendEntries(ctx context.Context, nodeId uint32, request *raft.AppendEntriesRequest) (*raft.AppendEntriesResponse, error)
	RequestVote(ctx context.Context, nodeId uint32, request *raft.RequestVoteRequest) (*raft.RequestVoteResponse, error)
	InstallSnapshot(ctx context.Context, nodeId uint32, request *raft.InstallSnapshotRequest) (*raft.InstallSnapshotResponse, error)
}

type RaftModuleService interface {
	AppendEntries(ctx context.Context, request *raft.AppendEntriesRequest) (*raft.AppendEntriesResponse, error)
	RequestVote(ctx context.Context, request *raft.RequestVoteRequest) (*raft.RequestVoteResponse, error)
	InstallSnapshot(ctx context.Context, request *raft.InstallSnapshotRequest) (*raft.InstallSnapshotResponse, error)
}

// TODO: add randomization by looking up an environment variable
// or by creating a second implmentation for testing
type server struct {
	raft.UnimplementedRaftModuleServiceServer
	rms RaftModuleService
}

var _ raft.RaftModuleServiceServer = &server{}

func (s *server) AppendEntries(ctx context.Context, request *raft.AppendEntriesRequest) (*raft.AppendEntriesResponse, error) {
	return s.rms.AppendEntries(ctx, request)
}
func (s *server) RequestVote(ctx context.Context, request *raft.RequestVoteRequest) (*raft.RequestVoteResponse, error) {
	return s.rms.RequestVote(ctx, request)
}
func (s *server) InstallSnapshot(ctx context.Context, request *raft.InstallSnapshotRequest) (*raft.InstallSnapshotResponse, error) {
	return s.rms.InstallSnapshot(ctx, request)
}
