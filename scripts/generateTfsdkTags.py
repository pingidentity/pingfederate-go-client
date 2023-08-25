#!/usr/bin/env python3
# coding: UTF-8
import glob
import re

# Find all files in the repo starting with "model_"
modelFiles = glob.glob('model_*.go')
jsonRegex = re.compile('`json:\"([0-9A-Za-z]*)[\",]')

for modelFile in modelFiles:
    updatedLines = []
    with open(modelFile, 'r') as model:
        for line in model:
            # Search for json tags
            match = re.search(jsonRegex, line)
            # Ensure we don't add tfsdk tags to a line that already has one
            if match != None and "tfsdk:" not in line:
                # Get the attribute name from the json tag
                attribute = match.group(1)
                # Create a tfsdk tag for the attribute
                underscoreAttribute = ""
                i = 0
                lastCharUpper = False
                while i < len(attribute):
                    if i > 0 and attribute[i].isupper() != attribute[i - 1].isupper():
                        underscoreAttribute += "_"
                        underscoreAttribute += attribute[i].lower()
                        i += 1
                    if i < len(attribute):
                        underscoreAttribute += attribute[i].lower()
                        i += 1
                # Add the tfsdk tag to the line
                line = line.replace("\"`", "\" tfsdk:\"{}\"`".format(underscoreAttribute))
            updatedLines.append(line)
    # Rewrite the file
    with open(modelFile, 'w') as model:
        for line in updatedLines:
            model.write(line)
