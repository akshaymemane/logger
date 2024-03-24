# Custom Logger Package

The Custom Logger package provides a flexible logging solution for Go applications, allowing users to configure log file, log level, and log format through environment variables.

## Installation

To use the Custom Logger package in your Go project, you can install it using the following command:

```bash
go get github.com/akshaymemane/logger
```
## Usage
Import the logger package in your Go code:
```bash
import "github.com/akshaymemane/logger"
```
Then, create a new instance of the logger using the New() function:
```bash
logger := logger.New()
```
You can then use the logger to log messages at different log levels:

```bash
logger.Debug("This is a debug message")
logger.Info("This is an info message")
logger.Warning("This is a warning message")
logger.Error("This is an error message")
```
## Configuration
The Custom Logger can be configured using the following environment variables:

* LOG_FILE: Specifies the path to the log file. Defaults to logfile.log if not set.
* LOG_LEVEL: Specifies the log level. Valid values are debug, info, warning, error, and all. Defaults to all.
* LOG_FORMAT: Specifies the log format using pipe-separated format specifiers. Available specifiers are date, time, microseconds, longfile, shortfile, utc, and stdflags. Defaults to date|time|shortfile.

### Example .env file:
```bash
LOG_FILE=application.log
LOG_LEVEL=info
LOG_FORMAT=date|time|shortfile
```
## Log Levels
The Custom Logger supports the following log levels:

* DebugLevel: Used for debugging information.
* InfoLevel: Used for informational messages.
* WarningLevel: Used for warning messages.
* ErrorLevel: Used for error messages.
* AllLevel: Log messages of all levels.

## Log Format
The log format determines how log entries are formatted. You can specify the log format using pipe-separated format specifiers. Available specifiers include date, time, microseconds, longfile, shortfile, utc, and stdflags.

## Contributing
If you encounter any issues or have suggestions for improvement, please feel free to open an issue or submit a pull request on GitHub.

## License
This project is licensed under the [MIT License](LICENSE) - see the LICENSE file for details.

