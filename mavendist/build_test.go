/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mavendist_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/buildpacks/libcnb"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"

	"github.com/paketo-buildpacks/maven-dist/v1/mavendist"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		ctx        libcnb.BuildContext
		mavenBuild mavendist.Build
	)

	it.Before(func() {
		var err error

		ctx.Application.Path, err = ioutil.TempDir("", "build-application")
		Expect(err).NotTo(HaveOccurred())

		ctx.Buildpack.Metadata = map[string]interface{}{
			"configurations": []map[string]interface{}{
				{"name": "BP_MAVEN_BUILD_ARGUMENTS", "default": "test-argument"},
			},
		}

		ctx.Layers.Path, err = ioutil.TempDir("", "build-layers")
		Expect(err).NotTo(HaveOccurred())
		mavenBuild = mavendist.Build{}

		//mvnwFilepath = filepath.Join(ctx.Application.Path, "mvnw")
	})

	it.After(func() {
		Expect(os.RemoveAll(ctx.Application.Path)).To(Succeed())
		Expect(os.RemoveAll(ctx.Layers.Path)).To(Succeed())
	})

	it("contributes distribution for API 0.7+", func() {
		ctx.Buildpack.Metadata["dependencies"] = []map[string]interface{}{
			{
				"id":      "maven",
				"version": "1.1.1",
				"stacks":  []interface{}{"test-stack-id"},
				"cpes":    []string{"cpe:2.3:a:apache:maven:3.8.3:*:*:*:*:*:*:*"},
				"purl":    "pkg:generic/apache-maven@3.8.3",
			},
		}
		ctx.StackID = "test-stack-id"
		planEntry := libcnb.BuildpackPlanEntry{
			Name: "maven",
			Metadata: map[string]interface{}{
				"command": "mvn",
			},
		}
		ctx.Plan.Entries = append(ctx.Plan.Entries, planEntry)

		result, err := mavenBuild.Build(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(result.Layers).To(HaveLen(2))
		Expect(result.Layers[0].Name()).To(Equal("maven"))

		Expect(result.BOM.Entries).To(HaveLen(1))
		Expect(result.BOM.Entries[0].Name).To(Equal("maven"))
		Expect(result.BOM.Entries[0].Build).To(BeTrue())
		Expect(result.BOM.Entries[0].Launch).To(BeFalse())
	})

	it("contributes distribution for API <=0.6", func() {
		ctx.Buildpack.Metadata["dependencies"] = []map[string]interface{}{
			{
				"id":      "maven",
				"version": "1.1.1",
				"stacks":  []interface{}{"test-stack-id"},
			},
		}
		ctx.StackID = "test-stack-id"
		ctx.Buildpack.API = "0.6"
		planEntry := libcnb.BuildpackPlanEntry{
			Name: "maven",
			Metadata: map[string]interface{}{
				"command": "mvn",
			},
		}
		ctx.Plan.Entries = append(ctx.Plan.Entries, planEntry)

		result, err := mavenBuild.Build(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(result.Layers).To(HaveLen(2))
		Expect(result.Layers[0].Name()).To(Equal("maven"))

		Expect(result.BOM.Entries).To(HaveLen(1))
		Expect(result.BOM.Entries[0].Name).To(Equal("maven"))
		Expect(result.BOM.Entries[0].Build).To(BeTrue())
		Expect(result.BOM.Entries[0].Launch).To(BeFalse())
	})

	it("contributes mvnd distribution for API 0.7+", func() {
		ctx.Buildpack.Metadata["dependencies"] = []map[string]interface{}{
			{
				"id":      "mvnd",
				"version": "1.1.1",
				"stacks":  []interface{}{"test-stack-id"},
				"cpes":    []string{"cpe:2.3:a:apache:mvnd:0.7.1:*:*:*:*:*:*:*"},
				"purl":    "pkg:generic/apache-mvnd@0.7.1",
			},
		}
		ctx.StackID = "test-stack-id"
		planEntry := libcnb.BuildpackPlanEntry{
			Name: "maven",
			Metadata: map[string]interface{}{
				"command": "mvnd",
			},
		}
		ctx.Plan.Entries = append(ctx.Plan.Entries, planEntry)

		result, err := mavenBuild.Build(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(result.Layers).To(HaveLen(2))
		Expect(result.Layers[0].Name()).To(Equal("mvnd"))

		Expect(result.BOM.Entries).To(HaveLen(1))
		Expect(result.BOM.Entries[0].Name).To(Equal("mvnd"))
		Expect(result.BOM.Entries[0].Build).To(BeTrue())
		Expect(result.BOM.Entries[0].Launch).To(BeFalse())
	})

	it("contributes mvnd distribution for API <=0.6", func() {
		ctx.Buildpack.Metadata["dependencies"] = []map[string]interface{}{
			{
				"id":      "mvnd",
				"version": "1.1.1",
				"stacks":  []interface{}{"test-stack-id"},
			},
		}
		ctx.StackID = "test-stack-id"
		ctx.Buildpack.API = "0.6"
		planEntry := libcnb.BuildpackPlanEntry{
			Name: "maven",
			Metadata: map[string]interface{}{
				"command": "mvnd",
			},
		}
		ctx.Plan.Entries = append(ctx.Plan.Entries, planEntry)

		result, err := mavenBuild.Build(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(result.Layers).To(HaveLen(2))
		Expect(result.Layers[0].Name()).To(Equal("mvnd"))

		Expect(result.BOM.Entries).To(HaveLen(1))
		Expect(result.BOM.Entries[0].Name).To(Equal("mvnd"))
		Expect(result.BOM.Entries[0].Build).To(BeTrue())
		Expect(result.BOM.Entries[0].Launch).To(BeFalse())
	})

	context("maven settings bindings exists", func() {
		var result libcnb.BuildResult

		it.Before(func() {
			var err error
			ctx.StackID = "test-stack-id"
			ctx.Platform.Path, err = ioutil.TempDir("", "maven-test-platform")
			ctx.Platform.Bindings = libcnb.Bindings{
				{
					Name:   "some-maven",
					Type:   "maven",
					Secret: map[string]string{"settings.xml": "maven-settings-content"},
					Path:   filepath.Join(ctx.Platform.Path, "bindings", "some-maven"),
				},
			}
			ctx.Buildpack.Metadata["dependencies"] = []map[string]interface{}{
				{
					"id":      "maven",
					"version": "1.1.1",
					"stacks":  []interface{}{"test-stack-id"},
				},
			}
			planEntry := libcnb.BuildpackPlanEntry{
				Name: "maven",
				Metadata: map[string]interface{}{
					"command": "mvn",
				},
			}
			ctx.Plan.Entries = append(ctx.Plan.Entries, planEntry)
			mavenSettingsPath, ok := ctx.Platform.Bindings[0].SecretFilePath("settings.xml")
			Expect(os.MkdirAll(filepath.Dir(mavenSettingsPath), 0777)).To(Succeed())
			Expect(ok).To(BeTrue())
			Expect(ioutil.WriteFile(
				mavenSettingsPath,
				[]byte("maven-settings-content"),
				0644,
			)).To(Succeed())

			result, err = mavenBuild.Build(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Layers).To(HaveLen(2))
		})

		it.After(func() {
			Expect(os.RemoveAll(ctx.Platform.Path)).To(Succeed())
		})

		it("adds does not fail to build", func() {
		})
	})

	context("maven settings incl. settings-security bindings exists", func() {
		var result libcnb.BuildResult

		it.Before(func() {
			var err error
			ctx.StackID = "test-stack-id"
			ctx.Platform.Path, err = ioutil.TempDir("", "maven-test-platform")
			ctx.Platform.Bindings = libcnb.Bindings{
				{
					Name: "some-maven",
					Type: "maven",
					Secret: map[string]string{
						"settings.xml":          "maven-settings-content",
						"settings-security.xml": "maven-settings-security-content",
					},
					Path: filepath.Join(ctx.Platform.Path, "bindings", "some-maven"),
				},
			}
			ctx.Buildpack.Metadata["dependencies"] = []map[string]interface{}{
				{
					"id":      "maven",
					"version": "1.1.1",
					"stacks":  []interface{}{"test-stack-id"},
				},
			}
			planEntry := libcnb.BuildpackPlanEntry{
				Name: "maven",
				Metadata: map[string]interface{}{
					"command": "mvn",
				},
			}
			ctx.Plan.Entries = append(ctx.Plan.Entries, planEntry)
			mavenSettingsPath, ok := ctx.Platform.Bindings[0].SecretFilePath("settings.xml")
			Expect(os.MkdirAll(filepath.Dir(mavenSettingsPath), 0777)).To(Succeed())
			Expect(ok).To(BeTrue())
			Expect(ioutil.WriteFile(
				mavenSettingsPath,
				[]byte("maven-settings-content"),
				0644,
			)).To(Succeed())

			mavenSettingsSecurityPath, ok := ctx.Platform.Bindings[0].SecretFilePath("settings-security.xml")
			Expect(os.MkdirAll(filepath.Dir(mavenSettingsSecurityPath), 0777)).To(Succeed())
			Expect(ok).To(BeTrue())
			Expect(ioutil.WriteFile(
				mavenSettingsSecurityPath,
				[]byte("maven-settings-security-content"),
				0644,
			)).To(Succeed())

			result, err = mavenBuild.Build(ctx)
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Layers).To(HaveLen(2))

		})

		it.After(func() {
			Expect(os.RemoveAll(ctx.Platform.Path)).To(Succeed())
		})

		it("does not faile to build", func() {
		})
	})
}
