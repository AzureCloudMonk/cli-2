package models

import (
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

type JobV1AlphaMetadata struct {
	Name       string      `json:"name,omitempty"`
	Id         string      `json:"id,omitempty"`
	CreateTime json.Number `json:"create_time,omitempty,string" yaml:"create_time,omitempty"`
	UpdateTime json.Number `json:"update_time,omitempty,string" yaml:"update_time,omitempty"`
	StartTime  json.Number `json:"start_time,omitempty,string" yaml:"start_time,omitempty"`
	FinishTime json.Number `json:"finish_time,omitempty,string" yaml:"finish_time,omitempty"`
}

type JobV1AlphaSpec struct {
	Agent struct {
		Machine struct {
			Type    string `json:"type,omitempty" yaml:"type,omitempty"`
			OsImage string `json:"os_image,omitempty" yaml:"type,omitempty"`
		} `json:"machine,omitempty" yaml:"machine,omitempty"`
	} `json:"agent,omitempty" yaml:"agent,omitempty"`

	ProjectId string `json:"project_id,omitempty" yaml:"project_id,omitempty"`
}

type JobV1AlphaStatus struct {
	State  string `json:"state" yaml:"state"`
	Result string `json:"result" yaml:"result"`
	Agent  struct {
		Ip    string `json:"ip" yaml:"ip"`
		Ports []struct {
			Name   string `json:"name" yaml:"name"`
			Number int32  `json:"number" yaml:"number"`
		} `json:"ports,omitempty"`
	} `json:"agent,omitempty"`
}

type JobV1Alpha struct {
	ApiVersion string `json:"apiVersion,omitempty" yaml:"apiVersion"`
	Kind       string `json:"kind,omitempty" yaml:"kind"`

	Metadata *JobV1AlphaMetadata `json:"metadata,omitempty"`
	Spec     *JobV1AlphaSpec     `json:"spec,omitempty"`
	Status   *JobV1AlphaStatus   `json:"status,omitempty"`
}

func NewJobV1Alpha(name string) JobV1Alpha {
	j := JobV1Alpha{}

	j.Metadata = &JobV1AlphaMetadata{}
	j.Metadata.Name = name
	j.setApiVersionAndKind()

	return j
}

func NewJobV1AlphaFromJson(data []byte) (*JobV1Alpha, error) {
	j := JobV1Alpha{}

	err := json.Unmarshal(data, &j)

	if err != nil {
		return nil, err
	}

	j.setApiVersionAndKind()

	return &j, nil
}

func NewJobV1AlphaFromYaml(data []byte) (*JobV1Alpha, error) {
	j := JobV1Alpha{}

	err := yaml.UnmarshalStrict(data, &j)

	if err != nil {
		return nil, err
	}

	j.setApiVersionAndKind()

	return &j, nil
}

func (j *JobV1Alpha) setApiVersionAndKind() {
	j.ApiVersion = "v1alpha"
	j.Kind = "Job"
}

func (j *JobV1Alpha) ObjectName() string {
	return fmt.Sprintf("Jobs/%s", j.Metadata.Id)
}

func (j *JobV1Alpha) ToJson() ([]byte, error) {
	return json.Marshal(j)
}

func (j *JobV1Alpha) ToYaml() ([]byte, error) {
	return yaml.Marshal(j)
}
