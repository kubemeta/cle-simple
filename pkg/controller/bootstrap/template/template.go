package template

import (
	"context"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"

	"github.com/kubemeta/cle-helper/esutil/model/component"
	"github.com/kubemeta/cle-helper/esutil/model/ilm"
)

const (
	DefaultShard            = 1
	DefaultNumberOfReplicas = 1
	ComponentTemplateSuffix = "_tpl"
)

// GetComponentTemplate returns the component template

// CreateComponentTemplate creates the component template
func CreateComponentTemplate(ctx context.Context, es *elasticsearch.Client, name string) (*esapi.Response, error) {
	template := GenComponentTemplate(component.Settings{
		NumberOfShards:   DefaultShard,
		NumberOfReplicas: DefaultNumberOfReplicas,
		Index: &ilm.IndexLevelSettings{
			Lifecycle: &ilm.Lifecycle{
				Name:          name + ComponentTemplateSuffix,
				RolloverAlias: name,
			},
		},
	}, component.Properties{}, component.Aliases{
		name: map[string]interface{}{
			"is_write_index": true,
		},
	}, &Option{AddonMeta: true})

	req := esapi.ClusterPutComponentTemplateRequest{
		Name: name,
		Body: esutil.NewJSONReader(template),
	}

	return req.Do(ctx, es)
}

// GetIndexTemplate returns the index template

// CreateIndexTemplate creates the index template
func CreateIndexTemplate(ctx context.Context, es *elasticsearch.Client, name string) (*esapi.Response, error) {
	template := GenIndexTemplate([]string{name + "*"}, []string{name + ComponentTemplateSuffix})
	req := esapi.IndicesPutIndexTemplateRequest{
		Name: name,
		Body: esutil.NewJSONReader(template),
	}

	return req.Do(ctx, es)
}
