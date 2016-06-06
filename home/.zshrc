#==============================================================================
# OH-MY-ZSH Configuration {{{
#==============================================================================
export ZSH=/home/wk/.oh-my-zsh                  #oh-my-zsh installation
ZSH_THEME="my-minimal"                          #zsh theme (custom) see: ~/.oh-my-zsh/themes/
DISABLE_AUTO_UPDATE="true"                      #disable bi-weekly auto-update checks
ZSH_CUSTOM=/home/wk/.zsh/                       #custom configurations

plugins=(git vi-mode)                           #load plugins, see: ~/.oh-my-zsh/plugins/

source $ZSH/oh-my-zsh.sh                        #note that this is run with the default path
# }}}

#==============================================================================
# Enironment Variables {{{
#==============================================================================
umask 007                                       #new files are private
export PATH="/home/wk/.local/bin:/home/wk/bin:$PATH"
export EDITOR="vim"

BASE16_SHELL="$HOME/.config/base16-shell/base16-default.dark.sh"
[[ -s $BASE16_SHELL ]] && source $BASE16_SHELL

export KEYTIMEOUT=1                             #shorten keytimeout for vi editing
# }}}

#==============================================================================
# Aliases {{{
#==============================================================================
alias ls="ls -F --color=auto"
alias sshb="ssh -Y wk@ssh.cs.brown.edu"
alias vpnb="sudo openvpn ~/misc/browncs-gateall.ovpn"
alias matlab="matlab -nodesktop -nosplash"
# }}}
