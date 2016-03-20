#file: .bashrc
umask 007 #New files are private by default
#ulimit -Sc unlimited #Programs can dump cores

#=============================================================
#Environment Variables
#=============================================================
export PATH="/local/bin:/usr/local/bin:/bin:/usr/bin:/sbin:/usr/sbin"
export PATH="$PATH:$HOME/bin"
export MANPATH="/local/share/man:/usr/local/man:/usr/share/man:/usr/local/share/man"
export EDITOR=gvim
export BROWSER=/usr/bin/chromium-browser

export TERM="xterm-256color"

BLUE="\[$(tput setaf 12)\]"
RESET="\[$(tput sgr0)\]"

export PS1="${BLUE}\u@\h: \w\$${RESET} "


#= CS138: Distributed Systems ===============================
export GOPATH="$HOME/cs138"
export PATH="$PATH:$GOPATH/bin"




#=============================================================
#Aliases
#=============================================================
alias ls="ls -F --color=auto"
alias rm="rm -i"
alias sshb="ssh -Y wk@ssh.cs.brown.edu"
alias vpnb="sudo openvpn ~/misc/browncs-gateall.ovpn"
alias matlab="matlab -nodesktop -nosplash"
alias vmatlab="$(which matlab)"


cd ~/ #Sunlab hack, get rid of /gpfs/main/home/wkochans



