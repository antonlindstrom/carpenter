# Carpenter - Image cleaner

This is a basically a copy of
[virt-sysprep](http://libguestfs.org/virt-sysprep.1.html), but scraches some of
my issues with it. The first is all the dependencies. To get virt-sysprep to
work on so many guests and platforms as possible, some dependencies need to be
met. For this project, I'll strictly focus on Linux guests with filesystems such
as ext[2-4].

The basic criteria is that a two dependencies is required, a Linux system with
standard SUS tools and that `qemu-nbd` is installed.

As we're using some exec calls, it's not as portable as it could be, for example
it won't work on Windows.

This is written because I tried out ZFS on Linux on a server which I also had
virt-sysprep on, this caused some problems which I didn't want. So I began
pondering on how to do a simple clean.

USE WITH CAUTION!

## Goals

In accordance to the top statement, the project will try to be as easy on
dependencies as possible.

One of the goals is to support the following features:
* Removing SSH Host keys
* Setting hostname

## Install

The following should do it (requires Go):

    make

## Example usage

Simplest example:

    ./bin/carpenter -hostname=vm.example.com -image=/tmp/template.qcow2

## Test

    make test

Unfortunately the tests are very uncomplete at the moment.

## Licence

See [LICENSE](LICENSE) file.
