# deci

`deci` is a terminal text editor like `nano`.

## Quickstart

> **I recommend downloading the latest binary from the [releases page](https://github.com/ethmarks/deci/releases).**

You can also build from source (make sure to [install Go](https://go.dev/doc/install)):

```sh
go install github.com/ethmarks/deci@latest
```

## Keybinds

`deci` has unique behavior for 112 keybinds in total.

Basic keybinds that every editor has but which still required a *lot* of code to implement:

- Typeable characters like "a" and "1" (94 in total): inserts the character at the cursor and moves the cursor right
- `enter`: splits the current line and moves the right half to a new line (which has the effect of creating a new line if the cursor is at the end of a line) and moves the cursor to the start of the next line.
- `backspace`: if the cursor is *not* at the start of a line, removes the character to the left of the cursor and moves the cursor left. If it *is* at the start of a line, it merges the current line with the previous line (which has the effect of removing it if the current line is blank) and moves the cursor up and to the end of the line
- `delete`: if the cursor is *not* at the end of a line, removes the character to the right of the cursor. If the cursor *is* at the end of a line, it merges the next line with the current line
- `tab`: inserts 4 spaces
- `ctrl+z`: reverts to the previous editor state
- `ctrl+y`: does the opposite of `ctrl+z`

Cursor movement keybinds, which every editor supports, but were so complicated that they get their own category:

- `up`: if the cursor is *not* on the first line, moves the cursor up one line and either to the end of the current line or the last x position, whichever is smaller. If the cursor *is* on the first line, moves the cursor to the start of the line
- `down`: if the cursor is *not* on the last line, moves the cursor down one line and either to the end of the current line or the last x position, whichever is smaller. If the cursor *is* on the last line, moves the cursor to the end of the line
- `left`: if the cursor is *not* at the start of the line, moves the cursor left. If the cursor *is* at the start of the line, moves the cursor up and to the end of the previous line
- `right`: if the cursor is *not* at the end of the line, moves the cursor right. If the cursor *is* at the end of the line, moves the cursor down and to the start of the line
- `ctrl+left`: moves left until either it reaches the start of the line or it reaches a character that *isn't* a delimiter (like a space or a punctuation mark), then continues moving left until it reaches a character that *is* a delimiter
- `ctrl+right`: moves right until either it reaches the end of the line or it reaches a delimiter, then continues moving right until it reaches a character that *isn't* a delimiter

Special keybinds:

- `ctrl+c`: closes `deci` without saving
- `ctrl+o`: writes the editor content out to a file
- `ctrl+h`: toggles the help screen
- `ctrl+p`: toggles the Markdown preview screen
- `alt+p`: changes the Markdown preview theme
- `alt+c`: changes the cursor shape

## Etymology

[Deci](https://en.wikipedia.org/wiki/Deci-) is the Metric prefix for 10^-1, like
how [Nano](https://en.wikipedia.org/wiki/Neci-) is the prefix for 10^-9.

| deci           | nano             |
| -------------- | ---------------- |
| pretty small   | _tiny_           |
| not well known | extremely common |

Written in [Go](https://go.dev/) with
[Bubble Tea](https://github.com/charmbracelet/bubbletea).

## License

This project is under an MIT License. See [LICENSE](LICENSE) for more
information.
