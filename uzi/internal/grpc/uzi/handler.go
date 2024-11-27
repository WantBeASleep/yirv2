package uzi

import (
	"context"

	"yirv2/pkg/gtc"
	"yirv2/uzi/internal/domain"
	pb "yirv2/uzi/internal/generated/grpc/service"
	"yirv2/uzi/internal/services/uzi"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UziHandler interface {
	CreateUzi(ctx context.Context, req *pb.CreateUziIn) (*pb.Id, error)
	UpdateUzi(ctx context.Context, req *pb.UpdateUziIn) (*pb.UpdateUziOut, error)
}

type handler struct {
	uziSrv uzi.Service
}

func New(
	uziSrv uzi.Service,
) UziHandler {
	return &handler{
		uziSrv: uziSrv,
	}
}

func (h *handler) CreateUzi(ctx context.Context, req *pb.CreateUziIn) (*pb.Id, error) {
	uuid, err := h.uziSrv.CreateUzi(ctx, domain.Uzi{
		Projection: req.Projection,
		PatientID:  uuid.MustParse(req.PatientId),
		DeviceID:   int(req.DeviceId),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.Id{Id: uuid.String()}, nil
}

func (h *handler) UpdateUzi(ctx context.Context, req *pb.UpdateUziIn) (*pb.UpdateUziOut, error) {
	uzi, err := h.uziSrv.UpdateUzi(ctx,
		uuid.MustParse(req.Id),
		uzi.OptionalUzi{
			Projection: req.Projection,
			PatientID:  gtc.Uuid.StringPToP(req.PatientId),
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateUziOut{
		Uzi: domainUziToPbUzi(&uzi),
	}, nil
}
