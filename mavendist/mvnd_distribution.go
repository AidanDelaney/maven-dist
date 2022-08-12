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

package mavendist

import (
	"fmt"
	"os"
	"strings"

	"github.com/buildpacks/libcnb"
	"github.com/paketo-buildpacks/libpak"
	"github.com/paketo-buildpacks/libpak/bard"
	"github.com/paketo-buildpacks/libpak/crush"
)

type MvndDistribution struct {
	LayerContributor libpak.DependencyLayerContributor
	Logger           bard.Logger
	Settings         []string
}

func NewMvndDistribution(dependency libpak.BuildpackDependency, cache libpak.DependencyCache, settings []string) (MvndDistribution, libcnb.BOMEntry) {
	contributor, entry := libpak.NewDependencyLayer(dependency, cache, libcnb.LayerTypes{
		Build: true,
		Cache: true,
	})
	return MvndDistribution{LayerContributor: contributor, Settings: settings}, entry
}

func (d MvndDistribution) Contribute(layer libcnb.Layer) (libcnb.Layer, error) {
	d.LayerContributor.Logger = d.Logger

	return d.LayerContributor.Contribute(layer, func(artifact *os.File) (libcnb.Layer, error) {
		d.Logger.Bodyf("Expanding to %s", layer.Path)
		if err := crush.ExtractZip(artifact, layer.Path, 1); err != nil {
			return libcnb.Layer{}, fmt.Errorf("unable to expand Maven\n%w", err)
		}
		if len(d.Settings) != 0 {
			args := strings.Join(d.Settings, " ")
			layer.BuildEnvironment.Override(BpMavenSecurityArgs, args)
		}
		return layer, nil
	})
}

func (d MvndDistribution) Name() string {
	return d.LayerContributor.LayerName()
}
