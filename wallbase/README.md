wallbase
========

The wallbase package facilitates requests to [wallbase.cc][]. Useful until the
public API is released.

[wallbase.cc]: http://wallbase.cc/

Documentation
-------------

Documentation provided by GoDoc.

- [wallbase][]

[wallbase]: http://godoc.org/github.com/mewmew/playground/wallbase

walls
=====

walls updates the desktop wallpaper at specified time intervals.

It performs a search for wallpapers on [wallbase.cc] based on the provided
search query. The wallpaper result order is random.

For persistent storage a custom download directory can be specified.

Installation
------------

	go get github.com/mewmew/playground/wallbase/cmd/walls

Dependencies
------------

The `hsetroot` command is used to set the wallpaper.

Documentation
-------------

Documentation provided by GoDoc.

- [walls][]

[walls]: http://godoc.org/github.com/mewmew/playground/wallbase/cmd/walls

Usage
-----

	walls [OPTION]... QUERY

Flags:

	-o (default="/tmp/wallbase")
		Output directory.
	-t (default="30m")
		Timeout interval between updates.

Examples
--------

1. Search for "nature waterfall" and update active wallpaper each 10s.

		wallbase -t 10s nature waterfall

2. Search for "nature" and store each wallpaper in "download/".

		wallbase -t 0s -o download/ nature

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/