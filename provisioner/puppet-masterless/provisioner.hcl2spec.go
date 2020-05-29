// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package puppetmasterless

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	CleanStagingDir     *bool             `mapstructure:"clean_staging_directory" cty:"clean_staging_directory" hcl:"clean_staging_directory"`
	GuestOSType         *string           `mapstructure:"guest_os_type" cty:"guest_os_type" hcl:"guest_os_type"`
	ExecuteCommand      *string           `mapstructure:"execute_command" cty:"execute_command" hcl:"execute_command"`
	ExtraArguments      []string          `mapstructure:"extra_arguments" cty:"extra_arguments" hcl:"extra_arguments"`
	Facter              map[string]string `cty:"facter" hcl:"facter"`
	HieraConfigPath     *string           `mapstructure:"hiera_config_path" cty:"hiera_config_path" hcl:"hiera_config_path"`
	IgnoreExitCodes     *bool             `mapstructure:"ignore_exit_codes" cty:"ignore_exit_codes" hcl:"ignore_exit_codes"`
	ModulePaths         []string          `mapstructure:"module_paths" cty:"module_paths" hcl:"module_paths"`
	ManifestFile        *string           `mapstructure:"manifest_file" cty:"manifest_file" hcl:"manifest_file"`
	ManifestDir         *string           `mapstructure:"manifest_dir" cty:"manifest_dir" hcl:"manifest_dir"`
	PreventSudo         *bool             `mapstructure:"prevent_sudo" cty:"prevent_sudo" hcl:"prevent_sudo"`
	PuppetBinDir        *string           `mapstructure:"puppet_bin_dir" cty:"puppet_bin_dir" hcl:"puppet_bin_dir"`
	StagingDir          *string           `mapstructure:"staging_directory" cty:"staging_directory" hcl:"staging_directory"`
	WorkingDir          *string           `mapstructure:"working_directory" cty:"working_directory" hcl:"working_directory"`
	ElevatedUser        *string           `mapstructure:"elevated_user" cty:"elevated_user" hcl:"elevated_user"`
	ElevatedPassword    *string           `mapstructure:"elevated_password" cty:"elevated_password" hcl:"elevated_password"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"clean_staging_directory":    &hcldec.AttrSpec{Name: "clean_staging_directory", Type: cty.Bool, Required: false},
		"guest_os_type":              &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
		"execute_command":            &hcldec.AttrSpec{Name: "execute_command", Type: cty.String, Required: false},
		"extra_arguments":            &hcldec.AttrSpec{Name: "extra_arguments", Type: cty.List(cty.String), Required: false},
		"facter":                     &hcldec.AttrSpec{Name: "facter", Type: cty.Map(cty.String), Required: false},
		"hiera_config_path":          &hcldec.AttrSpec{Name: "hiera_config_path", Type: cty.String, Required: false},
		"ignore_exit_codes":          &hcldec.AttrSpec{Name: "ignore_exit_codes", Type: cty.Bool, Required: false},
		"module_paths":               &hcldec.AttrSpec{Name: "module_paths", Type: cty.List(cty.String), Required: false},
		"manifest_file":              &hcldec.AttrSpec{Name: "manifest_file", Type: cty.String, Required: false},
		"manifest_dir":               &hcldec.AttrSpec{Name: "manifest_dir", Type: cty.String, Required: false},
		"prevent_sudo":               &hcldec.AttrSpec{Name: "prevent_sudo", Type: cty.Bool, Required: false},
		"puppet_bin_dir":             &hcldec.AttrSpec{Name: "puppet_bin_dir", Type: cty.String, Required: false},
		"staging_directory":          &hcldec.AttrSpec{Name: "staging_directory", Type: cty.String, Required: false},
		"working_directory":          &hcldec.AttrSpec{Name: "working_directory", Type: cty.String, Required: false},
		"elevated_user":              &hcldec.AttrSpec{Name: "elevated_user", Type: cty.String, Required: false},
		"elevated_password":          &hcldec.AttrSpec{Name: "elevated_password", Type: cty.String, Required: false},
	}
	return s
}
