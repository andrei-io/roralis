<h1 align=center>✨✨Go REST API Template✨✨

<br>
<img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/AndreiIonica/go-rest-api-template/Lint,%20Build%20and%20Test/main?style=for-the-badge" />
<img alt="License" src="https://img.shields.io/github/license/AndreiIonica/go-rest-api-template?style=for-the-badge" />
<img alt="go.mod go version" src="https://img.shields.io/github/go-mod/go-version/AndreiIonica/go-rest-api-template?style=for-the-badge" />
</h1>

## Get Started

```bash
git clone --depth=1 https://github.com/AndreiIonica/go-rest-api-template
cd go-rest-api-template
# Set up your own stuff, delete .git/ and init your own repo and then run
# You need to have bash installed for these to run. Git hooks weren't tested on Windows
bash scripts/setup_hooks.sh
bash scripts/install_dev_tools.sh
```

## Dependencies

-   Git Bash or some version of bash
-   Docker
-   Go
-   GolangCI Lint - Installed by script

## Why

When setting up a new Go REST API, there are a lot of repetitive steps and duplicated code from one project to another.
All the templates I found were either too simple or too complicated.
This repo serves as a starting point to Go REST API development

## Features

-   [x] Enviroment Config - viper
-   [x] Logger - default gin logger
-   [x] REST Framework - Gin
-   [x] Migrations - goose
-   [x] Dependency Injection - di(by sarulabs) is a dependecy injection container framework
-   [x] ORM - Gorm
-   [x] Tests - Only test business logic
-   [x] Swagger - Auto generated REST API docs
-   [x] Linter + CI - GolangCI Lint and Github Actions
-   [x] Commit Hooks - written in bash, not tested on windows

-   [x] Postgres - Persistent Store
-   [x] Docker - Container runtime
