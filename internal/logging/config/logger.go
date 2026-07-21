package config

import stdlog "log"

func Loading(file string) {
	stdlog.Printf("config: loading file=%s", file)
}

func FileLoaded(file string) {
	stdlog.Printf("config: file loaded file=%s", file)
}

func FileNotFound(file string) {
	stdlog.Printf("config: file not found, using environment variables file=%s", file)
}

func LoadingFailed(file string, err error) {
	stdlog.Printf("config: loading failed file=%s error=%q", file, err)
}

func Loaded(environment, httpHost string, httpPort int, repositoryGRPC, exchangeRateGRPC string) {
	stdlog.Printf(
		"config: loaded environment=%s http_host=%s http_port=%d repository_grpc=%s exchange_rate_grpc=%s",
		environment, httpHost, httpPort, repositoryGRPC, exchangeRateGRPC,
	)
}
