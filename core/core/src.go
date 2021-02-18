package core

import "github.com/thingworks/jarvis/common/domain"

type CodedCommonResource struct {
	Code          string
	ResourceId    string
	ResourceType  domain.ResourceType
	SubResourceId string
}
