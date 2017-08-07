# Improvement ideas

There is some improvements we can do.
Some of them are good ideas, some are very bad.
No order, no rules

## Have real parameters handling
Parameters are really barebone right now. We can improve by having
a more complete help, more parameters and better default.
Maybe sub-commands is the way (like every other tool in the docker world).

## Better parsing of Dockerfile
We should use/clone an existing dockerfile parsing instead of hacking a
simple parsing.

## TAG_AND_PUSH command
In case you need to publish automatically to a repository.
Propose to automatically do a docker push.

## Run Dockerfile
Instead of having to build it first, build and run.

## CACHE_FROM command
So fill the --cache-from parameter.
The catch, is that we can improve building by choosing the publish version
that is the nearest to our local one.
Looking at git history and present images. This only works if we use
git commit has tag in the repository in a coherent way.

## ASSERT command
Doing assertion commands, which is like RUN command, but the meaning is
explicit.

## Doing security and dockerfile check
Look for bad practices in dockerfile and send warnings to the user.
Like using ADD, overridden TAG of official repo, using MAINTAINER

## SQUASH layer command
A way to squash layers into one would be good.
For the case like that :
```
ADD ssh_key ~/.ssh/private
RUN fetch dependency behind security &&
    rm ~/.ssh/private
```

It's not possible right now to not have the ssh key in the layers.
(ok, now with template there is a hack for that, but that's behond the point)

It's feasible by looking at the hash the layer will get, and
doing the layer caching ourselves. It won't be easy, but it would be fun to try.

## Adding more commands that docker need
Like `clean` that remove stopped containers and dangling images.
It could become a wrapper around all the docker commands,
adding features to all part of it. (But that's a bigger scope than currently)

## Use Yaml instead of dockerfile
Who choose this format ? We should allow for more.
json, yaml, and maybe XML.
- The content will be easier to parse and be generated
- Type of values passed to commmands are explicit

## Caching checking informations
Looking at the cached layers and the dockerfile, telling us what's happening
even before trying to run the dockerfile. This could be useful for cache
optimization.
Telling what's in the context and what size it takes.

## Backporting some features
Like the new --squash, we can do it the old ways by exporting the image.
Some of the features like that can be emulated with older versions.


## INCLUDE Instruction
From a dockerfile, add another docker file. This allow more complex files to be used, or even compounds.
I know we can do that in the templating system, but having a docker command would be interesting for clarity

## Improve FROM Instruction
Instead of having a `FROM` and a `FROM_FILE`, let's use a FROM, but allow for multiples values to be added.
If one value is not present, then the next will be used.

`FROM tag` => load the tag
`FROM tag dockerfile` => if the dockerfile create the tag, it won't be called again if built
`FROM dockerifle` => always try to build the dockerfile

## HOST_RUN Instruction
It's just an idea, but maybe we can add the HOST_RUN command.
It would allow for exapmle automatically

## NO_CACHE_FROM_HERE Instruction
Cancel the cache system from a certain point in the dockerfile.
(instead of the all or nothing --no-cache)