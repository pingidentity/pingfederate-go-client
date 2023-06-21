#!/bin/bash
FILES=$(find . -type f -name '*.go')

for file in $FILES; do
  rm $file
done
  