// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aip0235

import (
	"github.com/googleapis/api-linter/lint"
	"github.com/googleapis/api-linter/rules/internal/utils"
	"github.com/jhump/protoreflect/desc"
)

var requestRequestsBehavior = &lint.FieldRule{
	Name: lint.NewRuleName(235, "request-requests-behavior"),
	OnlyIf: func(f *desc.FieldDescriptor) bool {
		return isBatchDeleteRequestMessage(f.GetOwner()) && f.GetName() == "requests"
	},
	LintField: func(f *desc.FieldDescriptor) []lint.Problem {
		if !utils.GetFieldBehavior(f).Contains("REQUIRED") {
			return []lint.Problem{{
				Message:    "Batch Delete requests: The `requests` field should include `(google.api.field_behavior) = REQUIRED`.",
				Descriptor: f,
			}}
		}
		return nil
	},
}
