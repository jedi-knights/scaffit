[//]: # ( ![scaffit logo]&#40;assets/scaffold.jpeg&#41;)

A Go utility to consistently scaffold Go projects.

[![Build](https://github.com/jedi-knights/scaffit/actions/workflows/build.yml/badge.svg)](https://github.com/jedi-knights/scaffit/actions/workflows/build.yml)
[![Lint](https://github.com/jedi-knights/scaffit/actions/workflows/lint.yml/badge.svg)](https://github.com/jedi-knights/scaffit/actions/workflows/lint.yml)
[![Test](https://github.com/jedi-knights/scaffit/actions/workflows/test.yml/badge.svg)](https://github.com/jedi-knights/scaffit/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jedi-knights/scaffit)](https://goreportcard.com/report/github.com/jedi-knights/scaffit)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/jedi-knights/scaffit)](https://pkg.go.dev/mod/github.com/jedi-knights/scaffit)

# Overview

Scaffit is an interactive Go utility providing users with a consistent way to scaffold Go projects.

It is intended to be suitable for use creating a variety of types of projects.  Due to its interactive nature it
should provide flexibility in the types of projects it can create.


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