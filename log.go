package eddnc

import (
	"git.fractalqb.de/fractalqb/qbsllm"
)

var (
	log    = qbsllm.New(qbsllm.Lnormal, "eddnc", nil, nil)
	LogCfg = qbsllm.NewConfig(log)
)
