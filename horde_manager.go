package horde

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// Manager keeps track of Nodes and Self.
// Each Manager is self managing.
// This is a holistic network.
type Manager struct {
	sync.Mutex

	// Nodes that are about to be added to the horde
	StagedNodes []Node

	// StageSha is a unique sha that represents the latest StagedNodes
	StageSha string

	// Nodes that have have been committed in the horde
	CommittedNodes []Node

	// CommitSha is a unique sha that represents the latest CommittedNodes
	CommitSha string

	// ReadEpoch is a simple mechanism to keep track of pending or current reads.
	// If the ReadEpoch is odd, the Manager is reading.
	// If the ReadEpoch is even, the Manager has finished reading.
	ReadEpoch int

	// WriteEpoch is a simple mechanism to keep track of pending or current writes.
	// If the WriteEpoch is odd, the Manager is writting.
	// If the WriteEpoch is even, the Manager has finished writting.
	WriteEpoch int

	// Self contains static information about the node.
	// This information is set on boot.
	Self Node
}

// Ping calls another node that is known to exist.
// If properly load balanced (round robin) it should not matter which node we ping.
// So we ping the first!
func (m *Manager) Ping() []byte {
	nodes := m.Nodes()
	committedNodesLen := len(nodes)

	if committedNodesLen > 0 {
		node := nodes[0]

		res, err := http.Get(node.LocalIP)

		if err != nil {
			fmt.Print("Target horde node is either unhealthy or down!", err)
		}

		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			_, err := ioutil.ReadAll(res.Body)

			if err != nil {
				fmt.Print("Failed to read body", err)

				return []byte("pang")
			}

			return []byte("pong")
		}
	}

	return []byte("pang")
}

// Nodes returns a list of known nodes
// and their network info in the horde.
// This method will not include Self
func (m *Manager) Nodes() []Node {
	return m.CommittedNodes
}

// RemoveSelfFromHorde is for when the http server fails or is shutdown.
// Something needs to happen. Gotta clean up the mess.
// This _will_ make a network call to a known node and remove itself from the horde!
func (m *Manager) RemoveSelfFromHorde() {
	log.Print("Horde has crashed - please send help")
}
