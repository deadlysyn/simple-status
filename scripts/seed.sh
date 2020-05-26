#!/usr/bin/env bash

set -eux -o pipefail

HOST="localhost:8080"

# Empty database
http ${HOST}/events

# Post events
http POST ${HOST}/events service=foo status=bar
http POST ${HOST}/events service=baz status=qux

# List events
http ${HOST}/events

# List event detail
ID=$(http ${HOST}/events | jq -r '.events[0]')
http ${HOST}/events/${ID}

# Update event
http PUT ${HOST}/events/${ID} service=foo status=CHANGED
