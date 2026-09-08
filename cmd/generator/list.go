package main

const (
	levelSO   = "SOL_SOCKET"
	levelIP   = "IPPROTO_IP"
	levelIPV6 = "IPPROTO_IPV6"
	levelTCP  = "IPPROTO_TCP"
	levelUDP  = "IPPROTO_UDP"
)

const (
	typeBool    = "Bool"
	typeInt     = "Int"
	typeLinger  = "Linger"
	typeTimeval = "Timeval"
	typeString  = "String"
)

type Value struct {
	Level  string
	Option string
	Type   string
}

var listLinux = []*Value{
	{levelSO, "SO_BINDTOIFINDEX", typeInt},
	{levelSO, "SO_BINDTODEVICE", typeString},
	{levelSO, "SO_DEBUG", typeBool},
	{levelSO, "SO_KEEPALIVE", typeBool},
	{levelSO, "SO_LINGER", typeLinger},
	{levelSO, "SO_MARK", typeInt},
	{levelSO, "SO_RCVBUF", typeInt},
	{levelSO, "SO_RCVBUFFORCE", typeInt},
	{levelSO, "SO_SNDBUF", typeInt},
	{levelSO, "SO_SNDBUFFORCE", typeInt},
	{levelSO, "SO_SNDTIMEO", typeTimeval},
	{levelSO, "SO_RCVTIMEO", typeTimeval},
	{levelSO, "SO_REUSEADDR", typeBool},
	{levelSO, "SO_REUSEPORT", typeBool},
	{levelIP, "IP_BIND_ADDRESS_NO_PORT", typeBool},
	{levelIP, "IP_FREEBIND", typeBool},
	{levelIP, "IP_LOCAL_PORT_RANGE", typeInt},
	{levelIP, "IP_TRANSPARENT", typeBool},
	{levelIP, "IP_TTL", typeInt},
	{levelIPV6, "IPV6_V6ONLY", typeBool},
	{levelTCP, "TCP_CORK", typeBool},
	{levelTCP, "TCP_DEFER_ACCEPT", typeInt},
	{levelTCP, "TCP_KEEPCNT", typeInt},
	{levelTCP, "TCP_KEEPIDLE", typeInt},
	{levelTCP, "TCP_KEEPINTVL", typeInt},
	{levelTCP, "TCP_LINGER2", typeLinger},
	{levelTCP, "TCP_MAXSEG", typeInt},
	{levelTCP, "TCP_NODELAY", typeBool},
	{levelTCP, "TCP_QUICKACK", typeBool},
	{levelTCP, "TCP_SYNCNT", typeInt},
	{levelTCP, "TCP_USER_TIMEOUT", typeInt},
	{levelTCP, "TCP_WINDOW_CLAMP", typeInt},
	{levelTCP, "TCP_FASTOPEN", typeBool},
	{levelTCP, "TCP_FASTOPEN_CONNECT", typeBool},
	{levelUDP, "UDP_CORK", typeBool},
	{levelUDP, "UDP_SEGMENT", typeInt},
	{levelUDP, "UDP_GRO", typeBool},
}

var listDarwin = []*Value{
	{levelSO, "SO_DEBUG", typeBool},
	{levelSO, "SO_KEEPALIVE", typeBool},
	{levelSO, "SO_LINGER", typeLinger},
	{levelSO, "SO_RCVBUF", typeInt},
	{levelSO, "SO_SNDBUF", typeInt},
	{levelSO, "SO_SNDTIMEO", typeTimeval},
	{levelSO, "SO_RCVTIMEO", typeTimeval},
	{levelSO, "SO_REUSEADDR", typeBool},
	{levelSO, "SO_REUSEPORT", typeBool},
	{levelIP, "IP_TTL", typeInt},
	{levelIPV6, "IPV6_V6ONLY", typeBool},
	{levelTCP, "TCP_KEEPCNT", typeInt},
	{levelTCP, "TCP_KEEPINTVL", typeInt},
	{levelTCP, "TCP_MAXSEG", typeInt},
	{levelTCP, "TCP_NODELAY", typeBool},
	{levelTCP, "TCP_FASTOPEN", typeBool},
}

var listWindows = []*Value{
	{levelSO, "SO_KEEPALIVE", typeBool},
	{levelSO, "SO_LINGER", typeLinger},
	{levelSO, "SO_RCVBUF", typeInt},
	{levelSO, "SO_SNDBUF", typeInt},
	{levelSO, "SO_REUSEADDR", typeBool},
	{levelIP, "IP_TTL", typeInt},
	{levelTCP, "TCP_NODELAY", typeBool},
}
