package hashes

import (
	"crypto/sha256"

	"golang.org/x/crypto/sha3"
	"lukechampine.com/blake3"
)

const (
	transcationHashDomain         = "TransactionHash"
	transcationIDDomain           = "TransactionID"
	transcationSigningDomain      = "TransactionSigningHash"
	transcationSigningECDSADomain = "TransactionSigningHashECDSA"
	blockDomain                   = "BlockHash"
	proofOfWorkDomain             = "ProofOfWorkHash"
	heavyHashDomain               = "HeavyHash"
	merkleBranchDomain            = "MerkleBranchHash"
)

// transactionSigningECDSADomainHash is a hashed version of transcationSigningECDSADomain that is used
// to make it a constant size. This is needed because this domain is used by sha256 hash writer, and
// sha256 doesn't support variable size domain separation.
var transactionSigningECDSADomainHash = sha256.Sum256([]byte(transcationSigningECDSADomain))

// NewTransactionHashWriter Returns a new HashWriter used for transaction hashes
func NewTransactionHashWriter() HashWriter {
	blake := blake3.New(32, nil)
	return HashWriter{blake}
}

// NewTransactionIDWriter Returns a new HashWriter used for transaction IDs
func NewTransactionIDWriter() HashWriter {
	blake := blake3.New(32, nil)
	return HashWriter{blake}
}

// NewTransactionSigningHashWriter Returns a new HashWriter used for signing on a transaction
func NewTransactionSigningHashWriter() HashWriter {
	blake := blake3.New(32, nil)
	return HashWriter{blake}
}

// NewTransactionSigningHashECDSAWriter Returns a new HashWriter used for signing on a transaction with ECDSA
func NewTransactionSigningHashECDSAWriter() HashWriter {
	hashWriter := HashWriter{sha256.New()}
	hashWriter.InfallibleWrite(transactionSigningECDSADomainHash[:])
	return hashWriter
}

// NewBlockHashWriter Returns a new HashWriter used for hashing blocks
func NewBlockHashWriter() HashWriter {
	blake := blake3.New(32, nil)
	return HashWriter{blake}
}

// NewPoWHashWriter Returns a new HashWriter used for the PoW function
func NewPoWHashWriter() Blake3HashWriter {
	blake := blake3.New(32, nil)
	return Blake3HashWriter{blake}
}

// NewHeavyHashWriter Returns a new HashWriter used for the HeavyHash function
func NewHeavyHashWriter() ShakeHashWriter {
	shake256 := sha3.NewCShake256(nil, []byte(heavyHashDomain))
	return ShakeHashWriter{shake256}
}

// NewMerkleBranchHashWriter Returns a new HashWriter used for a merkle tree branch
func NewMerkleBranchHashWriter() HashWriter {
	blake := blake3.New(32, nil)
	return HashWriter{blake}
}
