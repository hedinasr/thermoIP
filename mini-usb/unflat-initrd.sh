#!/bin/bash

rm -rf initrd
mkdir -p initrd && cd initrd
gzip -dc ../initrd-big.img | cpio -id
