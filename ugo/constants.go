package ugo

import (
	"time"
)

const (
	packetInit = 0
)

type encryptMethod byte

const (
	aesEncrypt = encryptMethod(iota)
)

const (
	fecHeaderSize = 6
	typeData      = 0xf1
	typeFEC       = 0xf2
	fecExpire     = 30000 // 30s
)

// IPv4 minimum reassembly buffer size is 576, IPv6 has it at 1500.
// Subtract header sizes from here

const minRetransmissionTime = 200 * time.Millisecond
const defaultRetransmissionTime = 500 * time.Millisecond
const maxCongestionWindow uint32 = 200

// DefaultMaxCongestionWindow is the default for the max congestion window
// Taken from Chrome
const DefaultMaxCongestionWindow uint32 = 107

// InitialCongestionWindow is the initial congestion window in QUIC packets
const InitialCongestionWindow uint32 = 32

// InitialIdleConnectionStateLifetime is the initial idle connection state lifetime
const InitialIdleConnectionStateLifetime = 30 * time.Second
