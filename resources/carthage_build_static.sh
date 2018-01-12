#!/bin/sh -e

xcconfig=$(mktemp /tmp/static.xcconfig.buildben)
trap 'rm -f "$xcconfig"' INT TERM HUP EXIT

echo "LD = $LD_PATH" >> $xcconfig
echo "DEBUG_INFORMATION_FORMAT = dwarf" >> $xcconfig

export XCODE_XCCONFIG_FILE="$xcconfig"

carthage build "$@"
