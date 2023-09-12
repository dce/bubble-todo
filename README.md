# Bubble Todo

_A Bubble Tea example app_

## Running the app:

```bash
asdf plugin add golang https://github.com/asdf-community/asdf-golang.git
asdf install golang 1.19
git clone git@github.com:dce/bubble-todo.git
cd bubble-todo
make
bin/todo
```

---

## Notes/Links

* https://charm.sh/
* https://github.com/charmbracelet/bubbletea
* https://github.com/charmbracelet/bubbletea/tree/master/examples

---

## Commands in Bubble Tea

<https://charm.sh/blog/commands-in-bubbletea/>

> - **Use commands for all I/O.** By doing so your program will stay responsive, snappy, and maintainable. Even something as simple as reading a file from disk could cause a small lock up in your program, and commands are build to handle such cases beautifully.
> - **Only use commands for I/O.** Sometimes it’s tempting to use a command simply to send a message to another part of the program, however due to the nature of the way data flows in Bubble Tea this is never actually necessary.
> - **Never use goroutines within a Bubble Tea program.** Bubble Tea works best when you use commands and messages for communication.
