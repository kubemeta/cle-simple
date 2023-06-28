package template

import (
	"github.com/kubemeta/cle-helper/esutil/model/component"
	"github.com/kubemeta/cle-helper/esutil/model/datastream"
	"github.com/kubemeta/cle-helper/esutil/model/field"
	"github.com/kubemeta/cle-helper/esutil/model/index"
)

type Option struct {
	// Add default log collection state meta information
	AddonMeta bool
}

func GenComponentTemplate(settings component.Settings, properties component.Properties, aliases component.Aliases, option *Option) *component.ComponentTemplate {
	if option.AddonMeta {
		CommonProperties.Extend(component.Properties{
			"state": map[string]component.Properties{
				"properties": {
					"pipeline": field.FieldObject{
						Type: string(field.Keyword),
					},
					"source": field.FieldObject{
						Type: string(field.Keyword),
					},
					"filename": field.FieldObject{
						Type: string(field.Keyword),
					},
					"timestamp": field.FieldObject{
						Type: string(field.Date),
						Parameters: field.Parameters{
							Format: "2006-01-02T15:04:05.000Z",
						},
					},
					"offset": field.FieldObject{
						Type: string(field.Integer),
					},
					"bytes": field.FieldObject{
						Type: string(field.Integer),
					},
					"hostname": field.FieldObject{
						Type: string(field.Keyword),
					},
				},
			},
		})
	}

	if properties != nil {
		CommonProperties.Extend(properties)
	}

	return &component.ComponentTemplate{
		Template: component.Template{
			Settings: component.NewIndexSetting(
				component.WithNumberOfShards(settings.NumberOfShards),
				component.WithNumberOfReplicas(settings.NumberOfReplicas),
				component.WithIndexSettings(settings.Index),
			),
			Mappings: component.NewMappings(component.WithProperties(CommonProperties)),
			Aliases:  aliases,
		},
	}
}

var CommonProperties = component.Properties{
	"metadata": map[string]component.Properties{
		"properties": {
			"container_runtime": field.FieldObject{
				Type: string(field.Keyword),
			},
			"node_name": field.FieldObject{
				Type: string(field.Keyword),
			},
			"namespace": field.FieldObject{
				Type: string(field.Keyword),
			},
			"pod_name": field.FieldObject{
				Type: string(field.Keyword),
			},
			"pod_ip": field.FieldObject{
				Type: string(field.IP),
			},
			"container_name": field.FieldObject{
				Type: string(field.Keyword),
			},
			"workload_kind": field.FieldObject{
				Type: string(field.Keyword),
			},
			"workload_name": field.FieldObject{
				Type: string(field.Keyword),
			},
		},
	},
	"@timestamp": field.FieldObject{
		Type: string(field.Date),
	},
	"message": field.FieldObject{
		Type: string(field.Text),
		Parameters: field.Parameters{
			Fields: map[string]field.FieldObject{
				"keyword": {
					Type: string(field.Keyword),
				},
			},
		},
	},
}

func GenIndexTemplate(indexPatterns []string, composedOf []string) *index.IndexTemplate {
	template := index.NewIndexTemplate(
		index.WithIndexPatterns(indexPatterns...),
		index.WithDataStream(&datastream.DataStream{}),
		index.WithComposedOf(composedOf...),
	)

	return template
}
