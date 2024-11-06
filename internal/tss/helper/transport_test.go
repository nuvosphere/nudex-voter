package helper_test

import (
	"testing"

	"github.com/bnb-chain/tss-lib/v2/tss"
	"github.com/nuvosphere/nudex-voter/internal/tss/helper"
)

func CreateAndConnectTransports(
	t *testing.T,
	partyIDs []*tss.PartyID,
) []*helper.MemoryTransporter {
	// Create transport between peers
	transports := make([]*helper.MemoryTransporter, 0)
	for _, partyID := range partyIDs {
		transports = append(transports, helper.NewMemoryTransporter(partyID))
	}

	t.Logf("transports: %+v", transports)

	// Add transport receivers to each other
	for _, transport := range transports {
		for _, otherTransport := range transports {
			transport.AddTarget(otherTransport.PartyID, otherTransport.GetReceiver())
		}
	}

	t.Logf("transports connected: %+v", transports)

	return transports
}

func CreateAndConnectReSharingTransports(
	t *testing.T,
	oldCommittee []*tss.PartyID,
	newCommittee []*tss.PartyID,
) ([]*helper.MemoryTransporter, []*helper.MemoryTransporter) {
	// Create transport between peers
	oldTransports := make([]*helper.MemoryTransporter, len(oldCommittee))

	for _, partyID := range oldCommittee {
		mt := helper.NewMemoryTransporter(partyID)
		oldTransports = append(oldTransports, mt)
	}

	newTransports := make([]*helper.MemoryTransporter, len(newCommittee))

	for _, partyID := range newCommittee {
		mt := helper.NewMemoryTransporter(partyID)
		newTransports = append(newTransports, mt)
	}

	t.Logf("old transports: %+v", oldTransports)
	t.Logf("new transports: %+v", newTransports)

	// Add old transport receivers to each other
	for _, transport := range oldTransports {
		for _, otherTransport := range oldTransports {
			transport.AddOldCommitteeTarget(otherTransport.PartyID, otherTransport.GetReceiver())
		}

		for _, otherTransport := range newTransports {
			transport.AddNewCommitteeTarget(otherTransport.PartyID, otherTransport.GetReceiver())
		}
	}

	t.Logf("old transports connected")

	// Add new transport receivers to each other
	for _, transport := range newTransports {
		for _, otherTransport := range oldTransports {
			transport.AddOldCommitteeTarget(otherTransport.PartyID, otherTransport.GetReceiver())
		}

		for _, otherTransport := range newTransports {
			transport.AddNewCommitteeTarget(otherTransport.PartyID, otherTransport.GetReceiver())
		}
	}

	t.Logf("new transports connected")

	return oldTransports, newTransports
}
