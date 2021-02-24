#!/bin/bash
# Builds module plugins.
#
# $1: source directory
# $2: source filename
# $3: absolute output path and filename
# $4: add debug flags
#
set -x

flags=$([[ -z $4 ]] && echo "" || echo -gcflags='all=-N -l')

echo -n "Building plugin $2..."
cd $1
if [[ -z "$flags" ]]
then
  go build -buildmode=plugin -o $3 $2
else
  go build -buildmode=plugin "$flags" -o $3 $2
fi
echo -n "done"
