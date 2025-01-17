package reputationrpc

import (
	"context"
	"crypto/ecdsa"

	"github.com/TrueCloudLab/frostfs-api-go/v2/reputation"
	"github.com/TrueCloudLab/frostfs-node/pkg/services/util"
)

type signService struct {
	sigSvc *util.SignService

	svc Server
}

func NewSignService(key *ecdsa.PrivateKey, svc Server) Server {
	return &signService{
		sigSvc: util.NewUnarySignService(key),
		svc:    svc,
	}
}

func (s *signService) AnnounceLocalTrust(ctx context.Context, req *reputation.AnnounceLocalTrustRequest) (*reputation.AnnounceLocalTrustResponse, error) {
	resp, err := s.sigSvc.HandleUnaryRequest(ctx, req,
		func(ctx context.Context, req any) (util.ResponseMessage, error) {
			return s.svc.AnnounceLocalTrust(ctx, req.(*reputation.AnnounceLocalTrustRequest))
		},
		func() util.ResponseMessage {
			return new(reputation.AnnounceLocalTrustResponse)
		},
	)
	if err != nil {
		return nil, err
	}

	return resp.(*reputation.AnnounceLocalTrustResponse), nil
}

func (s *signService) AnnounceIntermediateResult(ctx context.Context, req *reputation.AnnounceIntermediateResultRequest) (*reputation.AnnounceIntermediateResultResponse, error) {
	resp, err := s.sigSvc.HandleUnaryRequest(ctx, req,
		func(ctx context.Context, req any) (util.ResponseMessage, error) {
			return s.svc.AnnounceIntermediateResult(ctx, req.(*reputation.AnnounceIntermediateResultRequest))
		},
		func() util.ResponseMessage {
			return new(reputation.AnnounceIntermediateResultResponse)
		},
	)
	if err != nil {
		return nil, err
	}

	return resp.(*reputation.AnnounceIntermediateResultResponse), nil
}
