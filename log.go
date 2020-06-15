package eddn

import (
	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"
	"github.com/CmdrVasquess/goEDDNc/down"
)

var (
	log    = qbsllm.New(qbsllm.Lnormal, "eddn", nil, nil)
	LogCfg = c4hgol.Config(qbsllm.NewConfig(log), down.LogCfg)
)
