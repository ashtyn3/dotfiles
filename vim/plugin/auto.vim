fun! TrimWhitespace()
    let l:save = winsaveview()
    keeppatterns %s/\s\+$//e
    call winrestview(l:save)
endfun

autocmd BufWritePre *.go :silent call CocAction('runCommand', 'editor.action.organizeImport')
autocmd BufWritePost * :silent :Prettier
augroup MAIN
    autocmd!
    autocmd BufWritePost * :silent :call TrimWhitespace()
augroup END
