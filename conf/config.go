package conf

import (
	"bitbucket.org/juancolamendydevteam/golib/utils/logger"
	env "github.com/joeshaw/envdecode"
	"os"
)

type Config struct {
	AWS_ACCESS_KEY string `env:"AWS_ACCESS_KEY,default="`
	AWS_SECRET_KEY string `env:"AWS_SECRET_KEY,default="`
	AWS_REGION     string `env:"AWS_REGION,default="`

	DB_USERNAME string `env:"DB_USERNAME,default="`
	DB_PASSWORD string `env:"DB_PASSWORD,default="`
	DB_HOST string `env:"DB_HOST,default="`
	DB_PORT     int    `env:"DB_PORT,default="`
	DB_DATABASE string `env:"DB_DATABASE,default="`

	EMAIL_VERIFICATION_URL string `env:"EMAIL_VERIFICATION_URL,default="`
	ENVIRONMENT string `env:"ENVIRONMENT,default=dev"`

	FREE_PLAN_ID string `env:"FREE_PLAN_ID,default="`

	HTTP_PORT string `env:"HTTP_PORT,default=3000"`
	HTTP_REQUEST_TIMEOUT int64 `env:"HTTP_REQUEST_TIMEOUT,default=60"`

	MAX_INFLIGHT_REQUEST int `env:"MAX_INFLIGHT_REQUEST,default=10"`

	REDIS_HOSTNAME string `env:"REDIS_HOSTNAME,default="`
	REDIS_PORT     int    `env:"REDIS_PORT,default="`

	REQUEST_RETRIES      int `env:"REQUEST_RETRIES,default=1"`
	REQUEST_WAIT_TIME    int `env:"REQUEST_WAIT_TIME,default=10"`	

	SESSION_TIMEOUT int `env:"SESSION_TIMEOUT,default=1800"`

	STORAGE_ACCESS_ID     string `env:"STORAGE_ACCESS_ID,default="`
	STORAGE_SECRET_KEY    string `env:"STORAGE_SECRET_KEY,default="`
	STORAGE_PUBLIC_BUCKET string `env:"STORAGE_PUBLIC_BUCKET,default="`
	STORAGE_REGION        string `env:"STORAGE_REGION,default="`
	STORAGE_ENDPOINT_URL  string `env:"STORAGE_ENDPOINT_URL,default="`

	STRIPE_PRIVATE_KEY string `env:"STRIPE_PRIVATE_KEY,default="`	
	
	CWD         string
}

func GetConfig() *Config {
	var cwd string
	var err error

	cwd, err = os.Getwd()
	if err != nil {
		logger.Error.Printf("Cannot find working directory:\n%v\n", err)
	}

	config := &Config{CWD: cwd}
	err = env.Decode(config)
	if err != nil {
		logger.Error.Printf("Error on decoding the environment variables:\n%v\n", err)
	}

	return config
}

func (c *Config) Dump() {
	logger.Info.Printf("--- Config: AWS_ACCESS_KEY: %s", c.AWS_ACCESS_KEY)
	logger.Info.Printf("--- Config: AWS_SECRET_KEY: %s", c.AWS_SECRET_KEY)
	logger.Info.Printf("--- Config: AWS_REGION: %s", c.AWS_REGION)

	logger.Info.Printf("--- Config: DB_USERNAME: %s", c.DB_USERNAME)
	logger.Info.Printf("--- Config: DB_PASSWORD: %s", c.DB_PASSWORD)
	logger.Info.Printf("--- Config: DB_HOST: %s", c.DB_HOST)
	logger.Info.Printf("--- Config: DB_PORT: %d", c.DB_PORT)
	logger.Info.Printf("--- Config: DB_DATABASE: %s", c.DB_DATABASE)

	logger.Info.Printf("--- Config: EMAIL_VERIFICATION_URL: %s", c.EMAIL_VERIFICATION_URL)
	logger.Info.Printf("--- Config: ENVIRONMENT: %s", c.ENVIRONMENT)

	logger.Info.Printf("--- Config: FREE_PLAN_ID: %s", c.FREE_PLAN_ID)

	logger.Info.Printf("--- Config: MAX_INFLIGHT_REQUEST: %d", c.MAX_INFLIGHT_REQUEST)

	logger.Info.Printf("--- Config: HTTP_PORT: %s", c.HTTP_PORT)
	logger.Info.Printf("--- Config: HTTP_REQUEST_TIMEOUT: %d", c.HTTP_REQUEST_TIMEOUT)

	logger.Info.Printf("--- Config: REDIS_HOSTNAME: %s", c.REDIS_HOSTNAME)
	logger.Info.Printf("--- Config: REDIS_HOSTPORT: %d", c.REDIS_PORT)

	logger.Info.Printf("--- Config: REQUEST_RETRIES: %d", c.REQUEST_RETRIES)
	logger.Info.Printf("--- Config: REQUEST_WAIT_TIME: %d", c.REQUEST_WAIT_TIME)

	logger.Info.Printf("--- Config: SESSION_TIMEOUT: %d", c.SESSION_TIMEOUT)

	logger.Info.Printf("--- Config: STORAGE_ACCESS_ID: %s", c.STORAGE_ACCESS_ID)
	logger.Info.Printf("--- Config: STORAGE_SECRET_KEY: %s", c.STORAGE_SECRET_KEY)
	logger.Info.Printf("--- Config: STORAGE_PUBLIC_BUCKET: %s", c.STORAGE_PUBLIC_BUCKET)
	logger.Info.Printf("--- Config: STORAGE_REGION: %s", c.STORAGE_REGION)
	logger.Info.Printf("--- Config: STORAGE_ENDPOINT_URL: %s", c.STORAGE_ENDPOINT_URL)

	logger.Info.Printf("--- Config: STRIPE_PRIVATE_KEY: %s", c.STRIPE_PRIVATE_KEY)
}
