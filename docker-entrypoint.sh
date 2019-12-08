#!/bin/sh

echo "Running tests against $1"

# Run this first so we don't get fetching errors
go test -i $1
go test -count=1 -v "$1" | tee output.log
terratest_log_parser --testlog output.log
