# Extended configuration for Pulsar as a target (all options)

target {
  use "pulsar" {
    # Pulsar broker service connection string
    broker_service_url    = "pulsar://127.0.0.1:6650"

    # Pulsar topic name
    topic_name = "snowplow-enriched-good"

    # Pulsar default byte limit is 5MB (default: 5242880)
    byte_limit = 5242880
  }
}