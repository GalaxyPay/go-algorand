// Copyright (C) 2019-2022 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package ledgercore

import (
	"github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/data/basics"
	"github.com/algorand/go-algorand/protocol"
)

// SPVerificationContext represents the context provided by the ledger to verify a state proof transaction.
type SPVerificationContext struct {
	_struct struct{} `codec:",omitempty,omitemptyarray"`

	// LastAttestedRound is the last attested round of the state proof verified using this data.
	LastAttestedRound basics.Round `codec:"spround"`

	// VotersCommitment is the vector commitment root of the top N accounts to sign the next state proof.
	VotersCommitment crypto.GenericDigest `codec:"vc"`

	// OnlineTotalWeight is the total amount of stake attesting to the next state proof.
	OnlineTotalWeight basics.MicroAlgos `codec:"pw"`

	// Version is the protocol version that would be used to verify the state proof
	Version protocol.ConsensusVersion `codec:"v"`
}
