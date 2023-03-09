<div align="center" markdown="1">

# [`app-web-tianyi`][url-repo]

[![License][badge-license]][url-license]
[![Status][badge-status-abandoned]][url-repo]

SPA for CI

</div>

## About The Project

In 2022 I've decided to make an SPA for CI using Golang for backend and Vue.js for frontend. 

After several months of effort, I stopped working on it because there are too many great open source CI/CD tools out there to make another one.

### Architecture diagram

![overview](./docs/diagrams/app-web-tianyi.drawio.svg)

### Code structure

Clean architecture

![Clean architecture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

```
tianyi
├── api [Interface Adapters]
│   ├── controller
│   ├── middleware
│   └── presenter
├── cmd (CLI interface)
├── docs (OpenAPI Specification and other documentation)
│   └── diagrams
├── entity [Enterprise Business Rules]
├── infrastructure [Frameworks and Drivers]
│   ├── access
│   ├── branch
│   ├── config
│   ├── job
│   ├── jwt
│   ├── lifecycle
│   ├── pipeline
│   ├── pool
│   ├── project
│   ├── schedule
│   ├── session
│   └── user
├── pkg (Support packages)
│   └── error
└── usecase [Application Business Rules]
│   ├── access
│   ├── branch
│   ├── job
│   ├── jwt
│   ├── lifecycle
│   ├── pipeline
│   ├── pool
│   ├── project
│   ├── schedule
│   ├── session
│   └── user
└── web (Web interface)
```

### Logo

![black logo](./web/public/images/logo-dark.svg)

![white logo](./web/public/images/logo-white.svg)

they are made with [`google-font-to-svg-path`](https://danmarshall.github.io/google-font-to-svg-path/)
using font
[`zen-maru-gothic`](https://fonts.adobe.com/fonts/zen-maru-gothic#licensing-section)

### Frontend

[`vuejs`](https://vuejs.org/) SPA with
[`primevue`](https://www.primefaces.org/primevue/)

#### Theme

modified [`sakai-vue`](https://github.com/primefaces/sakai-vue) theme
with pieces from
[primevue official website](https://github.com/primefaces/primevue)

[`sakai-vue` demo](https://www.primefaces.org/sakai-vue/)

### Backend

[`golang`](https://go.dev/) with
[`fiber`](https://docs.gofiber.io/)

#### Database

[`gorm`](https://gorm.io/) is used as a driver, so any of
[these databases](https://gorm.io/docs/connecting_to_the_database.html)
can be used

only `postgresql` is configured

#### Passwords

Passwords are hashed with [Argon2id](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id),
as suggested in
[Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)

## Resources

- architecture

  - https://gitlab.com/redhatdemocentral/portfolio-architecture-examples
    https://redhatdemocentral.gitlab.io/portfolio-architecture-tooling/index.html?#/portfolio-architecture-examples/projects/retail-data-framework.drawio
  - https://medium.com/wesionary-team/a-clean-architecture-for-web-application-in-go-lang-4b802dd130bb
  - https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html
  - https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/
  - https://manakuro.medium.com/clean-architecture-with-go-bce409427d31
  - https://github.com/eminetto/clean-architecture-go
  - https://eltonminetto.dev/en/post/2018-03-05-clean-architecture-using-go/
  - https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
  - https://proandroiddev.com/why-you-need-use-cases-interactors-142e8a6fe576
  - https://itnext.io/golang-and-clean-architecture-19ae9aae5683
  - https://hackernoon.com/golang-clean-archithecture-efd6d7c43047

- golang

  - https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go
  - https://github.com/gofiber/recipes
  - https://ldej.nl/post/generating-swagger-docs-from-go/
  - https://github.com/hashicorp/hcl
  - https://github.com/gothinkster/golang-gin-realworld-example-app
  - https://github.com/alpody/golang-fiber-realworld-example-app
  - https://dev.to/aryaprakasa/serving-single-page-application-in-a-single-binary-file-with-go-12ij
  - https://codesahara.com/blog/golang-job-queue-with-redis-streams/
  - https://gorm.io/
  - https://docs.gofiber.io/
  - https://github.com/thomasvvugt/fiber-boilerplate
  - https://github.com/ansible-semaphore/semaphore
  - https://github.com/spf13/cobra
  - https://github.com/spf13/viper

- security

  - https://cheatsheetseries.owasp.org/cheatsheets/
  - https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go

- frontend

  - https://github.com/primefaces/sakai-vue
  - https://github.com/primefaces/primevue
  
<!-- relative links -->

<!-- project links -->

[url-repo]: https://github.com/shishifubing/app-web-tianyi
[url-license]: https://github.com/shishifubing/app-web-tianyi/blob/main/LICENSE

<!-- external links -->

<!-- badge links -->

[badge-license]: https://img.shields.io/github/license/shishifubing/app-web-tianyi.svg
[badge-status-abandoned]: https://img.shields.io/badge/status-abandoned-red
