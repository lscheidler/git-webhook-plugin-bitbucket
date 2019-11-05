git-webhook-plugin-bitbucket
============================

Plugin for [git-webhook-lambda](https://github.com/lscheidler/git-webhook-lambda),
which implements the [git-webhook-plugin](https://github.com/lscheidler/git-webhook-plugin)
interface.

Current supported bitbucket webhooks
------------------------------------

| Name      | Description           |
|-----------|-----------------------|
| repo:push | repository push event |

Features/Caveats
----------------

This plugin tries to validate incoming events and ignores invalid events. If
bitbucket changes the format for an event, this will break the validation and
the particular event will be ignored.
