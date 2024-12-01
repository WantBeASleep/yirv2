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
	UpdateEchographic(ctx context.Context, in *pb.UpdateEchographicIn) (*pb.UpdateEchographicOut, error)
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
			Checked:    in.Checked,
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
	uzi, echographic, err := h.uziSrv.GetUzi(ctx, uuid.MustParse(in.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	pbUzi := domainUziToPbUzi(&uzi)
	pbEchographic := domainEchographicToPb(&echographic)
	pbUzi.Echographic = pbEchographic

	return &pb.GetUziOut{
		Uzi: pbUzi,
	}, nil
}

func (h *handler) UpdateEchographic(ctx context.Context, in *pb.UpdateEchographicIn) (*pb.UpdateEchographicOut, error) {
	echographic, err := h.uziSrv.UpdateEchographic(ctx, uuid.MustParse(in.Id), uzi.OptionalEchographic{
		Contors:         in.Contors,
		LeftLobeLength:  in.LeftLobeLength,
		LeftLobeWidth:   in.LeftLobeWidth,
		LeftLobeThick:   in.LeftLobeThick,
		LeftLobeVolum:   in.LeftLobeVolum,
		RightLobeLength: in.RightLobeLength,
		RightLobeWidth:  in.RightLobeWidth,
		RightLobeThick:  in.RightLobeThick,
		RightLobeVolum:  in.RightLobeVolum,
		GlandVolum:      in.GlandVolum,
		Isthmus:         in.Isthmus,
		Struct:          in.Struct,
		Echogenicity:    in.Echogenicity,
		RegionalLymph:   in.RegionalLymph,
		Vascularization: in.Vascularization,
		Location:        in.Location,
		Additional:      in.Additional,
		Conclusion:      in.Conclusion,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Что то пошло не так: %s", err.Error())
	}

	return &pb.UpdateEchographicOut{
		Echographic: domainEchographicToPb(&echographic),
	}, nil
}
