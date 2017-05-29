# Mini-USB

Test a tiny bootable USB with `qemu`:
```
qemu-img create -f raw testing.img 10M
mkdosfs testing.img
syslinux testing.img
mkdir -p usb && mount testing.img usb # Now you can copy vmlinuz and initrd.gz
```

To modify the `initrd.img`, use the scripts `unflat-initrd.gz` and
`rebuild-initrd.sh`.
