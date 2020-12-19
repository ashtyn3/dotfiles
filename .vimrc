syntax on

set noerrorbells
set tabstop=4 softtabstop=4
set shiftwidth=4
set expandtab
set smartindent
set nu
set nowrap
set smartcase
set noswapfile
set nobackup
set undodir=~/.vim/undodir
set undofile
set incsearch

call plug#begin('~/.vim/plugged')

Plug 'jremmen/vim-ripgrep'
Plug 'tpope/vim-fugitive'
Plug 'leafgarland/typescript-vim'
Plug 'lyuts/vim-rtags'
Plug 'git@github.com:kien/ctrlp.vim.git'
Plug 'mbbill/undotree'
Plug 'neoclide/coc.nvim', {'branch': 'release'}
Plug 'alvan/vim-closetag' 
Plug 'neoclide/coc-prettier'
Plug 'mxw/vim-jsx'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'jacoborus/tender.vim'
Plug 'gruvbox-community/gruvbox'
Plug 'neovim/nvim-lspconfig'
Plug 'nvim-lua/completion-nvim'

Plug '~/code/gruber/'
call plug#end()

colorscheme gruvbox 
let g:gruvbox_italic=1
let g:gruvbox_bold=1
set background=dark
set number relativenumber
set nu rnu

let g:completion_matching_strategy_list = ['exact', 'substring', 'fuzzy']
lua require'lspconfig'.tsserver.setup{ on_attach=require'completion'.on_attach }
set completeopt=menuone,noinsert,noselect

if executable('rg')
    let g:rg_derive_root='true'
endif

let g:ctrlp_user_command = ['.git/', 'git --git-dir=%s/.git ls-files -oc --exclude-standard' ]
let mapleader = " "
let g:netrw_banner = 0
let g:netrw_winsize = 25
let g:netrw_localrmdir='rm -r'
set backspace=indent,eol,start
autocmd BufWritePre *.go :silent call CocAction('runCommand', 'editor.action.organizeImport')
let g:closetag_filenames = "*.html,*.xhtml,*.phtml,*.php,*.jsx,*.js"
command! -nargs=0 Prettier :CocCommand prettier.formatFile
let g:airline_solarized_bg='dark'
let g:airline_theme='solarized'
nnoremap <leader>h :wincmd h<CR>
nnoremap <leader>j :wincmd j<CR>
nnoremap <leader>k :wincmd k<CR>
nnoremap <leader>l :wincmd l<CR>
nnoremap <leader>u :UndotreeShow<CR>
nnoremap <leader>pv :wincmd v<bar> :Ex <bar> :vertical resize 30<CR>
nnoremap <Leader>= :vertical resize +5<CR>
nnoremap <Leader>- :vertical resize -5<CR>
nnoremap <Leader>ps :Rg<SPACE>
nnoremap <Leader>s :w<CR>
nnoremap <A-j> :m .+1<CR>==
nnoremap <A-k> :m .-2<CR>==
inoremap <A-j> <Esc>:m .+1<CR>==gi
inoremap <A-k> <Esc>:m .-2<CR>==gi
vnoremap <A-j> :m '>+1<CR>gv=gv
vnoremap <A-k> :m '<-2<CR>gv=gv

nnoremap <leader>va :lua vim.lsp.buf.definition()<CR>
nnoremap <leader>vd :lua vim.lsp.buf.definition()<CR>
nnoremap <leader>vi :lua vim.lsp.buf.implementation()<CR>
nnoremap <leader>vsh :lua vim.lsp.buf.signature_help()<CR>
nnoremap <leader>vrr :lua vim.lsp.buf.references()<CR>
nnoremap <leader>vrn :lua vim.lsp.buf.rename()<CR>
nnoremap <leader>vh :lua vim.lsp.buf.hover()<CR>
nnoremap <leader>vca :lua vim.lsp.buf.code_action()<CR>
nnoremap <leader>vsd :lua vim.lsp.util.show_line_diagnostics(); vim.lsp.util.show_line_diagnostics()<CR>

" Get rid of arrow keys in all modes in vim.
nnoremap <silent> <Leader>gd :YcmCompleter GoTo<CR>
nnoremap <silent> <Leader>gd :YcmCompleter FixIt<CR>
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

au BufRead,BufNewFile *.kpp set filetype=kpp
au! Syntax kpp source ~/.vim/syntax/kpp.vim


