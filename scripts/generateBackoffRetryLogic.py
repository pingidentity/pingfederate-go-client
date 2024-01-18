#!/usr/bin/env python3
# coding: UTF-8
import glob
import re

# Find all files in the repo starting with "api_"
modelFiles = glob.glob('api_*.go')
executeFunctionRegex = re.compile("func \(([a-zA-Z0-9\* ]+)\) ([a-zA-Z])([a-zA-Z0-9]+Execute)\(([a-zA-Z0-9 ]*)\) \(([\*a-zA-Z0-9]*), \*http\.Response, error\) {")

def backoffRetryLines():
    return ["//example haha\n"]

for modelFile in modelFiles:
    updatedLines = []
    with open(modelFile, 'r') as model:
        for line in model:
            # Search for matching lines
            match = re.search(executeFunctionRegex, line)
            #TODO Ensure this is idempotent
            if match != None:
                # add lines
                updatedLines.extend(backoffRetryLines())
            updatedLines.append(line)
    # Rewrite the file
    with open(modelFile, 'w') as model:
        for line in updatedLines:
            model.write(line)