package grpc

import (
	"yirv2/med/internal/generated/grpc/service"
	"yirv2/med/internal/grpc/patient"
)

type Handler struct {
	patient.PatientHandler

	service.UnsafeMedSrvServer
}

func New(
	patientHandler patient.PatientHandler,
) *Handler {
	return &Handler{
		PatientHandler: patientHandler,
	}
}
