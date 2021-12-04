package group

import (
	"context"

	"github.com/ozonmp/omp-bot/internal/model/logistic"
	grpc_api "github.com/ozonmp/omp-bot/internal/pkg/lgc-group-api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Service interface {
	Describe(ctx context.Context, groupID uint64) (*logistic.Group, error)
	List(ctx context.Context, cursor uint64, limit uint64) ([]logistic.Group, error)
	Create(ctx context.Context, grp logistic.Group) (uint64, error)
	Update(ctx context.Context, groupID uint64, group logistic.Group) error
	Remove(ctx context.Context, groupID uint64) (bool, error)
}

//  OmpGroupApiServiceClient - CRUD client
type OmpGroupApiServiceClient interface {
	DescribeGroupV1(ctx context.Context, in *grpc_api.DescribeGroupV1Request, opts ...grpc.CallOption) (*grpc_api.DescribeGroupV1Response, error)
	CreateGroupV1(ctx context.Context, in *grpc_api.CreateGroupV1Request, opts ...grpc.CallOption) (*grpc_api.CreateGroupV1Response, error)
	UpdateGroupV1(ctx context.Context, in *grpc_api.UpdateGroupV1Request, opts ...grpc.CallOption) (*grpc_api.UpdateGroupV1Response, error)
	RemoveGroupV1(ctx context.Context, in *grpc_api.RemoveGroupV1Request, opts ...grpc.CallOption) (*grpc_api.RemoveGroupV1Response, error)
}

// OmpGroupApiServiceMRClient - multy read client
type OmpGroupApiServiceMRClient interface {
	ListGroupV1(ctx context.Context, in *grpc_api.ListGroupV1Request, opts ...grpc.CallOption) (*grpc_api.ListGroupV1Response, error)
}

type grpcGroupService struct {
	groupApi    OmpGroupApiServiceClient
	groupFacade OmpGroupApiServiceMRClient
}

// NewGrpcGroupService - creates new grpc service to GroupApi & GroupFacade
func NewGrpcGroupService(grpcApi OmpGroupApiServiceClient, grpcFacade OmpGroupApiServiceMRClient) *grpcGroupService {
	return &grpcGroupService{
		groupApi:    grpcApi,
		groupFacade: grpcFacade,
	}
}

func (s *grpcGroupService) List(ctx context.Context, cursor uint64, limit uint64) ([]logistic.Group, error) {
	resp, err := s.groupFacade.ListGroupV1(ctx, &grpc_api.ListGroupV1Request{
		Offset: cursor,
		Limit:  limit,
	})
	if err != nil {
		return nil, errors.Wrap(err, "groupFacade.ListGroupV1()")
	}

	var result []logistic.Group
	for _, grp := range resp.GetGroups() {
		result = append(result, logistic.Group{
			Id:   grp.GetId(),
			Name: grp.GetName(),
		})
	}
	return result, nil
}

func (s *grpcGroupService) Describe(ctx context.Context, groupID uint64) (*logistic.Group, error) {
	resp, err := s.groupApi.DescribeGroupV1(ctx, &grpc_api.DescribeGroupV1Request{
		GroupId: groupID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "groupApi.DescribeGroupV1()")
	}
	return &logistic.Group{
		Id:   resp.GetGroup().GetId(),
		Name: resp.GetGroup().GetName(),
	}, nil
}

func (s *grpcGroupService) Create(ctx context.Context, grp logistic.Group) (uint64, error) {
	resp, err := s.groupApi.CreateGroupV1(ctx, &grpc_api.CreateGroupV1Request{
		Name: grp.Name,
	})
	if err != nil {
		return 0, errors.Wrap(err, "groupApi.CreateGroupV1()")
	}
	return resp.GetGroupId(), nil
}

func (s *grpcGroupService) Remove(ctx context.Context, groupID uint64) (bool, error) {
	_, err := s.groupApi.RemoveGroupV1(ctx, &grpc_api.RemoveGroupV1Request{
		GroupId: groupID,
	})
	if err != nil {
		return false, errors.Wrap(err, "groupApi.RemoveGroupV1()")
	}
	return true, nil
}

func (s *grpcGroupService) Update(ctx context.Context, groupID uint64, grp logistic.Group) error {
	_, err := s.groupApi.UpdateGroupV1(ctx, &grpc_api.UpdateGroupV1Request{
		GroupId: groupID,
		Name:    grp.Name,
	})
	if err != nil {
		return errors.Wrap(err, "groupApi.UpdateGroupV1()")
	}
	return nil
}
