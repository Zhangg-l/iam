package apiserver

type apiServer struct {
}

type preparedAPIServer struct {
}

type ExtraConfig struct {
	Addr         string
	MaxMsgSize   int
	ServerCert   string
	mysqlOptions string
}

func createAPIServer(cfg *string) (*apiServer, error) {
	return nil, nil
}
