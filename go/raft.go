package main

import (
	"context"

	raft "github.com/FrancescoLuzzi/raft/raft"
)

type RaftModuleService interface {
	AppendEntries(context.Context, *raft.AppendEntriesRequest) (*raft.AppendEntriesResponse, error)
	RequestVote(context.Context, *raft.RequestVoteRequest) (*raft.RequestVoteResponse, error)
	InstallSnapshot(context.Context, *raft.InstallSnapshotRequest) (*raft.InstallSnapshotResponse, error)
}

type server struct {
	raft.UnimplementedRaftModuleServiceServer
}

var _ RaftModuleService = &server{}

func (s *server) AppendEntries(context.Context, *raft.AppendEntriesRequest) (*raft.AppendEntriesResponse, error) {
	return nil, nil
}
func (s *server) RequestVote(context.Context, *raft.RequestVoteRequest) (*raft.RequestVoteResponse, error) {
	return nil, nil
}
func (s *server) InstallSnapshot(context.Context, *raft.InstallSnapshotRequest) (*raft.InstallSnapshotResponse, error) {
	return nil, nil
}
