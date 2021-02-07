package hraftdispatcher

import (
	"crypto/tls"
	"github.com/casbin/casbin/v2"
	"github.com/hashicorp/raft"
)

// Config holds dispatcher config.
type Config struct {
	// Enforcer is a enforcer of casbin.
	Enforcer casbin.IDistributedEnforcer
	// TLSConfig is used to configure a TLS server and client.
	TLSConfig *tls.Config
	// RaftAddress is a network address for raft server.
	RaftAddress string
	// DataDir holds raft data, default to the DefaultDataDir.
	DataDir string
	// ServerID is a unique string identifying this server for all time, default to the RaftAddress.
	ServerID string
	// RaftConfig is hashicorp-raft configuration, default to the raft.DefaultConfig().
	RaftConfig *raft.Config
	// HttpAddress is a network address for dispatcher backend, default to the DefaultHttpAddress.
	HttpAddress string
}
