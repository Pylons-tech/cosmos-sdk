package client

import (
	govclient "github.com/Pylons-tech/cosmos-sdk/x/gov/client"
	"github.com/Pylons-tech/cosmos-sdk/x/upgrade/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal)
