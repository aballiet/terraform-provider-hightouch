// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"hightouch/internal/sdk/pkg/models/shared"
	"math/big"
	"time"
)

func (r *SyncDataSourceModel) RefreshFromGetResponse(resp *shared.Sync) {
	if r.Configuration == nil && len(resp.Configuration) > 0 {
		r.Configuration = make(map[string]types.String)
		for key, value := range resp.Configuration {
			result, _ := json.Marshal(value)
			r.Configuration[key] = types.StringValue(string(result))
		}
	}
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339))
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.Disabled = types.BoolValue(resp.Disabled)
	r.ID = types.StringValue(resp.ID)
	r.LastRunAt = types.StringValue(resp.LastRunAt.Format(time.RFC3339))
	r.ModelID = types.StringValue(resp.ModelID)
	r.PrimaryKey = types.StringValue(resp.PrimaryKey)
	r.ReferencedColumns = nil
	for _, v := range resp.ReferencedColumns {
		r.ReferencedColumns = append(r.ReferencedColumns, types.StringValue(v))
	}
	if resp.Schedule.Schedule.IntervalSchedule != nil {
		r.Schedule.Schedule.IntervalSchedule = &IntervalSchedule{}
		r.Schedule.Schedule.IntervalSchedule.Interval.Quantity = types.NumberValue(big.NewFloat(float64(resp.Schedule.Schedule.IntervalSchedule.Interval.Quantity)))
		r.Schedule.Schedule.IntervalSchedule.Interval.Unit = types.StringValue(string(resp.Schedule.Schedule.IntervalSchedule.Interval.Unit))
	}
	if resp.Schedule.Schedule.CronSchedule != nil {
		r.Schedule.Schedule.CronSchedule = &CronSchedule{}
		r.Schedule.Schedule.CronSchedule.Expression = types.StringValue(resp.Schedule.Schedule.CronSchedule.Expression)
	}
	if resp.Schedule.Schedule.VisualCronSchedule != nil {
		r.Schedule.Schedule.VisualCronSchedule = &VisualCronSchedule{}
		r.Schedule.Schedule.VisualCronSchedule.Expressions = nil
		for _, expressionsItem := range resp.Schedule.Schedule.VisualCronSchedule.Expressions {
			var expressions1 VisualCronScheduleExpressions
			if expressionsItem.Days.Friday != nil {
				expressions1.Days.Friday = types.BoolValue(*expressionsItem.Days.Friday)
			} else {
				expressions1.Days.Friday = types.BoolNull()
			}
			if expressionsItem.Days.Monday != nil {
				expressions1.Days.Monday = types.BoolValue(*expressionsItem.Days.Monday)
			} else {
				expressions1.Days.Monday = types.BoolNull()
			}
			if expressionsItem.Days.Saturday != nil {
				expressions1.Days.Saturday = types.BoolValue(*expressionsItem.Days.Saturday)
			} else {
				expressions1.Days.Saturday = types.BoolNull()
			}
			if expressionsItem.Days.Sunday != nil {
				expressions1.Days.Sunday = types.BoolValue(*expressionsItem.Days.Sunday)
			} else {
				expressions1.Days.Sunday = types.BoolNull()
			}
			if expressionsItem.Days.Thursday != nil {
				expressions1.Days.Thursday = types.BoolValue(*expressionsItem.Days.Thursday)
			} else {
				expressions1.Days.Thursday = types.BoolNull()
			}
			if expressionsItem.Days.Tuesday != nil {
				expressions1.Days.Tuesday = types.BoolValue(*expressionsItem.Days.Tuesday)
			} else {
				expressions1.Days.Tuesday = types.BoolNull()
			}
			if expressionsItem.Days.Wednesday != nil {
				expressions1.Days.Wednesday = types.BoolValue(*expressionsItem.Days.Wednesday)
			} else {
				expressions1.Days.Wednesday = types.BoolNull()
			}
			expressions1.Time = types.StringValue(expressionsItem.Time)
			r.Schedule.Schedule.VisualCronSchedule.Expressions = append(r.Schedule.Schedule.VisualCronSchedule.Expressions, expressions1)
		}
	}
	if resp.Schedule.Schedule.DBTSchedule != nil {
		r.Schedule.Schedule.DBTSchedule = &DBTSchedule{}
		r.Schedule.Schedule.DBTSchedule.Account.ID = types.StringValue(resp.Schedule.Schedule.DBTSchedule.Account.ID)
		r.Schedule.Schedule.DBTSchedule.DbtCredentialID = types.StringValue(resp.Schedule.Schedule.DBTSchedule.DbtCredentialID)
		r.Schedule.Schedule.DBTSchedule.Job.ID = types.StringValue(resp.Schedule.Schedule.DBTSchedule.Job.ID)
	}
	r.Schedule.Type = types.StringValue(resp.Schedule.Type)
	r.Slug = types.StringValue(resp.Slug)
	r.Status = types.StringValue(string(resp.Status))
	r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339))
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
