#=========================================================================
# Environment Variables {{{
#=========================================================================
umask 007
export PATH="/home/wk/.local/bin:/home/wk/bin:$PATH"
export EDITOR="vim"
export PS1="\e[0;31m\w\$ \e[m"

setxkbmap -option ctrl:nocaps

#BASE16_SHELL="$HOME/.config/base16-shell/base16-default.dark.sh"
#[[ -s $BASE16_SHELL ]] && source $BASE16_SHELL


# }}}

#=========================================================================
# Aliases {{{
#=========================================================================
alias ls="ls -F --color=auto"
alias sshb="ssh -Y wk@ssh.cs.brown.edu"
alias vpnb="sudo openvpn ~/misc/browncs-gateall.ovpn"
alias matlab="matlab -nodesktop -nosplash"

# }}}
