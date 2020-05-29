// Code generated by "mapstructure-to-hcl2 -type AmiFilterOptions,SecurityGroupFilterOptions,SubnetFilterOptions,VpcFilterOptions,PolicyDocument,Statement"; DO NOT EDIT.
package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/hcl2template"
	"github.com/zclconf/go-cty/cty"
)

// FlatAmiFilterOptions is an auto-generated flat version of AmiFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatAmiFilterOptions struct {
	Filters    map[string]string           `cty:"filters" hcl:"filters"`
	Filter     []hcl2template.FlatKeyValue `cty:"filter" hcl:"filter"`
	Owners     []string                    `cty:"owners" hcl:"owners"`
	MostRecent *bool                       `mapstructure:"most_recent" cty:"most_recent" hcl:"most_recent"`
}

// FlatMapstructure returns a new FlatAmiFilterOptions.
// FlatAmiFilterOptions is an auto-generated flat version of AmiFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*AmiFilterOptions) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatAmiFilterOptions)
}

// HCL2Spec returns the hcl spec of a AmiFilterOptions.
// This spec is used by HCL to read the fields of AmiFilterOptions.
// The decoded values from this spec will then be applied to a FlatAmiFilterOptions.
func (*FlatAmiFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters":     &hcldec.AttrSpec{Name: "filters", Type: cty.Map(cty.String), Required: false},
		"filter":      &hcldec.BlockListSpec{TypeName: "filter", Nested: hcldec.ObjectSpec((*hcl2template.FlatKeyValue)(nil).HCL2Spec())},
		"owners":      &hcldec.AttrSpec{Name: "owners", Type: cty.List(cty.String), Required: false},
		"most_recent": &hcldec.AttrSpec{Name: "most_recent", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatPolicyDocument is an auto-generated flat version of PolicyDocument.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatPolicyDocument struct {
	Version   *string         `cty:"version" hcl:"version"`
	Statement []FlatStatement `cty:"statement" hcl:"statement"`
}

// FlatMapstructure returns a new FlatPolicyDocument.
// FlatPolicyDocument is an auto-generated flat version of PolicyDocument.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*PolicyDocument) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatPolicyDocument)
}

// HCL2Spec returns the hcl spec of a PolicyDocument.
// This spec is used by HCL to read the fields of PolicyDocument.
// The decoded values from this spec will then be applied to a FlatPolicyDocument.
func (*FlatPolicyDocument) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"version":   &hcldec.AttrSpec{Name: "version", Type: cty.String, Required: false},
		"statement": &hcldec.BlockListSpec{TypeName: "statement", Nested: hcldec.ObjectSpec((*FlatStatement)(nil).HCL2Spec())},
	}
	return s
}

// FlatSecurityGroupFilterOptions is an auto-generated flat version of SecurityGroupFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatSecurityGroupFilterOptions struct {
	Filters map[string]string            `cty:"filters" hcl:"filters"`
	Filter  []hcl2template.FlatNameValue `cty:"filter" hcl:"filter"`
}

// FlatMapstructure returns a new FlatSecurityGroupFilterOptions.
// FlatSecurityGroupFilterOptions is an auto-generated flat version of SecurityGroupFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*SecurityGroupFilterOptions) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatSecurityGroupFilterOptions)
}

// HCL2Spec returns the hcl spec of a SecurityGroupFilterOptions.
// This spec is used by HCL to read the fields of SecurityGroupFilterOptions.
// The decoded values from this spec will then be applied to a FlatSecurityGroupFilterOptions.
func (*FlatSecurityGroupFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters": &hcldec.AttrSpec{Name: "filters", Type: cty.Map(cty.String), Required: false},
		"filter":  &hcldec.BlockListSpec{TypeName: "filter", Nested: hcldec.ObjectSpec((*hcl2template.FlatNameValue)(nil).HCL2Spec())},
	}
	return s
}

// FlatStatement is an auto-generated flat version of Statement.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatStatement struct {
	Effect   *string  `cty:"effect" hcl:"effect"`
	Action   []string `cty:"action" hcl:"action"`
	Resource *string  `cty:"resource" hcl:"resource"`
}

// FlatMapstructure returns a new FlatStatement.
// FlatStatement is an auto-generated flat version of Statement.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Statement) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatStatement)
}

// HCL2Spec returns the hcl spec of a Statement.
// This spec is used by HCL to read the fields of Statement.
// The decoded values from this spec will then be applied to a FlatStatement.
func (*FlatStatement) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"effect":   &hcldec.AttrSpec{Name: "effect", Type: cty.String, Required: false},
		"action":   &hcldec.AttrSpec{Name: "action", Type: cty.List(cty.String), Required: false},
		"resource": &hcldec.AttrSpec{Name: "resource", Type: cty.String, Required: false},
	}
	return s
}

// FlatSubnetFilterOptions is an auto-generated flat version of SubnetFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatSubnetFilterOptions struct {
	Filters  map[string]string            `cty:"filters" hcl:"filters"`
	Filter   []hcl2template.FlatNameValue `cty:"filter" hcl:"filter"`
	MostFree *bool                        `mapstructure:"most_free" cty:"most_free" hcl:"most_free"`
	Random   *bool                        `mapstructure:"random" cty:"random" hcl:"random"`
}

// FlatMapstructure returns a new FlatSubnetFilterOptions.
// FlatSubnetFilterOptions is an auto-generated flat version of SubnetFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*SubnetFilterOptions) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatSubnetFilterOptions)
}

// HCL2Spec returns the hcl spec of a SubnetFilterOptions.
// This spec is used by HCL to read the fields of SubnetFilterOptions.
// The decoded values from this spec will then be applied to a FlatSubnetFilterOptions.
func (*FlatSubnetFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters":   &hcldec.AttrSpec{Name: "filters", Type: cty.Map(cty.String), Required: false},
		"filter":    &hcldec.BlockListSpec{TypeName: "filter", Nested: hcldec.ObjectSpec((*hcl2template.FlatNameValue)(nil).HCL2Spec())},
		"most_free": &hcldec.AttrSpec{Name: "most_free", Type: cty.Bool, Required: false},
		"random":    &hcldec.AttrSpec{Name: "random", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatVpcFilterOptions is an auto-generated flat version of VpcFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatVpcFilterOptions struct {
	Filters map[string]string            `cty:"filters" hcl:"filters"`
	Filter  []hcl2template.FlatNameValue `cty:"filter" hcl:"filter"`
}

// FlatMapstructure returns a new FlatVpcFilterOptions.
// FlatVpcFilterOptions is an auto-generated flat version of VpcFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*VpcFilterOptions) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatVpcFilterOptions)
}

// HCL2Spec returns the hcl spec of a VpcFilterOptions.
// This spec is used by HCL to read the fields of VpcFilterOptions.
// The decoded values from this spec will then be applied to a FlatVpcFilterOptions.
func (*FlatVpcFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters": &hcldec.AttrSpec{Name: "filters", Type: cty.Map(cty.String), Required: false},
		"filter":  &hcldec.BlockListSpec{TypeName: "filter", Nested: hcldec.ObjectSpec((*hcl2template.FlatNameValue)(nil).HCL2Spec())},
	}
	return s
}
