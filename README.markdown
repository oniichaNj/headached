Description
-----------
A simple system daemon that
- [ ] randomly corrupts files by writing random data to them.
- [X] causes random CPU usage spikes (maybe by doing crypto things in the background?)
- [X] drains the entropy pool constantly to make legit entropy gathering slow/impossible

You don't want to run this on a production system unless you're looking for a challenge and/or testing your backup system.

Installation
------------
Will be written once it actually works.

TODO
----

- [ ] implement the description
- [ ] write makefile
- [ ] Make `go get`-able
- [ ] write a sane default config


Suggested ideas
---------------
* (evilmode) add option to disable entropy gathering and instead constantly feed '1' into the entropy pool
* (maybe?) write a file to echo random bytes to the system sound output device? could break compability for some systems
* ...your ideas go here!