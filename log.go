package eddnc

import (
	"git.fractalqb.de/fractalqb/c4hgol"
	"git.fractalqb.de/fractalqb/qbsllm"
	"github.com/CmdrVasquess/eddnc/subscriber"
)

var (
	log    = qbsllm.New(qbsllm.Lnormal, "eddn", nil, nil)
	LogCfg = c4hgol.Config(qbsllm.NewConfig(log), subscriber.LogCfg)
)
