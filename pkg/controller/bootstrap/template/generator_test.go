package template

import (
	"fmt"
	"testing"

	"github.com/kubemeta/cle-helper/esutil/model/component"
	"github.com/kubemeta/cle-helper/esutil/model/datastream"
	"github.com/kubemeta/cle-helper/esutil/model/field"
	"github.com/kubemeta/cle-helper/esutil/model/ilm"
	"github.com/kubemeta/cle-helper/esutil/model/index"
)

func TestGenComponentTemplate(t *testing.T) {
	template, _ := GenComponentTemplate(
		component.Settings{
			NumberOfShards:   1,
			NumberOfReplicas: 1,
			Index: &ilm.IndexLevelSettings{
				Lifecycle: &ilm.Lifecycle{
					Name: "my_policy",
				},
			},
		},
		component.Properties{
			"extend": field.FieldObject{
				Type: string(field.Keyword),
			},
		}, component.Aliases{}, &Option{
			AddonMeta: false,
		}).Marshal()
	fmt.Println(string(template))
}

func TestGenIndexTemplate(t *testing.T) {
	indexPatterns := []string{"log-els-*"}
	template := index.NewIndexTemplate(
		index.WithIndexPatterns(indexPatterns...),
		index.WithDataStream(&datastream.DataStream{}),
		index.WithComposedOf([]string{"component_template_test"}...),
	)

	data, _ := template.Marshal()
	fmt.Printf(string(data))
}
