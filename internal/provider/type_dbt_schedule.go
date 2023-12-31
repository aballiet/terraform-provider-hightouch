// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DBTSchedule struct {
	Account         DBTScheduleAccount `tfsdk:"account"`
	DbtCredentialID types.String       `tfsdk:"dbt_credential_id"`
	Job             DBTScheduleAccount `tfsdk:"job"`
}
