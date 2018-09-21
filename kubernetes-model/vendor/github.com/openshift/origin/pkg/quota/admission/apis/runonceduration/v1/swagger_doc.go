/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package v1

// This file contains methods that can be used by the go-restful package to generate Swagger
// documentation for the object types found in 'types.go' This file is automatically generated
// by hack/update-generated-swagger-descriptions.sh and should be run after a full build of OpenShift.
// ==== DO NOT EDIT THIS FILE MANUALLY ====

var map_RunOnceDurationConfig = map[string]string{
	"": "RunOnceDurationConfig is the configuration for the RunOnceDuration plugin. It specifies a maximum value for ActiveDeadlineSeconds for a run-once pod. The project that contains the pod may specify a different setting. That setting will take precedence over the one configured for the plugin here.",
	"activeDeadlineSecondsOverride": "ActiveDeadlineSecondsOverride is the maximum value to set on containers of run-once pods Only a positive value is valid. Absence of a value means that the plugin won't make any changes to the pod It is kept this way for compatibility. Only change it in a new version of the API.",
}

func (RunOnceDurationConfig) SwaggerDoc() map[string]string {
	return map_RunOnceDurationConfig
}
