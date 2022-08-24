# logatee
Test log adapter for logrus

## Usage
Anywhere you would inject a `*logrus.Logger`, use `logatee.New()` to get a test-compatible logger. It will write all logs through `*testing.T`, providing nicely-formatted test output.
