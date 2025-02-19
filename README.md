# iMprovedkancli

This is my improved version of the [Charmbracelet's Kanban CLI demo repository](https://github.com/charmbracelet/kancli).

My goal is to deepen my understanding of Charmbracelet tools by modifying and extending this project.

## 📜 What's New

Here are the features and improvements I’ve added (listed from oldest to newest):

- [fix(board): allow setting default selected column](https://github.com/mustafa-ozturk/iMkancli/commit/a5a52b0)
- [feat(board): add status message for task creation](https://github.com/mustafa-ozturk/iMkancli/commit/42adfad)
- [feat(board): add status message for task deletion](https://github.com/mustafa-ozturk/iMkancli/commit/606017a)
- [feat(board): add status message for task movement](https://github.com/mustafa-ozturk/iMkancli/commit/5e8f27b)
- [refactor(board): add moveFocus() to simplify column navigation](https://github.com/mustafa-ozturk/iMkancli/commit/74501fa)
- [fix(board): ensure only the focused column has a selected task on startup](https://github.com/mustafa-ozturk/iMkancli/commit/4ea71e9)
- [refactor(column): remove unnecessary tea.Sequence in MoveToNext()](https://github.com/mustafa-ozturk/iMkancli/commit/14408d1)
- [fix(column): delete task when pressing Enter in 'Done' column](https://github.com/mustafa-ozturk/iMkancli/commit/b5af3ce)
- [feat(board): load tasks from JSON file](https://github.com/mustafa-ozturk/iMkancli/commit/578f118)
- [fix(board): replace list.SetItems() with list.InsertItem() in loadTasks()](https://github.com/mustafa-ozturk/iMkancli/commit/152dd89)
- [feat(board): save tasks to JSON file before quitting](https://github.com/mustafa-ozturk/iMkancli/commit/09f7c89)
- [feat(theme): add default theme colors](https://github.com/mustafa-ozturk/iMkancli/commit/0d96d19)
- [feat(column): Add custom styling for list](https://github.com/mustafa-ozturk/iMkancli/commit/2fbc68d)
- [feat(column): use color from DefaultTheme for foucsed column](https://github.com/mustafa-ozturk/iMkancli/commit/fd9fa1f)
- [feat(board): use color from DefaultTheme on status message](https://github.com/mustafa-ozturk/iMkancli/commit/72b1b26)
- [test(task): add initial tests](https://github.com/mustafa-ozturk/iMkancli/commit/e28b719)
- [test(board): add initials tests for board](https://github.com/mustafa-ozturk/iMkancli/commit/09f4718)
- [feat(board): show task count in column titles](https://github.com/mustafa-ozturk/iMkancli/commit/69d9a0e)
- [chore(script): add resetData.sh to restore default data.json](https://github.com/mustafa-ozturk/iMkancli/commit/9623572)


## 📋 TODO

- Record a demo video/gif for README
- Auto-save tasks on change
- Add basic settings (.config file)
- Optional PIN protection (encrypt/decrypt data)
- More themes
- Better test coverage

## 🤝 Contributing

I'm currently developing **iMkancli** as a personal project to learn more about the Charmbracelet ecosystem.
While I'm not looking for contributors right now, feel free to check out the code and experiment with it.

### 🔧 Setup & Running

1. Clone the repo
```sh
git clone https://github.com/mustafa-ozturk/iMkancli.git
cd iMkancli
```

2. Run the app
```sh
go run main.go
```

3. Run tests
```sh
go test ./...
```
