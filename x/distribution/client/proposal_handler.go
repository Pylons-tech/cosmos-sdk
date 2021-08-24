package client

import (
	"github.com/Pylons-tech/cosmos-sdk/x/distribution/client/cli"
	govclient "github.com/Pylons-tech/cosmos-sdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal)
)
