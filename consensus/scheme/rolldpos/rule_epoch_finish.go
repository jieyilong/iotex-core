// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided ‘as is’ and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package rolldpos

import (
	"github.com/iotexproject/iotex-core/consensus/fsm"
	"github.com/iotexproject/iotex-core/delegate"
)

type ruleEpochFinish struct {
	*RollDPoS
}

func (r ruleEpochFinish) Condition(event *fsm.Event) bool {
	if event.State != stateEpochStart {
		return false
	}
	height, err := r.bc.TipHeight()
	if err != nil {
		event.Err = err
		return false
	}
	// if the height of the last committed block is already the last one should be minted from this epochStart, go back to
	// epochStart start
	if height >= r.epochCtx.height+uint64(uint(len(r.epochCtx.delegates))*r.epochCtx.numSubEpochs)-1 {
		if err := startNewEpochIfNecessary(r.RollDPoS); err != nil {
			event.Err = err
			return false
		}
		return true
	}
	return false
}

func startNewEpochIfNecessary(c *RollDPoS) error {
	ok, err := c.epochStartCb(c.pool)
	if err != nil {
		return err
	}
	if ok {
		c.enqueueEvent(&fsm.Event{
			State: stateDKGGenerate,
		})
	}
	return nil
}

// NeverStartNewEpoch will never allow to start a new epochStart after the first one
func NeverStartNewEpoch(_ delegate.Pool) (bool, error) { return false, nil }

// PseudoStarNewEpoch will always allow to start a new epochStart after the first one
func PseudoStarNewEpoch(_ delegate.Pool) (bool, error) { return true, nil }
