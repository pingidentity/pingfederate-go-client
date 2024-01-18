#!/usr/bin/env python3
# coding: UTF-8
import glob
import re

# Find all files in the repo starting with "api_"
apiFiles = glob.glob('api_*.go')
executeFunctionRegex = re.compile("func \(([a-zA-Z0-9\* ]+)\) ([A-Z][a-zA-Z0-9]+Execute)\(([a-zA-Z0-9 ]*)\) \(([\*a-zA-Z0-9]*), \*http\.Response, error\) {")
internalFunctionRegex = re.compile(".* internal([a-zA-Z0-9]+Execute)\(.*")

# Add call to processResponse and a separate internal function to handle the normal request logic
def backoffRetryLines(receiver, prefix, parameter, returnType):
    return """func ({0}) {1}({2}) ({3}, *http.Response, error) {{
	var (
		err                  error
		response             *http.Response
		localVarReturnValue  {3}
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {{
			return r.ApiService.internal{1}(r)
		}},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}}
			
func ({0}) internal{1}({2}) ({3}, *http.Response, error) {{
    """.format(receiver, prefix, parameter, returnType)

for apiFile in apiFiles:
    updatedLines = []
    allLines = []
    internalMethodsFound = []
    with open(apiFile, 'r') as model:
        # Find internalMethods that were already created by this process
        for line in model:
            allLines.append(line)
            match = re.search(internalFunctionRegex, line)
            if match != None:
                internalMethodsFound.append(match.groups()[0])
        # Insert backoff retry logic
        for line in allLines:
            # Search for matching lines
            match = re.search(executeFunctionRegex, line)
            # Avoid adding internal method if one alreaady exists, so that this script is idempotent
            if match != None:
                groups = match.groups()
                if groups[1] not in internalMethodsFound:
                    # Call processResponse function to handle retries
                    updatedLines.append(backoffRetryLines(groups[0], groups[1], groups[2], groups[3]))
                else:
                    updatedLines.append(line)
            else:
                updatedLines.append(line)
    # Rewrite the file
    with open(apiFile, 'w') as model:
        for line in updatedLines:
            model.write(line)