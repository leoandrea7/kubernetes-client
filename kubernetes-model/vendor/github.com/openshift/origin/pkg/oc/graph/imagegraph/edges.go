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
package imagegraph

import (
	"github.com/gonum/graph"

	imageapi "github.com/openshift/origin/pkg/image/apis/image"
	osgraph "github.com/openshift/origin/pkg/oc/graph/genericgraph"
	imagegraph "github.com/openshift/origin/pkg/oc/graph/imagegraph/nodes"
)

const (
	// ReferencedImageStreamGraphEdgeKind is an edge that goes from an ImageStreamTag node back to an ImageStream
	ReferencedImageStreamGraphEdgeKind = "ReferencedImageStreamGraphEdge"
	// ReferencedImageStreamImageGraphEdgeKind is an edge that goes from an ImageStreamImage node back to an ImageStream
	ReferencedImageStreamImageGraphEdgeKind = "ReferencedImageStreamImageGraphEdgeKind"
)

// AddImageStreamTagRefEdge ensures that a directed edge exists between an IST Node and the IS it references
func AddImageStreamTagRefEdge(g osgraph.MutableUniqueGraph, node *imagegraph.ImageStreamTagNode) {
	isName, _, _ := imageapi.SplitImageStreamTag(node.Name)
	imageStream := &imageapi.ImageStream{}
	imageStream.Namespace = node.Namespace
	imageStream.Name = isName

	imageStreamNode := imagegraph.FindOrCreateSyntheticImageStreamNode(g, imageStream)
	g.AddEdge(node, imageStreamNode, ReferencedImageStreamGraphEdgeKind)
}

// AddImageStreamImageRefEdge ensures that a directed edge exists between an ImageStreamImage Node and the IS it references
func AddImageStreamImageRefEdge(g osgraph.MutableUniqueGraph, node *imagegraph.ImageStreamImageNode) {
	dockImgRef, _ := imageapi.ParseDockerImageReference(node.Name)
	imageStream := &imageapi.ImageStream{}
	imageStream.Namespace = node.Namespace
	imageStream.Name = dockImgRef.Name

	imageStreamNode := imagegraph.FindOrCreateSyntheticImageStreamNode(g, imageStream)
	g.AddEdge(node, imageStreamNode, ReferencedImageStreamImageGraphEdgeKind)
}

// AddAllImageStreamRefEdges calls AddImageStreamRefEdge for every ImageStreamTagNode in the graph
func AddAllImageStreamRefEdges(g osgraph.MutableUniqueGraph) {
	for _, node := range g.(graph.Graph).Nodes() {
		if istNode, ok := node.(*imagegraph.ImageStreamTagNode); ok {
			AddImageStreamTagRefEdge(g, istNode)
		}
	}
}

// AddAllImageStreamImageRefEdges calls AddImageStreamImageRefEdge for every ImageStreamImageNode in the graph
func AddAllImageStreamImageRefEdges(g osgraph.MutableUniqueGraph) {
	for _, node := range g.(graph.Graph).Nodes() {
		if isimageNode, ok := node.(*imagegraph.ImageStreamImageNode); ok {
			AddImageStreamImageRefEdge(g, isimageNode)
		}
	}
}
