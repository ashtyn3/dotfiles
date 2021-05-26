syntax on

set exrc
set noerrorbells
set tabstop=4 softtabstop=4
set shiftwidth=4
set expandtab
set hidden
set smartindent
set nu
set nowrap
set smartcase
set noswapfile
set nobackup
set undodir=~/.vim/undodir
set undofile
set incsearch
set nocompatible
set number relativenumber
set nu rnu
set nohlsearch
filetype off

set cmdheight=2
set updatetime=50
set shortmess+=c

call plug#begin('~/.vim/plugged')

Plug 'jremmen/vim-ripgrep'
Plug 'tpope/vim-fugitive'
Plug 'leafgarland/typescript-vim'
Plug 'lyuts/vim-rtags'
Plug 'ctrlpvim/ctrlp.vim'
Plug 'mbbill/undotree'
Plug 'neoclide/coc.nvim', {'branch': 'release'}
Plug 'alvan/vim-closetag'
Plug 'neoclide/coc-prettier'
Plug 'mxw/vim-jsx'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'gruvbox-community/gruvbox'
Plug 'neovim/nvim-lspconfig'
Plug 'nvim-lua/completion-nvim'
Plug 'mattn/emmet-vim'
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }
Plug 'cormacrelf/vim-colors-github'
Plug 'simlrh/neovim-node-plugin-example', { 'do': 'npm install' }
Plug 'google/vim-maktaba'
Plug 'google/vim-codefmt'
Plug 'google/vim-glaive'
Plug 'rakr/vim-two-firewatch'
Plug 'prettier/vim-prettier', { 'do': 'yarn install' }
Plug 'markeganfuller/vim-journeyman'
Plug 'ludovicchabant/vim-gutentags'
Plug 'kristijanhusak/vim-js-file-import', {'do': 'npm install'}
Plug 'Quramy/tsuquyomi'
Plug 'ayu-theme/ayu-vim'
Plug 'drsooch/gruber-darker-vim'
Plug 'wojciechkepka/bogster'
Plug 'Chiel92/vim-autoformat'



Plug 'nvim-telescope/telescope.nvim'
Plug 'nvim-lua/popup.nvim'
Plug 'nvim-lua/plenary.nvim'

call plug#end()



" colorscheme gruvbox
"colo gruvbox
"let g:gruvbox_italic=1
"let g:gruvbox_bold=1
"let g:gruvbox_contrast_dark = 'hard'
"set background=dark " or light if you prefer the light version

set termguicolors
"let ayucolor="mirage"   " for dark version of theme
colo bogster

let g:airline_theme='bogster' " if you have Airline installed and want the associated theme

hi Normal guibg=NONE ctermbg=NONE
hi SignColumn guibg=NONE
hi CursorLineNR guibg=NONE
autocmd vimenter * hi EndOfBuffer guibg=NONE ctermbg=NONE

autocmd BufRead,BufNewFile *.kl set filetype=javascript

" highlight Normal guibg=none ctermbg=none

" use the dark theme
" set background=dark
" colorscheme github
" let g:airline_theme = \"github\"
" let g:lightline = { 'colorscheme': 'github' }
au FileType * let b:prettier_exec_cmd = "prettier-stylelint"

let g:completion_matching_strategy_list = ['exact', 'substring', 'fuzzy']

" lsp-config stuff
lua require'lspconfig'.gopls.setup{ on_attach=require'completion'.on_attach }
lua require'lspconfig'.tsserver.setup{ on_attach=require'completion'.on_attach }
lua require'lspconfig'.rust_analyzer.setup{ on_attach=require'completion'.on_attach }
lua require'lspconfig'.clangd.setup{ on_attach=require'completion'.on_attach }



let g:go_fmt_autosave = 1
let g:go_fmt_command = 'goimports'

set completeopt=menuone,noinsert,noselect

if executable('rg')
    let g:rg_derive_root='true'
endif

let g:ctrlp_user_command = ['.git/', 'git --git-dir=%s/.git ls-files -oc --exclude-standard' ]
let mapleader = " "
let g:netrw_banner = 0
let g:netrw_winsize = 25
let g:netrw_localrmdir='rm -r'
let g:gutentags_enabled=0

set backspace=indent,eol,start

"let g:airline_solarized_bg='dark'
"let g:airline_theme='solarized'


tnoremap <Esc><Esc> <C-\><C-n>
nnoremap <leader>h :wincmd h<CR>
nnoremap <leader>nl :set nu! <bar> :set rnu!<CR>
nnoremap <leader>cs q:<C-n>
nnoremap <leader>j :wincmd j<CR>
nnoremap <leader>pp :set paste<CR>
nnoremap <leader>k :wincmd k<CR>
nnoremap <leader>l :wincmd l<CR>
nnoremap <leader>u :UndotreeShow<CR>
nnoremap <leader>pv :wincmd v<bar> :Ex <bar> :vertical resize 30<CR>
nnoremap <Leader>= :vertical resize +5<CR>
nnoremap <Leader>- :vertical resize -5<CR>
nnoremap <Leader>s :w<CR>
nnoremap <A-j> :m .+1<CR>==
nnoremap <A-k> :m .-2<CR>==
inoremap <A-j> <Esc>:m .+1<CR>==gi
inoremap <A-k> <Esc>:m .-2<CR>==gi
nnoremap J :m '>+1<CR>gv=gv
nnoremap K :m '<-2<CR>gv=gv

" Get rid of arrow keys in all modes in vim.
" Left arrow
nnoremap <Left> :echo "Stop!"<CR>
vnoremap <Left> :<C-u>echo "Stop!"<CR>
inoremap <Left> <C-o>:echo "Stop!"<CR>
" Right arrow
nnoremap <Right> :echo "Stop!"<CR>
vnoremap <Right> :<C-u>echo "Stop!"<CR>
inoremap <Right> <C-o>:echo "Stop!"<CR>
" Up arrow
nnoremap <Up> :echo "Stop!"<CR>
vnoremap <Up> :<C-u>echo "Stop!"<CR>
inoremap <Up> <C-o>:echo "Stop!"<CR>
" Down arrow
nnoremap <Down> :echo "Stop!"<CR>
vnoremap <Down> :<C-u>echo "Stop!"<CR>
inoremap <Down> <C-o>:echo "Stop!"<CR>
