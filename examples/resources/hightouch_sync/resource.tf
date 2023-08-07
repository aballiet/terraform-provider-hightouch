resource "hightouch_sync" "my_sync" {
    configuration = {
        "minus" = "{ \"see\": \"documentation\" }"
        "placeat" = "{ \"see\": \"documentation\" }"
    }
            destination_id = "...my_destinationId..."
            disabled = false
            model_id = "...my_modelId..."
            schedule = {
        schedule = {
            interval_schedule =     {
                    interval = {
                        quantity = 48
                        unit = "day"
                    }
                }
        }
        type = "...my_type..."
    }
            slug = "...my_slug..."
            type = "...my_type..."
        }