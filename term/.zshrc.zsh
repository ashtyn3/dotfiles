# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH
# Path to your oh-my-zsh installation.
#export ZSH="/home/ashtyn/.oh-my-zsh"


#MODE_INDICATOR="%F{yellow}+%f"
#VI_MODE_RESET_PROMPT_ON_MODE_CHANGE=true

# Autoload zsh add-zsh-hook and vcs_info functions (-U autoload w/o substition, -z use zsh style)
autoload -Uz add-zsh-hook vcs_info
# Enable substitution in the prompt.
KEYTIMEOUT=1

#VI_MODE_SET_CURSOR=true

setopt prompt_subst
# Run vcs_info just before a prompt is displayed (precmd)
add-zsh-hook precmd vcs_info
# add ${vcs_info_msg_0} to the prompt
# e.g. here we add the Git information in red
#PROMPT="%n %F{red}።%f %F{10}%~%f %F{red${vcs_info_msg_0_}%f %F{blue}»%f"
PROMPT='%n %F{red}።%f %B%F{10}%~%f%b%F{red}${vcs_info_msg_0_}%f %F{blue}» %f'
# Enable checking for (un)staged changes, enabling use of %u and %c
zstyle ':vcs_info:*' check-for-changes true
# Set custom strings for an unstaged vcs repo changes (*) and staged changes (+)
zstyle ':vcs_info:*' unstagedstr '*'
zstyle ':vcs_info:*' stagedstr '+'
# Set the format of the Git information for vcs_info
zstyle ':vcs_info:git:*' formats       ' %b%u%c'
zstyle ':vcs_info:git:*' actionformats ' %b|%a%u%c'

bindkey -v


#source $ZSH/oh-my-zsh.sh

alias grep='grep --colour=auto'
alias egrep='egrep --colour=auto'
alias fgrep='fgrep --colour=auto'

export editor='vim'
#export PATH=$PATH:/usr/local/go/bin;
#export PATH=$PATH:/home/ashtyn372/.local/bin
export GOPATH=~/code/go
export PATH=$PATH:$GOPATH/bin
alias ls="exa --color=always"
alias ll="exa -l"
alias lt="exa -T"
alias la="exa -a"
#source "$HOME/.cargo/env"
alias v="nvim"
alias vim="nvim"
alias code="cd ~/code"
alias pacman="sudo pacman"
code
#export PATH=$PATH:/home/ashtyn/.nvm/versions/node/v15.5.0/bin;
export PATH="$DENO_INSTALL/bin:$PATH"
#export NVM_DIR="$HOME/.config/nvm"
export NVM_DIR=~/.config/nvm
 [ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh"

export CHROMIUM_BUILDTOOLS_PATH="~/code/depot_tools/"
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$GOPATH/bin
export DENO_INSTALL="/home/ashtyn/.deno"
export PATH="$DENO_INSTALL/bin:$PATH"
export PATH=~/code/depot_tools:$PATH
export PATH=/var/lib/snapd/snap/bin:$PATH
alias python="python2"
alias clang-format="clang-format"
source $HOME/.cargo/env


export WORKSPACE="$HOME/.workspace"
export CT_PREFIX="$WORKSPACE/toolchains"
export PATH="$WORKSPACE/bin:$PATH"
