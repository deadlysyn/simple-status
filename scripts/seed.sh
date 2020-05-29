#!/usr/bin/env bash

set -eux -o pipefail

HOST="localhost:8080/api"

# Empty database
http ${HOST}/events

# Post events
http POST ${HOST}/events service=foo status=up description="doing ok"
http POST ${HOST}/events service=bar status=down description="kinda broke"

# List events
http ${HOST}/events

# List event detail
SVC=$(http ${HOST}/events | jq -r '.events[0].service')
http ${HOST}/events/${SVC}

# Update event
http PUT ${HOST}/events/${SVC} status=maintenance description="UPDATED"
