

# Current state

At fivestars we love using Docker. And docker orchestration. Our new local and production environments is now 80% kubernetes+helm, and improving.
The life of the DevOps has never been more interesting than when some tools takes some part of your job.

Sadly, there is still some corner to smooth and that's why we have an internal tool to build and deploy each project and dependency.
There is a lot of way to do that kind of project, and we try to have all we need inside the project using it. That mean, external dependencies, growth/shrinking rules, helm charts, and of course how to build a Docker image.

Now comes the complicated part. We need a development version of thoses images, with all the tool someone might need. Tools like vim, build-essentials, network tools.
We also need a debug version that will allow our favoride IDE to do step-by-step debugging. The ENTRYPOINT is different, the tools are different.
As good practices, we need all that inherit from the production dockerfile, nobody wants something breaking in prod because it's different from the development environment.
Some projects build multiples images, think server and its monitoring, both of them built on the same base.
And we want everything cache optimized, because who likes to wait ?
We ended up having a pretty complex build system.

Our tool can launch any command to build the images. Which mainly means a shell script running build with very specific parameters.
Thinking about it, most of the time, the parameters are the same, the context and the tags. Why the Dockefile doesn't know about it ?


# Problems

- caching optimization
- inheritance inside a project
- specific instructions for development/debugging
- cache friendly
- Extending the dockerfile with new instructions to simplify the building.


# Solution found

A templating system.

That's the first things that comes into mind. A template enable a lots of features :
- activate instruction or not from environments
- Insert into the dockerfiles some external or calculated values everywhere. Even in the FROM command (Kubernetes don't have the latest docker version).
- It's basically a programming language
- I like lists

A lot of peoples wrapped a templating system on top of dockerfile. It's not hard as template system are a dime in a dozen.
Why our version is a good one ? First it's using golang template. Golang is used in every project that touch docker. Docker, Kuber


About the instructions, we started by adding :
- `TAG` : add automatically a tag at the end of successful build.
- `CONTEXT` : set the context this dockerfile is going to be run in
- `FROM` : accepting dockrefile instead of just tag, for inheritance purposes.

# How it's better

Such an improvment, already our build.sh is cleaner and by ready the Dockerfile, we can see what's happpening.
We've added more 


# feature synergy
Now we have features to improve speed and keep 

For our project we can do :
```
FROM ./Dockerfile.project-base.ext
```
Everytime it's going to build the project-base, and use the cache if needed. Perfect for production building through jenkins.


For our debug instance, we can do : 
```
FROM project-prod:latest ./Dockerfile.project-prod.ext
```
Because the `project-prod` define the `project-prod:latest` tag, it's going to be built only if needed. Perfect for local environment when speed is needed.

# After that

We've gone crasy with the features. It's alive,
Now there is more features that help to have a cleaner dockerfile and build system :
- `INCLUDE` that include another dockerfile
- using yaml instead the dockerfile format.




