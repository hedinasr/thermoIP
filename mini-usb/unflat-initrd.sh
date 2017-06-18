#!/bin/bash
#
# usage: ./unflat-initrd.sh <initrd.gz>

rm -rf initrd
mkdir -p initrd && cd initrd
gzip -dc ../"$1" | cpio -id
