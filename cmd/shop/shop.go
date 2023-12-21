package shop

import (
	"context"
	"github.com/hubzaj/golang-component-test/pkg/runner"
	"github.com/spf13/cobra"
)

func NewShopCommand() *cobra.Command {
	ctx := runner.NewShopServiceContext(context.Background())
	return &cobra.Command{
		Use:   "start-shop-service",
		Short: "start shop service",
		Run: func(cmd *cobra.Command, args []string) {
			runner.StartShop(ctx)
			ctx.Wait()
		},
	}
}
