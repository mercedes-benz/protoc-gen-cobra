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

func NestedClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Nested"),
		Short: "Nested service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_NestedGetCommand(cfg),
		_NestedGetOptionalCommand(cfg),
		_NestedGetDeepCommand(cfg),
		_NestedGetOneOfCommand(cfg),
		_NestedGetOneOfDeepCommand(cfg),
	)
	return cmd
}

func _NestedGetCommand(cfg *client.Config) *cobra.Command {
	req := &NestedRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Get"),
		Short: "Get RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested", "Get"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewNestedClient(cc)
				v := &NestedRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Get(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	_Top := &Top{}
	cmd.PersistentFlags().StringVar(&_Top.Value, cfg.FlagNamer("Top Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Top Value"), func() { req.Top = _Top })
	_Inner := &NestedRequest_Inner{}
	cmd.PersistentFlags().StringVar(&_Inner.Value, cfg.FlagNamer("Inner Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Inner Value"), func() { req.Inner = _Inner })

	return cmd
}

func _NestedGetOptionalCommand(cfg *client.Config) *cobra.Command {
	req := &OptionalRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetOptional"),
		Short: "GetOptional RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested", "GetOptional"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewNestedClient(cc)
				v := &OptionalRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetOptional(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	_Top := &Top{}
	cmd.PersistentFlags().StringVar(&_Top.Value, cfg.FlagNamer("Top Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Top Value"), func() { req.Top = _Top })
	_Inner := &OptionalRequest_Inner{}
	flag.StringPointerVar(cmd.PersistentFlags(), &_Inner.Value, cfg.FlagNamer("Inner Value"), "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Inner Value"), func() { req.Inner = _Inner })

	return cmd
}

func _NestedGetDeepCommand(cfg *client.Config) *cobra.Command {
	req := &DeepRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetDeep"),
		Short: "GetDeep RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested", "GetDeep"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewNestedClient(cc)
				v := &DeepRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetDeep(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	_L0 := &DeepRequest_Outer{}
	_L0_L1 := &DeepRequest_Outer_Middle{}
	_L0_L1_L2 := &DeepRequest_Outer_Middle_Inner{}
	cmd.PersistentFlags().StringVar(&_L0_L1_L2.Value, cfg.FlagNamer("L0 L1 L2 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 L1 L2 Value"), func() { req.L0 = _L0; _L0.L1 = _L0_L1; _L0_L1.L2 = _L0_L1_L2 })

	return cmd
}

func _NestedGetOneOfCommand(cfg *client.Config) *cobra.Command {
	req := &OneOfRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetOneOf"),
		Short: "GetOneOf RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested", "GetOneOf"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewNestedClient(cc)
				v := &OneOfRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetOneOf(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	_Option1 := &OneOfRequest_First{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("Option1"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option1"), func() { req.Choose = &OneOfRequest_Option1{Option1: _Option1} })
	cmd.PersistentFlags().StringVar(&_Option1.Value, cfg.FlagNamer("Option1 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option1 Value"), func() { req.Choose = &OneOfRequest_Option1{Option1: _Option1} })
	_Option2 := &OneOfRequest_Second{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("Option2"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option2"), func() { req.Choose = &OneOfRequest_Option2{Option2: _Option2} })
	cmd.PersistentFlags().StringVar(&_Option2.Value, cfg.FlagNamer("Option2 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option2 Value"), func() { req.Choose = &OneOfRequest_Option2{Option2: _Option2} })
	_Option3 := &OneOfRequest_Third{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("Option3"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option3"), func() { req.Choose = &OneOfRequest_Option3{Option3: _Option3} })
	cmd.PersistentFlags().StringVar(&_Option3.Value, cfg.FlagNamer("Option3 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("Option3 Value"), func() { req.Choose = &OneOfRequest_Option3{Option3: _Option3} })

	return cmd
}

func _NestedGetOneOfDeepCommand(cfg *client.Config) *cobra.Command {
	req := &OneOfDeepRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetOneOfDeep"),
		Short: "GetOneOfDeep RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Nested", "GetOneOfDeep"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewNestedClient(cc)
				v := &OneOfDeepRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetOneOfDeep(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	_L0 := &OneOfDeepRequest_Outer{}
	_L0_Option1 := &OneOfDeepRequest_Outer_First{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option1"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1"), func() { req.L0 = _L0; _L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: _L0_Option1} })
	_L0_Option1_L1 := &OneOfDeepRequest_Outer_Middle{}
	_L0_Option1_L1_L2 := &OneOfDeepRequest_Outer_Middle_Inner{}
	_L0_Option1_L1_L2_Option1 := &OneOfDeepRequest_Outer_Middle_Inner_First{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option1 L1 L2 Option1"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option1"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: _L0_Option1}
		_L0_Option1.L1 = _L0_Option1_L1
		_L0_Option1_L1.L2 = _L0_Option1_L1_L2
		_L0_Option1_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: _L0_Option1_L1_L2_Option1}
	})
	cmd.PersistentFlags().StringVar(&_L0_Option1_L1_L2_Option1.Value, cfg.FlagNamer("L0 Option1 L1 L2 Option1 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option1 Value"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: _L0_Option1}
		_L0_Option1.L1 = _L0_Option1_L1
		_L0_Option1_L1.L2 = _L0_Option1_L1_L2
		_L0_Option1_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: _L0_Option1_L1_L2_Option1}
	})
	_L0_Option1_L1_L2_Option2 := &OneOfDeepRequest_Outer_Middle_Inner_Second{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option1 L1 L2 Option2"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option2"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: _L0_Option1}
		_L0_Option1.L1 = _L0_Option1_L1
		_L0_Option1_L1.L2 = _L0_Option1_L1_L2
		_L0_Option1_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: _L0_Option1_L1_L2_Option2}
	})
	cmd.PersistentFlags().StringVar(&_L0_Option1_L1_L2_Option2.Value, cfg.FlagNamer("L0 Option1 L1 L2 Option2 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option2 Value"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: _L0_Option1}
		_L0_Option1.L1 = _L0_Option1_L1
		_L0_Option1_L1.L2 = _L0_Option1_L1_L2
		_L0_Option1_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: _L0_Option1_L1_L2_Option2}
	})
	_L0_Option1_L1_L2_Option3 := &OneOfDeepRequest_Outer_Middle_Inner_Third{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option1 L1 L2 Option3"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option3"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: _L0_Option1}
		_L0_Option1.L1 = _L0_Option1_L1
		_L0_Option1_L1.L2 = _L0_Option1_L1_L2
		_L0_Option1_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: _L0_Option1_L1_L2_Option3}
	})
	cmd.PersistentFlags().StringVar(&_L0_Option1_L1_L2_Option3.Value, cfg.FlagNamer("L0 Option1 L1 L2 Option3 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option1 L1 L2 Option3 Value"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option1{Option1: _L0_Option1}
		_L0_Option1.L1 = _L0_Option1_L1
		_L0_Option1_L1.L2 = _L0_Option1_L1_L2
		_L0_Option1_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: _L0_Option1_L1_L2_Option3}
	})
	_L0_Option2 := &OneOfDeepRequest_Outer_Second{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option2"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2"), func() { req.L0 = _L0; _L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: _L0_Option2} })
	_L0_Option2_L1 := &OneOfDeepRequest_Outer_Middle{}
	_L0_Option2_L1_L2 := &OneOfDeepRequest_Outer_Middle_Inner{}
	_L0_Option2_L1_L2_Option1 := &OneOfDeepRequest_Outer_Middle_Inner_First{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option2 L1 L2 Option1"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option1"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: _L0_Option2}
		_L0_Option2.L1 = _L0_Option2_L1
		_L0_Option2_L1.L2 = _L0_Option2_L1_L2
		_L0_Option2_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: _L0_Option2_L1_L2_Option1}
	})
	cmd.PersistentFlags().StringVar(&_L0_Option2_L1_L2_Option1.Value, cfg.FlagNamer("L0 Option2 L1 L2 Option1 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option1 Value"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: _L0_Option2}
		_L0_Option2.L1 = _L0_Option2_L1
		_L0_Option2_L1.L2 = _L0_Option2_L1_L2
		_L0_Option2_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: _L0_Option2_L1_L2_Option1}
	})
	_L0_Option2_L1_L2_Option2 := &OneOfDeepRequest_Outer_Middle_Inner_Second{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option2 L1 L2 Option2"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option2"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: _L0_Option2}
		_L0_Option2.L1 = _L0_Option2_L1
		_L0_Option2_L1.L2 = _L0_Option2_L1_L2
		_L0_Option2_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: _L0_Option2_L1_L2_Option2}
	})
	cmd.PersistentFlags().StringVar(&_L0_Option2_L1_L2_Option2.Value, cfg.FlagNamer("L0 Option2 L1 L2 Option2 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option2 Value"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: _L0_Option2}
		_L0_Option2.L1 = _L0_Option2_L1
		_L0_Option2_L1.L2 = _L0_Option2_L1_L2
		_L0_Option2_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: _L0_Option2_L1_L2_Option2}
	})
	_L0_Option2_L1_L2_Option3 := &OneOfDeepRequest_Outer_Middle_Inner_Third{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option2 L1 L2 Option3"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option3"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: _L0_Option2}
		_L0_Option2.L1 = _L0_Option2_L1
		_L0_Option2_L1.L2 = _L0_Option2_L1_L2
		_L0_Option2_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: _L0_Option2_L1_L2_Option3}
	})
	cmd.PersistentFlags().StringVar(&_L0_Option2_L1_L2_Option3.Value, cfg.FlagNamer("L0 Option2 L1 L2 Option3 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option2 L1 L2 Option3 Value"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option2{Option2: _L0_Option2}
		_L0_Option2.L1 = _L0_Option2_L1
		_L0_Option2_L1.L2 = _L0_Option2_L1_L2
		_L0_Option2_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: _L0_Option2_L1_L2_Option3}
	})
	_L0_Option3 := &OneOfDeepRequest_Outer_Third{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option3"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3"), func() { req.L0 = _L0; _L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: _L0_Option3} })
	_L0_Option3_L1 := &OneOfDeepRequest_Outer_Middle{}
	_L0_Option3_L1_L2 := &OneOfDeepRequest_Outer_Middle_Inner{}
	_L0_Option3_L1_L2_Option1 := &OneOfDeepRequest_Outer_Middle_Inner_First{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option3 L1 L2 Option1"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option1"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: _L0_Option3}
		_L0_Option3.L1 = _L0_Option3_L1
		_L0_Option3_L1.L2 = _L0_Option3_L1_L2
		_L0_Option3_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: _L0_Option3_L1_L2_Option1}
	})
	cmd.PersistentFlags().StringVar(&_L0_Option3_L1_L2_Option1.Value, cfg.FlagNamer("L0 Option3 L1 L2 Option1 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option1 Value"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: _L0_Option3}
		_L0_Option3.L1 = _L0_Option3_L1
		_L0_Option3_L1.L2 = _L0_Option3_L1_L2
		_L0_Option3_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option1{Option1: _L0_Option3_L1_L2_Option1}
	})
	_L0_Option3_L1_L2_Option2 := &OneOfDeepRequest_Outer_Middle_Inner_Second{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option3 L1 L2 Option2"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option2"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: _L0_Option3}
		_L0_Option3.L1 = _L0_Option3_L1
		_L0_Option3_L1.L2 = _L0_Option3_L1_L2
		_L0_Option3_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: _L0_Option3_L1_L2_Option2}
	})
	cmd.PersistentFlags().StringVar(&_L0_Option3_L1_L2_Option2.Value, cfg.FlagNamer("L0 Option3 L1 L2 Option2 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option2 Value"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: _L0_Option3}
		_L0_Option3.L1 = _L0_Option3_L1
		_L0_Option3_L1.L2 = _L0_Option3_L1_L2
		_L0_Option3_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option2{Option2: _L0_Option3_L1_L2_Option2}
	})
	_L0_Option3_L1_L2_Option3 := &OneOfDeepRequest_Outer_Middle_Inner_Third{}
	cmd.PersistentFlags().Bool(cfg.FlagNamer("L0 Option3 L1 L2 Option3"), false, "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option3"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: _L0_Option3}
		_L0_Option3.L1 = _L0_Option3_L1
		_L0_Option3_L1.L2 = _L0_Option3_L1_L2
		_L0_Option3_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: _L0_Option3_L1_L2_Option3}
	})
	cmd.PersistentFlags().StringVar(&_L0_Option3_L1_L2_Option3.Value, cfg.FlagNamer("L0 Option3 L1 L2 Option3 Value"), "", "")
	flag.WithPostSetHook(cmd.PersistentFlags(), cfg.FlagNamer("L0 Option3 L1 L2 Option3 Value"), func() {
		req.L0 = _L0
		_L0.Choose = &OneOfDeepRequest_Outer_Option3{Option3: _L0_Option3}
		_L0_Option3.L1 = _L0_Option3_L1
		_L0_Option3_L1.L2 = _L0_Option3_L1_L2
		_L0_Option3_L1_L2.Choose = &OneOfDeepRequest_Outer_Middle_Inner_Option3{Option3: _L0_Option3_L1_L2_Option3}
	})

	return cmd
}
