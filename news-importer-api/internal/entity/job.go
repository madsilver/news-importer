package entity

import (
	"bytes"
	"encoding/json"
)

type Job struct {
	ID                int64
	Name              string
	Owner             string
	ProducerName      *string
	Active            int
	FunctionID        int
	Arguments         string
	ExtraArguments    *string
	Priority          int
	RunInterval       *string
	LastRun           *string
	LastRunMessage    *string
	NextRun           *string
	RunTimeout        *string
	PossibleRetries   int
	LastErrors        int
	ConcurrentRun     *string
	HighPriority      *string
	DispatchedAt      *string
	StartedAt         *string
	LastRunHash       *string
	Destination       *string
	SourceName        *string
	SourceUrl         *string
	SourceSection     *string
	SourceChannel     *string
	SourceProducer    *string
	RemovalPolicyType *string
	RemovalPolicyArgs int
	Comments          *string
	Deleted           int
	AlarmLimit        int
	Alarmed           int
}

func (j *Job) Serialize() ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(j)
	return b.Bytes(), err
}

func Deserialize(b []byte) (Job, error) {
	var job Job
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&job)
	return job, err
}
