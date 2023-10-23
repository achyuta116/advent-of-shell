#!/bin/bash
#
# Compute the number of animals whose vaccinations have passed the due date
jq '.[] | .animals | .[] | if (.last_vaccination_date | fromdate) > (.due_vaccination_date | fromdate) then .quantity else 0 end' | awk '{sum+=$1} END {print sum}'

