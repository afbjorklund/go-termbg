# go-termbg
A Go library for terminal background color detection.
The detected color is provided by RGB or theme ( dark or light ).

Based on https://github.com/dalance/termbg for Rust

## Verified terminals

* [Alacritty](https://github.com/alacritty/alacritty)
* GNOME Terminal
* GNU Screen
* [kitty](https://sw.kovidgoyal.net/kitty/)
* [iTerm2](https://iterm2.com)
* macOS terminal
* MATE Terminal
* [mintty](https://mintty.github.io)
* [RLogin](http://nanno.dip.jp/softlib/man/rlogin/)
* rxvt-unicode
* sakura
* [PuTTY PRIVATE PATCHES](https://ice.hotmint.com/putty/)
* [Tera Term](https://ttssh2.osdn.jp)
* [Terminator](https://terminator-gtk3.readthedocs.io/en/latest/)
* [tmux](https://github.com/tmux/tmux)
* xfce4-terminal
* xterm
* Win32 console

If you check other terminals, please report through [issue](https://github.com/dalance/termbg/issues).

## Unsupported terminals

* [LilyTerm](https://github.com/Tetralet/LilyTerm)
* [Poderosa](https://ja.poderosa-terminal.com)
* [PuTTY](https://www.putty.org)
* [QTerminal](https://github.com/lxqt/qterminal)
* [Visual Studio Code](https://code.visualstudio.com)
* [Windows Terminal](https://github.com/microsoft/terminal)

"Windows Terminal" may be supported in a future release: https://github.com/microsoft/terminal/issues/3718.

## Check program

This module provides a simple program to check.

```console
$ go run ./...
Check terminal background color
  Term : XtermCompatible
  Color: R=2e2e, G=3434, B=3636
  Theme: Dark
```

## Detecting mechanism

If the terminal is win32 console, WIN32API is used for detection.
If the terminal is xterm compatible, "Xterm Control Sequences" is used.

The detected RGB is converted to YCbCr.
If Y > 0.5, the theme is detected as "light", otherwise "dark".
