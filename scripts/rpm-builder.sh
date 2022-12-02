# Determine OS platform
UNAME=$(uname | tr "[:upper:]" "[:lower:]")
UNAMEP=$(uname -p | tr "[:upper:]" "[:lower:]")
BASEDIR=$(pwd)/specs
SPECBIN=$BASEDIR/chref.spec
SPECMAN=$BASEDIR/chref-man.spec

if [ -z "$1" ]; then
	echo "no source file specicifed as input parameter"
	echo "Usage: $0 <name>-<version>.tar.gz"
	exit 1
fi

SOURCEFILE=$1

el_prereq () {
    echo "++++ Install $UNAME $UNAMEP required packages ++++"
    echo ""
    echo "++++ Install git ++++"
    sudo dnf install git -y
    echo ""
    echo "++++ Install go toolset ++++"
    sudo dnf module install go-toolset -y
    echo ""
    echo "++++ Install RPM Development Tools ++++"
    sudo dnf groupinstall "RPM Development Tools" -y
    echo ""
}

el_rpmbuild ()  {
   if [ -f $SPECBIN ]; then
        echo "--> .spec file available execute rpmdev-setuptree"
        rpmdev-setuptree
        echo ""
        if [ -f $SOURCEFILE ]; then
                echo "---> move source file into ~/rpmbuild/SOURCES directory"
                mv $SOURCEFILE ~/rpmbuild/SOURCES
                echo ""
                echo "----> starting build..."
                rpmbuild -ba $SPECBIN
                echo ""
                echo "-----> rpm available under ~/rpmbuild/RPMS/$UNAMEP/"
                cd ~/rpmbuild/RPMS/$UNAMEP/ && ls
        else
                echo "invalid source file specified $1 not available"   
                exit 1
        fi

    else
        echo ".spec file unavailable"
        exit 1
    fi
}

# If Linux, try to determine specific distribution
if [ "$UNAME" == "linux" ]; then
case $UNAMEP in
  x86_64)
    el_prereq
    # check if specfile exist then build rpm's
    el_rpmbuild
    exit 0
    ;;

  aarch64)
    el_prereq
    # check if specfile exist then build rpm's
    el_rpmbuild
    exit 0
    ;;
  *)
    echo "unknown"
    exit 1
    ;;
esac
fi

# unset all variables
unset UNAME UNAMEP BASEDIR SPECBIN SPECMAN SOURCEFILE
