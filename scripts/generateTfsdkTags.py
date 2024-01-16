#!/usr/bin/env python3
# coding: UTF-8
import glob
import re

# Find all files in the repo starting with "model_"
modelFiles = glob.glob('model_*.go')
jsonRegex = re.compile('`json:\"([0-9A-Za-z]*)[\",]')

def removeUnderscoresFromGroup(group) -> str:
    return re.sub(r'_', '', group)

def buildAttributeValue(attribute) -> str:
    underscoreAttribute = ""
    match = re.findall(r'[0-9A-Z]', attribute)
    if match != None:
        underscoreAttribute = re.sub(r'([A-Z])', r'_\g<0>', attribute).lower()
        fourUpperChars = re.findall(r'[A-Z]{4,4}', attribute)
        threeUpperChars = re.findall(r'[A-Z]{3,3}', attribute)
        twoUpperChars = re.findall(r'[A-Z]{2,2}', attribute)
        if len(fourUpperChars) > 0:
            underscoreAttribute = re.sub(r'_[a-z]_[a-z]_[a-z]_[a-z]', lambda m: "_" + removeUnderscoresFromGroup(m.group()) + "_", underscoreAttribute).lower()
        elif len(threeUpperChars) > 0:
            underscoreAttribute = re.sub(r'_[a-z]_[a-z]_[a-z]', lambda m: "_" + removeUnderscoresFromGroup(m.group()) + "_", underscoreAttribute).lower()
        elif len(twoUpperChars) > 0 :
            underscoreAttribute = re.sub(r'_[a-z]_[a-z]', lambda m: "_" + removeUnderscoresFromGroup(m.group()), underscoreAttribute).lower()
        underscoreAttribute =  underscoreAttribute.replace("__", "_")
    else:
        underscoreAttribute = attribute
    if underscoreAttribute.endswith("_"):
        underscoreAttribute = underscoreAttribute[:-1]
    return underscoreAttribute

def cleanExceptionAttributes(attribute) -> str:
    # Handle OAuth attributes that start with "oAuth"
    if len(attribute) >= 6 and attribute[0:5] == "oAuth":
        return "oauth" + attribute[5:]
    # Handle "IDEncrypted" string
    if "IDEncrypted" in attribute:
        return attribute.replace("IDEncrypted", "IdEncrypted")
    return attribute

for modelFile in modelFiles:
    updatedLines = []
    with open(modelFile, 'r') as model:
        for line in model:
            # Search for json tags
            match = re.search(jsonRegex, line)
            # Ensure we don't add tfsdk tags to a line that already has one
            if match != None and "tfsdk:" not in line:
                # Get the attribute name from the json tag
                attribute = str(match.group(1))
                finalAttribute = buildAttributeValue(cleanExceptionAttributes(attribute))
                # Add the tfsdk tag to the line
                line = line.replace("\"`", "\" tfsdk:\"{}\"`".format(finalAttribute))
            updatedLines.append(line)
    # Rewrite the file
    with open(modelFile, 'w') as model:
        for line in updatedLines:
            model.write(line)