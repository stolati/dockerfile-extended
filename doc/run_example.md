
# Example or run command :
```
cd dockerfile-extended # go inside this project
dockerext --debug
```

For security reason, some path are replaced with `<path>`,
and some environments variables are removed


```
#####################
Dockerfile searching :  2  found
#####################
examples\Dockerfile.child.1.ext
examples\Dockerfile.child.2.ext

**********************************************************
**********************************************************
** Running dockerfile : examples\Dockerfile.child.1.ext **
**********************************************************
**********************************************************

#####################
Context of template :
#####################
Git:
    IS_STAGING: "false"
    IS_PORCELAIN: "false"
    PROJECT_PATH: "<path>/p/dockerfile-extended"
    PROJECT_NAME: "dockerfile-extended"
    HASH_FULL: "4f271055b85bdcbf8e060cce52134b25252a8a26"
    HASH_10: "4f271055b8"
    BRANCH: "master"
    IS_MASTER: "true"
Env:
    USERNAME: "GOD"
    GOPATH: "<path>/go;<path>/p/dockerfile-extended"
    PROCESSOR_IDENTIFIER: "Intel64 Family 6 Model 78 Stepping 3, GenuineIntel"
    ProgramFiles: "C:\Program Files"
    SystemDrive: "C:"
    TVT: "C:\Program Files (x86)\Lenovo"
    ProgramData: "C:\ProgramData"
    PSModulePath: "C:\Program Files\WindowsPowerShell\Modules;C:\WINDOWS\system32\WindowsPowerShell\v1.0\Modules"
    USERPROFILE: "C:\Users\Mickael.Kerbrat"
    OneDrive: "C:\Users\Mickael.Kerbrat\OneDrive"
    PROCESSOR_REVISION: "4e03"
    USERDOMAIN_ROAMINGPROFILE: "GOD
    ChocolateyToolsLocation: "C:\tools"
    COMPUTERNAME: "GOD
    DOCKER_HOST: "tcp://192.168.99.100:2376"
    HOMEPATH: "\Users\Mickael.Kerbrat"
    JAVA_HOME: "C:\Program Files\Java\jdk1.8.0_131"
    PROCESSOR_LEVEL: "6"
    PUBLIC: "C:\Users\Public"
    SESSIONNAME: "Console"
    CommonProgramW6432: "C:\Program Files\Common Files"
    DOCKER_API_VERSION: "1.23"
    DOCKER_CERT_PATH: "C:\Users\Mickael.Kerbrat\.minikube\certs"
    LOCALAPPDATA: "C:\Users\Mickael.Kerbrat\AppData\Local"
    PATHEXT: ".COM;.EXE;.BAT;.CMD;.VBS;.VBE;.JS;.JSE;.WSF;.WSH;.MSC;.PY;.PYW"
    TEMP: "C:\Users\MICKAE~1.KER\AppData\Local\Temp"
    TMP: "C:\Users\MICKAE~1.KER\AppData\Local\Temp"
    windir: "C:\WINDOWS"
    CLASSPATH: ".;"
    GOROOT: "C:\Go"
    NUMBER_OF_PROCESSORS: "4"
    PROCESSOR_ARCHITECTURE: "AMD64"
    ProgramFiles(x86): "C:\Program Files (x86)"
    USERDOMAIN: "GOD
    VBOX_MSI_INSTALL_PATH: "C:\Program Files\Oracle\VirtualBox\"
    APPDATA: "C:\Users\Mickael.Kerbrat\AppData\Roaming"
    ChocolateyInstall: "C:\ProgramData\chocolatey"
    CommonProgramFiles: "C:\Program Files\Common Files"
    OS: "Windows_NT"
    ProgramW6432: "C:\Program Files"
    DOCKER_TLS_VERIFY: "1"
    HOMEDRIVE: "C:"
    LOGONSERVER: "\\GOD
Local:
    HOSTNAME: "GOD
    RUN_CWD: "C:\Users\Mickael.Kerbrat\p\dockerfile-extended"
    DOCKER_CWD: "C:\Users\Mickael.Kerbrat\p\dockerfile-extended\examples"
    USERNAME: "GOD"
    OS_NAME: "windows"

#####################
Template output :
#####################

CONTEXT NONE

FROM_FILE ./Dockerfile.ext.root

RUN echo "Mickael.Kerbrat" is the best



#####################
Dockerfile to execute
#####################

# CONTEXT NONE

# FROM_FILE ./Dockerfile.ext.root
FROM docker_extended_tmp/4d65822107fcfd52

RUN echo "Mickael.Kerbrat" is the best


*******************************************************
*******************************************************
** Running dockerfile : examples\Dockerfile.ext.root **
*******************************************************
*******************************************************

#####################
Context of template :
#####################
Env:
    DOCKER_HOST: "tcp://192.168.99.100:2376"
    PATH: "C:\cygwin64\usr\local\bin\;C:\Program Files (x86)\Intel\iCLS Client\;C:\Program Files\Intel\iCLS Client\;C:\Python27\;C:\Python27\Scripts;C:\Python36\Scripts\;C:\Python36\;C:\WINDOWS\system32;C:\WINDOWS;C:\WINDOWS\System32\Wbem;C:\WINDOWS\System32\WindowsPowerShell\v1.0\;C:\ProgramData\chocolatey\bin;C:\Program Files (x86)\Boxcryptor\bin\;C:\Program Files\Git\cmd;C:\Program Files\Java\jdk1.8.0_131\bin;C:\Program Files (x86)\Intel\Intel(R) Management Engine Components\DAL;C:\Program Files\Intel\Intel(R) Management Engine Components\DAL;C:\Program Files (x86)\Intel\Intel(R) Management Engine Components\IPT;C:\Program Files\Intel\Intel(R) Management Engine Components\IPT;C:\Go\bin;C:\Program Files\Intel\WiFi\bin\;C:\Program Files\Common Files\Intel\WirelessCommon\;;C:\Program Files\Oracle\VirtualBox;C:\Users\Mickael.Kerbrat\AppData\Local\Microsoft\WindowsApps;C:\Users\Mickael.Kerbrat\.babun;C:\tools\cmder;C:\Program Files\Intel\WiFi\bin\;C:\Program Files\Common Files\Intel\WirelessCommon\;C:/Go/bin"
    PROCESSOR_ARCHITECTURE: "AMD64"
    ProgramFiles(x86): "C:\Program Files (x86)"
    ProgramW6432: "C:\Program Files"
    SystemRoot: "C:\WINDOWS"
    USERDOMAIN: "GOD
    ALLUSERSPROFILE: "C:\ProgramData"
    VBOX_MSI_INSTALL_PATH: "C:\Program Files\Oracle\VirtualBox\"
    windir: "C:\WINDOWS"
    USERDOMAIN_ROAMINGPROFILE: "GOD
    CommonProgramFiles: "C:\Program Files\Common Files"
    CommonProgramFiles(x86): "C:\Program Files (x86)\Common Files"
    COMPUTERNAME: "GOD
    OS: "Windows_NT"
    PATHEXT: ".COM;.EXE;.BAT;.CMD;.VBS;.VBE;.JS;.JSE;.WSF;.WSH;.MSC;.PY;.PYW"
    TMP: "C:\Users\MICKAE~1.KER\AppData\Local\Temp"
    ChocolateyToolsLocation: "C:\tools"
    DOCKER_CERT_PATH: "C:\Users\Mickael.Kerbrat\.minikube\certs"
    JAVA_HOME: "C:\Program Files\Java\jdk1.8.0_131"
    PUBLIC: "C:\Users\Public"
    USERNAME: "Mickael.Kerbrat"
    ChocolateyLastPathUpdate: "Wed Jul 12 12:33:01 2017"
    HOMEDRIVE: "C:"
    LOGONSERVER: "\\GOD
    OneDrive: "C:\Users\Mickael.Kerbrat\OneDrive"
    PSModulePath: "C:\Program Files\WindowsPowerShell\Modules;C:\WINDOWS\system32\WindowsPowerShell\v1.0\Modules"
    GOPATH: "<path>/go;<path>/p/dockerfile-extended"
    NUMBER_OF_PROCESSORS: "4"
    HOMEPATH: "\Users\Mickael.Kerbrat"
    ChocolateyInstall: "C:\ProgramData\chocolatey"
    CommonProgramW6432: "C:\Program Files\Common Files"
    ProgramFiles: "C:\Program Files"
    TVT: "C:\Program Files (x86)\Lenovo"
    USERPROFILE: "C:\Users\Mickael.Kerbrat"
    APPDATA: "C:\Users\Mickael.Kerbrat\AppData\Roaming"
    DOCKER_API_VERSION: "1.23"
    DOCKER_TLS_VERIFY: "1"
    GOROOT: "C:\Go"
    LOCALAPPDATA: "C:\Users\Mickael.Kerbrat\AppData\Local"
    PROCESSOR_LEVEL: "6"
    ProgramData: "C:\ProgramData"
    SystemDrive: "C:"
    ComSpec: "C:\WINDOWS\system32\cmd.exe"
    PROCESSOR_IDENTIFIER: "Intel64 Family 6 Model 78 Stepping 3, GenuineIntel"
    PROCESSOR_REVISION: "4e03"
    SESSIONNAME: "Console"
    TEMP: "C:\Users\MICKAE~1.KER\AppData\Local\Temp"
    CLASSPATH: ".;"
Local:
    HOSTNAME: "GOD
    RUN_CWD: "C:\Users\Mickael.Kerbrat\p\dockerfile-extended"
    DOCKER_CWD: "C:\Users\Mickael.Kerbrat\p\dockerfile-extended\examples"
    USERNAME: "GOD"
    OS_NAME: "windows"
Git:
    HASH_FULL: "4f271055b85bdcbf8e060cce52134b25252a8a26"
    HASH_10: "4f271055b8"
    BRANCH: "master"
    IS_MASTER: "true"
    IS_STAGING: "false"
    IS_PORCELAIN: "false"
    PROJECT_PATH: "<path>/p/dockerfile-extended"
    PROJECT_NAME: "dockerfile-extended"

#####################
Template output :
#####################

TAG toto
TAG master

CONTEXT ..

FROM alpine

# THis is a very simple example of value getting from

RUN echo "<path>/go;<path>/p/dockerfile-extended"
RUN echo "master"
RUN echo "true"

COPY .gitignore .

CMD echo "I was build on host GODELK\Mickael.Kerbrat"


#####################
Dockerfile to execute
#####################

# TAG toto
# TAG master

# CONTEXT ..

FROM alpine

# THis is a very simple example of value getting from
```