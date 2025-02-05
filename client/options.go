package client

import (
	"github.com/WuKongIM/wkrpc/proto"
)

type Options struct {
	Addr  string
	Uid   string
	Token string

	HeartbeatTick        int  // 心跳间隔tick数
	HeartbeatTimeoutTick int  // 心跳超时tick数
	MessagePoolOn        bool // Whether to open the message pool, the default is true
	LogDetailOn          bool // 是否开启详细日志
	OnMessage            func(msg *proto.Message)
	MessagePoolSize      int // The size of the message pool, the default is 10000
}

func NewOptions(opt ...Option) *Options {
	return &Options{
		HeartbeatTick:        10,
		HeartbeatTimeoutTick: 20,
		MessagePoolOn:        true,
		MessagePoolSize:      2000,
	}
}

type Option func(*Options)

func WithUid(uid string) Option {
	return func(opts *Options) {
		opts.Uid = uid
	}
}

func WithToken(token string) Option {
	return func(opts *Options) {
		opts.Token = token
	}
}

func WithHeartbeatTick(tick int) Option {
	return func(opts *Options) {
		opts.HeartbeatTick = tick
	}
}

func WithHeartbeatTimeoutTick(tick int) Option {
	return func(opts *Options) {
		opts.HeartbeatTimeoutTick = tick
	}
}

func WithLogDetailOn(on bool) Option {
	return func(opts *Options) {
		opts.LogDetailOn = on
	}
}

func WithOnMessage(on func(msg *proto.Message)) Option {
	return func(opts *Options) {
		opts.OnMessage = on
	}
}

func WithMessagePoolOn(on bool) Option {
	return func(opts *Options) {
		opts.MessagePoolOn = on
	}
}

func WithMessagePoolSize(size int) Option {
	return func(opts *Options) {
		opts.MessagePoolSize = size
	}
}
