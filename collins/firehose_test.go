package collins

import (
	"testing"
)

func TestFirehoseService_Consume_single(t *testing.T) {
	setup()
	defer teardown()

	SetupFirehose("../tests/firehose/single_update.txt", "text/event-stream", t)

	events, err := client.Firehose.Consume()
	if err != nil {
		t.Errorf("Firehose.Consume return error: %v", err)
	}
	event := <-events

	tag := event.(*AssetUpdateEvent).Tag
	if tag != "tag30" {
		t.Errorf("FirehoseService.Consume() returned asset tag %s, want tag30", tag)
	}
}

func TestFirehoseService_Consume_multiple(t *testing.T) {
	setup()
	defer teardown()

	SetupFirehose("../tests/firehose/multiple_updates.txt", "text/event-stream", t)

	events, err := client.Firehose.Consume()
	if err != nil {
		t.Errorf("Firehose.Consume return error: %v", err)
	}

	tags := []string{}
	for event := range events {
		tags = append(tags, event.(*AssetUpdateEvent).Tag)
	}

	if len(tags) != 3 {
		t.Errorf("FirehoseService.Consume() %d assets, want 3", len(tags))
	}
}

func TestFirehoseService_Consume_multiline(t *testing.T) {
	setup()
	defer teardown()

	SetupFirehose("../tests/firehose/multiline_data.txt", "text/event-stream", t)

	events, err := client.Firehose.Consume()
	if err != nil {
		t.Errorf("Firehose.Consume return error: %v", err)
	}
	asset := <-events

	tag := asset.(*AssetUpdateEvent).Tag
	if tag != "tag30" {
		t.Errorf("FirehoseService.Consume() returned asset tag %s, want tag30", tag)
	}
}

func TestFirehoseService_Consume_unauthorized(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(403, "/api/firehose", "../tests/firehose/single_update.txt", "text/plain;", t)

	_, err := client.Firehose.Consume()
	if err == nil {
		t.Errorf("Firehose.Consume did not return error when unauthorized")
	}
}

func TestFirehoseService_Consume_wrong_contentype(t *testing.T) {
	setup()
	defer teardown()

	SetupFirehose("../tests/firehose/single_update.txt", "application/json", t)

	_, err := client.Firehose.Consume()
	if err == nil {
		t.Errorf("Firehose.Consume did not return error when unauthorized")
	}
}
