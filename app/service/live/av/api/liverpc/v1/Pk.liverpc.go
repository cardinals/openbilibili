// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/Pk.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// ============
// Pk Interface
// ============

type Pk interface {
	// * 根据id获取PK基础信息
	//
	GetInfoById(context.Context, *PkGetInfoByIdReq) (*PkGetInfoByIdResp, error)

	// * 根据id获取PK基础信息
	//
	GetPkStatus(context.Context, *PkGetPkStatusReq) (*PkGetPkStatusResp, error)

	// * 批量获取pkIds
	//
	GetPkIdsByRoomIds(context.Context, *PkGetPkIdsByRoomIdsReq) (*PkGetPkIdsByRoomIdsResp, error)
}

// ==================
// Pk Live Rpc Client
// ==================

type pkRpcClient struct {
	client *liverpc.Client
}

// NewPkRpcClient creates a Rpc client that implements the Pk interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewPkRpcClient(client *liverpc.Client) Pk {
	return &pkRpcClient{
		client: client,
	}
}

func (c *pkRpcClient) GetInfoById(ctx context.Context, in *PkGetInfoByIdReq) (*PkGetInfoByIdResp, error) {
	out := new(PkGetInfoByIdResp)
	err := doRpcRequest(ctx, c.client, 1, "Pk.getInfoById", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pkRpcClient) GetPkStatus(ctx context.Context, in *PkGetPkStatusReq) (*PkGetPkStatusResp, error) {
	out := new(PkGetPkStatusResp)
	err := doRpcRequest(ctx, c.client, 1, "Pk.getPkStatus", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pkRpcClient) GetPkIdsByRoomIds(ctx context.Context, in *PkGetPkIdsByRoomIdsReq) (*PkGetPkIdsByRoomIdsResp, error) {
	out := new(PkGetPkIdsByRoomIdsResp)
	err := doRpcRequest(ctx, c.client, 1, "Pk.getPkIdsByRoomIds", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}