# chref

 [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
 [![OS - Linux](https://img.shields.io/badge/OS-Linux-blue?logo=linux&logoColor=white)](https://www.linux.org/ "Go to Linux homepage")
 [![contributions - welcome](https://img.shields.io/badge/contributions-welcome-blue)](/CONTRIBUTING.md "Go to contributions doc")


`chref` is an utility that allows you to combine the commands `chmod --reference` and `chown --reference` into one command.

This is early development version. I am currently considering:

## Manual Installation

Remove any previous `chref` installation by deleting the /usr/local/chref folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh `chref` tree in /usr/local/chref

```bash
sudo tar -C /usr/local -zxf chref.tar.gz
```
You can do this by adding the following line to your $HOME/.profile (or $HOME/.bashrc) or /etc/profile (for a system-wide installation)

```bash
export PATH=$PATH:/usr/local/chref/bin
```

## Roadmap:

- [ ] On a Linux system, changing symbolic link, by default changes only the target of the symbolic link. I'd like to change ownership of symbolic link itself
- [ ] man pages
- [ ] DEB package
- [ ] RPM package
- [ ] ARM porting

If you want to tip me:

[![DogecoinBadge](https://img.shields.io/badge/Doge-Coin-yellow.svg)](https://dogecoin.com) **DAqFTEzxSxaNZ1HpVJsEV5XL2a73g3ucio**



