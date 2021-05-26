let g:closetag_filenames = "*.html,*.xhtml,*.phtml,*.php,*.jsx,*.js"
command! -nargs=0 Prettier :CocCommand prettier.formatFile

call glaive#Install()

augroup autoformat_settings
    autocmd FileType bzl AutoFormatBuffer buildifier
    autocmd FileType c,cpp,proto,javascript AutoFormatBuffer clang-format
    autocmd FileType dart AutoFormatBuffer dartfmt
    autocmd FileType go AutoFormatBuffer gofmt
    autocmd FileType gn AutoFormatBuffer gn
    autocmd FileType html,css,sass,scss,less,json AutoFormatBuffer js-beautify
    autocmd FileType java AutoFormatBuffer google-java-format
    autocmd FileType python AutoFormatBuffer yapf
    " Alternative: autocmd FileType python AutoFormatBuffer autopep8
    autocmd FileType vue AutoFormatBuffer prettier
augroup END
"au BufNewFile,BufRead *.go setlocal noet ts=4 sw=4 sts=4
let g:prettier#autoformat_require_pragma = 0
"au BufWrite * :Autoformat
