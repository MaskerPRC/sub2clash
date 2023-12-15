package model

type SmuxStruct struct {
	Enabled bool `yaml:"enable"`
}

type Proxy struct {
	Name                string         `yaml:"name,omitempty"`
	Server              string         `yaml:"server,omitempty"`
	Port                int            `yaml:"port,omitempty"`
	Type                string         `yaml:"type,omitempty"`
	Cipher              string         `yaml:"cipher,omitempty"`
	Password            string         `yaml:"password,omitempty"`
	UDP                 bool           `yaml:"udp,omitempty"`
	UUID                string         `yaml:"uuid,omitempty"`
	Network             string         `yaml:"network,omitempty"`
	Flow                string         `yaml:"flow,omitempty"`
	TLS                 bool           `yaml:"tls,omitempty"`
	ClientFingerprint   string         `yaml:"client-fingerprint,omitempty"`
	Plugin              string         `yaml:"plugin,omitempty"`
	PluginOpts          map[string]any `yaml:"plugin-opts,omitempty"`
	Smux                SmuxStruct     `yaml:"smux,omitempty"`
	Sni                 string         `yaml:"sni,omitempty"`
	AllowInsecure       bool           `yaml:"allow-insecure,omitempty"`
	Fingerprint         string         `yaml:"fingerprint,omitempty"`
	SkipCertVerify      bool           `yaml:"skip-cert-verify,omitempty"`
	Alpn                []string       `yaml:"alpn,omitempty"`
	XUDP                bool           `yaml:"xudp,omitempty"`
	Servername          string         `yaml:"servername,omitempty"`
	WSOpts              WSOptions      `yaml:"ws-opts,omitempty"`
	AlterID             int            `yaml:"alterId,omitempty"`
	GrpcOpts            GrpcOptions    `yaml:"grpc-opts,omitempty"`
	RealityOpts         RealityOptions `yaml:"reality-opts,omitempty"`
	Protocol            string         `yaml:"protocol,omitempty"`
	Obfs                string         `yaml:"obfs,omitempty"`
	ObfsParam           string         `yaml:"obfs-param,omitempty"`
	ProtocolParam       string         `yaml:"protocol-param,omitempty"`
	Remarks             []string       `yaml:"remarks,omitempty"`
	HTTPOpts            HTTPOptions    `yaml:"http-opts,omitempty"`
	HTTP2Opts           HTTP2Options   `yaml:"h2-opts,omitempty"`
	PacketAddr          bool           `yaml:"packet-addr,omitempty"`
	PacketEncoding      string         `yaml:"packet-encoding,omitempty"`
	GlobalPadding       bool           `yaml:"global-padding,omitempty"`
	AuthenticatedLength bool           `yaml:"authenticated-length,omitempty"`
	UDPOverTCP          bool           `yaml:"udp-over-tcp,omitempty"`
	UDPOverTCPVersion   int            `yaml:"udp-over-tcp-version,omitempty"`
	SubName             string         `yaml:"-"`
	GroupTags           map[string]bool
}

func (p Proxy) MarshalYAML() (interface{}, error) {
	switch p.Type {
	case "vmess":
		return ProxyToVmess(p), nil
	case "ss":
		return ProxyToShadowSocks(p), nil
	case "ssr":
		return ProxyToShadowSocksR(p), nil
	case "vless":
		return ProxyToVless(p), nil
	case "trojan":
		return ProxyToTrojan(p), nil
	}
	return nil, nil
}
