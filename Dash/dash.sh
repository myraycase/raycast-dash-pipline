#!/bin/bash

# Required parameters:
# @raycast.title Dash document
# @raycast.author moli
# @raycast.authorURL https://github.com/molizz
# @raycast.description Find document

# @raycast.mode silent
# @raycast.schemaVersion 1

# Optional parameters:
# @raycast.icon images/dash-logo.png

# @raycast.argument1 {"type":"text", "placeholder": "Query" }

open $($(dirname $0)/dash $1)
