resource "hightouch_sync" "my_sync" {
    configuration = {
        "minus" = "{ \"see\": \"documentation\" }"
        "placeat" = "{ \"see\": \"documentation\" }"
    }
            destination_id = "...my_destination_id..."
            disabled = false
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