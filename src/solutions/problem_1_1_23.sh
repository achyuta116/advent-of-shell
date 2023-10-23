#!/bin/bash
#
# Find the most common last name for farmers
jq '.[] | .last_name // (.name | split(" ")[-1])' | sort | uniq -c | sort -r | head -n 1 | awk '{ print $1 }'
