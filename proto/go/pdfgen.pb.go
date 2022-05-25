// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: pdfgen.proto

package _go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message
type GenerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Dpi          uint64  `protobuf:"varint,2,opt,name=dpi,proto3" json:"dpi,omitempty"`
	Zoom         float64 `protobuf:"fixed64,3,opt,name=zoom,proto3" json:"zoom,omitempty"`
	PageSize     string  `protobuf:"bytes,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Orientation  string  `protobuf:"bytes,5,opt,name=orientation,proto3" json:"orientation,omitempty"`
	Grayscale    bool    `protobuf:"varint,6,opt,name=grayscale,proto3" json:"grayscale,omitempty"`
	MarginLeft   string  `protobuf:"bytes,7,opt,name=marginLeft,proto3" json:"marginLeft,omitempty"`
	MarginRight  string  `protobuf:"bytes,8,opt,name=marginRight,proto3" json:"marginRight,omitempty"`
	MarginTop    string  `protobuf:"bytes,9,opt,name=marginTop,proto3" json:"marginTop,omitempty"`
	MarginBottom string  `protobuf:"bytes,10,opt,name=marginBottom,proto3" json:"marginBottom,omitempty"`
	HtmlBody     string  `protobuf:"bytes,11,opt,name=htmlBody,proto3" json:"htmlBody,omitempty"`
	HtmlHeader   string  `protobuf:"bytes,12,opt,name=htmlHeader,proto3" json:"htmlHeader,omitempty"`
	HtmlFooter   string  `protobuf:"bytes,13,opt,name=htmlFooter,proto3" json:"htmlFooter,omitempty"`
}

func (x *GenerateRequest) Reset() {
	*x = GenerateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pdfgen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRequest) ProtoMessage() {}

func (x *GenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pdfgen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRequest.ProtoReflect.Descriptor instead.
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return file_pdfgen_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GenerateRequest) GetDpi() uint64 {
	if x != nil {
		return x.Dpi
	}
	return 0
}

func (x *GenerateRequest) GetZoom() float64 {
	if x != nil {
		return x.Zoom
	}
	return 0
}

func (x *GenerateRequest) GetPageSize() string {
	if x != nil {
		return x.PageSize
	}
	return ""
}

func (x *GenerateRequest) GetOrientation() string {
	if x != nil {
		return x.Orientation
	}
	return ""
}

func (x *GenerateRequest) GetGrayscale() bool {
	if x != nil {
		return x.Grayscale
	}
	return false
}

func (x *GenerateRequest) GetMarginLeft() string {
	if x != nil {
		return x.MarginLeft
	}
	return ""
}

func (x *GenerateRequest) GetMarginRight() string {
	if x != nil {
		return x.MarginRight
	}
	return ""
}

func (x *GenerateRequest) GetMarginTop() string {
	if x != nil {
		return x.MarginTop
	}
	return ""
}

func (x *GenerateRequest) GetMarginBottom() string {
	if x != nil {
		return x.MarginBottom
	}
	return ""
}

func (x *GenerateRequest) GetHtmlBody() string {
	if x != nil {
		return x.HtmlBody
	}
	return ""
}

func (x *GenerateRequest) GetHtmlHeader() string {
	if x != nil {
		return x.HtmlHeader
	}
	return ""
}

func (x *GenerateRequest) GetHtmlFooter() string {
	if x != nil {
		return x.HtmlFooter
	}
	return ""
}

// Response message
type GenerateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pdf   []byte `protobuf:"bytes,1,opt,name=pdf,proto3" json:"pdf,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GenerateResponse) Reset() {
	*x = GenerateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pdfgen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResponse) ProtoMessage() {}

func (x *GenerateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pdfgen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResponse.ProtoReflect.Descriptor instead.
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return file_pdfgen_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateResponse) GetPdf() []byte {
	if x != nil {
		return x.Pdf
	}
	return nil
}

func (x *GenerateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pdfgen_proto protoreflect.FileDescriptor

var file_pdfgen_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x64, 0x66, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x64, 0x66, 0x67, 0x65, 0x6e, 0x22, 0x87, 0x03, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x70, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x64, 0x70, 0x69,
	0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x7a, 0x6f, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x61, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x67, 0x72, 0x61, 0x79, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x4c, 0x65, 0x66, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x4c, 0x65, 0x66, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x70, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x54, 0x6f, 0x70,
	0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x42, 0x6f, 0x74, 0x74, 0x6f, 0x6d,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x42, 0x6f,
	0x74, 0x74, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x74, 0x6d, 0x6c, 0x42, 0x6f, 0x64, 0x79,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x74, 0x6d, 0x6c, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x74, 0x6d, 0x6c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x74, 0x6d, 0x6c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x74, 0x6d, 0x6c, 0x46, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x74, 0x6d, 0x6c, 0x46, 0x6f, 0x6f, 0x74, 0x65, 0x72,
	0x22, 0x3a, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x64, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x70, 0x64, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x49, 0x0a, 0x06,
	0x50, 0x64, 0x66, 0x47, 0x65, 0x6e, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x64, 0x66, 0x67, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x64,
	0x66, 0x67, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x64, 0x61, 0x74,
	0x61, 0x2d, 0x63, 0x7a, 0x2f, 0x70, 0x64, 0x66, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pdfgen_proto_rawDescOnce sync.Once
	file_pdfgen_proto_rawDescData = file_pdfgen_proto_rawDesc
)

func file_pdfgen_proto_rawDescGZIP() []byte {
	file_pdfgen_proto_rawDescOnce.Do(func() {
		file_pdfgen_proto_rawDescData = protoimpl.X.CompressGZIP(file_pdfgen_proto_rawDescData)
	})
	return file_pdfgen_proto_rawDescData
}

var file_pdfgen_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pdfgen_proto_goTypes = []interface{}{
	(*GenerateRequest)(nil),  // 0: pdfgen.GenerateRequest
	(*GenerateResponse)(nil), // 1: pdfgen.GenerateResponse
}
var file_pdfgen_proto_depIdxs = []int32{
	0, // 0: pdfgen.PdfGen.Generate:input_type -> pdfgen.GenerateRequest
	1, // 1: pdfgen.PdfGen.Generate:output_type -> pdfgen.GenerateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pdfgen_proto_init() }
func file_pdfgen_proto_init() {
	if File_pdfgen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pdfgen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRequest); i {
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
		file_pdfgen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pdfgen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pdfgen_proto_goTypes,
		DependencyIndexes: file_pdfgen_proto_depIdxs,
		MessageInfos:      file_pdfgen_proto_msgTypes,
	}.Build()
	File_pdfgen_proto = out.File
	file_pdfgen_proto_rawDesc = nil
	file_pdfgen_proto_goTypes = nil
	file_pdfgen_proto_depIdxs = nil
}
