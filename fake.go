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

// line buffers

// stuff

var configPath = flag.String("config", "", "Path to the config file to use.")

type config struct {
		Authentication authenticationsvc.Config `yaml:"authentication"`
		Sso            ssosvc.Config            `yaml:"sso"`
		OAuth          oauthsvc.Config          `yaml:"oauth"`
}

func main() {
		flag.Parse()

		var cfg config
		ctx, bp, err := baseplate.New(context.Background(), baseplate.NewArgs{
				ConfigPath: *configPath,
				ServiceCfg: &cfg,
				EdgeContextFactory: edgecontext.Factory(edgecontext.Config{
						Logger: log.ErrorWithSentryWrapper(),
				})
		})
}