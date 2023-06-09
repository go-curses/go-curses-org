//go:build prd

// Copyright (c) 2023  The Go-Curses Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"embed"

	"github.com/go-enjin/be/features/fs/embeds/content"
	publicEmbed "github.com/go-enjin/be/features/fs/embeds/public"
	"github.com/go-enjin/be/pkg/feature"
	"github.com/go-enjin/be/pkg/log"
	"github.com/go-enjin/be/pkg/theme"
)

//go:embed themes/**
//go:embed themes/go-curses/layouts/_default/**
var themeFs embed.FS

func getTheme() (t *theme.Theme) {
	var err error
	if t, err = theme.NewEmbed("themes/go-curses", themeFs); err != nil {
		log.FatalF("error loading embedded go-curses theme: %v", err)
	}
	return
}

//go:embed public/**
var publicFs embed.FS

func getPublicFeature() (f feature.Feature) {
	f = publicEmbed.New().
		MountPathFs("/", "public", publicFs).
		Make()
	return
}

//go:embed content/**
var contentFs embed.FS

func getContentFeature() (f feature.Feature) {
	f = content.New().
		MountPathFs("/", "content", contentFs).
		Make()
	return
}
