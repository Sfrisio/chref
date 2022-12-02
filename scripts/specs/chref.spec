Name:           chref
Version:        1.0.3
Release:        1%{?dist}
Summary:        chref is an utility that allows you to combine the commands chmod --reference and chown --reference into one command.
Group:          chref Team
Vendor:         chref Team
Packager:       olelbis
License:        GPLv3
URL:            https://github.com/Sfrisio/chref
Source0:        %{name}-%{version}.tar.gz

BuildRequires:  golang

Provides:       %{name} = %{version}

%description
chref is an utility that allows you to combine the commands chmod --reference and chown --reference into one command.

%global debug_package %{nil}

%prep
%autosetup


%build
#go build -v -o %{name}
go build -v -ldflags="-linkmode=external -X 'github.com/Sfrisio/chref/build.Version=$(cat VERSION)' -X 'github.com/Sfrisio/chref/build.BuildUser=$(id -u -n)' -X 'github.com/Sfrisio/chref/build.BuildTime=$(date)'" -o %{name}


%install
install -Dpm 0755 %{name} %{buildroot}%{_bindir}/%{name}
#install -Dpm 0755 config.json %{buildroot}%{_sysconfdir}/%{name}/config.json
#install -Dpm 644 %{name}.service %{buildroot}%{_unitdir}/%{name}.service

%check
# go test should be here... :)

%post
#%systemd_post %{name}.service

%preun
#%systemd_preun %{name}.service

%files
%license LICENSE
%{_bindir}/%{name}

%changelog
* Fri Dec 2 2022 chref Team - 1.0.3
- apply recursively if destination is a directory with -R flag
- if -R flag is not specified will ask for user confirmation
- first basic building script%changelog
