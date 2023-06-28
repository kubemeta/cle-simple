package ilm

import (
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"

	"github.com/kubemeta/cle-helper/esutil/model/ilm"
	"github.com/kubemeta/cle-helper/esutil/model/unit"
)

const (
	DefaultPrefix = "k8s"
	PolicySuffix  = "_policy"
)

func GenDefaultName(name string) string {
	return fmt.Sprintf("%s-%s-%s", DefaultPrefix, name, PolicySuffix)
}

func CreateIlmPolicy(ctx context.Context, es *elasticsearch.Client, policy string, obj ilm.IlmObject) (*esapi.Response, error) {
	req := esapi.ILMPutLifecycleRequest{
		Policy: GenDefaultName(policy),
		Body:   esutil.NewJSONReader(obj),
	}

	return req.Do(ctx, es)
}

var DefaultHot = ilm.IlmObject{
	Policy: ilm.Policy{
		Phases: ilm.Phases{
			Hot: ilm.NewPhaseSetting(ilm.WithPriority(100),
				ilm.WithRollover(ilm.DefaultRollover()),
			),
			Delete: ilm.NewPhaseSetting().SetMinAge(7, unit.Days),
		},
	},
}

var DefaultHotWarm = ilm.IlmObject{
	Policy: ilm.Policy{
		Phases: ilm.Phases{
			Hot: ilm.NewPhaseSetting(ilm.WithPriority(100),
				ilm.WithRollover(ilm.DefaultRollover()),
			),
			Warm: ilm.NewPhaseSetting(ilm.WithPriority(50),
				ilm.WithAllocate(nil, 0, ilm.Require, map[string]interface{}{
					"data": "warm",
				}),
			).SetMinAge(15, unit.Days),
			//Cold:   ilm.NewPhaseSetting(ilm.WithPriority(0)).SetMinAge(phaseAge.Cold, phaseAge.ColdUnit),
			Delete: ilm.NewPhaseSetting().SetMinAge(30, unit.Days),
		},
	},
}

var DefaultHotWarmCold = ilm.IlmObject{}
