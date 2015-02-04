Description
===========
A simple system daemon that
- [X] randomly corrupts files by writing random data to them.
- [X] causes random CPU usage spikes (by doing crypto things in the background, as this proved to use more user time than just blank for{}'s)
- [X] drains the entropy pool constantly to make legit entropy gathering slow/impossible

You don't want to run this on a production system unless you're looking for a challenge and/or testing your backup system.

Installation
============
(This section will be changed once a Makefile is written.)
Dependencies
------------
* `go` (the Google Go compiler)
* a GNU/Linux or (untested) *BSD operating system

Compiling
---------
Grab the code by running `go get github.com/oniichaNj/headached`.
Enter your `$GOPATH/src/github.com/oniichaNj/headached/` and run `go build headached.go`.
The resulting binary is available in that same directory as just `headached` and can be moved to a `$PATH`.

Until it's fully daemonised via the most common daemon managers, running it as `headached &` willl have to do.

Bare in mind that `headached` needs to be ran as root in order to use the file corruption features.


TODO
====

- [X] implement the description
- [ ] write makefile
- [ ] Make `go get`-able
- [X] write a sane default config


Suggested ideas
===============
* (evilmode) add option to disable entropy gathering and instead constantly feed '1' into the entropy pool
* (maybe?) write a file to echo random bytes to the system sound output device? could break compability for some systems
* randomly allocate a bunch of memory
* (maybe) create a bunch of files here and there?
* ...your ideas go here!