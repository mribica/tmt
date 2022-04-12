# 🍅 tmt

Command-line application for Pomodoro technique time tracking.

![tmt demo](examples/demo.gif)

## Overview

tmt is a Pomodoro timer command-line application written in Go. It's simple to use, configure, and integrate with your everyday command line flow.

By default, tmt is set for 25 minutes of work followed by a short 5-minute break. Every four completed work sessions are followed by, a long, 15 minutes break. 

A work session is indivisible so there is no option to pause the timer. If you are interrupted you need to stop the application and abandon the incompleted work session.

There is an option to set up notifications for work or pause sessions. See configuration for more details.

## Build and Run

```
$ make build
$ bin/tmt
```

Use `Ctrl+C` to stop timer.

## Custom Configuration

tmt custom options are configurable and can be defined in `$HOME/.tmt/config.json`. Use [Example config](example/config.json) as a reference. This example is using macOS `say` command for work and pause notifications.

**config.json**
```json
{
    "pomodoroLength":   25,
    "shortBreakLength":  5,
    "longBreakLength":  15,
    "pomodoroCmd": {
        "cmd": "say",
        "args": ["Start working"]
    },
    "shortBreakCmd": {
        "cmd": "say",
        "args": ["Take a 5 minute break"]
    },
    "longBreakCmd": {
        "cmd": "say",
        "args": ["Take a 15 minutes break"]
    }
}
```
### Duration Configuration

- `pomodoroLength` duration of the work item in minutes (default 25)
- `shortBreakLength` duration of a short break in minutes (default 5)
- `longBreakLength` duration of a long break after each four work items are completed (default 15)

### Commands Configuration

- `pomodoroCmd` command is executed before each work item
- `shortBreakCmd`command is executed before a short break
- `longBreakCmd` command is executed before a long break
    - `cmd` is the name of the command you want to execute
    - `args` is a list of arguments passed to the command

You can use [https://github.com/julienXX/terminal-notifier](terminal-notifier) to show macOS notifications.

```json
"pomodoroCmd": {
    "cmd": "terminal-notifier",
    "args": ["-title", "🍅 TMT", "-message", "Start working for 25 minutes.", "-group", "tmt"]
},
```