#!/bin/bash

# mapfile -d $'\0' array < <(find . -name *.go -print0) && echo "${array[*]}"
# # Get file name: sed 's|.*/||'
outputDir=$1
echo $outputDir
pwd
ls -al
ls -al other
mapfile -d $'\0' gofiles < <(find . -name "*.go" -print0)
for file in "${gofiles[@]}"
do
   echo $file
done
go build -buildmode=plugin -gcflags='all=-N -l' -o ../.plugins/page.so page.go
