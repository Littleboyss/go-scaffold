package greet

import (
	"github.com/casbin/casbin/v2"
	"github.com/spf13/cobra"
)

type Service struct {
	Enforcer *casbin.Enforcer
}

func (h handler) AddPolicy(cmd *cobra.Command, args []string) {
	h.logger.Infof("%s 命令调用成功", cmd.Use)
}
