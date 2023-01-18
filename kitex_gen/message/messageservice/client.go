// Code generated by Kitex v0.3.4. DO NOT EDIT.

package messageservice

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/message"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessageChat(ctx context.Context, Req *message.DouyinMessageChatRequest, callOptions ...callopt.Option) (r *message.DouyinMessageChatResponse, err error)
	MessageAction(ctx context.Context, Req *message.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *message.DouyinRelationActionResponse, err error)
}

// NewClient creates a rpc for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kMessageServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a rpc for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMessageServiceClient struct {
	*kClient
}

func (p *kMessageServiceClient) MessageChat(ctx context.Context, Req *message.DouyinMessageChatRequest, callOptions ...callopt.Option) (r *message.DouyinMessageChatResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageChat(ctx, Req)
}

func (p *kMessageServiceClient) MessageAction(ctx context.Context, Req *message.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *message.DouyinRelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageAction(ctx, Req)
}