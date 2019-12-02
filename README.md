# terratest-helm-tester

A prebuilt Docker container for running Helm chart spec tests.

## Usage

Create a directory in your project folder for your tests and write a test as per the Terratest helm
[standard](https://github.com/gruntwork-io/terratest/blob/master/test/helm_basic_example_template_test.go).

`docker run -v ${PWD}/tests: -v ${PWD}/test:/go/src/github.com/seanson/terratest-helm-tester/test seanson/terratest-helm-tester`

## TODO

- Sort out better mounting path for tests
- Document testing tags
- Show real CI usage
