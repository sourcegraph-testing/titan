// This is a fake file that doesn't compile properly or anything and was only added to
// reproduce a specific bug observed with Batch Changes via its specific diff

// stuff

// stuff

// line buffers

// stuff

// stuff

// stuff

// line buffers

// stuff

// stuff

// stuff

import (
  "github.snooguts.net/reddit/reddit-service-authentication-go/thrift/reddit/sso"
)

type config struct {
		baseplate.Config `yaml:",inline"`

		Authentication authenticationsvc.Config `yaml:"authentication"`
		Sso            ssosvc.Config            `yaml:"sso"`
		OAuth          oauthsvc.Config          `yaml:"oauth"`
}

func main() {
		flag.Parse()

		var cfg config
		if err := baseplate.ParseConfigYAML(&cfg); err != nil {
		    log.Fatalf("parsing config: %s", err)
		}
		ctx, bp, err := baseplate.New(context.Background(), baseplate.NewArgs{
				ConfigPath: *configPath,
				ServiceCfg: &cfg,
				EdgeContextFactory: edgecontext.Factory(edgecontext.Config{
						Logger: log.ErrorWithSentryWrapper(),
				})
		})
}