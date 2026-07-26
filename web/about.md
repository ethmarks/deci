---
title: About Me
description: I am a general-purpose Lume theme based on Tufte CSS.
layout: layouts/base.vto
---

# About Me

Hi! I'm Tufte, a general-purpose [Lume theme](https://lume.land/themes/) based
on [Tufte CSS](https://edwardtufte.github.io/tufte-css/).

## My Skills

Here are a few of the things that I can do. For a more comprehensive list, check
out the [Using this theme](/blog/usage/#syntax) post.

### Sidenotes

Tufte CSS recommends the use of
[sidenotes](https://edwardtufte.github.io/tufte-css/#sidenotes) rather than
footnotes. I have a custom markdown-it plugin that lets me automatically render
inline Markdown footnotes into Tufte-style sidenotes[^"Sidenotes" in this case
meaning footnotes except lacking in feet and sufficient in sides.].

### Code Highlighting

I can highlight code blocks using [Nueglow](https://nuejs.org/docs/nueglow)[^Fun
fact: Lume's Nueglow plugin was created by Ethan Marks, the same developer who
created me!].

```ts
function fibonacci(num: number): number {
  if (num == 1) return 0;
  if (num == 2) return 1;
  return fibonacci(num - 1) + fibonacci(num - 2);
}
```

### KaTeX

I can automatically render complex math using Lume's
[KaTeX plugin](https://lume.land/plugins/katex/).

```math
i\hbar \frac{\partial}{\partial t}\Psi(\mathbf{r},t) = \left[-\frac{\hbar^2}{2m}\nabla^2 + V(\mathbf{r},t)\right]\Psi(\mathbf{r},t)
```

### Collapsibles

I can automatically render collapsible blocks using the
[markdown-it-collapsible](https://www.npmjs.com/package/markdown-it-collapsible)
plugin.

+++ Click here to open

This is the inside of the collapsible.

There's lots of space in here, so make yourself at home.

+++

## My History

I was created by [Ethan Marks](https://github.com/ethmarks) in June 2026, when
he was bored on vacation and decided on a whim to make a new Lume theme. He
based it on Tufte CSS, a fairly-popular CSS micro-framework created in 2014 by
[David Liepmann](https://github.com/daveliepmann) based on
[Edward Tufte](https://github.com/edwardtufte)'s distinctively-styled
handouts[^For example,
<http://rmarkdown.rstudio.com/examples/tufte-handout.pdf>.].
