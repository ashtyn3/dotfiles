# ~/.bashrc: executed by bash(1) for non-login shells.
# see /usr/share/doc/bash/examples/startup-files (in the package bash-doc)
# for examples

# If not running interactively, don't do anything
case $- in
    *i*) ;;
      *) return;;
esac

# don't put duplicate lines or lines starting with space in the history.
# See bash(1) for more options
HISTCONTROL=ignoreboth

# append to the history file, don't overwrite it
shopt -s histappend

# for setting history length see HISTSIZE and HISTFILESIZE in bash(1)
HISTSIZE=1000
HISTFILESIZE=2000

# check the window size after each command and, if necessary,
# update the values of LINES and COLUMNS.
shopt -s checkwinsize

# If set, the pattern "**" used in a pathname expansion context will
# match all files and zero or more directories and subdirectories.
#shopt -s globstar

# make less more friendly for non-text input files, see lesspipe(1)
#[ -x /usr/bin/lesspipe ] && eval "$(SHELL=/bin/sh lesspipe)"

# set variable identifying the chroot you work in (used in the prompt below)
if [ -z "${debian_chroot:-}" ] && [ -r /etc/debian_chroot ]; then
    debian_chroot=$(cat /etc/debian_chroot)
fi

# set a fancy prompt (non-color, unless we know we "want" color)
case "$TERM" in
    xterm-color|*-256color) color_prompt=yes;;
esac

# uncomment for a colored prompt, if the terminal has the capability; turned
# off by default to not distract the user: the focus in a terminal window
# should be on the output of commands, not on the prompt
#force_color_prompt=yes

if [ -n "$force_color_prompt" ]; then
    if [ -x /usr/bin/tput ] && tput setaf 1 >&/dev/null; then
	# We have color support; assume it's compliant with Ecma-48
	# (ISO/IEC-6429). (Lack of such support is extremely rare, and such
	# a case would tend to support setf rather than setaf.)
	color_prompt=yes
    else
	color_prompt=
    fi
fi

if [ "$color_prompt" = yes ]; then
    PS1='${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]\$ '
else
    PS1='${debian_chroot:+($debian_chroot)}\u@\h:\w\$ '
fi
unset color_prompt force_color_prompt

# If this is an xterm set the title to user@host:dir
case "$TERM" in
xterm*|rxvt*)
    PS1="\[\e]0;${debian_chroot:+($debian_chroot)}\u@\h: \w\a\]$PS1"
    ;;
*)
    ;;
esac

# enable color support of ls and also add handy aliases
if [ -x /usr/bin/dircolors ]; then
    test -r ~/.dircolors && eval "$(dircolors -b ~/.dircolors)" || eval "$(dircolors -b)"
    alias ls='ls --color=auto'
    alias dir='dir --color=auto'
    alias vdir='vdir --color=auto'

    alias grep='grep --color=auto'
    #alias fgrep='fgrep --color=auto'
    #alias egrep='egrep --color=auto'
fi

# colored GCC warnings and errors
#export GCC_COLORS='error=01;31:warning=01;35:note=01;36:caret=01;32:locus=01:quote=01'

# some more ls aliases
alias ll='ls -l'
alias la='ls -a'
alias lss='ls --tree --long'
alias lt='ls --tree'

# Alias definitions.
# You may want to put all your additions into a separate file like
# ~/.bash_aliases, instead of adding them here directly.
# See /usr/share/doc/bash-doc/examples in the bash-doc package.

if [ -f ~/.bash_aliases ]; then
    . ~/.bash_aliases
fi

# enable programmable completion features (you don't need to enable
# this, if it's already enabled in /etc/bash.bashrc and /etc/profile
# sources /etc/bash.bashrc).
if ! shopt -oq posix; then
  if [ -f /usr/share/bash-completion/bash_completion ]; then
    . /usr/share/bash-completion/bash_completion
  elif [ -f /etc/bash_completion ]; then
    . /etc/bash_completion
  fi
fi

export NVM_DIR="$HOME/.config/nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

# export GOOGLE_APPLICATION_CREDENTIALS="/home/ashtyn372/.firebase-auth.json";
# get current branch in git repo
# get current branch in git repo
function parse_git_branch() {
	BRANCH=`git branch 2> /dev/null | sed -e '/^[^*]/d' -e 's/* \(.*\)/\1/'`
	if [ ! "${BRANCH}" == "" ]
	then
		STAT=`parse_git_dirty`
		if [ "${BRANCH}" == "master" ]
		then
			# ✓
		if [ ! "${STAT}" == "" ]
		then
			echo -e "\b${STAT}"
		else
			echo -e "✓"
		fi
		else
			echo "${BRANCH}${STAT}"
		fi
	else
		echo -e "\b"
	fi
}

# get current status of git repo
function parse_git_dirty {
	status=`git status 2>&1 | tee`
	dirty=`echo -n "${status}" 2> /dev/null | grep "modified:" &> /dev/null; echo "$?"`
	untracked=`echo -n "${status}" 2> /dev/null | grep "Untracked files" &> /dev/null; echo "$?"`
	ahead=`echo -n "${status}" 2> /dev/null | grep "Your branch is ahead of" &> /dev/null; echo "$?"`
	newfile=`echo -n "${status}" 2> /dev/null | grep "new file:" &> /dev/null; echo "$?"`
	renamed=`echo -n "${status}" 2> /dev/null | grep "renamed:" &> /dev/null; echo "$?"`
	deleted=`echo -n "${status}" 2> /dev/null | grep "deleted:" &> /dev/null; echo "$?"`
	bits=''
	if [ "${renamed}" == "0" ]; then
		bits=">${bits}"
	fi
	if [ "${ahead}" == "0" ]; then
		bits="*${bits}"
	fi
	if [ "${newfile}" == "0" ]; then
		bits="+${bits}"
	fi
	if [ "${untracked}" == "0" ]; then
		bits="?${bits}"
	fi
	if [ "${deleted}" == "0" ]; then
		bits="x${bits}"
	fi
	if [ "${dirty}" == "0" ]; then
		bits="!${bits}"
	fi
	if [ ! "${bits}" == "" ]; then
		echo " ${bits}"
	else
		echo ""
	fi
}

# \[\e[0;1;91m\] red and bold text \[\e[0;94m\]
# export PS1="\[\e[33m\]\`parse_git_branch\`\[\e[m\]\[\e[30;46m\]\w\[\e[m\]\[\e[30;46m\] λ\[\e[m\] "
export PS1="\[\e[0m\]\u\[\e[m\] \[\e[0;91m\]።\[\e[m\] \[\e[0;1;92m\]\w\[\e[m\] \[\e[0;91m\]\`parse_git_branch\`\[\e[0;94m\] \[\e[0;94m\]»\[\e[m\] \[\e[0m\]$()\[\e0"
export PATH=$PATH:~/.cargo/bin
# tabtab source for electron-forge package
# uninstall by removing these lines or running `tabtab uninstall electron-forge`
[ -f /home/ashtyn372/xi-electron/node_modules/tabtab/.completions/electron-forge.bash ] && . /home/ashtyn372/xi-electron/node_modules/tabtab/.completions/electron-forge.bash

 alias g='git'

 if [ -f /etc/bash_completion ] && ! shopt -oq posix; then
     . /etc/bash_completion
 fi


 function_exists() {
     declare -f -F $1 > /dev/null
     return $?
 }
export editor='vim'
export PATH=$PATH:/usr/local/go/bin;
export PATH=$PATH:/home/ashtyn372/.local/bin
export JAVA_HOME="/usr/lib/jvm/java-1.11.0-openjdk-amd64"
export GOLAND_JDK="/usr/lib/jvm/java-1.11.0-openjdk-amd64"
export BOOST_ROOT=/usr/include/boost
export GOPATH=~/code
export PATH=$PATH:$GOPATH/bin
alias ls="exa --color=always"
export coin_address="0cb1f2e59e7e45c7ef8c7a6f3e1a2e58bace7b0d7f98841cc0f0faa758fb0f73"
tmux attach-session -t 0
source "$HOME/.cargo/env"
alias v="nvim"
alias vim="nvim"
alias code="cd ~/code"
code

export PATH=$PATH:/home/ashtyn/.nvm/versions/node/v15.5.0/bin;
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
export PATH=$PATH:$JAVA_HOME/bin
export ANDROID_HOME=~/Android
export PATH=$ANDROID_HOME/tools:$PATH
export PATH=$ANDROID_HOME/tools/bin:$PATH
export PATH=$ANDROID_HOME/platform-tools:$PATH
source "$HOME/.sdkman/bin/sdkman-init.sh"

export DISPLAY=:0
export LIBGL_ALWAYS_INDIRECT=1
export GO111MODULE=on

#THIS MUST BE AT THE END OF THE FILE FOR SDKMAN TO WORK!!!
export SDKMAN_DIR="/home/ashtyn/.sdkman"
[[ -s "/home/ashtyn/.sdkman/bin/sdkman-init.sh" ]] && source "/home/ashtyn/.sdkman/bin/sdkman-init.sh"
export DENO_INSTALL="/home/ashtyn/.deno"
export PATH="$DENO_INSTALL/bin:$PATH"
