# pushgover
Hobby implementation of a small subset of the [pushover](https://pushover.net) API

## Usage
Given the environment variables PUSHGOVER_APPTOKEN and PUSHGOVER_USERKEY are set to one's pushover apptoken and userkey respectively, one can send a message to the declared pushover application by invoking `pushgover $message`. This can be used for alerting via push messages from cron jobs or simply to send a custom push message to one's phone.