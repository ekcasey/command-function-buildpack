/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package command

import (
	"fmt"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/libcfbuildpack/build"
	"github.com/cloudfoundry/libcfbuildpack/detect"
	"github.com/projectriff/riff-buildpack/function"
)

type CommandBuildpack struct {
	id string
}

func (bp *CommandBuildpack) Id() string {
	return bp.id
}

func (bp *CommandBuildpack) Detect(d detect.Detect, m function.Metadata) (*buildplan.BuildPlan, error) {
	if detected, err := bp.detect(d, m); err != nil {
		return nil, err
	} else if detected {
		plan := BuildPlanContribution(d, m)
		return &plan, nil
	}
	// didn't detect
	return nil, nil
}

func (*CommandBuildpack) detect(d detect.Detect, m function.Metadata) (bool, error) {
	// Try command
	return DetectCommand(d, m)
}

func (*CommandBuildpack) Build(b build.Build) error {
	invoker, ok, err := NewCommandInvoker(b)
	if err != nil {
		return err
	} else if !ok {
		return fmt.Errorf("buildpack passed detection but did not know how to actually build")
	}
	return invoker.Contribute()
}

func NewBuildpack() function.Buildpack {
	return &CommandBuildpack{
		id: "command",
	}
}
