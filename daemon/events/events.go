package events

const (
	NewTopoheight             string = `new_topo_height`
	NewBlock                  string = `new_block`
	BlockOrdered              string = `block_ordered`
	BlockOrphaned             string = `block_orphaned`
	StableHeightChanged       string = `stable_height_changed`
	StableTopoheightChanged   string = `stable_topo_height_changed`
	TransactionOrphaned       string = `transaction_orphaned`
	TransactionAddedInMempool string = `transaction_added_in_mempool`
	TransactionExecuted       string = `transaction_executed`
	InvokeContract            string = `contract_invoke`
	DeployContract            string = `contract_deploy`
	ContractTransfers         string = `contract_transfers`
	ContractEvent             string = `contract_event`
	NewAsset                  string = `new_asset`
	PeerConnected             string = `peer_connected`
	PeerDisconnected          string = `peer_disconnected`
	PeerStateUpdated          string = `peer_state_updated`
	PeerPeerListUpdated       string = `peer_peer_list_updated`
	PeerPeerDisconnected      string = `peer_peer_disconnected`
	NewBlockTemplate          string = `new_block_template`
)
