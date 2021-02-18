package domain

import "github.com/thingworks/jarvis-go/common/domain"

type AutoIncreasable interface {
	GetContainerResourceId() string
	GetContainer() domain.ResourceType
	GetTarget() domain.ResourceType
	GetSubContainer() interface{}
	OnIndexGenerated()
}
