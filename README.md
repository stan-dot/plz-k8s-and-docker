# this repo follows this tutorial

<https://please.build/codelabs/k8s/index.html#0>

## for devcontainer setup used this as inspiration

<https://github.com/qdm12/godevcontainer/blob/master/.devcontainer/devcontainer.json>

## ran into an error with the go service creation

https://please.build/codelabs/k8s/index.html#1
`GoTool = //third_party/go:toolchain|go`
the pipe in the end was missing
to be honest I was considering why is is that looking odd

## dagger - discovered when reading about earthly vs dagger

earthly vs dagger
https://news.ycombinator.com/item?id=30858896

https://github.com/dagger/dagger/tree/v0.9.11

latest dagger release is v0.9.11
not sure if production ready

https://docs.dagger.io/quickstart/593914/hello

definitely an intestering idea to have a python script there

comparison with non-portable devkits
https://archive.docs.dagger.io/0.2/1220/vs/

https://cuelang.org/
https://github.com/cue-lang/cue
does not look mature and understandable dagger made a pivot

on a monorepo migration
https://dagger.io/blog/adopting-monorepo-strategy

ok it seems that it's for running CI jobs

CI engine agnostc - good between github and gitlab
https://www.reddit.com/r/devops/comments/tvrfvc/dagger_a_portable_devkit_for_cicd_pipelines/

all written in 2 weeks, impressive
https://www.reddit.com/r/devops/comments/185tey1/whats_your_experience_with_dagger/

some say that more issues

the alternative is the gitlab runner

but it still requires another Gitlab repo
https://docs.dagger.io/759201/gitlab-google-cloud/#step-5-create-a-gitlab-cicd-pipeline

potentially gcloud push might be easier
but we do not have such complex pipelines to be honest

## locally building the image might cause issues if plz not in path

https://please.build/codelabs/k8s/index.html#3

might run manually with /home/<user>/.please/bin/plz

> podman build -t hub.docker.com/image:37901176dc15983224d6edb168970a528c13ae978c1505726bc620b4165fc211 -f Dockerfile - < plz-out/gen/hello_service/k8s/\_image#docker_context.tar.gz

## deployinig isn't plug and play straightforward

minikube not feasible locally beacuse

> system policy prevents local management of virtualized systems
> as the minicube isn't there so it's running on argus

documentation isn't searchable

the k8s isn't as plug and play as epxected.
might need another layer for this to play nicely with helm

## they mention themselves that it's not trivial so coordinate this with docker build

https://please.build/codelabs/k8s/index.html#7
because the build process is happening in the background

there is a number of pre-build scripts

- :image_fqn - fully qualified name - uniquely identifies
- :image and :image_load
- :image_push - into docker registry --> we should change to gcloud registry 
- :image_save - into `plz-out/gen/{image}.tar`
- :image_run - will run in the local docker env


ideally we'd like to override:
`docker_image(
    name = "image",
    srcs= ["//hello_service"],
    dockerfile = "Dockerfile",
    base_image = "//common/docker:base",
    run_args = "-p 8000:8000",
    visibility = ["//k8s/hello_service:all"],
    # see https://please.build/codelabs/k8s/index.html#7
)`

with 'podman-image' to generate better scripts
here a list of plugins
https://github.com/please-build/please-rules?tab=readme-ov-file

https://github.com/please-build/docker-rules

shouldn't be too difficult
https://github.com/please-build/docker-rules/blob/master/build_defs/docker.build_defs

then email this person - https://github.com/Tatskaari

but theoretically can make local target for podman rules and work with that
jpoole@thoughtmachine.net

potentially also for helm

