package model

import "github.com/Hoosat-Oy/HTND/domain/consensus/model/externalapi"

// BlockProcessor is responsible for processing incoming blocks
type BlockProcessor interface {
	ValidateAndInsertBlock(block *externalapi.DomainBlock, shouldValidateAgainstUTXO bool, powSkip bool) (*externalapi.VirtualChangeSet, externalapi.BlockStatus, error)
	ValidateAndInsertImportedPruningPoint(newPruningPoint *externalapi.DomainHash) error
	ValidateAndInsertBlockWithTrustedData(block *externalapi.BlockWithTrustedData, validateUTXO bool) (*externalapi.VirtualChangeSet, externalapi.BlockStatus, error)
}
