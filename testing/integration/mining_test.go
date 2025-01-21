package integration

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Hoosat-Oy/HTND/app/appmessage"
	"github.com/Hoosat-Oy/HTND/domain/consensus/model/externalapi"
	"github.com/Hoosat-Oy/HTND/domain/consensus/utils/mining"
)

func mineNextBlock(t *testing.T, harness *appHarness) *externalapi.DomainBlock {
	blockTemplate, err := harness.rpcClient.GetBlockTemplate(harness.miningAddress, "integration")
	if err != nil {
		t.Fatalf("Error getting block template: %+v", err)
	}

	powHash, _ := externalapi.NewDomainHashFromString("REAL_MAIN_POW_HASH")
	block, err := appmessage.RPCBlockToDomainBlock(blockTemplate.Block, powHash)
	if err != nil {
		t.Fatalf("Error converting block: %s", err)
	}

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	_, powHash = mining.SolveBlock(block, rd)
	block.PoWHash = powHash
	_, err = harness.rpcClient.SubmitBlockAlsoIfNonDAA(block, block.PoWHash)
	if err != nil {
		t.Fatalf("Error submitting block: %s", err)
	}

	return block
}
