if [ -x /usr/bin/dircolors ]; then
    test -r ~/.dircolors && eval "$(dircolors -b ~/.dircolors)" || eval "$(dircolors -b)"
    alias ls='ls --color=auto'
    alias dir='dir --color=auto'
    alias vdir='vdir --color=auto'

    alias grep='grep --color=auto'
    #alias fgrep='fgrep --color=auto'
    #alias egrep='egrep --color=auto'
fi

alias ll='ls -l'
alias la='ls -a'
alias lss='ls --tree --long'
alias lt='ls --tree'

if [ -f ~/.bash_aliases ]; then
    . ~/.bash_aliases
fi


export NVM_DIR="$HOME/.config/nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

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
#export PATH=$PATH:/usr/local/go/bin;
#export PATH=$PATH:/home/ashtyn372/.local/bin
export GOPATH=~/code/go
#export PATH=$PATH:$GOPATH/bin
alias ls="exa --color=always"
export coin_address="0cb1f2e59e7e45c7ef8c7a6f3e1a2e58bace7b0d7f98841cc0f0faa758fb0f73"
#source "$HOME/.cargo/env"
alias v="nvim"
alias vim="nvim"
alias code="cd ~/code"
code
#export PATH=$PATH:/home/ashtyn/.nvm/versions/node/v15.5.0/bin;
export PATH="$DENO_INSTALL/bin:$PATH"
export NVM_DIR="$HOME/.config/nvm"
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$GOPATH/bin
export DENO_INSTALL="/home/ashtyn/.deno"
export PATH="$DENO_INSTALL/bin:$PATH"
export PATH=~/code/depot_tools:$PATH
export PATH=/var/lib/snapd/snap/bin:$PATH
alias python="python2"
alias clang-format="clang-format -style='{IndentWidth: 8}'"
