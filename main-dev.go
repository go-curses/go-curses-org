//go:build !production

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
	"github.com/go-enjin/be/features/fs/content"
	"github.com/go-enjin/be/features/fs/public"
	"github.com/go-enjin/be/features/fs/themes"
	"github.com/go-enjin/be/pkg/log"
)

func init() {
	// locals environment, early startup debug logging
	log.Config.LogLevel = log.LevelDebug
	log.Config.Apply()

	fPublic = public.New().
		MountLocalPath("/", "public").
		Make()

	fContent = content.New().
		MountLocalPath("/", "content").
		AddToIndexProviders(gPagesPqlFeature).
		Make()

	fThemes = themes.New().
		LocalTheme("themes/go-curses").
		SetTheme("go-curses").
		Make()

	hotReload = true
}