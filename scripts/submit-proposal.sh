#!/bin/bash
set -e

CHAIN_ID="velora-testnet-1"
FROM="validatorA"
HOME_DIR="$HOME/.velorachain"

# 1. Submit proposal
echo "Submitting proposal..."
velorachaind tx gov submit-proposal scripts/proposals/text-proposal.json   --from $FROM --chain-id $CHAIN_ID --home $HOME_DIR --keyring-backend test -y

# 2. Deposit
echo "Depositing..."
velorachaind tx gov deposit 1 1000000aphoton   --from $FROM --chain-id $CHAIN_ID --home $HOME_DIR --keyring-backend test -y

# 3. Vote
echo "Voting YES..."
velorachaind tx gov vote 1 yes   --from $FROM --chain-id $CHAIN_ID --home $HOME_DIR --keyring-backend test -y

echo "Done. Query with: velorachaind query gov proposal 1"
