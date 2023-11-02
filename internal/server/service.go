// Copyright (c) 2022 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bow/iris/api"
	"github.com/bow/iris/internal"
	"github.com/bow/iris/internal/store"
)

// service implements the iris service API.
type service struct {
	api.UnimplementedIrisServer

	store store.FeedStore
}

// AddFeed satisfies the service API.
func (svc *service) AddFeed(
	ctx context.Context,
	req *api.AddFeedRequest,
) (*api.AddFeedResponse, error) {

	created, err := svc.store.AddFeed(
		ctx,
		req.GetUrl(),
		req.Title,
		req.Description,
		req.GetTags(),
		req.IsStarred,
	)
	if err != nil {
		return nil, err
	}

	payload, err := created.Proto()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	rsp := api.AddFeedResponse{Feed: payload}

	return &rsp, nil
}

// ListFeeds satisfies the service API.
func (svc *service) ListFeeds(
	ctx context.Context,
	_ *api.ListFeedsRequest,
) (*api.ListFeedsResponse, error) {

	feeds, err := svc.store.ListFeeds(ctx)
	if err != nil {
		return nil, err
	}

	rsp := api.ListFeedsResponse{}
	for _, feed := range feeds {
		proto, err := feed.Proto()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		rsp.Feeds = append(rsp.Feeds, proto)
	}

	return &rsp, nil
}

// EditFeeds satisfies the service API.
func (svc *service) EditFeeds(
	ctx context.Context,
	req *api.EditFeedsRequest,
) (*api.EditFeedsResponse, error) {

	ops := make([]*store.FeedEditOp, len(req.Ops))
	for i, op := range req.GetOps() {
		ops[i] = store.NewFeedEditOp(op)
	}

	feeds, err := svc.store.EditFeeds(ctx, ops)
	if err != nil {
		return nil, err
	}

	rsp := api.EditFeedsResponse{}
	for _, feed := range feeds {
		fp, err := feed.Proto()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		rsp.Feeds = append(rsp.Feeds, fp)
	}

	return &rsp, nil
}

// DeleteFeeds satisfies the service API.
func (svc *service) DeleteFeeds(
	ctx context.Context,
	req *api.DeleteFeedsRequest,
) (*api.DeleteFeedsResponse, error) {

	ids := make([]store.DBID, len(req.GetFeedIds()))
	for i, id := range req.GetFeedIds() {
		ids[i] = id
	}

	err := svc.store.DeleteFeeds(ctx, ids)

	rsp := api.DeleteFeedsResponse{}

	return &rsp, err
}

// PullFeeds satisfies the service API.
func (svc *service) PullFeeds(
	req *api.PullFeedsRequest,
	stream api.Iris_PullFeedsServer,
) error {

	convert := func(pr store.PullResult) (*api.PullFeedsResponse, error) {
		if err := pr.Error(); err != nil {
			url := pr.URL()
			if url == "" {
				return nil, err
			}
			rspErr := err.Error()
			rsp := api.PullFeedsResponse{Url: url, Error: &rspErr}
			return &rsp, nil
		}
		feed := pr.Feed()
		if feed == nil {
			return nil, nil
		}
		fp, err := feed.Proto()
		if err != nil {
			return nil, err
		}
		rsp := api.PullFeedsResponse{Url: pr.URL(), Feed: fp}

		return &rsp, nil
	}

	ids := make([]store.DBID, len(req.GetFeedIds()))
	for i, id := range req.GetFeedIds() {
		ids[i] = id
	}

	ch := svc.store.PullFeeds(stream.Context(), ids)

	for pr := range ch {
		payload, err := convert(pr)
		if err != nil {
			return err
		}
		if payload == nil {
			continue
		}
		if err := stream.Send(payload); err != nil {
			return err
		}
	}

	return nil
}

// ListEntries satisfies the service API.
func (svc *service) ListEntries(
	ctx context.Context,
	req *api.ListEntriesRequest,
) (*api.ListEntriesResponse, error) {

	entries, err := svc.store.ListEntries(ctx, req.GetFeedId())
	if err != nil {
		return nil, err
	}

	rsp := api.ListEntriesResponse{}
	for _, entry := range entries {
		proto, err := entry.Proto()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		rsp.Entries = append(rsp.Entries, proto)
	}

	return &rsp, nil
}

// EditEntries satisfies the service API.
func (svc *service) EditEntries(
	ctx context.Context,
	req *api.EditEntriesRequest,
) (*api.EditEntriesResponse, error) {

	ops := make([]*store.EntryEditOp, len(req.Ops))
	for i, op := range req.GetOps() {
		ops[i] = store.NewEntryEditOp(op)
	}

	entries, err := svc.store.EditEntries(ctx, ops)
	if err != nil {
		return nil, err
	}

	rsp := api.EditEntriesResponse{}
	for _, entry := range entries {
		ep, err := entry.Proto()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		rsp.Entries = append(rsp.Entries, ep)
	}

	return &rsp, nil
}

// GetEntry satisfies the service API.
func (svc *service) GetEntry(
	ctx context.Context,
	req *api.GetEntryRequest,
) (*api.GetEntryResponse, error) {

	entry, err := svc.store.GetEntry(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	ep, err := entry.Proto()
	if err != nil {
		return nil, err
	}
	rsp := api.GetEntryResponse{Entry: ep}

	return &rsp, nil
}

// ExportOPML satisfies the service API.
func (svc *service) ExportOPML(
	ctx context.Context,
	req *api.ExportOPMLRequest,
) (*api.ExportOPMLResponse, error) {

	payload, err := svc.store.ExportOPML(ctx, req.Title)
	if err != nil {
		return nil, err
	}

	rsp := api.ExportOPMLResponse{Payload: payload}

	return &rsp, nil
}

// ImportOPML satisfies the service API.
func (svc *service) ImportOPML(
	ctx context.Context,
	req *api.ImportOPMLRequest,
) (*api.ImportOPMLResponse, error) {

	nproc, nimp, err := svc.store.ImportOPML(ctx, req.Payload)
	if err != nil {
		return nil, err
	}

	rsp := api.ImportOPMLResponse{
		NumProcessed: uint32(nproc),
		NumImported:  uint32(nimp),
	}

	return &rsp, nil
}

// GetInfo satisfies the service API.
func (svc *service) GetInfo(
	_ context.Context,
	_ *api.GetInfoRequest,
) (*api.GetInfoResponse, error) {

	rsp := api.GetInfoResponse{
		Name:      internal.AppName(),
		Version:   internal.Version(),
		GitCommit: internal.GitCommit(),
		BuildTime: internal.BuildTime(),
	}

	return &rsp, nil
}
