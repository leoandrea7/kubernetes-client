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
package image_ecosystem

import (
	"fmt"
	"time"

	g "github.com/onsi/ginkgo"
	o "github.com/onsi/gomega"

	e2e "k8s.io/kubernetes/test/e2e/framework"

	exutil "github.com/openshift/origin/test/extended/util"
)

var _ = g.Describe("[image_ecosystem][php][Slow] hot deploy for openshift php image", func() {
	defer g.GinkgoRecover()
	var (
		cakephpTemplate = "https://raw.githubusercontent.com/openshift/cakephp-ex/master/openshift/templates/cakephp-mysql.json"
		oc              = exutil.NewCLI("s2i-php", exutil.KubeConfigPath())
		hotDeployParam  = "OPCACHE_REVALIDATE_FREQ=0"
		modifyCommand   = []string{"sed", "-ie", `s/\$result\['c'\]/1337/`, "app/View/Layouts/default.ctp"}
		pageCountFn     = func(count int) string { return fmt.Sprintf(`<span class="code" id="count-value">%d</span>`, count) }
		dcName          = "cakephp-mysql-example-1"
		dcLabel         = exutil.ParseLabelsOrDie(fmt.Sprintf("deployment=%s", dcName))
	)

	g.Context("", func() {
		g.BeforeEach(func() {
			exutil.DumpDockerInfo()
		})

		g.AfterEach(func() {
			if g.CurrentGinkgoTestDescription().Failed {
				exutil.DumpPodStates(oc)
				exutil.DumpPodLogsStartingWith("", oc)
			}
		})

		g.Describe("CakePHP example", func() {
			g.It(fmt.Sprintf("should work with hot deploy"), func() {
				oc.SetOutputDir(exutil.TestContext.OutputDir)

				exutil.CheckOpenShiftNamespaceImageStreams(oc)
				g.By(fmt.Sprintf("calling oc new-app -f %q -p %q", cakephpTemplate, hotDeployParam))
				err := oc.Run("new-app").Args("-f", cakephpTemplate, "-p", hotDeployParam).Execute()
				o.Expect(err).NotTo(o.HaveOccurred())

				g.By("waiting for build to finish")
				err = exutil.WaitForABuild(oc.BuildClient().Build().Builds(oc.Namespace()), dcName, nil, nil, nil)
				if err != nil {
					exutil.DumpBuildLogs("cakephp-mysql-example", oc)
				}
				o.Expect(err).NotTo(o.HaveOccurred())

				err = exutil.WaitForDeploymentConfig(oc.KubeClient(), oc.AppsClient().Apps(), oc.Namespace(), "cakephp-mysql-example", 1, oc)
				o.Expect(err).NotTo(o.HaveOccurred())

				g.By("waiting for endpoint")
				err = e2e.WaitForEndpoint(oc.KubeFramework().ClientSet, oc.Namespace(), "cakephp-mysql-example")
				o.Expect(err).NotTo(o.HaveOccurred())

				assertPageCountIs := func(i int) {
					_, err := exutil.WaitForPods(oc.KubeClient().Core().Pods(oc.Namespace()), dcLabel, exutil.CheckPodIsRunningFn, 1, 4*time.Minute)
					o.ExpectWithOffset(1, err).NotTo(o.HaveOccurred())

					result, err := CheckPageContains(oc, "cakephp-mysql-example", "", pageCountFn(i))
					o.ExpectWithOffset(1, err).NotTo(o.HaveOccurred())
					o.ExpectWithOffset(1, result).To(o.BeTrue())
				}

				g.By("checking page count")

				assertPageCountIs(1)
				assertPageCountIs(2)

				g.By("modifying the source code with hot deploy enabled")
				err = RunInPodContainer(oc, dcLabel, modifyCommand)
				o.Expect(err).NotTo(o.HaveOccurred())
				g.By("checking page count after modifying the source code")
				assertPageCountIs(1337)
			})
		})
	})
})
