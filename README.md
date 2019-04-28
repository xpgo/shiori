# Copy declaration

This project is forked from [Shiori](https://github.com/RadhiFadlillah/shiori) created by RadhiFadlillah.

Currently, the main changes include: 

1. Since Mercury has open-sourced, I try to use self-hosted Mercury for page rendering which shows images in each cached page.
2. Use env MERCURY_KEY and MERCURY_API if you want to change the Mercury API.
3. Changed some css style for better mobile read. (Change css: modify lss, lessc xx xx, go run assets-generator.go)
4. Add archive button for each bookmark and auto hide bookmarks with tag "archive"
5. Support search with chinese character
6. Resize thumb image to small size, and customize leading image url in edit dialog.

# Shiori

[![Travis CI](https://travis-ci.org/xpgo/shiori.svg?branch=master)](https://travis-ci.org/xpgo/shiori)
[![Go Report Card](https://goreportcard.com/badge/github.com/xpgo/shiori)](https://goreportcard.com/report/github.com/xpgo/shiori)
[![Docker Build Status](https://img.shields.io/docker/build/xpgo/shiori.svg)](https://hub.docker.com/r/xpgo/shiori/)

Shiori is a simple bookmarks manager written in Go language. Intended as a simple clone of [Pocket](https://getpocket.com//). You can use it as command line application or as web application. This application is distributed as a single binary, which means it can be installed and used easily.

![Screenshot](https://raw.githubusercontent.com/xpgo/shiori/master/screenshot/pc-grid.png)

## Features

- Simple and clean command line interface.
- Basic bookmarks management i.e. add, edit and delete.
- Search bookmarks by their title, tags, url and page content.
- Import and export bookmarks from and to Netscape Bookmark file.
- Portable, thanks to its single binary format and sqlite3 database
- Simple web interface for those who don't want to use a command line app.
- Where possible, by default `shiori` will download a static copy of the webpage in simple text and HTML format, which later can be used as an offline archive for that page.

## Documentation

All documentation is available in [wiki](https://github.com/xpgo/shiori/wiki). If you think there are incomplete or incorrect information, feels free to edit it.

## License

Shiori is distributed using [MIT license](https://choosealicense.com/licenses/mit/), which means you can use and modify it however you want. However, if you make an enhancement for it, if possible, please send a pull request.
