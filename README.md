# Dotfile Configuration

Personal configuration management, including scripting tools.

Note that currently, none of this is real.

## Usage

To track items, use `dotfiles add`:

    dotfiles add .bashrc

This adds the file or directory to `.dotfiles/home` and creates a symbolic link
in `$HOME` referencing the file or directory.

Once you've added a file, you need to commit it to the repository.

    dotfiles commit

Which will `git add` the changed files and do a `git commit`.

If you want to stop tracking an item, run `dotfiles restore`:

    dotfiles restore .bashrc

This removes the symbolic link in `$HOME` and copies the original file back to
`$HOME` from `.dotfiles/home`.


#### bin

To track items from `$HOME/bin`, use `dotfiles bin add`:

    dotfiles bin add script-name

This adds the script to `.dotfiles/bin` and creates a symbolic link in
`$HOME/bin` referencing the file.

To save this change, run `dotfiles commit` as before.

To stop tracking an item, use `dotfiles bin restore`:

    dotfiles bin restore script-name


#### Sync Remote Changes

Remote changes can be applied using `dotfiles sync` which first checks
the remote repository, then creates a symbolic link in `$HOME` or `$HOME/bin`
for each file tracked under dotfiles.


#### Git Commands
Several commands are included to make it easier to perform git operations
on the `./dotfiles` repository. These include:

- `dotfiles pull`
- `dotfiles push`
- `dotfiles git`

Use the help command (`dotfiles help <command>`) to learn more about these.


### Configuration Scripts

`scripts` contains configuration scripts, including the above commands.

`bootstrap` should be run once to perform initial setup.

`install` contains installation scripts which can either be run individually:

    dotfiles install browser

Or all at once using the `--all` flag:

    dotfiles install --all


## Installation

    git clone https://github.com/wspk/dotfiles ~/.dotfiles
    ~/.dotfiles/scripts/bootstrap
    dofiles install --all
