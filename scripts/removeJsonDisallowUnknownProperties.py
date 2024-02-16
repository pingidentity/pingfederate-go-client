#!/usr/bin/env python3
# coding: UTF-8
import glob
import re

# Find all files in the repo starting with "model_"
modelFiles = glob.glob('model_*.go')

for modelFile in modelFiles:
    updatedLines = []
    with open(modelFile, 'r') as model:
        for line in model:
            # Remove DisallowUnknownFields call for JSON unmarshalling
            if "DisallowUnknownFields" not in line:
                updatedLines.append(line)
    # Rewrite the file
    with open(modelFile, 'w') as model:
        for line in updatedLines:
            model.write(line)