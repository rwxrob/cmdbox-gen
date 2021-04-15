# KN `gen` Static Site Generator CmdBox Tool


![WIP](https://img.shields.io/badge/status-wip-red.svg)
[![GoDoc](https://godoc.org/cmdbox-gen?status.svg)](https://godoc.org/cmdbox-gen)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/cmdbox-gen)](https://goreportcard.com/report/cmdbox-gen)
[![Coverage](https://gocover.io/_badge/cmdbox-gen)](https://gocover.io/cmdbox-gen)

The `gen` [CmdBox](https://github.com/rwxrob/cmdbox) tool generates the
`index.html` and sometimes the `README.md` from [KEG](https://keg.sh)
`data.yml` and `README.md` content and other optional external sources
using the same high-performance
[Goldmark](https://github.com/yuin/goldmark) Markdown engine and [Go
templating](https://golang.org/pkg/text/template/) that
[Hugo](https://gohugo.com) uses. Unlike Hugo, however, each KEG node is
independently generated in parallel making `gen` the fastest and most
extensible static site generator in the world. (But that might just be our
little secret. We aren't into silly branding and marketing hype.)

## Also Known As

Most will install `gen` as just that, but since it is included
in `kn`, the [KEG](https://keg.sh) tool as `kn gen` is has garnered the
moniker *Kngen* pronounced "KEN-gen" (with a schwa between the `k` and
`n` with heavy emphasis on the first syllable, almost like "Cajun" but
starting with "ken" instead).

## Summary

* `<node>/data.yml` ONLY data (no `gen` config)
* `<node>/data.yml` ONLY uppercase names
* `<node>/data.yml` no `---` by convention
* `<node>/data.yml` can be JSON by specification
* `<node>/.gen/config` YAML configuration
* `<node>/.gen/plugins` Go `template.funcMap` plugins
* `<node>/.gen/hooks/pre` will be run before generation
* `<node>/.gen/hooks/gen` overrides base generation
* `<node>/.gen/hooks/post` will be run after generation
* `<node>/.gen/tmpl/html` templates for `index.html`
* `<node>/.gen/tmpl/md` templates for `README.md`
* `<node>/.gen/tmpl/<type>` overrides type defaults
* `html` and `md` reserved within `tmpl`
* `html/body.*` replace default template
* `html/top.*` content before body
* `html/bottom.*` content after body
* `html/_.html` take over all templating
* `type` field changes templates
* Changes to any `.gen` apply to all children
* `<root>/.gen` generally has defaults
* `<root>` = `$GENROOT` || grandparent with `.git` || `$PWD`

## Examples

[Examples](testdata/examples) can be found within the `testdata`
directory.

## Reserved Fields `.gen/config`

Note that usually reserved fields will be in the root-level `data.yml`
file.

* `genVersion` version of the `gen` API (default `latest`)
* `kegVersion` version of KEG specification (default `latest`)

## Design Decisions

* Data driven by design, views over blogs
* Zero configuration, just works by convention
* No support for PDF since web browser support covers it
* Sustainable and extensible without sacrificing performance

## Legal Considerations

* [Apache License](LICENSE)
* [Contributors Notification](CONTRIBUTING)
* [Developer Certificate of Origin](DCO)

Copyright (c) 2021 Rob Muhlestein <mailto:rob@rwx.gg>

The content within this Git repository is released under the permissive
Apache 2 license which includes the granting of patent rights.

Contributors guarantee they have complied with the terms of the
[Developer Certificate of Origin](DCO) before making *any* contribution.
