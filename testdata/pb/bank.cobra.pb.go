// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	flag "github.com/NathanBaulch/protoc-gen-cobra/flag"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func BankClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Bank"),
		Short: "Bank service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_BankDepositCommand(cfg),
	)
	return cmd
}

func _BankDepositCommand(cfg *client.Config) *cobra.Command {
	req := &DepositRequest{
		ClusterWithNamespaces: &DepositRequest_ClusterWithNamespaces{
			Cluster: &Cluster{},
		},
	}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Deposit"),
		Short: "Deposit RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Bank"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Bank", "Deposit"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewBankClient(cc)
				v := &DepositRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Deposit(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Parent, cfg.FlagNamer("Parent"), "", "")
	cmd.PersistentFlags().StringVar(&req.Tenant, cfg.FlagNamer("Tenant"), "", "")
	cmd.PersistentFlags().StringVar(&req.Environment, cfg.FlagNamer("Environment"), "", "")
	flag.SliceVar(cmd.PersistentFlags(), flag.ParseMessageE[*DepositRequest_ClusterWithNamespaces], &req.Clusters, cfg.FlagNamer("Clusters"), "")
	flag.SliceVar(cmd.PersistentFlags(), flag.ParseMessageE[*DepositRequest_NamespaceWithDeployments], &req.ClusterWithNamespaces.Namespaces, cfg.FlagNamer("ClusterWithNamespaces Namespaces"), "")

	return cmd
}
