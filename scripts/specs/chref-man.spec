Name:           chref-man
Version:        1.0
Release:        1%{?dist}
Summary:        Chref Documentation
Group:          chref Team
License:        GPLv3
Source0:        %{name}-%{version}.tar.gz
BuildArch:      noarch
Vendor:         chref Team
Packager:       olelbis
URL:            https://github.com/Sfrisio/chref

Provides:       %{name} = %{version}

%description
This is a test. We are trying to install a
sysadmin generated man page from a rpm file.
Just to prove we can.

%prep
%setup -q

%build

%install
mkdir -p $RPM_BUILD_ROOT/usr/local/share/man/man8
install chref.8.gz $RPM_BUILD_ROOT/usr/local/share/man/man8

%clean
rm -rf $RPM_BUILD_ROOT

%files
%defattr(-,root,root,-)
/usr/local/share/man/man8/chref.8.gz

%changelog
* Fri Dec 2 2022 chref Team 
- Initial build.
