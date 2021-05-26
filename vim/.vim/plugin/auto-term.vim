function TermandResize()
    :vsplit
    :wincmd l
    :term
    :vertical resize -10
endfunction
command -nargs=1 -complete=file Do :call TermandResize() <bar> :term <args>
function PipeCmdToTerm()
    let INPUT = input("? ")
    :execute ':Do '.INPUT
endfunction

nnoremap <leader>t :call TermandResize() <CR>
nnoremap <Leader>r :call PipeCmdToTerm()<CR>
