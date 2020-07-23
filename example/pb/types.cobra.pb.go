// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	context "context"
	tls "crypto/tls"
	x509 "crypto/x509"
	fmt "fmt"
	flag "github.com/NathanBaulch/protoc-gen-cobra/flag"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	oauth2 "golang.org/x/oauth2"
	grpc "google.golang.org/grpc"
	credentials "google.golang.org/grpc/credentials"
	oauth "google.golang.org/grpc/credentials/oauth"
	ioutil "io/ioutil"
	net "net"
	os "os"
	filepath "path/filepath"
	strconv "strconv"
	strings "strings"
	time "time"
)

var TypesClientDefaultConfig = &_TypesClientConfig{
	ServerAddr:     "localhost:8080",
	ResponseFormat: "json",
	Timeout:        10 * time.Second,
	AuthTokenType:  "Bearer",
}

type _TypesClientConfig struct {
	ServerAddr         string
	RequestFile        string
	Stdin              bool
	ResponseFormat     string
	Timeout            time.Duration
	TLS                bool
	ServerName         string
	InsecureSkipVerify bool
	CACertFile         string
	CertFile           string
	KeyFile            string
	AuthToken          string
	AuthTokenType      string
	JWTKey             string
	JWTKeyFile         string
}

func (o *_TypesClientConfig) addFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.ServerAddr, "server-addr", "s", o.ServerAddr, "server address in form of host:port")
	fs.StringVarP(&o.RequestFile, "request-file", "f", o.RequestFile, "client request file (must be json, yaml, or xml); use \"-\" for stdin + json")
	fs.BoolVar(&o.Stdin, "stdin", o.Stdin, "read client request from STDIN; alternative for '-f -'")
	fs.StringVarP(&o.ResponseFormat, "response-format", "o", o.ResponseFormat, "response format (json, prettyjson, xml, prettyxml, or yaml)")
	fs.DurationVar(&o.Timeout, "timeout", o.Timeout, "client connection timeout")
	fs.BoolVar(&o.TLS, "tls", o.TLS, "enable tls")
	fs.StringVar(&o.ServerName, "tls-server-name", o.ServerName, "tls server name override")
	fs.BoolVar(&o.InsecureSkipVerify, "tls-insecure-skip-verify", o.InsecureSkipVerify, "INSECURE: skip tls checks")
	fs.StringVar(&o.CACertFile, "tls-ca-cert-file", o.CACertFile, "ca certificate file")
	fs.StringVar(&o.CertFile, "tls-cert-file", o.CertFile, "client certificate file")
	fs.StringVar(&o.KeyFile, "tls-key-file", o.KeyFile, "client key file")
	fs.StringVar(&o.AuthToken, "auth-token", o.AuthToken, "authorization token")
	fs.StringVar(&o.AuthTokenType, "auth-token-type", o.AuthTokenType, "authorization token type")
	fs.StringVar(&o.JWTKey, "jwt-key", o.JWTKey, "jwt key")
	fs.StringVar(&o.JWTKeyFile, "jwt-key-file", o.JWTKeyFile, "jwt key file")
}

func TypesClientCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "types",
		Short: "Types service client",
		Long:  "",
	}
	TypesClientDefaultConfig.addFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_TypesEchoCommand(),
	)
	return cmd
}

func _TypesDial(ctx context.Context) (*grpc.ClientConn, TypesClient, error) {
	cfg := TypesClientDefaultConfig
	opts := []grpc.DialOption{grpc.WithBlock()}
	if cfg.TLS {
		tlsConfig := &tls.Config{InsecureSkipVerify: cfg.InsecureSkipVerify}
		if cfg.CACertFile != "" {
			caCert, err := ioutil.ReadFile(cfg.CACertFile)
			if err != nil {
				return nil, nil, fmt.Errorf("ca cert: %v", err)
			}
			certPool := x509.NewCertPool()
			certPool.AppendCertsFromPEM(caCert)
			tlsConfig.RootCAs = certPool
		}
		if cfg.CertFile != "" {
			if cfg.KeyFile == "" {
				return nil, nil, fmt.Errorf("missing key file")
			}
			pair, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
			if err != nil {
				return nil, nil, fmt.Errorf("cert/key: %v", err)
			}
			tlsConfig.Certificates = []tls.Certificate{pair}
		}
		if cfg.ServerName != "" {
			tlsConfig.ServerName = cfg.ServerName
		} else {
			addr, _, _ := net.SplitHostPort(cfg.ServerAddr)
			tlsConfig.ServerName = addr
		}
		cred := credentials.NewTLS(tlsConfig)
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	if cfg.AuthToken != "" {
		cred := oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: cfg.AuthToken,
			TokenType:   cfg.AuthTokenType,
		})
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	if cfg.JWTKey != "" {
		cred, err := oauth.NewJWTAccessFromKey([]byte(cfg.JWTKey))
		if err != nil {
			return nil, nil, fmt.Errorf("jwt key: %v", err)
		}
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	if cfg.JWTKeyFile != "" {
		cred, err := oauth.NewJWTAccessFromFile(cfg.JWTKeyFile)
		if err != nil {
			return nil, nil, fmt.Errorf("jwt key file: %v", err)
		}
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	if cfg.Timeout > 0 {
		var done context.CancelFunc
		ctx, done = context.WithTimeout(ctx, cfg.Timeout)
		defer done()
	}
	conn, err := grpc.DialContext(ctx, cfg.ServerAddr, opts...)
	if err != nil {
		return nil, nil, err
	}
	return conn, NewTypesClient(conn), nil
}

type _TypesRoundTripFunc func(cli TypesClient, in iocodec.Decoder, out iocodec.Encoder) error

func _TypesRoundTrip(ctx context.Context, fn _TypesRoundTripFunc) error {
	cfg := TypesClientDefaultConfig
	var dm iocodec.DecoderMaker
	r := os.Stdin
	if cfg.Stdin || cfg.RequestFile == "-" {
		dm = iocodec.DefaultDecoders["json"]
	} else if cfg.RequestFile != "" {
		f, err := os.Open(cfg.RequestFile)
		if err != nil {
			return fmt.Errorf("request file: %v", err)
		}
		defer f.Close()
		if ext := strings.TrimLeft(filepath.Ext(cfg.RequestFile), "."); ext != "" {
			dm = iocodec.DefaultDecoders[ext]
		}
		if dm == nil {
			dm = iocodec.DefaultDecoders["json"]
		}
		r = f
	} else {
		dm = iocodec.DefaultDecoders["noop"]
	}
	var em iocodec.EncoderMaker
	if cfg.ResponseFormat == "" {
		em = iocodec.DefaultEncoders["json"]
	} else if em = iocodec.DefaultEncoders[cfg.ResponseFormat]; em == nil {
		return fmt.Errorf("invalid response format: %q", cfg.ResponseFormat)
	}
	conn, client, err := _TypesDial(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	return fn(client, dm.NewDecoder(r), em.NewEncoder(os.Stdout))
}

func _TypesEchoCommand() *cobra.Command {
	req := &Sound{}

	cmd := &cobra.Command{
		Use:   "echo",
		Short: "Echo RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return _TypesRoundTrip(cmd.Context(), func(cli TypesClient, in iocodec.Decoder, out iocodec.Encoder) error {
				v := &Sound{}

				if err := in.Decode(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Echo(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out.Encode(res)

			})
		},
	}

	_GlobalEnumVar(cmd.PersistentFlags(), &req.GlobalEnum, "globalenum", "")
	_Sound_NestedEnumSliceVar(cmd.PersistentFlags(), &req.ListEnum, "listenum", "")
	_Sound_NestedEnumVar(cmd.PersistentFlags(), &req.NestedEnum, "nestedenum", "")
	cmd.PersistentFlags().BoolSliceVar(&req.ListBool, "listbool", nil, "")
	cmd.PersistentFlags().BoolVar(&req.Bool, "bool", false, "")
	cmd.PersistentFlags().BytesBase64Var(&req.Bytes, "bytes", nil, "")
	cmd.PersistentFlags().Float32SliceVar(&req.ListFloat, "listfloat", nil, "")
	cmd.PersistentFlags().Float32Var(&req.Float, "float", 0, "")
	cmd.PersistentFlags().Float64SliceVar(&req.ListDouble, "listdouble", nil, "")
	cmd.PersistentFlags().Float64Var(&req.Double, "double", 0, "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListInt32, "listint32", nil, "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListSfixed32, "listsfixed32", nil, "")
	cmd.PersistentFlags().Int32SliceVar(&req.ListSint32, "listsint32", nil, "")
	cmd.PersistentFlags().Int32Var(&req.Int32, "int32", 0, "")
	cmd.PersistentFlags().Int32Var(&req.Sfixed32, "sfixed32", 0, "")
	cmd.PersistentFlags().Int32Var(&req.Sint32, "sint32", 0, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListInt64, "listint64", nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListSfixed64, "listsfixed64", nil, "")
	cmd.PersistentFlags().Int64SliceVar(&req.ListSint64, "listsint64", nil, "")
	cmd.PersistentFlags().Int64Var(&req.Int64, "int64", 0, "")
	cmd.PersistentFlags().Int64Var(&req.Sfixed64, "sfixed64", 0, "")
	cmd.PersistentFlags().Int64Var(&req.Sint64, "sint64", 0, "")
	cmd.PersistentFlags().StringSliceVar(&req.ListString, "liststring", nil, "")
	cmd.PersistentFlags().StringToInt64Var(&req.MapStringInt64, "mapstringint64", nil, "")
	cmd.PersistentFlags().StringToStringVar(&req.MapStringString, "mapstringstring", nil, "")
	cmd.PersistentFlags().StringVar(&req.String_, "string_", "", "")
	cmd.PersistentFlags().Uint32Var(&req.Fixed32, "fixed32", 0, "")
	cmd.PersistentFlags().Uint32Var(&req.Uint32, "uint32", 0, "")
	cmd.PersistentFlags().Uint64Var(&req.Fixed64, "fixed64", 0, "")
	cmd.PersistentFlags().Uint64Var(&req.Uint64, "uint64", 0, "")
	cmd.PersistentFlags().Var(flag.NewBoolWrapperValue(func(v *wrappers.BoolValue) { req.WrapperBool = v }), "wrapperbool", "")
	cmd.PersistentFlags().Var(flag.NewBytesBase64WrapperValue(func(v *wrappers.BytesValue) { req.WrapperBytes = v }), "wrapperbytes", "")
	cmd.PersistentFlags().Var(flag.NewDoubleWrapperValue(func(v *wrappers.DoubleValue) { req.WrapperDouble = v }), "wrapperdouble", "")
	cmd.PersistentFlags().Var(flag.NewDurationValue(func(v *duration.Duration) { req.Duration = v }), "duration", "")
	cmd.PersistentFlags().Var(flag.NewFloatWrapperValue(func(v *wrappers.FloatValue) { req.WrapperFloat = v }), "wrapperfloat", "")
	cmd.PersistentFlags().Var(flag.NewInt32WrapperValue(func(v *wrappers.Int32Value) { req.WrapperInt32 = v }), "wrapperint32", "")
	cmd.PersistentFlags().Var(flag.NewInt64WrapperValue(func(v *wrappers.Int64Value) { req.WrapperInt64 = v }), "wrapperint64", "")
	cmd.PersistentFlags().Var(flag.NewStringWrapperValue(func(v *wrappers.StringValue) { req.WrapperString = v }), "wrapperstring", "")
	cmd.PersistentFlags().Var(flag.NewTimestampValue(func(v *timestamp.Timestamp) { req.Timestamp = v }), "timestamp", "")
	cmd.PersistentFlags().Var(flag.NewUInt32WrapperValue(func(v *wrappers.UInt32Value) { req.WrapperUint32 = v }), "wrapperuint32", "")
	cmd.PersistentFlags().Var(flag.NewUInt64WrapperValue(func(v *wrappers.UInt64Value) { req.WrapperUint64 = v }), "wrapperuint64", "")
	flag.BytesBase64SliceVar(cmd.PersistentFlags(), &req.ListBytes, "listbytes", "")
	flag.Uint32SliceVar(cmd.PersistentFlags(), &req.ListFixed32, "listfixed32", "")
	flag.Uint32SliceVar(cmd.PersistentFlags(), &req.ListUint32, "listuint32", "")
	flag.Uint64SliceVar(cmd.PersistentFlags(), &req.ListFixed64, "listfixed64", "")
	flag.Uint64SliceVar(cmd.PersistentFlags(), &req.ListUint64, "listuint64", "")

	return cmd
}

type _GlobalEnumValue GlobalEnum

func _GlobalEnumVar(fs *pflag.FlagSet, p *GlobalEnum, name, usage string) {
	fs.Var((*_GlobalEnumValue)(p), name, usage)
}

func (v *_GlobalEnumValue) Set(val string) error {
	if e, err := parseGlobalEnum(val); err != nil {
		return err
	} else {
		*v = _GlobalEnumValue(e)
		return nil
	}
}

func (v *_GlobalEnumValue) Type() string { return "GlobalEnum" }

func (v *_GlobalEnumValue) String() string { return (GlobalEnum)(*v).String() }

func parseGlobalEnum(s string) (GlobalEnum, error) {
	if i, ok := GlobalEnum_value[s]; ok {
		return GlobalEnum(i), nil
	} else if i, err := strconv.ParseInt(s, 0, 32); err == nil {
		return GlobalEnum(i), nil
	} else {
		return 0, err
	}
}

type _Sound_NestedEnumValue Sound_NestedEnum

func _Sound_NestedEnumVar(fs *pflag.FlagSet, p *Sound_NestedEnum, name, usage string) {
	fs.Var((*_Sound_NestedEnumValue)(p), name, usage)
}

func (v *_Sound_NestedEnumValue) Set(val string) error {
	if e, err := parseSound_NestedEnum(val); err != nil {
		return err
	} else {
		*v = _Sound_NestedEnumValue(e)
		return nil
	}
}

func (v *_Sound_NestedEnumValue) Type() string { return "Sound_NestedEnum" }

func (v *_Sound_NestedEnumValue) String() string { return (Sound_NestedEnum)(*v).String() }

type _Sound_NestedEnumSliceValue struct {
	value   *[]Sound_NestedEnum
	changed bool
}

func _Sound_NestedEnumSliceVar(fs *pflag.FlagSet, p *[]Sound_NestedEnum, name, usage string) {
	fs.Var(&_Sound_NestedEnumSliceValue{value: p}, name, usage)
}

func (s *_Sound_NestedEnumSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]Sound_NestedEnum, len(ss))
	for i, s := range ss {
		var err error
		if out[i], err = parseSound_NestedEnum(s); err != nil {
			return err
		}
	}
	if !s.changed {
		*s.value = out
		s.changed = true
	} else {
		*s.value = append(*s.value, out...)
	}
	return nil
}

func (s *_Sound_NestedEnumSliceValue) Type() string { return "Sound_NestedEnumSlice" }

func (s *_Sound_NestedEnumSliceValue) String() string { return "[]" }

func (s *_Sound_NestedEnumSliceValue) Append(val string) error {
	var e Sound_NestedEnum
	if err := (*_Sound_NestedEnumValue)(&e).Set(val); err != nil {
		return err
	}
	*s.value = append(*s.value, e)
	return nil
}

func (s *_Sound_NestedEnumSliceValue) Replace(val []string) error {
	out := make([]Sound_NestedEnum, len(val))
	for i, s := range val {
		if err := (*_Sound_NestedEnumValue)(&out[i]).Set(s); err != nil {
			return err
		}
	}
	*s.value = out
	return nil
}

func (s *_Sound_NestedEnumSliceValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, v := range *s.value {
		out[i] = v.String()
	}
	return out
}

func parseSound_NestedEnum(s string) (Sound_NestedEnum, error) {
	if i, ok := Sound_NestedEnum_value[s]; ok {
		return Sound_NestedEnum(i), nil
	} else if i, err := strconv.ParseInt(s, 0, 32); err == nil {
		return Sound_NestedEnum(i), nil
	} else {
		return 0, err
	}
}
