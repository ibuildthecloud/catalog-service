package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Version struct {
	Catalog        string `json:"catalog"`
	EnvironmentId  string `json:"environmentId"`
	Template       string `json:"template"`
	Revision       int    `json:"revision"`
	DockerCompose  string `json:"dockerCompose"`
	RancherCompose string `json:"rancherCompose"`

	// TODO: can move to Resource?
	Category       string `json:"category"`
	IsSystem       string `json:"isSystem"`
	Description    string `json:"description"`
	Version        string `json:"version"`
	DefaultVersion string `json:"defaultVersion"`
	IconLink       string `json:"iconLink"`
	//VersionLinks        map[string]string `json:"versionLinks"`
	//UpgradeVersionLinks map[string]string `json:"upgradeVersionLinks"`
	//Files               map[string]string `json:"files"`
	//Questions                        []Question        `json:"questions"`
	Path                  string `json:"path"`
	MinimumRancherVersion string `json:"minimumRancherVersion"`
	//TemplateVersionRancherVersion    map[string]string
	//TemplateVersionRancherVersionGte map[string]string
	Maintainer string `json:"maintainer"`
	License    string `json:"license"`
	ProjectURL string `json:"projectURL"`
	ReadmeLink string `json:"readmeLink"`
	//Output                           Output                 `json:"output" yaml:"output,omitempty"`
	TemplateBase string `json:"templateBase"`
	//Labels                map[string]string      `json:"labels"`
	UpgradeFrom string `json:"upgradeFrom"`
	//Bindings              map[string]interface{} `json:"bindings"`
	MaximumRancherVersion string `json:"maximumRancherVersion"`

	// TODO
	//	FolderName     string `json:"revision"`
}

type VersionModel struct {
	gorm.Model
	Version
}