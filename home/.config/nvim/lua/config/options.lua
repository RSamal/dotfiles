-- Options are automatically loaded before lazy.nvim startup
-- Default options that are always set: https://github.com/LazyVim/LazyVim/blob/main/lua/lazyvim/config/options.lua
-- Add any additional options here

-- Additions (not set by LazyVim)
vim.opt.scrolloff = 10 -- keep 10 lines above/below cursor (LazyVim: 4)
vim.opt.sidescrolloff = 10 -- keep 10 columns left/right of cursor (LazyVim: 8)
vim.opt.softtabstop = 2 -- soft tab stop on tab/backspace
vim.opt.colorcolumn = "100" -- show a column at 100 chars
vim.opt.showmatch = true -- highlight matching brackets
vim.opt.concealcursor = "" -- do not conceal on the cursor line
vim.opt.synmaxcol = 300 -- syntax highlighting column limit
vim.opt.swapfile = false -- no swapfile
vim.opt.ttimeoutlen = 50 -- key code timeout
vim.opt.iskeyword:append("-") -- include - in words
vim.opt.path:append("**") -- include subdirs in :find
vim.opt.wildmode = "longest:full,full" -- command-line completion behaviour
vim.opt.diffopt:append("linematch:60") -- improve diff display
vim.opt.redrawtime = 10000 -- redraw tolerance
vim.opt.maxmempattern = 20000 -- pattern matching memory
vim.opt.guicursor =
  "n-v-c:block,i-ci-ve:block,r-cr:hor20,o:hor50,a:blinkwait700-blinkoff400-blinkon250-Cursor/lCursor,sm:block-blinkwait175-blinkoff150-blinkon175"

-- Custom undo directory (LazyVim default: ~/.local/state/nvim/undo)
local undodir = vim.fn.expand("~/.vim/undodir")
if vim.fn.isdirectory(undodir) == 0 then
  vim.fn.mkdir(undodir, "p")
end
vim.opt.undodir = undodir

-- Deliberate overrides of LazyVim defaults
vim.opt.laststatus = 2 -- per-window statusline (LazyVim: 3; pair with lualine globalstatus = false)
vim.opt.autowrite = false -- no auto-save (LazyVim: true)
vim.opt.timeoutlen = 500 -- mapping timeout (LazyVim: 300; also delays which-key)
vim.opt.updatetime = 300 -- CursorHold/swap interval (LazyVim: 200)
