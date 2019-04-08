package define

var (
	Config BuzConfig
)

type BuzConfig struct {
	AddrThriftHost    string
	AddrThriftPort    string
}

type LinkSlotAddrResponse struct {
	PublishAddr   string `json:"publish_addr"`
	StreamAddr    string `json:"stream_addr"`
	StreamAddrCdn string `json:"stream_addr_cdn"`
}

