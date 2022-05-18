// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

type IntegrationsUsage struct {
	Count int64 `json:"count"`
}

var InstalledIntegrationsIgnoredPlugins = map[string]struct{}{
	"playbooks":                             {},
	"focalboard:":                           {},
	"com.mattermost.apps:":                  {},
	"com.mattermost.nps:":                   {},
	"com.mattermost.plugin-channel-export:": {},
	"com.mattermost.calls:":                 {},
}

type InstalledIntegration struct {
	Type    string `json:"type"` // "plugin", "app", or "plugin-app" if it is an app installed as a plugin.
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Enabled bool   `json:"enabled"`
}
