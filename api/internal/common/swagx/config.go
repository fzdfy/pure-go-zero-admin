package swagx

type Opts func(*FileConfig)

func NewConfig(opts ...Opts) *FileConfig {
	c := &FileConfig{
		generalInfo: "main.go",
		outputTypes: "json",
		output:      "./docs",
		env:         "dev",
		swaggerConfig: swaggerConfig{
			SpecURL:     "swagger-json",
			SwaggerHost: "https://petstore.swagger.io",
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithSpecURL(specURL string) Opts {
	return func(c *FileConfig) {
		c.SpecURL = specURL
	}
}

func WithSwaggerHost(swaggerHost string) Opts {
	return func(c *FileConfig) {
		c.SwaggerHost = swaggerHost
	}
}

func WithGeneralInfo(generalInfo string) Opts {
	return func(c *FileConfig) {
		c.generalInfo = generalInfo
	}
}

func WithOutputTypes(outputTypes string) Opts {
	return func(c *FileConfig) {
		c.outputTypes = outputTypes
	}
}

func WithOutput(output string) Opts {
	return func(c *FileConfig) {
		c.output = output
	}
}

func WithEnv(env string) Opts {
	return func(c *FileConfig) {
		c.env = env
	}
}
