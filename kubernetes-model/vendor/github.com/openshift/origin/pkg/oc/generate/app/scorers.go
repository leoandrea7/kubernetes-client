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
package app

import (
	"strings"

	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	templateapi "github.com/openshift/origin/pkg/template/apis/template"
)

func templateScorer(template templateapi.Template, term string) (float32, bool) {
	score := stringProximityScorer(template.Name, term)
	return score, score < 0.3
}

func imageStreamScorer(imageStream imageapi.ImageStream, term string) (float32, bool) {
	score := stringProximityScorer(imageStream.Name, term)
	return score, score < 0.3
}

func stringProximityScorer(s, query string) float32 {
	sLower := strings.ToLower(s)
	queryLower := strings.ToLower(query)

	var score float32
	switch {
	case query == "*":
		score = 0.0
	case s == query:
		score = 0.0
	case strings.EqualFold(s, query):
		score = 0.02
	case strings.HasPrefix(s, query):
		score = 0.1
	case strings.HasPrefix(sLower, queryLower):
		score = 0.12
	case strings.Contains(s, query):
		score = 0.2
	case strings.Contains(sLower, queryLower):
		score = 0.22
	default:
		score = 1.0
	}

	return score
}

func partialScorer(a, b string, prefix bool, partial, none float32) (bool, float32) {
	switch {
	// If either one is empty, it's a partial match because the values do not conflict.
	case len(a) == 0 && len(b) != 0, len(a) != 0 && len(b) == 0:
		return true, partial
	case a != b:
		if prefix {
			if strings.HasPrefix(a, b) || strings.HasPrefix(b, a) {
				return true, partial
			}
		}
		return false, none
	default:
		return true, 0.0
	}
}
