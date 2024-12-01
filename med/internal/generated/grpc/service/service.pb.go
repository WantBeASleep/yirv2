// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: service.proto

package service

import (
	reflect "reflect"
	sync "sync"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Patient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
	Fullname    string               `protobuf:"bytes,200,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email       string               `protobuf:"bytes,300,opt,name=email,proto3" json:"email,omitempty"`
	Policy      string               `protobuf:"bytes,400,opt,name=policy,proto3" json:"policy,omitempty"`
	Active      bool                 `protobuf:"varint,500,opt,name=active,proto3" json:"active,omitempty"`
	Malignancy  bool                 `protobuf:"varint,600,opt,name=malignancy,proto3" json:"malignancy,omitempty"`
	LastUziDate *timestamp.Timestamp `protobuf:"bytes,700,opt,name=last_uzi_date,json=lastUziDate,proto3,oneof" json:"last_uzi_date,omitempty"`
}

func (x *Patient) Reset() {
	*x = Patient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patient) ProtoMessage() {}

func (x *Patient) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patient.ProtoReflect.Descriptor instead.
func (*Patient) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *Patient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Patient) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *Patient) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Patient) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *Patient) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Patient) GetMalignancy() bool {
	if x != nil {
		return x.Malignancy
	}
	return false
}

func (x *Patient) GetLastUziDate() *timestamp.Timestamp {
	if x != nil {
		return x.LastUziDate
	}
	return nil
}

type CreatePatientIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fullname   string `protobuf:"bytes,200,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email      string `protobuf:"bytes,300,opt,name=email,proto3" json:"email,omitempty"`
	Policy     string `protobuf:"bytes,400,opt,name=policy,proto3" json:"policy,omitempty"`
	Active     bool   `protobuf:"varint,500,opt,name=active,proto3" json:"active,omitempty"`
	Malignancy bool   `protobuf:"varint,600,opt,name=malignancy,proto3" json:"malignancy,omitempty"`
}

func (x *CreatePatientIn) Reset() {
	*x = CreatePatientIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePatientIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePatientIn) ProtoMessage() {}

func (x *CreatePatientIn) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePatientIn.ProtoReflect.Descriptor instead.
func (*CreatePatientIn) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePatientIn) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *CreatePatientIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreatePatientIn) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *CreatePatientIn) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *CreatePatientIn) GetMalignancy() bool {
	if x != nil {
		return x.Malignancy
	}
	return false
}

type CreatePatientOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePatientOut) Reset() {
	*x = CreatePatientOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePatientOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePatientOut) ProtoMessage() {}

func (x *CreatePatientOut) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePatientOut.ProtoReflect.Descriptor instead.
func (*CreatePatientOut) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePatientOut) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPatientIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPatientIn) Reset() {
	*x = GetPatientIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPatientIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPatientIn) ProtoMessage() {}

func (x *GetPatientIn) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPatientIn.ProtoReflect.Descriptor instead.
func (*GetPatientIn) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetPatientIn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPatientOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patient *Patient `protobuf:"bytes,100,opt,name=patient,proto3" json:"patient,omitempty"`
}

func (x *GetPatientOut) Reset() {
	*x = GetPatientOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPatientOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPatientOut) ProtoMessage() {}

func (x *GetPatientOut) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPatientOut.ProtoReflect.Descriptor instead.
func (*GetPatientOut) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetPatientOut) GetPatient() *Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

type GetDoctorPatientsIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoctorId string `protobuf:"bytes,100,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`
}

func (x *GetDoctorPatientsIn) Reset() {
	*x = GetDoctorPatientsIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDoctorPatientsIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDoctorPatientsIn) ProtoMessage() {}

func (x *GetDoctorPatientsIn) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDoctorPatientsIn.ProtoReflect.Descriptor instead.
func (*GetDoctorPatientsIn) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetDoctorPatientsIn) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

type GetDoctorPatientsOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patients []*Patient `protobuf:"bytes,100,rep,name=patients,proto3" json:"patients,omitempty"`
}

func (x *GetDoctorPatientsOut) Reset() {
	*x = GetDoctorPatientsOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDoctorPatientsOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDoctorPatientsOut) ProtoMessage() {}

func (x *GetDoctorPatientsOut) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDoctorPatientsOut.ProtoReflect.Descriptor instead.
func (*GetDoctorPatientsOut) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetDoctorPatientsOut) GetPatients() []*Patient {
	if x != nil {
		return x.Patients
	}
	return nil
}

// TODO: перевести timestamp в других протиках
type UpdatePatientIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DoctorId    string               `protobuf:"bytes,100,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`
	Id          string               `protobuf:"bytes,200,opt,name=id,proto3" json:"id,omitempty"`
	Active      *bool                `protobuf:"varint,300,opt,name=active,proto3,oneof" json:"active,omitempty"`
	Malignancy  *bool                `protobuf:"varint,400,opt,name=malignancy,proto3,oneof" json:"malignancy,omitempty"`
	LastUziDate *timestamp.Timestamp `protobuf:"bytes,500,opt,name=last_uzi_date,json=lastUziDate,proto3,oneof" json:"last_uzi_date,omitempty"`
}

func (x *UpdatePatientIn) Reset() {
	*x = UpdatePatientIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePatientIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePatientIn) ProtoMessage() {}

func (x *UpdatePatientIn) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePatientIn.ProtoReflect.Descriptor instead.
func (*UpdatePatientIn) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePatientIn) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

func (x *UpdatePatientIn) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePatientIn) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

func (x *UpdatePatientIn) GetMalignancy() bool {
	if x != nil && x.Malignancy != nil {
		return *x.Malignancy
	}
	return false
}

func (x *UpdatePatientIn) GetLastUziDate() *timestamp.Timestamp {
	if x != nil {
		return x.LastUziDate
	}
	return nil
}

type UpdatePatientOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patient *Patient `protobuf:"bytes,100,opt,name=patient,proto3" json:"patient,omitempty"`
}

func (x *UpdatePatientOut) Reset() {
	*x = UpdatePatientOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePatientOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePatientOut) ProtoMessage() {}

func (x *UpdatePatientOut) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePatientOut.ProtoReflect.Descriptor instead.
func (*UpdatePatientOut) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatePatientOut) GetPatient() *Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf8, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x17, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x90, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x17, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0xf4, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x6d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x61, 0x6e, 0x63, 0x79,
	0x18, 0xd8, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x61,
	0x6e, 0x63, 0x79, 0x12, 0x44, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x7a, 0x69, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0xbc, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x7a, 0x69, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x7a, 0x69, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x12,
	0x1b, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xc8, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x90, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x17, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0xf4, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x6d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x61,
	0x6e, 0x63, 0x79, 0x18, 0xd8, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6d, 0x61, 0x6c, 0x69,
	0x67, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22,
	0x32, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x74, 0x6f,
	0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x08, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x64, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0xf5, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x0f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0xac, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x24, 0x0a, 0x0a, 0x6d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x18,
	0x90, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x0a, 0x6d, 0x61, 0x6c, 0x69, 0x67, 0x6e,
	0x61, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x75, 0x7a, 0x69, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0xf4, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x0b, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x7a, 0x69, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x61, 0x6c,
	0x69, 0x67, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x75, 0x7a, 0x69, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0x36, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x22, 0x0a,
	0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x32, 0xe3, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x64, 0x53, 0x72, 0x76, 0x12, 0x34, 0x0a, 0x0d,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x1a,
	0x11, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4f,
	0x75, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x1a,
	0x0e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x12,
	0x40, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x6f, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x4f, 0x75,
	0x74, 0x12, 0x34, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x10, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x1a, 0x11, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x42, 0x21, 0x5a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_proto_goTypes = []any{
	(*Patient)(nil),              // 0: Patient
	(*CreatePatientIn)(nil),      // 1: CreatePatientIn
	(*CreatePatientOut)(nil),     // 2: CreatePatientOut
	(*GetPatientIn)(nil),         // 3: GetPatientIn
	(*GetPatientOut)(nil),        // 4: GetPatientOut
	(*GetDoctorPatientsIn)(nil),  // 5: GetDoctorPatientsIn
	(*GetDoctorPatientsOut)(nil), // 6: GetDoctorPatientsOut
	(*UpdatePatientIn)(nil),      // 7: UpdatePatientIn
	(*UpdatePatientOut)(nil),     // 8: UpdatePatientOut
	(*timestamp.Timestamp)(nil),  // 9: google.protobuf.Timestamp
}
var file_service_proto_depIdxs = []int32{
	9, // 0: Patient.last_uzi_date:type_name -> google.protobuf.Timestamp
	0, // 1: GetPatientOut.patient:type_name -> Patient
	0, // 2: GetDoctorPatientsOut.patients:type_name -> Patient
	9, // 3: UpdatePatientIn.last_uzi_date:type_name -> google.protobuf.Timestamp
	0, // 4: UpdatePatientOut.patient:type_name -> Patient
	1, // 5: MedSrv.createPatient:input_type -> CreatePatientIn
	3, // 6: MedSrv.getPatient:input_type -> GetPatientIn
	5, // 7: MedSrv.getDoctorPatients:input_type -> GetDoctorPatientsIn
	7, // 8: MedSrv.updatePatient:input_type -> UpdatePatientIn
	2, // 9: MedSrv.createPatient:output_type -> CreatePatientOut
	4, // 10: MedSrv.getPatient:output_type -> GetPatientOut
	6, // 11: MedSrv.getDoctorPatients:output_type -> GetDoctorPatientsOut
	8, // 12: MedSrv.updatePatient:output_type -> UpdatePatientOut
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Patient); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePatientIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePatientOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetPatientIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetPatientOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetDoctorPatientsIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetDoctorPatientsOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdatePatientIn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*UpdatePatientOut); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_service_proto_msgTypes[0].OneofWrappers = []any{}
	file_service_proto_msgTypes[7].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}