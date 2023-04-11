// Copyright (c) 2022  The Go-Enjin Authors
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
	"fmt"
	"os"

	"github.com/go-enjin/be/features/outputs/htmlify"
	pgc "github.com/go-enjin/be/features/pages/caching/stock-pgc"
	"github.com/go-enjin/golang-org-x-text/language"

	"github.com/go-enjin/be"
	"github.com/go-enjin/be/features/log/papertrail"
	"github.com/go-enjin/be/features/pages/formats"
	"github.com/go-enjin/be/features/pages/robots"
	"github.com/go-enjin/be/features/requests/headers/proxy"
	"github.com/go-enjin/be/pkg/lang"
)

var (
	SiteTag     = "GCO"
	SiteName    = "Go-Curses"
	SiteTagLine = "The official Go-Curses project website."
)

func main() {
	enjin := be.New().
		SiteTag(SiteTag).
		SiteName(SiteName).
		SiteTagLine(SiteTagLine).
		SiteCopyrightName(SiteName).
		SiteCopyrightNotice("All rights reserved").
		AddFeature(proxy.New().Enable().Make()).
		AddFeature(formats.New().Defaults().Make()).
		AddFeature(pgc.New().Make()).
		AddFeature(htmlify.New().Make()).
		SiteDefaultLanguage(language.English).
		SiteSupportedLanguages(language.English).
		SiteLanguageMode(lang.NewPathMode().Make()).
		AddTheme(getTheme()).
		SetTheme("go-curses").
		AddFeature(papertrail.Make()).
		AddFeature(robots.New().
			AddRuleGroup(robots.NewRuleGroup().
				AddUserAgent("*").AddAllowed("/").Make(),
			).Make(),
		).
		AddFeature(getPublicFeature()).
		AddFeature(getContentFeature()).
		SetStatusPage(404, "/").
		SetStatusPage(500, "/")
	// add content and status pages
	if err := enjin.Build().Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "enjin.Run error: %v\n", err)
		os.Exit(1)
	}
}
