package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// ViperConfig is a wrapper for viper.Viper.
type ViperConfig struct {
	*viper.Viper
}

// NewWrapper returns a *ViperConfig for the provided *viper.Viper.
func NewWrapper(v *viper.Viper) *ViperConfig {
	return &ViperConfig{Viper: v}
}

// BindConfig is a shortcut for the viper.SetDefault/viper.BindEnv pattern. It uses the provided key to perform the
// following on the underlying viper.Viper:
// * The default is set to nil.
// * The key is bound to an environment variable that matches the key, uppercased.
// Panics if the binding fails.
func (v *ViperConfig) BindConfig(key string) {
	v.BindConfigAndOverrideEnvVar(key, "")
}

// BindConfigAndOverrideEnvVar is a shortcut for the viper.SetDefault/viper.BindEnv pattern. It uses the provided key
// to perform the following on the underlying viper.Viper:
// * The default is set to nil.
// * The key is bound to an environment variable, using the provided override. If the override is a zero-value string,
// it will use the envvar name matching the key, uppercased (i.e. the same behavior as BindConfig).
// Panics if the binding fails.
func (v *ViperConfig) BindConfigAndOverrideEnvVar(key string, envVarNameOverride string) {

	v.SetDefault(key, nil)

	var err error

	if envVarNameOverride != "" {
		err = v.BindEnv(key, envVarNameOverride)
	} else {
		err = v.BindEnv(key)
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to bind environment variable for key '%s' (name override: '%s'): %v", key, envVarNameOverride, err))
	}
}
