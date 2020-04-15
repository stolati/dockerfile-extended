# dockerfile-extended

The goal of this project is to have Dockerfile as self-sufficient as possible.

We do this by adding more feature to the default Dockerfile through a cli command.
This tool is run instead of `docker build` and use the available `docker` command.

Very often, we create scripts to launch the building with specifics parameters.

On very complex project, we can have inheritance system, where a main dockerfile can be a base,
and then more dockerfile can build things like server, monitoring, unittest, loadtest.
This doesn't take into account when for thoses one we need a development version with added features (or removed ones).

The content of thoses dockerfile can be pretty complex too, and there is no good way to template a dockerfile,
to activate or not some parts, to get the git hash or project version.

# Usage

The command of `dockerfile-extended` is `dockerext`

Launching the dockerfile extended on a single file :
`dockerext --dockerfile ./Dockerfile.my.ext`

Searching for all the "Dockerfile.*.ext" in the directory and run thoses :
`dockerext --dir .` (also the default)

Launching a non-execution command and looking at the verbose output :
`dockerext --dockerfile ./Dockerfile.my.ext --dry-run --debug`

Passing more parameter to the build (here a custom host mapping) :
`dockerext --dockerfile ./Dockerfile.my.ext -- --add-host toto:127.0.0.1`



# Features

## Templating system

The main feature of this project is a templating system.
We use the golang `text/template` on the dockerfile.
And we add the [sprig](https://github.com/Masterminds/sprig) functions to it.
The informations comes from different sources.
Right now, it's mainly environment and git ones.


ex (Thoses are subject to changes) :
```
Env:
    DOCKER_API_VERSION: "1.23"
    GOROOT: "/goroot"
    DOCKER_HOST: "tcp://192.168.99.100:2376"
    GOPATH: "/gopath"
    HOME: "/notroot"
    USER: "god"
    DOCKER_TLS_VERIFY: "1"
    NUMBER_OF_PROCESSORS: "4"
    PATH: "/usr/bin"
Local:
    OS_NAME: "darwin"
    HOSTNAME: "mycomputer"
    RUN_CWD: "/myprojects/docker-extended"
    DOCKER_CWD: "/myprojects/docker-extended/examples"
    USERNAME: "god"
Git:
    HASH_FULL: "0ca0b7beaf3dc20e1e5044d741734bf6a568277d"
    HASH_10: "0ca0b7beaf"
    BRANCH: "master"
    IS_MASTER: "true"
    IS_STAGING: "false"
    IS_PORCELAIN: "false"
    PROJECT_NAME: "dockerfile-extended"
    PROJECT_PATH: "/myprojects/docker-extended"
```


So you can create a dockerfile like this :
```
{{ if ! .Git.IS_PORCELAIN }}
RUN false # make the build fail if uncommited files
{{end}}
```

Or like this :
```
# Easy way to get the version when something happens in prod
# And get used by the endpoint /monitor_info
RUN echo {{.Git.HASH_10 | quote}} >> /infos/hash
ENTRYPOINT entrypoint
CMD runserver
```


## `TAG` command

A new `TAG` command has been created.
The purpose of this command is to tag automatically the image just after a successful build.
It goes very well with templating values.

```
TAG {{.ENV.DOCKER_REGISTERY}}/myproject:{{.Git.HASH_10}}

{{if .Git.IS_MASTER }}
TAG myproject:latest
{{end}}
```



## `CONTEXT` command

Sometimes, especially when there is multiples dockerfile,
the context on which the dockerfile get called is not its directory.

Other times, we might need an empty context. Especially when the project
is large and uploading the context is long.

The `CONTEXT` command wait for either :
- a path relative to the dockerfile directory.
- the string `NONE` (not case sensitive)

examples:

```
CONTEXT .. # the context is the directory one up
```

```
CONTEXT NONE # We don't need context here
```

```
# This dockerfile is only for building static files
CONTEXT ../statics
# Or the absolute path version
CONTEXT {{.Git.PROJECT_PATH}}/statics
```


## `FROM_FILE` command

When doing docker inheritance,
this command replace the `FROM` to point to a file instead of pointing to a tag.

The goal is to have a local inheritance without having to tag in a script the parent image,
and then having the child image using only this tag.
This can feel a little bit hacky.

Dockerfile extended do the hack for you. It's creating a temporary tag,
and build the parent dockerfile before the child one.

This new command waits for a path, relative to the dockerfile directory.

examples :
```
FROM_FILE ./Dockerfile.parent
```
```
FROM_FILE {{.Git.PROJECT_PATH}}/docker_builds/Dockerfile.parent
```


# Limitations on commands

The commands inside the dockerfile are parsed by custom code.
They are some limitation :
- the command should be used before the FROM or FROM_FILE command
- Command don't allow multilines, escaping or quotes

# Futures Improvements

- Possibility to `--set` or `--values` to add values and values files for the template.
Follow helm lead.
- Having Improved commands, to use the same parsing system as docker.
- Being able to create yaml dockerfile instead of the default format
- Adding a CACHED_FROM command, with the corresponding parameter
The catched, is that cached from is only from a specific tag on the registry.
So we may need to have a smart way of choosing it.
- We use recursive programming to build the parent. That mean double building currently.
If two child need the same parent, we are trying to build the parent twice.
It's not a big problem with the caching system, but it's still avoidable waiting.
- Checking for looping dependency. No check is made for that.
But something's wrong if you have so many inheritance that you left some loop somewhere.
- adding some security measure :
  - TAG : don't override existing image
  - CONTEXT : don't allow absolute path, don't grab more than the git project
  - ENV : remove everything that can be used against people (PATH, PGPASSWORD)
  - parameters/parser command to activate/deactivate features (ex : only template)
- SQUASH command, to squash multiples layers into one


# Good practices

We're not telling you what to do,
but some of thoses features are ready to fire you in the foot.

### Think twice about using values that are local

With this version, it's easy to use a value that is specific to the computer
you're using. The `.Env` is filled with that.
Everything should be possible without using thoses.
We provide the env variable because it's an easy way to give values to your template.
But the preferred way will be passing values/values' file (once it's implemented).

### Templating can break your cache layers

Templating is a great tool, but the output of it can break you cache layer.
Between having values that are calculated and


### Context are meant in the project

Don't use the context outside of the current project.
By going up `../../../..` or starting from root `/home`
it's easy to select more than just the project you are working on.

### Thinking security

You're supposed to know what you are doing, and the `--debug` command
give you plenty of information for that.
This project has been created with the though that you own your building process and your dockerfiles.
Don't build a dockerfile.ext without looking inside it if it's doing something fishy.

There is many way adding thoses feature can lead to security flaws:
- tag can override an existing image (ex: ubuntu, apline, or any database)
- context can grab all disk (don't publish it publicly after that)
- someone env can be printed and saved inside the docker image (PGPASSWORD anyone?)

### Dockerfile is already powerfull

Just who choose this format ? Why creating a brand new format when you
can choose json, yaml, xml and many more that you don't have to parse in
a strange way ? I digress.

Dockerfile has already a lot of features that make it powerfull. The features
dockerfile-extended add are not necessary to have a sane build environment.
It's just hell more powerfull than adding some ARG.



### TODO
Here is a list of TODO without order :
- copy file with template
- pull open source project from github and see if we can 80% test+build it without configuration.
- docker-compose wrapper
- Change path on volumes (usefull for wsl on windows)
- Template on the dockerfile.
- Change to python for faster development.
- Use most of the docker parameters from inside the dockerfile.
- Reference volume by name from outside ?


