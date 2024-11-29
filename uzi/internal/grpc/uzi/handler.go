package uzi

import (
	"context"

	"yirv2/pkg/gtclib"
	"yirv2/uzi/internal/domain"
	pb "yirv2/uzi/internal/generated/grpc/service"
	"yirv2/uzi/internal/services/uzi"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UziHandler interface {
	CreateUzi(ctx context.Context, req *pb.CreateUziIn) (*pb.CreateUziOut, error)
	UpdateUzi(ctx context.Context, req *pb.UpdateUziIn) (*pb.UpdateUziOut, error)
	GetUzi(ctx context.Context, in *pb.GetUziIn) (*pb.GetUziOut, error)
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

func (h *handler) CreateUzi(ctx context.Context, in *pb.CreateUziIn) (*pb.CreateUziOut, error) {
	uuid, err := h.uziSrv.CreateUzi(ctx, domain.Uzi{
		Projection: in.Projection,
		PatientID:  uuid.MustParse(in.PatientId),
		DeviceID:   int(in.DeviceId),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.CreateUziOut{Id: uuid.String()}, nil
}

func (h *handler) UpdateUzi(ctx context.Context, in *pb.UpdateUziIn) (*pb.UpdateUziOut, error) {
	uzi, err := h.uziSrv.UpdateUzi(ctx,
		uuid.MustParse(in.Id),
		uzi.OptionalUzi{
			Projection: in.Projection,
			PatientID:  gtclib.Uuid.StringPToP(in.PatientId),
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateUziOut{
		Uzi: domainUziToPbUzi(&uzi),
	}, nil
}

func (h *handler) GetUzi(ctx context.Context, in *pb.GetUziIn) (*pb.GetUziOut, error) {
	uzi, err := h.uziSrv.GetUzi(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.GetUziOut{
		Uzi: domainUziToPbUzi(&uzi),
	}, nil
}
