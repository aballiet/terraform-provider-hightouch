// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
)

type SyncCreateScheduleScheduleType string

const (
	SyncCreateScheduleScheduleTypeIntervalSchedule   SyncCreateScheduleScheduleType = "IntervalSchedule"
	SyncCreateScheduleScheduleTypeCronSchedule       SyncCreateScheduleScheduleType = "CronSchedule"
	SyncCreateScheduleScheduleTypeVisualCronSchedule SyncCreateScheduleScheduleType = "VisualCronSchedule"
	SyncCreateScheduleScheduleTypeDBTSchedule        SyncCreateScheduleScheduleType = "DBTSchedule"
)

type SyncCreateScheduleSchedule struct {
	IntervalSchedule   *IntervalSchedule
	CronSchedule       *CronSchedule
	VisualCronSchedule *VisualCronSchedule
	DBTSchedule        *DBTSchedule

	Type SyncCreateScheduleScheduleType
}

func CreateSyncCreateScheduleScheduleIntervalSchedule(intervalSchedule IntervalSchedule) SyncCreateScheduleSchedule {
	typ := SyncCreateScheduleScheduleTypeIntervalSchedule

	return SyncCreateScheduleSchedule{
		IntervalSchedule: &intervalSchedule,
		Type:             typ,
	}
}

func CreateSyncCreateScheduleScheduleCronSchedule(cronSchedule CronSchedule) SyncCreateScheduleSchedule {
	typ := SyncCreateScheduleScheduleTypeCronSchedule

	return SyncCreateScheduleSchedule{
		CronSchedule: &cronSchedule,
		Type:         typ,
	}
}

func CreateSyncCreateScheduleScheduleVisualCronSchedule(visualCronSchedule VisualCronSchedule) SyncCreateScheduleSchedule {
	typ := SyncCreateScheduleScheduleTypeVisualCronSchedule

	return SyncCreateScheduleSchedule{
		VisualCronSchedule: &visualCronSchedule,
		Type:               typ,
	}
}

func CreateSyncCreateScheduleScheduleDBTSchedule(dbtSchedule DBTSchedule) SyncCreateScheduleSchedule {
	typ := SyncCreateScheduleScheduleTypeDBTSchedule

	return SyncCreateScheduleSchedule{
		DBTSchedule: &dbtSchedule,
		Type:        typ,
	}
}

func (u *SyncCreateScheduleSchedule) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	intervalSchedule := new(IntervalSchedule)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&intervalSchedule); err == nil {
		u.IntervalSchedule = intervalSchedule
		u.Type = SyncCreateScheduleScheduleTypeIntervalSchedule
		return nil
	}

	cronSchedule := new(CronSchedule)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&cronSchedule); err == nil {
		u.CronSchedule = cronSchedule
		u.Type = SyncCreateScheduleScheduleTypeCronSchedule
		return nil
	}

	visualCronSchedule := new(VisualCronSchedule)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&visualCronSchedule); err == nil {
		u.VisualCronSchedule = visualCronSchedule
		u.Type = SyncCreateScheduleScheduleTypeVisualCronSchedule
		return nil
	}

	dbtSchedule := new(DBTSchedule)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&dbtSchedule); err == nil {
		u.DBTSchedule = dbtSchedule
		u.Type = SyncCreateScheduleScheduleTypeDBTSchedule
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SyncCreateScheduleSchedule) MarshalJSON() ([]byte, error) {
	if u.IntervalSchedule != nil {
		return json.Marshal(u.IntervalSchedule)
	}

	if u.CronSchedule != nil {
		return json.Marshal(u.CronSchedule)
	}

	if u.VisualCronSchedule != nil {
		return json.Marshal(u.VisualCronSchedule)
	}

	if u.DBTSchedule != nil {
		return json.Marshal(u.DBTSchedule)
	}

	return nil, nil
}

// SyncCreateSchedule - The scheduling configuration. It can be triggerd based on several ways:
//
// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
//
// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
//
// Visual: the sync will be trigged based a visual cron configuration on UI
//
// DBT-cloud: the sync will be trigged based on a dbt cloud job
type SyncCreateSchedule struct {
	Schedule SyncCreateScheduleSchedule `json:"schedule"`
	Type     string                     `json:"type"`
}

// SyncCreate - The input for creating a Sync
type SyncCreate struct {
	// The sync's configuration. This specifies how data is mapped, among other
	// configuration.
	//
	// The schema depends on the destination type.
	//
	// Consumers should NOT make assumptions on the contents of the
	// configuration. It may change as Hightouch updates its internal code.
	Configuration map[string]interface{} `json:"configuration"`
	// The id of the Destination that sync is connected to
	DestinationID string `json:"destinationId"`
	// Whether the sync has been disabled by the user.
	Disabled bool `json:"disabled"`
	// The id of the Model that sync is connected to
	ModelID string `json:"modelId"`
	// The scheduling configuration. It can be triggerd based on several ways:
	//
	// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
	//
	// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
	//
	// Visual: the sync will be trigged based a visual cron configuration on UI
	//
	// DBT-cloud: the sync will be trigged based on a dbt cloud job
	Schedule SyncCreateSchedule `json:"schedule"`
	// The sync's slug
	Slug string `json:"slug"`
}
