package option

import "flag"

type Option struct {
	Listen string
	Cert   string
	Key    string
}

func GetOption() Option {
	opt := Option{}
	flag.StringVar(&opt.Listen, "port", "", "The port to serve on")
	flag.StringVar(&opt.Cert, "cert", "", "TLS certificate file")
	flag.StringVar(&opt.Key, "key", "", "TLS certificate key file")
	flag.Parse()

	if len(opt.Listen) > 0 {
		opt.Listen = ":" + opt.Listen
	} else if len(opt.Cert) > 0 && len(opt.Key) > 0 {
		opt.Listen = ":8443"
	} else {
		opt.Listen = ":8080"
	}

	return opt
}
