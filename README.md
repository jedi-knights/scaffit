[//]: # ( ![scaffit logo]&#40;assets/scaffold.jpeg&#41;)

A Go utility to consistently scaffold Go projects.

[![CI](https://github.com/jedi-knights/scaffit/actions/workflows/ci.yml/badge.svg)](https://github.com/jedi-knights/scaffit/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jedi-knights/scaffit)](https://goreportcard.com/report/github.com/jedi-knights/scaffit)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/jedi-knights/scaffit)](https://pkg.go.dev/mod/github.com/jedi-knights/scaffit)

# Overview

Scaffit is an interactive Go utility providing users with a consistent way to scaffold Go projects.

It is intended to be suitable for use creating a variety of types of projects.  Due to its interactive nature it
should provide flexibility in the types of projects it can create.


## What does it take to make a module go gettable?

While working on this I started thinking intently about what it takes to make a module go gettable.

I would like Scaffit to always create modules and CLI packages that are go gettable. That being said
I am going back to do some more research on that topic.  I will update this readme with what I find.


> Do Go workspaces have any impact on how projects should be created?

## Usage

Creating a new project

```bash
scaffit new --module-path=<module-path> --dir=<project-dir>
```

This command will create a new empty Go project directory.

Where should it create the project directory?

If the target directory is not specified, it will set the project directory to a subdirectory 
of ~/go/src/github.com/<username>/<project-name>

this assumes your module path was github.com/<username>/<project-name>

So for example if you are creating a new project file that is supposed to go to your GitHub repo (assuming your
GitHub username is jdoe) 

```shell
scaffit new module --location ~/temp
```

Project layout (~/temp/github.com/jdoe/myproject)
--------------

    ├─ .husky/          Husky hooks
    ├─ assets/          an assets directory to drop things like images into
    ├─ cmd/             a cmd directory to house generic commands (or cobra commands)
    ├─ data/            a directory to house data access related files and/or models
    ├─ docs/            a directory to house documentation related files
    ├─ internal/        a directory to house internal packages
    ├─ pkg/             a directory to house public packages
    ├─ scripts/         a directory to house scripts
    ├─ test/            a directory to house tests
    ├─ .gitignore       a gitignore file
    ├─ .golangci.yml    a golangci.yml file
    ├─ .goreleaser.yml  a goreleaser.yml file
    ├─ .editorconfig    an editorconfig file
    ├─ .github/         a .github directory
    │  └─ workflows/    a workflows directory for your actions
    ├─ node_modules     a node_modules directory
    └─ test/            tests (see test/README.md)


> scaffit new --module-path=github.com/jdoe/myproject

Case 1: I'm trying to create a new project that's destined for my personal GitHub repo.

> scaffit new --module-path=github.com/jdoe/myproject

Case 2: I'm trying to create a new project that has a custom module path that is different from my GitHub repo.

> scaffit new --module-path=something-different

Case 3: I'm trying to create a new project that is targeted at a GitHub organization instead of my personal repo.

> scaffit new --module-path=github.com/myorg/myproject

Each of these cases must be handled differently, yet must be very easy to do.

In all cases, unless the directory path is specified the new project directory should be created in ~/go/src.


Adding a subcommand in Cobra

```bash
cobra add child -p 'parentCmd'
```

This should add a new subcommand named child to the parent command.
It assumes you have already created the parent command via `cobra add parent` first

## References

- [Go Project Layout](https://www.medium.com/golang-learn/go-project-layout-e5213cdcfaa2)
- [How to Write Go Code](https://golang.org/doc/code.html)

- [Prompt UI](https://github.com/manifoldco/promptui)
- [Reusing Workflows](https://docs.github.com/en/actions/using-workflows/reusing-workflows)
- [Using Semantic-Release with GitHub Actions](https://levelup.gitconnected.com/using-semantic-release-with-github-actions-c30d197829f1)
- [Semantic Release Configuration](https://semantic-release.gitbook.io/semantic-release/usage/configuration)
- [How to Use Cobra and Viper to Create Your First Golang CLI Tool](https://betterprogramming.pub/step-by-step-using-cobra-and-viper-to-create-your-first-golang-cli-tool-8050d7675093)
- [Building and Testing Go with GitHub Actions](https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go)
- [Control permissions for GITHUB_TOKEN](https://github.blog/changelog/2021-04-20-github-actions-control-permissions-for-github_token/)
- [Automatic Token Authentication](https://docs.github.com/en/actions/security-guides/automatic-token-authentication)
- [Permissions for the GitHub Token](https://docs.github.com/en/actions/security-guides/automatic-token-authentication#permissions-for-the-github_token)
- [How to Upgrade GoLang Dependencies](https://golang.cafe/blog/how-to-upgrade-golang-dependencies.html)
- [Learn Go with tests](https://quii.gitbook.io/learn-go-with-tests/)
- [Golang's Mocking Techniques](https://www.youtube.com/watch?v=LEnXBueFBzk&t=1401s)
- [Generics in Go](https://www.youtube.com/watch?v=F8Gl8-3ZW0E)
- [How to structure New Projects in Golang](https://www.youtube.com/watch?v=dJIUxvfSg6A)
- [Accept Interfaces and Return Structs](https://subscription.packtpub.com/book/programming/9781838647940/7/ch07lvl1sec45/accepting-interfaces-and-returning-structs#:~:text=There%20is%20a%20Go%20proverb,be%20structs%20or%20concrete%20types.)
- [Variables in Makefiles](http://aggregate.org/rfisher/Tutorials/Make/make5.html#:~:text=Variables%20in%20a%20makefile%20work,is%20replaced%20by%20the%20string.)
- [50 Shades of Go](http://golang50shad.es/)