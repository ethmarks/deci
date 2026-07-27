---
title: Deci
description: deci is a text editor like nano
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

Deci is a terminal text editor like [nano](https://nano-editor.org/).

![](/uploads/screenshot_0.2.0.png)

## Installation

The easiest way to install deci is to download a [pre-built binary](https://github.com/ethmarks/deci/releases).

You can also build from source (requires [Go](https://go.dev/dl/)):

```sh
git clone https://github.com/ethmarks/deci.git
cd deci
go build
```

## Usage

Just run `deci` followed by the name of the file you want to edit. If it doesn't exist, deci will create it.

```sh
deci myfile.txt
```
