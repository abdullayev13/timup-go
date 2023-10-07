package utill

import "abdullayev13/timeup/internal/config"

func PutDomain(path string) string {
	if path == "" {
		return ""
	}

	if path[0] == '/' {
		return config.Domain + path
	}

	return path
}

func PutMediaDomain(path string) string {
	if path == "" {
		return ""
	}

	if path[0] == '/' {
		return config.Domain + "/api/v1" + path
	}

	return path
}
