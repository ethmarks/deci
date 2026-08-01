---
title: Deci
description: deci is a terminal text editor like nano
layout: layouts/home.vto
blogList:
  filter: layout=layouts/blog-single.vto
  order: date=desc
  limit: 5
  showAuthor: true
  showDate: true
  showMins: false
---

# Deci

Your second-to-last next editor {.subtitle}

<div>
	<a href="https://github.com/ethmarks/deci/releases"><img src="https://img.shields.io/github/v/release/ethmarks/deci" alt="GitHub Release" /></a>
	<a href="https://pkg.go.dev/github.com/ethmarks/deci"><img src="https://pkg.go.dev/badge/github.com/ethmarks/deci.svg" alt="Go Reference" /></a>
	<a href="https://github.com/ethmarks/deci"><img src="https://img.shields.io/badge/github-repo-blue?logo=github" alt="GitHub" /></a>
	<a href="https://ethmarks.github.io/deci/"><img src="https://img.shields.io/badge/demo-live-green" alt="Demo" /></a>
</div>

Deci is a terminal text editor like [nano](https://nano-editor.org/).

![](/uploads/demo.gif)

## Installation

The easiest way to install deci is to download a
[pre-built binary](https://github.com/ethmarks/deci/releases).

You can also install it with Go:

```sh
go install github.com/ethmarks/deci@latest
```

## Usage

Just run `deci` followed by the name of the file you want to edit. If it doesn't
exist, deci will create it once you save.

```sh
deci myfile.txt
```

## Gallery

Check out the [Gallery](/gallery) to see more screenshots.
