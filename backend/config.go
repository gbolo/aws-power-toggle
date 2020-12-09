package backend

import (
	"strings"

	"github.com/spf13/viper"
)

// ConfigInit instantiates and validates the configuration options
// optionally it can print out a configuration summary
func ConfigInit(cfgFile string, printConfig bool) {

	// init viper
	initViper(cfgFile)

	// Print config if required
	if printConfig {
		printConfigSummary()
	}

	// Sanity checks
	sanityChecks()

	// assign variable values to config values
	awsRegions = viper.GetStringSlice("aws.regions")
	envNameIgnore = viper.GetStringSlice("aws.ignore_environments")
	instanceTypeIgnore = viper.GetStringSlice("aws.ignore_instance_types")
	maxInstancesToShutdown = viper.GetInt("aws.max_instances_to_shutdown")
	requiredTagKey = viper.GetString("aws.required_tag_key")
	requiredTagValue = viper.GetString("aws.required_tag_value")
	environmentTagKey = viper.GetString("aws.environment_tag_key")
	slackEnabled = viper.GetBool("slack.enabled")
	slackWebHooks = viper.GetStringSlice("slack.webhook_urls")
	mockEnabled = viper.GetBool("mock.enabled")
	experimentalEnabled = viper.GetBool("experimental.enabled")
	asgEnabled = viper.GetBool("aws.enable_asg_support")

	return
}

// setup viper
func initViper(cfgFile string) {

	// Set some defaults
	viper.SetDefault("log_level", "DEBUG")
	viper.SetDefault("server.bind_address", "127.0.0.1")
	viper.SetDefault("server.bind_port", "8080")
	viper.SetDefault("server.access_log", true)

	// Configuring and pulling overrides from environmental variables
	viper.SetEnvPrefix("POWER_TOGGLE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// set default config name and paths to look for it
	viper.SetConfigType("yaml")
	viper.SetConfigName("power-toggle-config")
	viper.AddConfigPath("./")

	// if the user provides a config file in a flag, lets use it
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	// If a config file is found, read it in.
	err := viper.ReadInConfig()

	// Kick-off the logging module
	loggingInit(viper.GetString("log_level"))

	if err == nil {
		log.Infof("using config file: %s", viper.ConfigFileUsed())
	} else {
		log.Warningf("no config file found: using environment variables and hard-coded defaults: %v", err)
	}
}

// prints the config options
func printConfigSummary() {

	log.Debugf("Configuration:\n")
	for _, c := range []string{
		"log_level",
		"server.bind_address",
		"server.bind_port",
		"server.tls.enabled",
		"server.access_log",
		"server.compression",
		"aws.polling_interval",
		"aws.required_tag_key",
		"aws.required_tag_value",
		"aws.environment_tag_key",
		"aws.max_instances_to_shutdown",
		"aws.enable_asg_support",
		"slack.enabled",
		"mock.enabled",
		"mock.delay",
		"mock.errors",
		"experimental.enabled",
	} {
		log.Debugf("%s: %s\n", c, viper.GetString(c))
	}

	for _, c := range []string{
		"aws.regions",
		"aws.ignore_instance_types",
		"aws.ignore_environments",
	} {
		log.Debugf("%s: %v\n", c, viper.GetStringSlice(c))
	}
}

// checks that the config is correctly defined
func sanityChecks() {

	for _, k := range []string{
		"aws.required_tag_key",
		"aws.required_tag_value",
		"aws.environment_tag_key",
	} {
		if viper.GetString(k) == "" {
			log.Fatalf("%s MUST be defined and not empty", k)
		}
	}

	if len(viper.GetStringSlice("aws.regions")) == 0 {
		log.Fatal("aws.regions MUST be defined and not empty")
	}

	for _, k := range []string{
		"aws.max_instances_to_shutdown",
		"aws.polling_interval",
	} {
		if !(viper.GetInt(k) > 0) {
			log.Fatal("polling_interval MUST be defined and greater than 0")
		}
	}

	if !(viper.GetInt("aws.max_instances_to_shutdown") > 0) {
		log.Fatal("max_instances_to_shutdown MUST be defined and greater than 0")
	}

	if !(viper.GetInt("aws.polling_interval") > 0) {
		log.Fatal("polling_interval MUST be defined and greater than 0")
	}

	if viper.GetBool("slack.enabled") && len(viper.GetStringSlice("slack.webhook_urls")) == 0 {
		log.Warning("slack is ENABLED but slack.webhook_urls is empty")
	}
}
