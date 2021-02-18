package event

import (
	"github.com/thingworks/jarvis/core/indicator/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IndicatorDeleted struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Indicator     domain.Indicator
	UserId        string
	EventTime     int64
	DomainVersion int64
	IndicatorId   string
}

func (indicatorDel *IndicatorDeleted) CollectionName() string {
	return "indicator_event"
}

func (indicatorDel *IndicatorDeleted) ObjectId() primitive.ObjectID {
	return indicatorDel.Id
}

func (indicatorDel *IndicatorDeleted) Init() {
}
