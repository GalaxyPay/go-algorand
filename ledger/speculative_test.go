// Copyright (C) 2019-2021 Algorand, Inc.
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

package ledger

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/data/basics"
	"github.com/algorand/go-algorand/data/bookkeeping"
	"github.com/algorand/go-algorand/data/transactions"
	ledgertesting "github.com/algorand/go-algorand/ledger/testing"
	"github.com/algorand/go-algorand/logging"
	"github.com/algorand/go-algorand/protocol"
)

func TestSpeculative(t *testing.T) {
	genesisInitState, _ := ledgertesting.GenerateInitState(t, protocol.ConsensusCurrentVersion, 1000)
	const inMem = true
	cfg := config.GetDefaultLocal()
	log := logging.TestingLog(t)
	l, err := OpenLedger(log, t.Name(), inMem, genesisInitState, cfg)
	require.NoError(t, err, "could not open ledger")
	defer l.Close()

	blk0, err := l.BlockHdr(l.Latest())
	require.NoError(t, err)

	var blk1 bookkeeping.Block
	blk1.CurrentProtocol = protocol.ConsensusCurrentVersion
	blk1.Branch = blk0.Hash()
	blk1.RewardsPool = testPoolAddr
	blk1.FeeSink = testSinkAddr
	blk1.BlockHeader.GenesisHash = genesisInitState.GenesisHash
	blk1.BlockHeader.Round = l.Latest() + 1

	//sl, err := MakeSpeculativeLedger(l)
	//require.NoError(t, err)

	require.NoError(t, err)

	blk1aslfe, statedelta, err := MakeBlockAsLFE(blk1, l)
	require.NoError(t, err)
	require.Equal(t, blk1aslfe.vb.Delta(), statedelta)

	blk2 := blk1
	blk2.BlockHeader.Round++
	blk2.Branch = blk1.Hash()

	// Pick some accounts at random
	var addr1, addr2 basics.Address
	for a := range genesisInitState.Accounts {
		if addr1 == (basics.Address{}) {
			addr1 = a
		} else if addr2 == (basics.Address{}) {
			addr2 = a
		} else {
			break
		}
	}

	var tx21 transactions.Transaction
	tx21.Type = protocol.PaymentTx
	tx21.Sender = addr1
	tx21.Receiver = addr2
	tx21.FirstValid = blk2.BlockHeader.Round
	tx21.LastValid = blk2.BlockHeader.Round
	tx21.Amount.Raw = 1000000
	blk2.Payset = append(blk2.Payset, transactions.SignedTxnInBlock{
		SignedTxnWithAD: transactions.SignedTxnWithAD{
			SignedTxn: transactions.SignedTxn{
				Txn: tx21,
			},
		},
		HasGenesisID: true,
	})

	blk2aslfe, statedelta, err := MakeBlockAsLFE(blk2, blk1aslfe)
	require.NoError(t, err)
	require.Equal(t, blk2aslfe.vb.Delta(), statedelta)

	ad11, rnd, err := blk2aslfe.LookupWithoutRewards(blk1.Round(), addr1)
	require.NoError(t, err)
	// account was never changed
	require.Equal(t, rnd, basics.Round(0))

	ad22, rnd, err := blk2aslfe.LookupWithoutRewards(blk2.Round(), addr1)
	require.NoError(t, err)
	// account changed at blk2
	require.Equal(t, rnd, blk2.Round())

	require.Equal(t, ad22.MicroAlgos.Raw, ad11.MicroAlgos.Raw-1000000)
}
