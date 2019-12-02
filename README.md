# terratest-helm-tester

<img src="https://img.shields.io/docker/build/seanson/terratest-helm-tester?style=for-the-badge" />

A prebuilt Docker container for running Helm chart spec tests.

## Usage

Create a directory in your project folder for your tests and write a test as per the [example helm test](./test/helm_test.go).

`docker run -v ${PWD}/tests: -v ${PWD}:/app/test seanson/terratest-helm-tester go test -v ./test/test/helm_test.go`

## TODO

- [x] Sort out better mounting path for tests
- [ ] Document testing tags
- [ ] Show real CI usage
