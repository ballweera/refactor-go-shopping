# Overview
This repo is to demonstrate refactoring legacy Go project from MVC structure style to the proper Go's structure.

# How to read this repo
- `master` branch contains the legacy version
- Refactoring has 5 steps. Each step contains commits that explain how I refactor Go structure
- `refactor-step1` branch is to move `user` domain to `user` package
- `refactor-step2` branch is to move `product` domain to `product` package
- `refactor-step3` branch is to rename components to proper name
- `refactor-step4` branch is to remove all legacy structure
- `refactor-step5` branch is to introduce `cmd` and `internal` package

# How to build
- Create directory `dist` at root project
- On `master` branch to step 1-4, you can build with `go build -o dist/shopping-api .`
- On `refactor-step5` branch, you can build with `go build -o dist/shopping-api ./cmd/shopping-api/`