# terratest-helm-tester

[![Docker Build](https://img.shields.io/docker/cloud/automated/seanson/terratest-helm-tester?style=for-the-badge)](https://hub.docker.com/r/seanson/terratest-helm-tester)
[![Docker Build](https://img.shields.io/docker/cloud/build/seanson/terratest-helm-tester?style=for-the-badge)](https://hub.docker.com/r/seanson/terratest-helm-tester)
[![CircleCI Build](https://img.shields.io/circleci/build/github/seanson/terratest-helm-tester/master?style=for-the-badge)](https://circleci.com/gh/seanson/terratest-helm-tester)

A prebuilt Docker container for running Helm chart spec tests.

## Usage

Create a directory in your project folder for your tests and write a test as per the [example helm test](./test/helm_test.go).

`docker run -v ${PWD}/tests: -v ${PWD}:/app/test seanson/terratest-helm-tester`

This will run the tests found in the `./tests/...` path with the Terratest log wrapper and generate text and XML results.

## CI Support

### CircleCI

CircleCI does not let you copy images in and out of your container in Docker mode so machine mode is required.

See [.cicleci/config.yml](.circleci/config.yml) for details on running CI and copying tests out afterwards.

## TODO

- [x] Sort out better mounting path for tests
- [x] Add wrapper and log parser for generating test outputs
- [x] Show real CI usage
- [ ] Document testing tags
