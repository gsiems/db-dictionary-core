package model

import (
	"path"
	"strings"
	"time"

	"github.com/gsiems/db-dictionary/config"
	m "github.com/gsiems/go-db-meta/model"
)

// MetaData is the metadata structure used for generating a data dictionary
type MetaData struct {
	TmspGenerated     string
	DBEngine          string
	CommentsFormat    string
	OutputDir         string
	File              string
	Version           string
	CharacterSet      string
	Name              string
	Alias             string
	Owner             string
	Comment           string
	Cfg               config.Config
	Schemas           []Schema
	Tables            []Table
	Columns           []Column
	Domains           []Domain
	Indexes           []Index
	CheckConstraints  []CheckConstraint
	UniqueConstraints []UniqueConstraint
	ForeignKeys       []ForeignKey
	PrimaryKeys       []PrimaryKey
	Dependencies      []Dependency
	Dependents        []Dependency
	UserTypes         []UserType
}

// Init initializes, and returns, a dictionary metadata structure
func Init(cfg config.Config) *MetaData {

	var md MetaData

	md.Cfg = cfg

	t := time.Now()
	md.TmspGenerated = t.Format("2006-01-02 15:04:05")

	if cfg.DbName != "" {
		md.Alias = cfg.DbName
	} else {
		tn := path.Base(cfg.File)
		te := path.Ext(tn)
		if te != "" {
			tn = strings.TrimSuffix(tn, te)
		}
		md.Alias = tn
	}

	switch cfg.CommentsFormat {
	case "markdown":
		md.CommentsFormat = cfg.CommentsFormat
	default:
		md.CommentsFormat = "none"
	}

	if cfg.OutputDir != "" {
		md.OutputDir = cfg.OutputDir
	} else {
		md.OutputDir = "."
	}

	return &md
}

// LoadCatalog loads the catalog information from go-db-meta into the
// dictionary metadata structure
func (md *MetaData) LoadCatalog(x *m.Catalog) {
	md.File = x.CatalogName.String
	md.Version = x.DBMSVersion.String
	md.CharacterSet = x.DefaultCharacterSetName.String
	md.Name = x.CatalogName.String
	md.Owner = x.CatalogOwner.String
	md.Comment = x.Comment.String
}