# chref

 [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
 [![OS - Linux](https://img.shields.io/badge/OS-Linux-blue?logo=linux&logoColor=white)](https://www.linux.org/ "Go to Linux homepage")
 [![OS - MacOS](https://img.shields.io/badge/OS-macOS-blue?logo=Apple&logoColor=white)](https://apple.com/ "Go to Apple homepage")
 [![contributions - welcome](https://img.shields.io/badge/contributions-welcome-blue)](/CONTRIBUTING.md "Go to contributions doc")
 ![visitor badge](https://visitor-badge.glitch.me/badge?page_id=Sfrisio.chref)

`chref` is an utility that allows you to combine the commands `chmod --reference` and `chown --reference` into one command.

This is early development version. I am currently considering:

## Roadmap:

- [ ] On a Linux system, changing symbolic link, by default changes only the target of the symbolic link. I'd like to change ownership of symbolic link itself
- [ ] man pages
- [ ] DEB package
- [ ] RPM package
- [ ] MacOS PKG
- [ ] ARM porting

## Building from source

If you want to build `chref` from source, please verify to have already installed **go1.19.x** or higher.

Then run this command:

```bash
go build -v -ldflags="-X 'chref/build.Version=$(cat VERSION)' -X 'chref/build.BuildUser=$(id -u -n)' -X 'chref/build.BuildTime=$(date)'"
```

If you want to build automatically `chref` for all the supported platform consider to use `binary-builder.sh` provided in *scripts* folder.

Clone this repository firs, then follow these steps

```bash
cd chref
chmod +x scripts/binary-builder.sh
./scripts/binary-builder.sh
```

## Manual Installation

Remove any previous `chref` installation by deleting the /usr/local/chref folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh `chref` tree in /usr/local/chref

```bash
sudo tar -C /usr/local -zxf chref.tar.gz
```
You can do this by adding the following line to your $HOME/.profile (or $HOME/.bashrc) or /etc/profile (for a system-wide installation)

```bash
export PATH=$PATH:/usr/local/chref/bin
```
**Note**: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, run the following command (may differ if you are using .profile instead of .bashrc):

```bash
source $HOME/.bashrc
```

If you want to tip me:

[![DogecoinBadge](https://img.shields.io/badge/Doge-Coin-yellow.svg)](https://dogecoin.com) **DAqFTEzxSxaNZ1HpVJsEV5XL2a73g3ucio**
