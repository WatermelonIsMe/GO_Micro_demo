// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userStatus.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Session struct {
	Id      int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=Address" json:"Address,omitempty"`
	Phone   string `protobuf:"bytes,4,opt,name=Phone" json:"Phone,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Session) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Session) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Session) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Session) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type GetSessionByUIDReq struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *GetSessionByUIDReq) Reset()                    { *m = GetSessionByUIDReq{} }
func (m *GetSessionByUIDReq) String() string            { return proto.CompactTextString(m) }
func (*GetSessionByUIDReq) ProtoMessage()               {}
func (*GetSessionByUIDReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GetSessionByUIDReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type GetSessionByUIDRep struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *GetSessionByUIDRep) Reset()                    { *m = GetSessionByUIDRep{} }
func (m *GetSessionByUIDRep) String() string            { return proto.CompactTextString(m) }
func (*GetSessionByUIDRep) ProtoMessage()               {}
func (*GetSessionByUIDRep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetSessionByUIDRep) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type GetSessionByTokenReq struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetSessionByTokenReq) Reset()                    { *m = GetSessionByTokenReq{} }
func (m *GetSessionByTokenReq) String() string            { return proto.CompactTextString(m) }
func (*GetSessionByTokenReq) ProtoMessage()               {}
func (*GetSessionByTokenReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetSessionByTokenReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UpdateConnectorAddrReq struct {
	Uid                int64  `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	ConnectorAddr      string `protobuf:"bytes,2,opt,name=connectorAddr" json:"connectorAddr,omitempty"`
	ConnectorSessionID uint64 `protobuf:"varint,3,opt,name=connectorSessionID" json:"connectorSessionID,omitempty"`
}

func (m *UpdateConnectorAddrReq) Reset()                    { *m = UpdateConnectorAddrReq{} }
func (m *UpdateConnectorAddrReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateConnectorAddrReq) ProtoMessage()               {}
func (*UpdateConnectorAddrReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *UpdateConnectorAddrReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateConnectorAddrReq) GetConnectorAddr() string {
	if m != nil {
		return m.ConnectorAddr
	}
	return ""
}

func (m *UpdateConnectorAddrReq) GetConnectorSessionID() uint64 {
	if m != nil {
		return m.ConnectorSessionID
	}
	return 0
}

type GetSessionByTokenRep struct {
	Session *Session `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
}

func (m *GetSessionByTokenRep) Reset()                    { *m = GetSessionByTokenRep{} }
func (m *GetSessionByTokenRep) String() string            { return proto.CompactTextString(m) }
func (*GetSessionByTokenRep) ProtoMessage()               {}
func (*GetSessionByTokenRep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *GetSessionByTokenRep) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type GetConnectorAddrReq struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *GetConnectorAddrReq) Reset()                    { *m = GetConnectorAddrReq{} }
func (m *GetConnectorAddrReq) String() string            { return proto.CompactTextString(m) }
func (*GetConnectorAddrReq) ProtoMessage()               {}
func (*GetConnectorAddrReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *GetConnectorAddrReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type GetConnectorAddrRep struct {
	ConnectorAddr      string `protobuf:"bytes,1,opt,name=connectorAddr" json:"connectorAddr,omitempty"`
	ConnectorSessionID uint64 `protobuf:"varint,2,opt,name=connectorSessionID" json:"connectorSessionID,omitempty"`
}

func (m *GetConnectorAddrRep) Reset()                    { *m = GetConnectorAddrRep{} }
func (m *GetConnectorAddrRep) String() string            { return proto.CompactTextString(m) }
func (*GetConnectorAddrRep) ProtoMessage()               {}
func (*GetConnectorAddrRep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *GetConnectorAddrRep) GetConnectorAddr() string {
	if m != nil {
		return m.ConnectorAddr
	}
	return ""
}

func (m *GetConnectorAddrRep) GetConnectorSessionID() uint64 {
	if m != nil {
		return m.ConnectorSessionID
	}
	return 0
}

type UpdateConnectorAddrRep struct {
}

func (m *UpdateConnectorAddrRep) Reset()                    { *m = UpdateConnectorAddrRep{} }
func (m *UpdateConnectorAddrRep) String() string            { return proto.CompactTextString(m) }
func (*UpdateConnectorAddrRep) ProtoMessage()               {}
func (*UpdateConnectorAddrRep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

type NewSessionReq struct {
	Id      int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=Address" json:"Address,omitempty"`
	Phone   string `protobuf:"bytes,4,opt,name=Phone" json:"Phone,omitempty"`
}

func (m *NewSessionReq) Reset()                    { *m = NewSessionReq{} }
func (m *NewSessionReq) String() string            { return proto.CompactTextString(m) }
func (*NewSessionReq) ProtoMessage()               {}
func (*NewSessionReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *NewSessionReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NewSessionReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewSessionReq) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NewSessionReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type NewSessionRep struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *NewSessionRep) Reset()                    { *m = NewSessionRep{} }
func (m *NewSessionRep) String() string            { return proto.CompactTextString(m) }
func (*NewSessionRep) ProtoMessage()               {}
func (*NewSessionRep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *NewSessionRep) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RemoveSessionReq struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *RemoveSessionReq) Reset()                    { *m = RemoveSessionReq{} }
func (m *RemoveSessionReq) String() string            { return proto.CompactTextString(m) }
func (*RemoveSessionReq) ProtoMessage()               {}
func (*RemoveSessionReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *RemoveSessionReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RemoveSessionRep struct {
}

func (m *RemoveSessionRep) Reset()                    { *m = RemoveSessionRep{} }
func (m *RemoveSessionRep) String() string            { return proto.CompactTextString(m) }
func (*RemoveSessionRep) ProtoMessage()               {}
func (*RemoveSessionRep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

type RefreshSessionReq struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *RefreshSessionReq) Reset()                    { *m = RefreshSessionReq{} }
func (m *RefreshSessionReq) String() string            { return proto.CompactTextString(m) }
func (*RefreshSessionReq) ProtoMessage()               {}
func (*RefreshSessionReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *RefreshSessionReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RefreshSessionRep struct {
	Token   string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Session *Session `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
}

func (m *RefreshSessionRep) Reset()                    { *m = RefreshSessionRep{} }
func (m *RefreshSessionRep) String() string            { return proto.CompactTextString(m) }
func (*RefreshSessionRep) ProtoMessage()               {}
func (*RefreshSessionRep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *RefreshSessionRep) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RefreshSessionRep) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type UserConnectedReq struct {
	UserID             int64  `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	ConnectorAddress   string `protobuf:"bytes,2,opt,name=connectorAddress" json:"connectorAddress,omitempty"`
	ConnectorSessionID uint64 `protobuf:"varint,3,opt,name=connectorSessionID" json:"connectorSessionID,omitempty"`
}

func (m *UserConnectedReq) Reset()                    { *m = UserConnectedReq{} }
func (m *UserConnectedReq) String() string            { return proto.CompactTextString(m) }
func (*UserConnectedReq) ProtoMessage()               {}
func (*UserConnectedReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *UserConnectedReq) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *UserConnectedReq) GetConnectorAddress() string {
	if m != nil {
		return m.ConnectorAddress
	}
	return ""
}

func (m *UserConnectedReq) GetConnectorSessionID() uint64 {
	if m != nil {
		return m.ConnectorSessionID
	}
	return 0
}

type UserConnectedRep struct {
}

func (m *UserConnectedRep) Reset()                    { *m = UserConnectedRep{} }
func (m *UserConnectedRep) String() string            { return proto.CompactTextString(m) }
func (*UserConnectedRep) ProtoMessage()               {}
func (*UserConnectedRep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

type UserDisonnectedReq struct {
	UserID             int64  `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	ConnectorAddress   string `protobuf:"bytes,2,opt,name=connectorAddress" json:"connectorAddress,omitempty"`
	ConnectorSessionID uint64 `protobuf:"varint,3,opt,name=connectorSessionID" json:"connectorSessionID,omitempty"`
	OnlineTime         int32  `protobuf:"varint,4,opt,name=onlineTime" json:"onlineTime,omitempty"`
}

func (m *UserDisonnectedReq) Reset()                    { *m = UserDisonnectedReq{} }
func (m *UserDisonnectedReq) String() string            { return proto.CompactTextString(m) }
func (*UserDisonnectedReq) ProtoMessage()               {}
func (*UserDisonnectedReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

func (m *UserDisonnectedReq) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *UserDisonnectedReq) GetConnectorAddress() string {
	if m != nil {
		return m.ConnectorAddress
	}
	return ""
}

func (m *UserDisonnectedReq) GetConnectorSessionID() uint64 {
	if m != nil {
		return m.ConnectorSessionID
	}
	return 0
}

func (m *UserDisonnectedReq) GetOnlineTime() int32 {
	if m != nil {
		return m.OnlineTime
	}
	return 0
}

type UserDisonnectedRep struct {
}

func (m *UserDisonnectedRep) Reset()                    { *m = UserDisonnectedRep{} }
func (m *UserDisonnectedRep) String() string            { return proto.CompactTextString(m) }
func (*UserDisonnectedRep) ProtoMessage()               {}
func (*UserDisonnectedRep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{18} }

type RemoveSessionByUIDReq struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *RemoveSessionByUIDReq) Reset()                    { *m = RemoveSessionByUIDReq{} }
func (m *RemoveSessionByUIDReq) String() string            { return proto.CompactTextString(m) }
func (*RemoveSessionByUIDReq) ProtoMessage()               {}
func (*RemoveSessionByUIDReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{19} }

func (m *RemoveSessionByUIDReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type RemoveSessionByUIDRep struct {
}

func (m *RemoveSessionByUIDRep) Reset()                    { *m = RemoveSessionByUIDRep{} }
func (m *RemoveSessionByUIDRep) String() string            { return proto.CompactTextString(m) }
func (*RemoveSessionByUIDRep) ProtoMessage()               {}
func (*RemoveSessionByUIDRep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{20} }

type GetUserIDByTokenReq struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetUserIDByTokenReq) Reset()                    { *m = GetUserIDByTokenReq{} }
func (m *GetUserIDByTokenReq) String() string            { return proto.CompactTextString(m) }
func (*GetUserIDByTokenReq) ProtoMessage()               {}
func (*GetUserIDByTokenReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{21} }

func (m *GetUserIDByTokenReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetUserIDByTokenRep struct {
	Uid int64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *GetUserIDByTokenRep) Reset()                    { *m = GetUserIDByTokenRep{} }
func (m *GetUserIDByTokenRep) String() string            { return proto.CompactTextString(m) }
func (*GetUserIDByTokenRep) ProtoMessage()               {}
func (*GetUserIDByTokenRep) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{22} }

func (m *GetUserIDByTokenRep) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func init() {
	proto.RegisterType((*Session)(nil), "pb.Session")
	proto.RegisterType((*GetSessionByUIDReq)(nil), "pb.GetSessionByUIDReq")
	proto.RegisterType((*GetSessionByUIDRep)(nil), "pb.GetSessionByUIDRep")
	proto.RegisterType((*GetSessionByTokenReq)(nil), "pb.GetSessionByTokenReq")
	proto.RegisterType((*UpdateConnectorAddrReq)(nil), "pb.UpdateConnectorAddrReq")
	proto.RegisterType((*GetSessionByTokenRep)(nil), "pb.GetSessionByTokenRep")
	proto.RegisterType((*GetConnectorAddrReq)(nil), "pb.GetConnectorAddrReq")
	proto.RegisterType((*GetConnectorAddrRep)(nil), "pb.GetConnectorAddrRep")
	proto.RegisterType((*UpdateConnectorAddrRep)(nil), "pb.UpdateConnectorAddrRep")
	proto.RegisterType((*NewSessionReq)(nil), "pb.NewSessionReq")
	proto.RegisterType((*NewSessionRep)(nil), "pb.NewSessionRep")
	proto.RegisterType((*RemoveSessionReq)(nil), "pb.RemoveSessionReq")
	proto.RegisterType((*RemoveSessionRep)(nil), "pb.RemoveSessionRep")
	proto.RegisterType((*RefreshSessionReq)(nil), "pb.RefreshSessionReq")
	proto.RegisterType((*RefreshSessionRep)(nil), "pb.RefreshSessionRep")
	proto.RegisterType((*UserConnectedReq)(nil), "pb.UserConnectedReq")
	proto.RegisterType((*UserConnectedRep)(nil), "pb.UserConnectedRep")
	proto.RegisterType((*UserDisonnectedReq)(nil), "pb.UserDisonnectedReq")
	proto.RegisterType((*UserDisonnectedRep)(nil), "pb.UserDisonnectedRep")
	proto.RegisterType((*RemoveSessionByUIDReq)(nil), "pb.RemoveSessionByUIDReq")
	proto.RegisterType((*RemoveSessionByUIDRep)(nil), "pb.RemoveSessionByUIDRep")
	proto.RegisterType((*GetUserIDByTokenReq)(nil), "pb.GetUserIDByTokenReq")
	proto.RegisterType((*GetUserIDByTokenRep)(nil), "pb.GetUserIDByTokenRep")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserStatus service

type UserStatusClient interface {
	// 通过uid获取session
	GetSessionByUID(ctx context.Context, in *GetSessionByUIDReq, opts ...client.CallOption) (*GetSessionByUIDRep, error)
	// 通过token获取session
	GetSessionByToken(ctx context.Context, in *GetSessionByTokenReq, opts ...client.CallOption) (*GetSessionByTokenRep, error)
	// 获取用户的长连接地址
	GetConnectorAddr(ctx context.Context, in *GetConnectorAddrReq, opts ...client.CallOption) (*GetConnectorAddrRep, error)
	// 更新用户长连接地址（用户建立长连接时调用)（备注：目前有没在用要问下其它服务的开发）
	UpdateConnectorAddr(ctx context.Context, in *UpdateConnectorAddrReq, opts ...client.CallOption) (*UpdateConnectorAddrRep, error)
	// 构建session用户登录时调用，此接口会清除旧session
	NewSession(ctx context.Context, in *NewSessionReq, opts ...client.CallOption) (*NewSessionRep, error)
	// 移除session登出时会调用
	RemoveSession(ctx context.Context, in *RemoveSessionReq, opts ...client.CallOption) (*RemoveSessionRep, error)
	// token续期
	RefreshSession(ctx context.Context, in *RefreshSessionReq, opts ...client.CallOption) (*RefreshSessionRep, error)
	// 更新用户长连接地址（用户建立长连接时调用)
	UserConnected(ctx context.Context, in *UserConnectedReq, opts ...client.CallOption) (*UserConnectedRep, error)
	// 删除用户的长连接地址（用户长连接断开时调用）
	UserDisonnected(ctx context.Context, in *UserDisonnectedReq, opts ...client.CallOption) (*UserDisonnectedRep, error)
	// 通过uid来移除session
	RemoveSessionByUID(ctx context.Context, in *RemoveSessionByUIDReq, opts ...client.CallOption) (*RemoveSessionByUIDRep, error)
	// 通过token找uid
	GetUserIDByToken(ctx context.Context, in *GetUserIDByTokenReq, opts ...client.CallOption) (*GetUserIDByTokenRep, error)
}

type userStatusClient struct {
	c           client.Client
	serviceName string
}

func NewUserStatusClient(serviceName string, c client.Client) UserStatusClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "pb"
	}
	return &userStatusClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userStatusClient) GetSessionByUID(ctx context.Context, in *GetSessionByUIDReq, opts ...client.CallOption) (*GetSessionByUIDRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserStatus.GetSessionByUID", in)
	out := new(GetSessionByUIDRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStatusClient) GetSessionByToken(ctx context.Context, in *GetSessionByTokenReq, opts ...client.CallOption) (*GetSessionByTokenRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserStatus.GetSessionByToken", in)
	out := new(GetSessionByTokenRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStatusClient) GetConnectorAddr(ctx context.Context, in *GetConnectorAddrReq, opts ...client.CallOption) (*GetConnectorAddrRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserStatus.GetConnectorAddr", in)
	out := new(GetConnectorAddrRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStatusClient) UpdateConnectorAddr(ctx context.Context, in *UpdateConnectorAddrReq, opts ...client.CallOption) (*UpdateConnectorAddrRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserStatus.UpdateConnectorAddr", in)
	out := new(UpdateConnectorAddrRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStatusClient) NewSession(ctx context.Context, in *NewSessionReq, opts ...client.CallOption) (*NewSessionRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserStatus.NewSession", in)
	out := new(NewSessionRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStatusClient) RemoveSession(ctx context.Context, in *RemoveSessionReq, opts ...client.CallOption) (*RemoveSessionRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserStatus.RemoveSession", in)
	out := new(RemoveSessionRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStatusClient) RefreshSession(ctx context.Context, in *RefreshSessionReq, opts ...client.CallOption) (*RefreshSessionRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserStatus.RefreshSession", in)
	out := new(RefreshSessionRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStatusClient) UserConnected(ctx context.Context, in *UserConnectedReq, opts ...client.CallOption) (*UserConnectedRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserStatus.UserConnected", in)
	out := new(UserConnectedRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStatusClient) UserDisonnected(ctx context.Context, in *UserDisonnectedReq, opts ...client.CallOption) (*UserDisonnectedRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserStatus.UserDisonnected", in)
	out := new(UserDisonnectedRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStatusClient) RemoveSessionByUID(ctx context.Context, in *RemoveSessionByUIDReq, opts ...client.CallOption) (*RemoveSessionByUIDRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserStatus.RemoveSessionByUID", in)
	out := new(RemoveSessionByUIDRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userStatusClient) GetUserIDByToken(ctx context.Context, in *GetUserIDByTokenReq, opts ...client.CallOption) (*GetUserIDByTokenRep, error) {
	req := c.c.NewRequest(c.serviceName, "UserStatus.GetUserIDByToken", in)
	out := new(GetUserIDByTokenRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserStatus service

type UserStatusHandler interface {
	// 通过uid获取session
	GetSessionByUID(context.Context, *GetSessionByUIDReq, *GetSessionByUIDRep) error
	// 通过token获取session
	GetSessionByToken(context.Context, *GetSessionByTokenReq, *GetSessionByTokenRep) error
	// 获取用户的长连接地址
	GetConnectorAddr(context.Context, *GetConnectorAddrReq, *GetConnectorAddrRep) error
	// 更新用户长连接地址（用户建立长连接时调用)（备注：目前有没在用要问下其它服务的开发）
	UpdateConnectorAddr(context.Context, *UpdateConnectorAddrReq, *UpdateConnectorAddrRep) error
	// 构建session用户登录时调用，此接口会清除旧session
	NewSession(context.Context, *NewSessionReq, *NewSessionRep) error
	// 移除session登出时会调用
	RemoveSession(context.Context, *RemoveSessionReq, *RemoveSessionRep) error
	// token续期
	RefreshSession(context.Context, *RefreshSessionReq, *RefreshSessionRep) error
	// 更新用户长连接地址（用户建立长连接时调用)
	UserConnected(context.Context, *UserConnectedReq, *UserConnectedRep) error
	// 删除用户的长连接地址（用户长连接断开时调用）
	UserDisonnected(context.Context, *UserDisonnectedReq, *UserDisonnectedRep) error
	// 通过uid来移除session
	RemoveSessionByUID(context.Context, *RemoveSessionByUIDReq, *RemoveSessionByUIDRep) error
	// 通过token找uid
	GetUserIDByToken(context.Context, *GetUserIDByTokenReq, *GetUserIDByTokenRep) error
}

func RegisterUserStatusHandler(s server.Server, hdlr UserStatusHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserStatus{hdlr}, opts...))
}

type UserStatus struct {
	UserStatusHandler
}

func (h *UserStatus) GetSessionByUID(ctx context.Context, in *GetSessionByUIDReq, out *GetSessionByUIDRep) error {
	return h.UserStatusHandler.GetSessionByUID(ctx, in, out)
}

func (h *UserStatus) GetSessionByToken(ctx context.Context, in *GetSessionByTokenReq, out *GetSessionByTokenRep) error {
	return h.UserStatusHandler.GetSessionByToken(ctx, in, out)
}

func (h *UserStatus) GetConnectorAddr(ctx context.Context, in *GetConnectorAddrReq, out *GetConnectorAddrRep) error {
	return h.UserStatusHandler.GetConnectorAddr(ctx, in, out)
}

func (h *UserStatus) UpdateConnectorAddr(ctx context.Context, in *UpdateConnectorAddrReq, out *UpdateConnectorAddrRep) error {
	return h.UserStatusHandler.UpdateConnectorAddr(ctx, in, out)
}

func (h *UserStatus) NewSession(ctx context.Context, in *NewSessionReq, out *NewSessionRep) error {
	return h.UserStatusHandler.NewSession(ctx, in, out)
}

func (h *UserStatus) RemoveSession(ctx context.Context, in *RemoveSessionReq, out *RemoveSessionRep) error {
	return h.UserStatusHandler.RemoveSession(ctx, in, out)
}

func (h *UserStatus) RefreshSession(ctx context.Context, in *RefreshSessionReq, out *RefreshSessionRep) error {
	return h.UserStatusHandler.RefreshSession(ctx, in, out)
}

func (h *UserStatus) UserConnected(ctx context.Context, in *UserConnectedReq, out *UserConnectedRep) error {
	return h.UserStatusHandler.UserConnected(ctx, in, out)
}

func (h *UserStatus) UserDisonnected(ctx context.Context, in *UserDisonnectedReq, out *UserDisonnectedRep) error {
	return h.UserStatusHandler.UserDisonnected(ctx, in, out)
}

func (h *UserStatus) RemoveSessionByUID(ctx context.Context, in *RemoveSessionByUIDReq, out *RemoveSessionByUIDRep) error {
	return h.UserStatusHandler.RemoveSessionByUID(ctx, in, out)
}

func (h *UserStatus) GetUserIDByToken(ctx context.Context, in *GetUserIDByTokenReq, out *GetUserIDByTokenRep) error {
	return h.UserStatusHandler.GetUserIDByToken(ctx, in, out)
}

func init() { proto.RegisterFile("userStatus.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xae, 0xdd, 0x3c, 0xd4, 0xa9, 0xd2, 0x3a, 0xd3, 0x34, 0x35, 0x7b, 0x40, 0xd5, 0x8a, 0xd2,
	0xf0, 0x50, 0x0e, 0x85, 0x1b, 0xaa, 0x04, 0x24, 0xa2, 0x8a, 0x84, 0xa0, 0x72, 0x9b, 0x23, 0x87,
	0x24, 0x5e, 0x54, 0xab, 0xc4, 0xde, 0xda, 0x0e, 0x88, 0x1b, 0xa7, 0xfe, 0x16, 0x7e, 0x26, 0x1a,
	0xdb, 0x49, 0x6d, 0xef, 0xba, 0x29, 0x07, 0xc4, 0xcd, 0xfb, 0xcd, 0x37, 0xbb, 0xf3, 0xfc, 0x64,
	0xb0, 0x16, 0x91, 0x08, 0x2f, 0xe2, 0x49, 0xbc, 0x88, 0xfa, 0x32, 0x0c, 0xe2, 0x00, 0x4d, 0x39,
	0xe5, 0x5f, 0xa0, 0x79, 0x21, 0xa2, 0xc8, 0x0b, 0x7c, 0xdc, 0x01, 0xd3, 0x73, 0x6d, 0xe3, 0xd0,
	0xe8, 0xd5, 0x1d, 0xd3, 0x73, 0x11, 0xa1, 0xe6, 0x4f, 0xe6, 0xc2, 0x36, 0x0f, 0x8d, 0xde, 0x96,
	0x93, 0x7c, 0xa3, 0x0d, 0xcd, 0x77, 0xae, 0x1b, 0x8a, 0x28, 0xb2, 0x37, 0x13, 0x78, 0x79, 0xc4,
	0x0e, 0xd4, 0xcf, 0xaf, 0x02, 0x5f, 0xd8, 0xb5, 0x04, 0x4f, 0x0f, 0xfc, 0x29, 0xe0, 0x99, 0x88,
	0xb3, 0x17, 0xde, 0xff, 0x1c, 0x8f, 0x86, 0x8e, 0xb8, 0x41, 0x0b, 0x36, 0x17, 0xd9, 0x53, 0x9b,
	0x0e, 0x7d, 0xf2, 0x37, 0x1a, 0x9e, 0xc4, 0x23, 0x68, 0x46, 0x29, 0x94, 0x70, 0xb7, 0x4f, 0xb6,
	0xfb, 0x72, 0xda, 0xcf, 0x58, 0xce, 0xd2, 0xc6, 0x5f, 0x42, 0x27, 0xef, 0x7c, 0x19, 0x5c, 0x0b,
	0x9f, 0x9e, 0xe9, 0x40, 0x3d, 0xa6, 0xef, 0xc4, 0x79, 0xcb, 0x49, 0x0f, 0xfc, 0x97, 0x01, 0xdd,
	0xb1, 0x74, 0x27, 0xb1, 0x18, 0x04, 0xbe, 0x2f, 0x66, 0x71, 0x10, 0x52, 0x0e, 0xda, 0xb8, 0xf0,
	0x09, 0xb4, 0x66, 0x79, 0x56, 0x56, 0x8c, 0x22, 0x88, 0x7d, 0xc0, 0x15, 0x90, 0x85, 0x31, 0x1a,
	0x26, 0x05, 0xaa, 0x39, 0x1a, 0x0b, 0x3f, 0xd5, 0x06, 0xfc, 0xe0, 0x7c, 0x8f, 0x61, 0xef, 0x4c,
	0xc4, 0xeb, 0xa3, 0xe7, 0xd7, 0x3a, 0xa2, 0x54, 0x93, 0x32, 0x1e, 0x9e, 0x94, 0x59, 0x99, 0x94,
	0x5d, 0x51, 0x56, 0xc9, 0x67, 0xd0, 0xfa, 0x24, 0x7e, 0x2c, 0xd3, 0x10, 0x37, 0xff, 0x64, 0xd2,
	0x8e, 0x8a, 0x8f, 0xc8, 0x8a, 0xee, 0xf7, 0xc0, 0x72, 0xc4, 0x3c, 0xf8, 0x2e, 0x72, 0xe1, 0xe8,
	0x99, 0xa8, 0x30, 0x25, 0x7f, 0x06, 0x6d, 0x47, 0x7c, 0x0d, 0x45, 0x74, 0xb5, 0xd6, 0xfd, 0x5c,
	0xa5, 0x56, 0xc4, 0x94, 0x6f, 0xbb, 0x79, 0x4f, 0xdb, 0x6f, 0x0d, 0xb0, 0xc6, 0x91, 0x08, 0xb3,
	0xfa, 0x0a, 0x97, 0x1e, 0xef, 0x42, 0x83, 0xf6, 0x7a, 0x34, 0xcc, 0xfa, 0x9e, 0x9d, 0xf0, 0x39,
	0x58, 0x85, 0x76, 0x52, 0x1d, 0xd3, 0xf2, 0x2a, 0xf8, 0x5f, 0x8f, 0x2f, 0x2a, 0x71, 0x48, 0xfe,
	0xdb, 0x00, 0x24, 0x70, 0xe8, 0x45, 0xff, 0x39, 0x3c, 0x7c, 0x0c, 0x10, 0xf8, 0xdf, 0x3c, 0x5f,
	0x5c, 0x7a, 0xf3, 0x74, 0x48, 0xea, 0x4e, 0x0e, 0xe1, 0x1d, 0x4d, 0xa4, 0xd4, 0xda, 0xfd, 0x42,
	0xbb, 0xef, 0x11, 0xab, 0x03, 0x3d, 0x55, 0xf2, 0x17, 0xc9, 0xbe, 0x8d, 0x93, 0x14, 0xd7, 0xea,
	0xd0, 0xb1, 0x8e, 0x2c, 0xd5, 0xe7, 0x4e, 0x6e, 0x1b, 0x00, 0xe3, 0x95, 0x76, 0xe3, 0x00, 0x76,
	0x4b, 0x52, 0x89, 0x5d, 0x9a, 0x17, 0x55, 0x67, 0x99, 0x1e, 0x97, 0x7c, 0x03, 0x47, 0xd0, 0x56,
	0x14, 0x08, 0xed, 0x32, 0x7d, 0x99, 0x01, 0xab, 0xb2, 0xd0, 0x55, 0x1f, 0xc0, 0x2a, 0x8b, 0x0c,
	0x1e, 0x64, 0xfc, 0xb2, 0x46, 0xb1, 0x0a, 0x03, 0xdd, 0xf3, 0x19, 0xf6, 0x34, 0xfa, 0x81, 0x8c,
	0x3c, 0xf4, 0x7a, 0xcd, 0xaa, 0x6d, 0x74, 0xe1, 0x6b, 0x80, 0x3b, 0x45, 0xc0, 0x36, 0x71, 0x0b,
	0x32, 0xc4, 0x14, 0x88, 0xbc, 0x4e, 0xa1, 0x55, 0x68, 0x2e, 0x76, 0x88, 0x55, 0xd6, 0x0c, 0xa6,
	0x43, 0xc9, 0xfd, 0x2d, 0xec, 0x14, 0xd7, 0x1e, 0xf7, 0x53, 0x66, 0x49, 0x35, 0x98, 0x16, 0xce,
	0x02, 0x28, 0x6c, 0x57, 0x1a, 0x40, 0x79, 0xf1, 0x99, 0x0e, 0x25, 0xf7, 0x01, 0xec, 0x96, 0xa6,
	0x3b, 0x1d, 0x0f, 0x75, 0x39, 0x99, 0x1e, 0xa7, 0x4b, 0x3e, 0x02, 0xaa, 0x13, 0x8e, 0x8f, 0x94,
	0x9c, 0x57, 0x93, 0x56, 0x69, 0xba, 0x9b, 0x90, 0xc2, 0xa4, 0xaf, 0x26, 0xa4, 0xbc, 0x2c, 0xac,
	0xc2, 0x20, 0xf9, 0xc6, 0xb4, 0x91, 0xfc, 0xb6, 0xbc, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0xed,
	0xec, 0x53, 0x2a, 0xca, 0x08, 0x00, 0x00,
}
