#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

# Copyright (c) 2023  The Go-Curses Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

BE_LOCAL_PATH ?= ../be

APP_NAME    ?= be-go-curses-org
APP_SUMMARY ?= go-curses.org

DENY_DURATION ?= 600

COMMON_TAGS = htmlify,stock_pgc,page_robots,header_proxy,papertrail
BUILD_TAGS = prd,embeds,$(COMMON_TAGS)
DEV_BUILD_TAGS = dev,locals,$(COMMON_TAGS)
EXTRA_PKGS =

include ./Enjin.mk
