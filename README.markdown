Description
===========
A simple system daemon that
* randomly corrupts files by writing random data to them.
* causes random CPU usage spikes (by doing crypto things in the background, as this proved to use more user time than just blank for{}'s)
* drains the entropy pool constantly to make legit entropy gathering slow/impossible

You don't want to run this on a production system unless you're looking for a challenge and/or testing your backup system.

Installation
============
(This section is subject to change.)
Dependencies
------------
* `go` (the Google Go compiler)
* a GNU/Linux or (untested) *BSD operating system

Compiling 
---------
If you don't haev a GOPATH set up, do so by executing `mkdir ~/go` (replacing `~/go` with where you want your GOPATH), followed by `export GOPATH=$HOME/go` (assuming you used that directory).

You might want to add the line `export GOPATH=$HOME/go` to your shell rc file, so you don't have to run it every time you restart the shell.

Grab the code by running `go get github.com/oniichaNj/headached`.
Enter your `$GOPATH/src/github.com/oniichaNj/headached/` and run `make`.

`sudo make install` can be ran afterwards, to install the compiled file to the system. 


Bare in mind that `headached` needs to be run as root in order to use the file corruption features (well, not really, but you probably want to use it on paths that require it anyways).

Gentoo install
--------------
Run `sudo make gentoo-install` to install the init file to the system.

Debian install
--------------
Run `sudo make debian-install` to install the init file to the system.

systemd install
---------------
Run `sudo make systemd-install` to install the unit file to the system.

Other systems
-------------
Running it as `/usr/sbin/headached &> /var/log/headached.log &` works fine.
If your system isn't supported, just pull request what you need and I'll look at it.



TODO
====

- [X] implement the description
- [X] write makefile
- [X] write a sane default config


Suggested ideas
===============
* (evilmode) add option to disable entropy gathering and instead constantly feed '1' into the entropy pool
* (maybe?) write a file to echo random bytes to the system sound output device? could break compability for some systems
* randomly allocate a bunch of memory
* (maybe) create a bunch of files here and there?
* ...your ideas go here!