package _go

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/enumflags"
	"github.com/metamatex/metamate/metactl/pkg/v0/types"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils"
	"github.com/metamatex/metamate/metactl/pkg/v0/utils/ptr"
	"github.com/pkg/errors"
	"os"
)

const (
	TaskSetTypes           = "types"
	TaskSetService         = "server"
	TaskSetClient          = "client"
	TaskSetHttpJson        = "httpjson"
	DependecyRelationNames = "dependencyRelationNames"
	DependecyLookupService = "dependencyLookupService"
)

const (
	SdkService = "go-service"
	SdkClient  = "go-client"
)

const (
	DataPackage = "package"
	DataName    = "name"
	RootPath    = "gen/"
	Path        = RootPath + "v0/mql/"
)

var tasks = map[string]types.RenderTask{}

func initSdk(g *types.SdkGenerator, c types.SdkConfig) (err error) {
	_, ok := c.Args[DataPackage]
	if !ok {
		err = errors.New("data.package is missing")

		return
	}

	prepareTasks(c.Args, g.Dependencies)
	prepareTasks(c.Args, g.Tasks)

	return
}

func prepareTasks(data map[string]interface{}, tasks []types.RenderTask) {
	for i, _ := range tasks {
		tasks[i].TemplateData = ptr.String("// generated by metactl sdk gen \n" + *tasks[i].TemplateData)
		tasks[i].Data = data
		tasks[i].Out = ptr.String(Path + *tasks[i].Out)
	}
}

func initServiceSdk(g *types.SdkGenerator, c types.SdkConfig) (err error) {
	_, ok := c.Args[DataPackage]
	if !ok {
		err = errors.New("data.package is missing")

		return
	}

	_, ok = c.Args[DataName]
	if !ok {
		err = errors.New("data.name is missing")

		return
	}

	prepareTasks(c.Args, g.Dependencies)
	prepareTasks(c.Args, g.Tasks)

	return
}

func resetSdk(c types.SdkConfig) (err error) {
	err = os.RemoveAll(RootPath)
	if err != nil {
		return
	}

	return
}

func GetSdks() []types.SdkGenerator {
	taskSets := map[string][]types.RenderTask{}

	taskSets[TaskSetTypes] = []types.RenderTask{
		tasks[TaskTypes],
		tasks[TaskEnums],
		tasks[TaskUtils],
		tasks[TaskUtilsPtr],
		tasks[TaskVersion],
	}

	taskSets[TaskSetClient] = []types.RenderTask{
		tasks[TaskTypedClient],
	}

	taskSets[TaskSetService] = []types.RenderTask{
		tasks[TaskServiceInterface],
		tasks[TaskTypedServer],
	}

	taskSets[TaskSetHttpJson] = []types.RenderTask{
		tasks[TaskHeader],
	}

	taskSets[DependecyRelationNames] = []types.RenderTask{
		{
			TemplateData: tasks[TaskEnums].TemplateData,
			Out:          tasks[TaskEnums].Out,
			Filter: &types.NodeFilters{
				Enums: &graph.Filter{
					Flags: &graph.FlagsSubset{
						Or: []string{enumflags.IsRelationNames},
					},
				},
			},
			Iterate: ptr.String(graph.ENUM),
		},
	}

	taskSets[DependecyLookupService] = []types.RenderTask{
		{
			TemplateData: tasks[TaskTypes].TemplateData,
			Out:          tasks[TaskTypes].Out,
			Dependencies: &types.RenderTaskDependencies{
				Endpoints: &graph.Filter{
					Names: &graph.NamesSubset{
						Or: []string{endpointnames.LookupService},
					},
				},
			},
			Iterate: ptr.String(graph.TYPE),
		},
	}

	return []types.SdkGenerator{
		{
			Name:         SdkClient,
			Description:  "go client sdk",
			Tasks:        utils.ConcatTaskSets(taskSets[TaskSetTypes], taskSets[TaskSetHttpJson], taskSets[TaskSetClient]),
			Init:         initSdk,
			Reset:        resetSdk,
			Dependencies: append(taskSets[DependecyRelationNames], taskSets[DependecyLookupService]...),
		},
		{
			Name:         SdkService,
			Description:  "go server sdk",
			Tasks:        utils.ConcatTaskSets(taskSets[TaskSetTypes], taskSets[TaskSetHttpJson], taskSets[TaskSetService]),
			Init:         initServiceSdk,
			Reset:        resetSdk,
			Dependencies: append(taskSets[DependecyRelationNames], taskSets[DependecyLookupService]...),
		},
	}
}
