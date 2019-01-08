package horde

// Manager keeps track of Nodes and Self.
// Each Manager is self managing.
// This is a holistic network.
type Manager struct {
	// Nodes that are about to be added to the network
	StagedNodes []Node

	// Nodes that all Nodes are aware of
	CommitedNodes []Node

	// Own information - set on boot
	Self Node
}

// Node represents a node in the mesh
type Node struct {
	Domain string
	IP     string
}

// Ping calls another node that is known to exist
func (m *Manager) Ping() {

}

// Nodes returns a list of known nodes
// and their network info in the horde.
// This method will not include Self
func (m *Manager) Nodes() {

}
