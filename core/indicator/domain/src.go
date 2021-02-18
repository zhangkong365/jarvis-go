package domain

import (
	"github.com/thingworks/jarvis-go/common"
	"github.com/thingworks/jarvis-go/core/core"
	"github.com/thingworks/jarvis-go/core/process/domain/counter"
	"github.com/thingworks/jarvis-go/label"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IndicatorType int

const (
	DOUBLE  IndicatorType = 0
	FORMULA IndicatorType = 1
)

type Indicator struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IndicatorId    string
	Name           string
	Code           string
	DomainVersion  int64
	TenantId       string
	NamespaceId    string
	Formula        string
	Unit           string
	Active         string
	Transitional   string
	Type           IndicatorType
	LastCacheTime  time.Time
	Dependencies   []core.CodedCommonResource
	FinalCounter   int64
	ChangeCounter  int64
	LatestCounter  counter.DependencyCounter
	CurrentCounter counter.DependencyCounter
	Labels         []label.Label
	Folder         common.FolderInfo
}

func (i *Indicator) CollectionName() string {
	return "indicator"
}

func (i *Indicator) ObjectId() primitive.ObjectID {
	return i.Id
}

func (i *Indicator) Init() {
}
