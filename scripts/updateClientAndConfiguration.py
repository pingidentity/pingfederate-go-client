#!/usr/bin/env python3
# coding: UTF-8

# Update client.go
updatedClientLines = []
with open("client.go", 'r') as clientFile:
    for line in clientFile:
        # Skip unused imports
        if "\"log\"" in line or "\"net/http/httputil\"" in line:
            continue
        # Remove debug lines for printing request contents
        if "c.cfg.Debug" in line:
            # advance iterator to skip the expected lines
            for i in range(7):
                line = clientFile.readline()
        # Use configuration.go method for UserAgent
        if "User-Agent" in line:
            line = line.replace("c.cfg.UserAgent", "c.cfg.UserAgent()")
        updatedClientLines.append(line)

with open("client.go", 'w') as clientFile:
    for line in updatedClientLines:
        clientFile.write(line)


# Update configuration.go
updatedConfigurationLines = []
with open("configuration.go", 'r') as configurationFile:
    for line in configurationFile:
        # Split UserAgent into two fields in the Configuration struct
        if "`json:\"userAgent,omitempty\"`" in line:
            updatedConfigurationLines.append("	UserAgentSuffix   *string           `json:\"userAgentSuffix,omitempty\"`\n")
            updatedConfigurationLines.append("	UserAgentOverride *string           `json:\"userAgentOverride,omitempty\"`\n")
            continue
        # Remove UserAgent from the default configuration struct
        if "UserAgent:" in line:
            continue
        updatedConfigurationLines.append(line)

# Add new UserAgent() method with default that handle override and suffix
updatedConfigurationLines.append("""
func (c *Configuration) UserAgent() string {
	if c.UserAgentOverride != nil {
		return *c.UserAgentOverride
	}

	result := "PingFederate-GOLANG-SDK/1200.0.3"
	if c.UserAgentSuffix != nil {
		result += " "
		result += *c.UserAgentSuffix
	}
	return result
}
""")

with open("configuration.go", 'w') as configurationFile:
    for line in updatedConfigurationLines:
        configurationFile.write(line)

