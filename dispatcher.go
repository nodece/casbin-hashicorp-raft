package hraftdispatcher

import (
	"context"
	"crypto/tls"
	"github.com/hashicorp/go-multierror"

	"github.com/casbin/casbin/v2/persist"
	"github.com/hashicorp/raft"
	"github.com/nodece/casbin-hraft-dispatcher/command"
	"github.com/nodece/casbin-hraft-dispatcher/http"
	"github.com/nodece/casbin-hraft-dispatcher/store"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var _ persist.Dispatcher = &HRaftDispatcher{}

// HRaftDispatcher implements the persist.Dispatcher interface.
type HRaftDispatcher struct {
	store       http.Store
	tlsConfig   *tls.Config
	httpService *http.Service
	shutdownFn  func() error

	logger *zap.Logger
}

// NewHRaftDispatcher returns a HRaftDispatcher.
func NewHRaftDispatcher(config *Config) (*HRaftDispatcher, error) {
	if config == nil {
		return nil, errors.New("config is not provided")
	}

	if config.Enforcer == nil {
		return nil, errors.New("Enforcer is not provided in config")
	}

	if len(config.DataDir) == 0 {
		return nil, errors.New("DataDir is not provided in config")
	}

	if len(config.RaftListenAddress) == 0 {
		return nil, errors.New("RaftListenAddress is not provided in config")
	}

	if config.TLSConfig == nil {
		return nil, errors.New("TLSConfig is not provided in config")
	}

	httpListenAddress, err := http.ConvertRaftAddressToHTTPAddress(config.RaftListenAddress)
	if err != nil {
		return nil, err
	}

	if len(config.ServerID) == 0 {
		config.ServerID = config.RaftListenAddress
	}

	logger := zap.NewExample()

	streamLayer, err := store.NewTCPStreamLayer(config.RaftListenAddress, config.TLSConfig)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	storeConfig := &store.Config{
		ID:  config.ServerID,
		Dir: config.DataDir,
		NetworkTransportConfig: &raft.NetworkTransportConfig{
			Stream:  streamLayer,
			MaxPool: 5,
			Logger:  nil,
		},
		Enforcer: config.Enforcer,
	}
	s, err := store.NewStore(storeConfig)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	isNewCluster := !s.IsInitializedCluster()
	enableBootstrap := false

	if isNewCluster == true {
		enableBootstrap = true
	}

	if len(config.JoinAddress) != 0 {
		enableBootstrap = false
	}

	if enableBootstrap {
		logger.Info("bootstrapping a new cluster")
	} else {
		logger.Info("skip bootstrapping a new cluster")
	}

	err = s.Start(enableBootstrap)
	if err != nil {
		logger.Error("failed to start raft service", zap.Error(err))
		return nil, err
	}

	if enableBootstrap {
		err = s.WaitLeader()
		if err != nil {
			logger.Error(err.Error())
		}
	}

	if isNewCluster && config.JoinAddress != config.RaftListenAddress && len(config.JoinAddress) != 0 {
		entryAddress, err := http.ConvertRaftAddressToHTTPAddress(config.JoinAddress)
		if err != nil {
			logger.Error("failed to convert the Raft address to HTTP address", zap.String("nodeID", config.ServerID), zap.String("nodeAddress", config.RaftListenAddress), zap.String("clusterAddress", config.JoinAddress), zap.Error(err))
			return nil, err
		}
		err = http.DoJoinNodeRequest(entryAddress, config.ServerID, config.RaftListenAddress, config.TLSConfig)
		if err != nil {
			logger.Error("failed to join the current node to existing cluster", zap.String("nodeID", config.ServerID), zap.String("nodeAddress", config.RaftListenAddress), zap.String("clusterAddress", config.JoinAddress), zap.Error(err))
			return nil, err
		}
	}

	httpService, err := http.NewService(httpListenAddress, config.TLSConfig, s)
	if err != nil {
		return nil, err
	}

	err = httpService.Start()
	if err != nil {
		return nil, err
	}

	h := &HRaftDispatcher{
		store:       s,
		tlsConfig:   config.TLSConfig,
		httpService: httpService,
		logger:      logger,
	}

	h.shutdownFn = func() error {
		var ret error
		err := s.Stop()
		if err != nil {
			ret = multierror.Append(ret, err)
		}
		err = httpService.Stop(context.Background())
		if err != nil {
			ret = multierror.Append(ret, err)
		}
		return ret
	}

	return h, nil
}

//AddPolicies implements the persist.Dispatcher interface.
func (h *HRaftDispatcher) AddPolicies(sec string, pType string, rules [][]string) error {
	var items []*command.StringArray
	for _, rule := range rules {
		var item = &command.StringArray{Items: rule}
		items = append(items, item)
	}

	addPolicyRequest := &command.AddPoliciesRequest{
		Sec:   sec,
		PType: pType,
		Rules: items,
	}
	return h.httpService.DoAddPolicyRequest(addPolicyRequest)
}

// RemovePolicies implements the persist.Dispatcher interface.
func (h *HRaftDispatcher) RemovePolicies(sec string, pType string, rules [][]string) error {
	var items []*command.StringArray
	for _, rule := range rules {
		var item = &command.StringArray{Items: rule}
		items = append(items, item)
	}

	request := &command.RemovePoliciesRequest{
		Sec:   sec,
		PType: pType,
		Rules: items,
	}
	return h.httpService.DoRemovePolicyRequest(request)
}

// RemoveFilteredPolicy implements the persist.Dispatcher interface.
func (h *HRaftDispatcher) RemoveFilteredPolicy(sec string, pType string, fieldIndex int, fieldValues ...string) error {
	request := &command.RemoveFilteredPolicyRequest{
		Sec:         sec,
		PType:       pType,
		FieldIndex:  int32(fieldIndex),
		FieldValues: fieldValues,
	}
	return h.httpService.DoRemoveFilteredPolicyRequest(request)
}

// ClearPolicy implements the persist.Dispatcher interface.
func (h *HRaftDispatcher) ClearPolicy() error {
	return h.httpService.DoClearPolicyRequest()
}

// UpdatePolicy implements the persist.Dispatcher interface.
func (h *HRaftDispatcher) UpdatePolicy(sec string, pType string, oldRule, newRule []string) error {
	request := &command.UpdatePolicyRequest{
		Sec:     sec,
		PType:   pType,
		OldRule: oldRule,
		NewRule: newRule,
	}
	return h.httpService.DoUpdatePolicyRequest(request)
}

// UpdatePolicies implements the persist.Dispatcher interface.
func (h *HRaftDispatcher) UpdatePolicies(sec string, pType string, oldRules, newRules [][]string) error {
	var olds []*command.StringArray
	for _, rule := range oldRules {
		var item = &command.StringArray{Items: rule}
		olds = append(olds, item)
	}

	var news []*command.StringArray
	for _, rule := range newRules {
		var item = &command.StringArray{Items: rule}
		news = append(news, item)
	}

	request := &command.UpdatePoliciesRequest{
		Sec:      sec,
		PType:    pType,
		OldRules: olds,
		NewRules: news,
	}
	return h.httpService.DoUpdatePoliciesRequest(request)
}

// JoinNode joins a node to the current cluster.
func (h *HRaftDispatcher) JoinNode(serverID, serverAddress string) error {
	request := &command.AddNodeRequest{
		Id:      serverID,
		Address: serverAddress,
	}
	return h.httpService.DoJoinNodeRequest(request)
}

// JoinNode joins a node from the current cluster.
func (h *HRaftDispatcher) RemoveNode(serverID string) error {
	request := &command.RemoveNodeRequest{
		Id: serverID,
	}
	return h.httpService.DoRemoveNodeRequest(request)
}

// Shutdown is used to close the http and raft service.
func (h *HRaftDispatcher) Shutdown() error {
	return h.shutdownFn()
}
