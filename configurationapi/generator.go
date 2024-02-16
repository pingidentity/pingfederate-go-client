package configurationapi

// This file contains any go generate comments used by the client

//go:generate ../scripts/generateTfsdkTags.py

//go:generate ../scripts/generateBackoffRetryLogic.py

//go:generate ../scripts/updateClientAndConfiguration.py

//go:generate ../scripts/removeJsonDisallowUnknownProperties.py
