#!/bin/bash
#
# ./rebuild-initrd <my-initrd>

cd initrd
find ./ | cpio -H newc -o > ../"$1".cpio
cd ..
gzip -9 "$1".cpio
mv "$1".cpio.gz "$1".img
rm -f "$1".cpio
