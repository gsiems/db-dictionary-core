package view

import (
	"path"

	m "github.com/gsiems/db-dictionary-core/model"
	t "github.com/gsiems/db-dictionary-core/template"
)

// domainsView contains the data used for generating the schema domains page
type domainsView struct {
	Title         string
	DBMSVersion   string
	DBName        string
	DBComment     string
	SchemaName    string
	SchemaComment string
	TmspGenerated string
	Domains       []m.Domain
}

// makeDomainsList marshals the data needed for, and then creates, a schema domains page
func makeDomainsList(md *m.MetaData) (err error) {

	for _, vs := range md.Schemas {

		context := domainsView{
			Title:         "Domains for " + md.Alias + "." + vs.Name,
			TmspGenerated: md.TmspGenerated,
			DBName:        md.Name,
			DBComment:     md.Comment,
			SchemaName:    vs.Name,
			SchemaComment: vs.Comment,
		}

		var tmplt t.T
		tmplt.AddPageHeader(1, md)

		context.Domains = md.FindDomains(vs.Name)
		if len(context.Domains) > 0 {
			tmplt.AddSnippet("SchemaDomains")
			md.SortDomains(context.Domains)
		} else {
			return nil
		}

		tmplt.AddPageFooter(1, md)

		dirName := path.Join(md.OutputDir, vs.Name)
		err = tmplt.RenderPage(dirName, "domains", context, md.Cfg.Minify)
		if err != nil {
			return err
		}
	}

	return err
}
