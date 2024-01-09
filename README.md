# Advent of Code 2022 Workspace - Golang Edition


## Badges

[![Go Report Card](https://goreportcard.com/badge/github.com/vpayno/adventofcode-2022-golang-workspace)](https://goreportcard.com/report/github.com/vpayno/adventofcode-2022-golang-workspace)
[![Maintainability](https://api.codeclimate.com/v1/badges/605e8e2d133f7093cddf/maintainability)](https://codeclimate.com/github/vpayno/adventofcode-2022-golang-workspace/maintainability)

[![Test Coverage](https://api.codeclimate.com/v1/badges/605e8e2d133f7093cddf/test_coverage)](https://codeclimate.com/github/vpayno/adventofcode-2022-golang-workspace/test_coverage)

[![Go Checks](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/go.yml)
[![CodeQL Analysis](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/codeql-analysis-go.yml/badge.svg?branch=main)](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/codeql-analysis-go.yml)

[![Bash Checks](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/bash.yml/badge.svg?branch=main)](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/bash.yml)
[![Git Checks](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/git.yml/badge.svg?branch=main)](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/git.yml)
[![Link Check](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/links.yml/badge.svg?branch=main)](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/links.yml)
[![Woke Check](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/woke.yml/badge.svg?branch=main)](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/woke.yml)

[![Spelling Checks](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/misspell.yml/badge.svg?branch=main)](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/misspell.yml)
[![Yaml Workflow](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/yaml.yml/badge.svg?branch=main)](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/yaml.yml)
[![GH Actions Workflow](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/gh-actions.yml/badge.svg?branch=main)](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/gh-actions.yml)
[![Json Checks](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/json.yml/badge.svg?branch=main)](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/json.yml)
[![Markdown Checks](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/markdown.yml/badge.svg?branch=main)](https://github.com/vpayno/adventofcode-2022-golang-workspace/actions/workflows/markdown.yml)


## Directory Structure

- `cmd`: `main()` application and cli bootstrap
- `internal`: application libraries not accessible as a module externally
- `internal/aocshared`: internal application libraries shared by the challenges

- `data`: where the challenge and test data is stored

- `reports`: coverage reports

- `.github/workflows`: ci jobs for performing tests on the code base
- `scripts`: user and ci scripts


### Reports

└── reports \
    ├── [coverage-annotations.txt](./reports/coverage-annotations.txt) \
    └── [coverage-summary.txt](./reports/coverage-summary.txt)


### Code Shared Between Challenges

└── internal \
    └── aocshared \
        ├── [aocshared.go](./internal/aocshared/aocshared.go) \
        └── [aocshared_test.go](./internal/aocshared/aocshared_test.go)


## Daily Challenges


### Day 01 - Calorie Counting

- [Challenge Description](./calendar/day01-description.md)

├── cmd \
│   └── day01 \
│       ├── [aoc-day01.go](./cmd/day01/aoc-day01.go) \
│       └── [aoc-day01_test.go](./cmd/day01/aoc-day01_test.go) \
├── data \
│   └── [day01](./data/day01/) \
└── internal \
    └── day01 \
        ├── [app.go](./internal/day01/app.go) \
        ├── [app_test.go](./internal/day01/app_test.go) \
        ├── [init.go](./internal/day01/init.go) \
        └── [init_test.go](./internal/day01/init_test.go)


### Day 02 - Rock, Paper, Scissors

- [Challenge Description](./calendar/day02-description.md)

├── cmd \
│   └── day02 \
│       ├── [aoc-day02.go](./cmd/day02/aoc-day02.go) \
│       └── [aoc-day02_test.go](./cmd/day02/aoc-day02_test.go) \
├── data \
│   └── [day02](./data/day02/) \
└── internal \
    └── day02 \
        ├── [app.go](./internal/day02/app.go) \
        ├── [app_test.go](./internal/day02/app_test.go) \
        ├── [init.go](./internal/day02/init.go) \
        └── [init_test.go](./internal/day02/init_test.go)


### Day 03 - Rucksack Management

- [Challenge Description](./calendar/day03-description.md)

├── cmd \
│   └── day03 \
│       ├── [aoc-day03.go](./cmd/day03/aoc-day03.go) \
│       └── [aoc-day03_test.go](./cmd/day03/aoc-day03_test.go) \
├── data \
│   └── [day03](./data/day03/) \
└── internal \
    └── day03 \
        ├── [app.go](./internal/day03/app.go) \
        ├── [app_test.go](./internal/day03/app_test.go) \
        ├── [init.go](./internal/day03/init.go) \
        └── [init_test.go](./internal/day03/init_test.go)


### [Day 04 - Camp Cleanup](https://adventofcode.com/2022/day/4)

- [Challenge Description](./calendar/day04-description.md)

├── cmd \
│   └── day04 \
│       ├── [aoc-day04.go](./cmd/day04/aoc-day04.go) \
│       └── [aoc-day04_test.go](./cmd/day04/aoc-day04_test.go) \
├── data \
│   └── [day04](./data/day04/) \
└── internal \
    └── day04 \
        ├── [app.go](./internal/day04/app.go) \
        ├── [app_test.go](./internal/day04/app_test.go) \
        ├── [init.go](./internal/day04/init.go) \
        └── [init_test.go](./internal/day04/init_test.go)


### [Day 05 - Supply Stacks](https://adventofcode.com/2022/day/5)

- [Challenge Description](./calendar/day05-description.md)

├── cmd \
│   └── day05 \
│       ├── [aoc-day05.go](./cmd/day05/aoc-day05.go) \
│       └── [aoc-day05_test.go](./cmd/day05/aoc-day05_test.go) \
├── data \
│   └── [day05](./data/day05/) \
└── internal \
    └── day05 \
        ├── [app.go](./internal/day05/app.go) \
        ├── [app_test.go](./internal/day05/app_test.go) \
        ├── [init.go](./internal/day05/init.go) \
        └── [init_test.go](./internal/day05/init_test.go)
