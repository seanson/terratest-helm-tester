#!/bin/bash

echo "Running tests with args: $@"

# Run this first so we don't get fetching errors
go test -i $@

go test -count=1 -v $@  | tee output.log
EXIT_CODE=${PIPESTATUS[0]}

echo "Parsing logs and writing test output to ./out"

terratest_log_parser --testlog output.log >/dev/null 2>&1

if [ ! -z "${USER_ID}" ]; then
    echo "USER_ID set, changing ownership of out directory."
    chown -R ${USER_ID} out
fi

exit $EXIT_CODE
