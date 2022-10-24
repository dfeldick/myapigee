package apigee

import (
	"github.com/Axway/agent-sdk/pkg/jobs"
)

type registerAPIValidatorJob struct {
	jobs.Job
	proxiesReady      jobFirstRunDone
	registerValidator func()
}

func newRegisterAPIValidatorJob(proxiesReady jobFirstRunDone, registerValidator func()) *registerAPIValidatorJob {
	job := &registerAPIValidatorJob{
		proxiesReady:      proxiesReady,
		registerValidator: registerValidator,
	}
	return job
}

func (j *registerAPIValidatorJob) Ready() bool {
	return j.proxiesReady()
}

func (j *registerAPIValidatorJob) Status() error {
	return nil
}

func (j *registerAPIValidatorJob) Execute() error {
	j.registerValidator()
	return nil
}
