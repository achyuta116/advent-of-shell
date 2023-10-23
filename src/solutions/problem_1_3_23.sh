#!/bin/bash
#
# Compute the number of farmers eligible for a loan
# Given eligibility criteria being the number of species in the farm < 40 the farm size should be at least 'Small'
# and if the number of species is >= 40 the farm size should be 'Large'

jq 'map(. + { "sum" : .animals | map(.quantity) | add }) | map({sum, size}) | map(.sum >= 40 and .size == "Large" or .sum <= 40) | .[]' | sort -r | uniq -c | head -n 1 | awk '{print $1}'
