// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type RecordDayBooleanOrUndefined struct {
	Friday    types.Bool `tfsdk:"friday"`
	Monday    types.Bool `tfsdk:"monday"`
	Saturday  types.Bool `tfsdk:"saturday"`
	Sunday    types.Bool `tfsdk:"sunday"`
	Thursday  types.Bool `tfsdk:"thursday"`
	Tuesday   types.Bool `tfsdk:"tuesday"`
	Wednesday types.Bool `tfsdk:"wednesday"`
}
