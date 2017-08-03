
# Building
On windows, the building process is not stable.
And the owner of this project is not familliar enough with golang builing process.

Command to launch :

```
TAG=dockerfile-extended:latest
docker build . -t $TAG
mkdir -p output
# TODO use docker copy instead of this hack
docker run $TAG cat /srv/output/main_windows > output/main_windows.exe
docker run $TAG cat /srv/output/main_linux > output/main_linux
docker run $TAG cat /srv/output/main_darwin > output/main_darwin

chmod a+x output/*
```

There is a way to mount volumes in docker,
but the version that use kubernetes don't.














