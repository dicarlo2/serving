/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file contains logic to encapsulate flags which are needed to specify
// what cluster, etc. to use for e2e tests.

package test

import (
	"flag"
	"os"

	"github.com/knative/pkg/test"
	"github.com/knative/pkg/test/logging"
)

const (
	// ServingNamespace is the default namespace for serving e2e tests
	ServingNamespace = "serving-tests"

	// E2EMetricExporter is the name for the metrics exporter logger
	E2EMetricExporter = "e2e-metrics"
)

// ServingFlags holds the flags or defaults for knative/serving settings in the user's environment.
var ServingFlags = initializeServingFlags()

// ServingEnvironmentFlags holds the e2e flags needed only by the serving repo.
type ServingEnvironmentFlags struct {
	ResolvableDomain bool   // Resolve Route controller's `domainSuffix`
	DockerRepo       string // Docker repo (defaults to $KO_DOCKER_REPO)
	Tag              string // Test images version tag
}

func initializeServingFlags() *ServingEnvironmentFlags {
	var f ServingEnvironmentFlags

	flag.BoolVar(&f.ResolvableDomain, "resolvabledomain", false,
		"Set this flag to true if you have configured the `domainSuffix` on your Route controller to a domain that will resolve to your test cluster.")

	flag.StringVar(&f.DockerRepo, "dockerrepo", os.Getenv("KO_DOCKER_REPO"),
		"Provide the uri of the docker repo you have uploaded the test image to using `upload-test-images.sh`. Defaults to $KO_DOCKER_REPO")

	flag.StringVar(&f.Tag, "tag", "latest",
		"Provide the version tag for the test images.")

	flag.Parse()
	flag.Set("alsologtostderr", "true")
	logging.InitializeLogger(test.Flags.LogVerbose)

	if test.Flags.EmitMetrics {
		logging.InitializeMetricExporter(E2EMetricExporter)
	}

	return &f
}
