package uzi

import (
	"yirv2/uzi/internal/domain"
	pb "yirv2/uzi/internal/generated/grpc/service"
)

func domainUziToPbUzi(d *domain.Uzi) *pb.Uzi {
	return &pb.Uzi{
		Id:         d.Id.String(),
		Projection: d.Projection,
		PatientId:  d.PatientID.String(),
		DeviceId:   int64(d.DeviceID),
	}
}
