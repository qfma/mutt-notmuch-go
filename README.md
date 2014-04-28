mutt-notmuch-go
===============

This program integrates notmuch search into mutt via a macro. This tool assumes that you have a working notmuch installation; otherwise head over [here](http://notmuchmail.org/).

You can either clone the directory and compile it in your local repository or you just install it using go get: `go get github.com/qfma/mutt-notmuch-go`

The search results are written to this Maildir folder:
`$HOME.cache/mutt_results` 

In order to enable the search in mutt add this to your .muttrc (Search is enabled by pressing /):
``` bash
macro index / "<enter-command>unset wait_key<enter><shell-escape>mutt-notmuch-go<enter><change-folder-readonly>~/.cache/mutt_results<enter>" \
"search mail (using notmuch)"
```

This script is inspired by another Python version available [here](https://github.com/honza/mutt-notmuch-py/)
