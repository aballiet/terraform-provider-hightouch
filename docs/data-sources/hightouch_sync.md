---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hightouch_sync Data Source - aballiet-terraform-provider-hightouch"
subcategory: ""
description: |-
  Sync DataSource
---

# hightouch_sync (Data Source)

Sync DataSource

## Example Usage

```terraform
data "hightouch_sync" "my_sync" {
    id = "10faaa23-52c5-4955-907a-ff1a3a2fa946"
        }
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the sync

### Read-Only

- `configuration` (Map of String) The sync's configuration. This specifies how data is mapped, among other
configuration.

The schema depends on the destination type.

Consumers should NOT make assumptions on the contents of the
configuration. It may change as Hightouch updates its internal code.
- `created_at` (String) The timestamp when the sync was created
- `destination_id` (String) The id of the Destination that sync is connected to
- `disabled` (Boolean) Whether the sync has been disabled by the user.
- `last_run_at` (String) The timestamp of the last sync run
- `model_id` (String) The id of the Model that sync is connected to
- `primary_key` (String) The primary key that sync uses to identify data from source
- `referenced_columns` (List of String) The reference column that sync depends on to sync data from source
- `schedule` (Attributes) The scheduling configuration. It can be triggerd based on several ways:

Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)

Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.

Visual: the sync will be trigged based a visual cron configuration on UI

DBT-cloud: the sync will be trigged based on a dbt cloud job (see [below for nested schema](#nestedatt--schedule))
- `slug` (String) The sync's slug
- `status` (String) must be one of ["disabled", "pending", "cancelled", "failed", "queued", "success", "warning", "querying", "processing", "reporting", "interrupted"]
SyncStatus
- `updated_at` (String) The timestamp when the sync was last updated
- `workspace_id` (String) The id of the workspace that the sync belongs to

<a id="nestedatt--schedule"></a>
### Nested Schema for `schedule`

Read-Only:

- `schedule` (Attributes) (see [below for nested schema](#nestedatt--schedule--schedule))
- `type` (String)

<a id="nestedatt--schedule--schedule"></a>
### Nested Schema for `schedule.schedule`

Read-Only:

- `cron_schedule` (Attributes) (see [below for nested schema](#nestedatt--schedule--schedule--cron_schedule))
- `dbt_schedule` (Attributes) (see [below for nested schema](#nestedatt--schedule--schedule--dbt_schedule))
- `interval_schedule` (Attributes) (see [below for nested schema](#nestedatt--schedule--schedule--interval_schedule))
- `visual_cron_schedule` (Attributes) (see [below for nested schema](#nestedatt--schedule--schedule--visual_cron_schedule))

<a id="nestedatt--schedule--schedule--cron_schedule"></a>
### Nested Schema for `schedule.schedule.cron_schedule`

Read-Only:

- `expression` (String)


<a id="nestedatt--schedule--schedule--dbt_schedule"></a>
### Nested Schema for `schedule.schedule.dbt_schedule`

Read-Only:

- `account` (Attributes) (see [below for nested schema](#nestedatt--schedule--schedule--dbt_schedule--account))
- `dbt_credential_id` (String)
- `job` (Attributes) (see [below for nested schema](#nestedatt--schedule--schedule--dbt_schedule--job))

<a id="nestedatt--schedule--schedule--dbt_schedule--account"></a>
### Nested Schema for `schedule.schedule.dbt_schedule.job`

Read-Only:

- `id` (String)


<a id="nestedatt--schedule--schedule--dbt_schedule--job"></a>
### Nested Schema for `schedule.schedule.dbt_schedule.job`

Read-Only:

- `id` (String)



<a id="nestedatt--schedule--schedule--interval_schedule"></a>
### Nested Schema for `schedule.schedule.interval_schedule`

Read-Only:

- `interval` (Attributes) (see [below for nested schema](#nestedatt--schedule--schedule--interval_schedule--interval))

<a id="nestedatt--schedule--schedule--interval_schedule--interval"></a>
### Nested Schema for `schedule.schedule.interval_schedule.interval`

Read-Only:

- `quantity` (Number)
- `unit` (String) must be one of ["minute", "hour", "day", "week"]



<a id="nestedatt--schedule--schedule--visual_cron_schedule"></a>
### Nested Schema for `schedule.schedule.visual_cron_schedule`

Read-Only:

- `expressions` (Attributes List) (see [below for nested schema](#nestedatt--schedule--schedule--visual_cron_schedule--expressions))

<a id="nestedatt--schedule--schedule--visual_cron_schedule--expressions"></a>
### Nested Schema for `schedule.schedule.visual_cron_schedule.expressions`

Read-Only:

- `days` (Attributes) Construct a type with a set of properties K of type T (see [below for nested schema](#nestedatt--schedule--schedule--visual_cron_schedule--expressions--days))
- `time` (String)

<a id="nestedatt--schedule--schedule--visual_cron_schedule--expressions--days"></a>
### Nested Schema for `schedule.schedule.visual_cron_schedule.expressions.days`

Read-Only:

- `friday` (Boolean)
- `monday` (Boolean)
- `saturday` (Boolean)
- `sunday` (Boolean)
- `thursday` (Boolean)
- `tuesday` (Boolean)
- `wednesday` (Boolean)

