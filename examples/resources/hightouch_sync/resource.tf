resource "hightouch_sync" "my_sync" {
    configuration = {
        "perferendis" = "{ \"see\": \"documentation\" }"
        "ipsam" = "{ \"see\": \"documentation\" }"
    }
            destination_id = "...my_destination_id..."
            disabled = true
            model_id = "...my_model_id..."
            schedule = {
        schedule = {
            cron_schedule =     {
                    expression = "...my_expression..."
                }
        }
        type = "...my_type..."
    }
            slug = "...my_slug..."
        }