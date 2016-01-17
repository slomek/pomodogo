# PomodoGo - CLI Pomodoro Technique® tool

## Installation

In order to install the app you need to run

    go get -u github.com/slomek/pomodogo
    
In order to have a default notification icon, you may want to build it manually:

    make
    make pack
    
## Usage

For the most basic usage you just need to call the app's name:

    pomodogo
    
and the counter should start immediately, by default counting down 25 minutes. If the length is not good for you, it can be changed via a parameter, eg. the following

    pomodogo -duration 10m
    
will count down 10 minutes.

## The idea

The tool is based on a time management technique called Pomodoro Technique®. You can read more on it at [its official webpage](http://pomodorotechnique.com/).
  
## The author

You can contact me via many channels, just look me up at [pslomka.com](http://pslomka.com). You might also want to take a look at [My Code Smells](http://mycodesmells.com) blog.