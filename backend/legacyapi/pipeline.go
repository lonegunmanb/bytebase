package api

import (
	"encoding/json"
)

// PipelineStatus is the status for pipelines.
type PipelineStatus string

const (
	// PipelineOpen is the pipeline status for OPEN.
	PipelineOpen PipelineStatus = "OPEN"
	// PipelineDone is the pipeline status for DONE.
	PipelineDone PipelineStatus = "DONE"
	// PipelineCanceled is the pipeline status for CANCELED.
	PipelineCanceled PipelineStatus = "CANCELED"
)

// Pipeline is the API message for pipelines.
type Pipeline struct {
	ID int `jsonapi:"primary,pipeline"`

	// Related fields
	StageList []*Stage `jsonapi:"relation,stage"`

	// Domain specific fields
	Name   string `jsonapi:"attr,name"`
	Status PipelineStatus
}

// PipelineCreate is the API message for creating a pipeline.
type PipelineCreate struct {
	// Standard fields
	// Value is assigned from the jwt subject field passed by the client.
	CreatorID int

	// Related fields
	StageList []StageCreate `jsonapi:"attr,stageList"`

	// Domain specific fields
	Name string `jsonapi:"attr,name"`
}

// PipelineFind is the API message for finding pipelines.
type PipelineFind struct {
	ID *int

	// Domain specific fields
	Active *bool
}

func (find *PipelineFind) String() string {
	str, err := json.Marshal(*find)
	if err != nil {
		return err.Error()
	}
	return string(str)
}