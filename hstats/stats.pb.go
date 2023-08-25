// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: stats.proto

package hstats

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

type HostStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId   string     `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Hostname string     `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Ip       string     `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Os       string     `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	Cpu      *CpuStats  `protobuf:"bytes,7,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Net      *NetStats  `protobuf:"bytes,8,opt,name=net,proto3" json:"net,omitempty"`
	Mem      *MemStats  `protobuf:"bytes,9,opt,name=mem,proto3" json:"mem,omitempty"`
	Disk     *DiskStats `protobuf:"bytes,10,opt,name=disk,proto3" json:"disk,omitempty"`
	Ts       int64      `protobuf:"varint,11,opt,name=ts,proto3" json:"ts,omitempty"`
}

func (x *HostStats) Reset() {
	*x = HostStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostStats) ProtoMessage() {}

func (x *HostStats) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostStats.ProtoReflect.Descriptor instead.
func (*HostStats) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{0}
}

func (x *HostStats) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *HostStats) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *HostStats) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *HostStats) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *HostStats) GetCpu() *CpuStats {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *HostStats) GetNet() *NetStats {
	if x != nil {
		return x.Net
	}
	return nil
}

func (x *HostStats) GetMem() *MemStats {
	if x != nil {
		return x.Mem
	}
	return nil
}

func (x *HostStats) GetDisk() *DiskStats {
	if x != nil {
		return x.Disk
	}
	return nil
}

func (x *HostStats) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

type CpuStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhysicalCores uint32  `protobuf:"varint,1,opt,name=physical_cores,json=physicalCores,proto3" json:"physical_cores,omitempty"`
	LogicalCores  uint32  `protobuf:"varint,2,opt,name=logical_cores,json=logicalCores,proto3" json:"logical_cores,omitempty"`
	Percent       float64 `protobuf:"fixed64,3,opt,name=percent,proto3" json:"percent,omitempty"`
}

func (x *CpuStats) Reset() {
	*x = CpuStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuStats) ProtoMessage() {}

func (x *CpuStats) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuStats.ProtoReflect.Descriptor instead.
func (*CpuStats) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{1}
}

func (x *CpuStats) GetPhysicalCores() uint32 {
	if x != nil {
		return x.PhysicalCores
	}
	return 0
}

func (x *CpuStats) GetLogicalCores() uint32 {
	if x != nil {
		return x.LogicalCores
	}
	return 0
}

func (x *CpuStats) GetPercent() float64 {
	if x != nil {
		return x.Percent
	}
	return 0
}

type NetStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RxPkt   uint64 `protobuf:"varint,1,opt,name=rx_pkt,json=rxPkt,proto3" json:"rx_pkt,omitempty"`
	TxPkt   uint64 `protobuf:"varint,2,opt,name=tx_pkt,json=txPkt,proto3" json:"tx_pkt,omitempty"`
	RxBytes uint64 `protobuf:"varint,3,opt,name=rx_bytes,json=rxBytes,proto3" json:"rx_bytes,omitempty"`
	TxBytes uint64 `protobuf:"varint,4,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
}

func (x *NetStats) Reset() {
	*x = NetStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetStats) ProtoMessage() {}

func (x *NetStats) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetStats.ProtoReflect.Descriptor instead.
func (*NetStats) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{2}
}

func (x *NetStats) GetRxPkt() uint64 {
	if x != nil {
		return x.RxPkt
	}
	return 0
}

func (x *NetStats) GetTxPkt() uint64 {
	if x != nil {
		return x.TxPkt
	}
	return 0
}

func (x *NetStats) GetRxBytes() uint64 {
	if x != nil {
		return x.RxBytes
	}
	return 0
}

func (x *NetStats) GetTxBytes() uint64 {
	if x != nil {
		return x.TxBytes
	}
	return 0
}

type MemStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     uint64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Used      uint64 `protobuf:"varint,2,opt,name=used,proto3" json:"used,omitempty"`
	Available uint64 `protobuf:"varint,3,opt,name=available,proto3" json:"available,omitempty"`
}

func (x *MemStats) Reset() {
	*x = MemStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemStats) ProtoMessage() {}

func (x *MemStats) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemStats.ProtoReflect.Descriptor instead.
func (*MemStats) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{3}
}

func (x *MemStats) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MemStats) GetUsed() uint64 {
	if x != nil {
		return x.Used
	}
	return 0
}

func (x *MemStats) GetAvailable() uint64 {
	if x != nil {
		return x.Available
	}
	return 0
}

type DiskStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReadCnt    uint64 `protobuf:"varint,1,opt,name=read_cnt,json=readCnt,proto3" json:"read_cnt,omitempty"`
	WriteCnt   uint64 `protobuf:"varint,2,opt,name=write_cnt,json=writeCnt,proto3" json:"write_cnt,omitempty"`
	ReadBytes  uint64 `protobuf:"varint,3,opt,name=read_bytes,json=readBytes,proto3" json:"read_bytes,omitempty"`
	WriteBytes uint64 `protobuf:"varint,4,opt,name=write_bytes,json=writeBytes,proto3" json:"write_bytes,omitempty"`
	Iops       uint64 `protobuf:"varint,5,opt,name=iops,proto3" json:"iops,omitempty"`
}

func (x *DiskStats) Reset() {
	*x = DiskStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskStats) ProtoMessage() {}

func (x *DiskStats) ProtoReflect() protoreflect.Message {
	mi := &file_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskStats.ProtoReflect.Descriptor instead.
func (*DiskStats) Descriptor() ([]byte, []int) {
	return file_stats_proto_rawDescGZIP(), []int{4}
}

func (x *DiskStats) GetReadCnt() uint64 {
	if x != nil {
		return x.ReadCnt
	}
	return 0
}

func (x *DiskStats) GetWriteCnt() uint64 {
	if x != nil {
		return x.WriteCnt
	}
	return 0
}

func (x *DiskStats) GetReadBytes() uint64 {
	if x != nil {
		return x.ReadBytes
	}
	return 0
}

func (x *DiskStats) GetWriteBytes() uint64 {
	if x != nil {
		return x.WriteBytes
	}
	return 0
}

func (x *DiskStats) GetIops() uint64 {
	if x != nil {
		return x.Iops
	}
	return 0
}

var File_stats_proto protoreflect.FileDescriptor

var file_stats_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0x8f, 0x02, 0x0a, 0x09, 0x48, 0x6f, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x25, 0x0a, 0x03, 0x63,
	0x70, 0x75, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x43, 0x70, 0x75, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x03, 0x63,
	0x70, 0x75, 0x12, 0x25, 0x0a, 0x03, 0x6e, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x4e, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x03, 0x6e, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x65, 0x6d,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x03, 0x6d, 0x65, 0x6d,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x22, 0x70, 0x0a, 0x08, 0x43, 0x70,
	0x75, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63,
	0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x72,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x22, 0x6e, 0x0a, 0x08,
	0x4e, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x78, 0x5f, 0x70,
	0x6b, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x72, 0x78, 0x50, 0x6b, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x74, 0x78, 0x5f, 0x70, 0x6b, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x74, 0x78, 0x50, 0x6b, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x78, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x78, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x08,
	0x4d, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x22, 0x97, 0x01, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x43, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6f, 0x70, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x69, 0x6f, 0x70, 0x73, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x61, 0x77, 0x6e, 0x77, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x68, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stats_proto_rawDescOnce sync.Once
	file_stats_proto_rawDescData = file_stats_proto_rawDesc
)

func file_stats_proto_rawDescGZIP() []byte {
	file_stats_proto_rawDescOnce.Do(func() {
		file_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_stats_proto_rawDescData)
	})
	return file_stats_proto_rawDescData
}

var file_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_stats_proto_goTypes = []interface{}{
	(*HostStats)(nil), // 0: hoststats.HostStats
	(*CpuStats)(nil),  // 1: hoststats.CpuStats
	(*NetStats)(nil),  // 2: hoststats.NetStats
	(*MemStats)(nil),  // 3: hoststats.MemStats
	(*DiskStats)(nil), // 4: hoststats.DiskStats
}
var file_stats_proto_depIdxs = []int32{
	1, // 0: hoststats.HostStats.cpu:type_name -> hoststats.CpuStats
	2, // 1: hoststats.HostStats.net:type_name -> hoststats.NetStats
	3, // 2: hoststats.HostStats.mem:type_name -> hoststats.MemStats
	4, // 3: hoststats.HostStats.disk:type_name -> hoststats.DiskStats
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_stats_proto_init() }
func file_stats_proto_init() {
	if File_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostStats); i {
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
		file_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuStats); i {
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
		file_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetStats); i {
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
		file_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemStats); i {
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
		file_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskStats); i {
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
			RawDescriptor: file_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stats_proto_goTypes,
		DependencyIndexes: file_stats_proto_depIdxs,
		MessageInfos:      file_stats_proto_msgTypes,
	}.Build()
	File_stats_proto = out.File
	file_stats_proto_rawDesc = nil
	file_stats_proto_goTypes = nil
	file_stats_proto_depIdxs = nil
}
