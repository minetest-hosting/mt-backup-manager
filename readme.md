
## mt-backup-manager

creates ad-hoc btrfs-snapshots and streams the snapshot-content as zip-file to a connecting http-client


## test env preparation

```sh
sudo apt install btrfs-progs

fallocate -l 1G data.img
mkfs.btrfs data.img

sudo mkdir /mnt/btrfs
sudo mount -t btrfs -o loop ./data.img /mnt/btrfs/


```