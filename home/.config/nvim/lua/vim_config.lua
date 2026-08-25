local o = vim.opt
vim.g.mapleader = ' '          -- space is the leader key
-- indentation
o.tabstop = 2                  -- tab width
o.shiftwidth = 2               -- indent width
o.softtabstop = 2              -- spaces per tab/backspace, not tabs
o.expandtab = true             -- use spaces instead of tabs
o.smartindent = true           -- smart auto-indent
o.autoindent = true            -- copy indent from current line

-- line numbers & cursor display
o.number = true                -- line number
o.relativenumber = true        -- relative line numbers
o.cursorline = true            -- highlight current line
o.wrap = false                 -- do not wrap lines by default
o.scrolloff = 10               -- keep 10 lines above/below cursor
o.sidescrolloff = 10           -- keep 10 columns to left/right of cursor

-- search
o.ignorecase = true            -- case insensitive search
o.smartcase = true             -- case sensitive if uppercase in string
o.hlsearch = true              -- highlight search matches
o.incsearch = true             -- show matches as you type

-- ui & appearance
o.signcolumn = 'yes'                              -- always show a sign column
o.colorcolumn = '100'                             -- show a column at 100 chars
o.showmatch = true                                -- highlight matching brackets
o.cmdheight = 1                                   -- single line command line
o.completeopt = 'menuone,noinsert,noselect'       -- completion options
o.showmode = false                                -- hide mode, shown in statusline instead
o.pumheight = 10                                  -- popup menu height
o.pumblend = 10                                   -- popup menu transparency
o.winblend = 0                                    -- floating window transparency
o.conceallevel = 2                                -- obsidian requirement
o.concealcursor = ''                              -- do not hide cursorline in markup
o.synmaxcol = 300                                 -- syntax highlighting column limit
o.fillchars = { eob = ' ' }                       -- hide "~" on empty lines

-- files & behavior
o.encoding = 'utf-8'           -- use UTF-8 encoding
local undodir = vim.fn.expand('~/.vim/undodir')
if vim.fn.isdirectory(undodir) == 0 then  -- create undodir if nonexistent
  vim.fn.mkdir(undodir, 'p')
end
o.backup = false               -- do not create a backup file
o.writebackup = false          -- do not write to a backup file
o.swapfile = false             -- do not create a swapfile
o.undofile = true              -- do create an undo file
o.undodir = undodir            -- set the undo directory
o.updatetime = 300             -- faster completion
o.timeoutlen = 500             -- timeout duration
o.ttimeoutlen = 50             -- key code timeout
o.autoread = true              -- auto-reload changes if outside of neovim
o.autowrite = false            -- do not auto-save

-- buffers & interaction
o.hidden = true                -- allow hidden buffers
o.errorbells = false           -- no error sounds
o.backspace = 'indent,eol,start' -- better backspace behaviour
o.autochdir = false            -- do not autochange directories
o.iskeyword:append('-')        -- include - in words
o.path:append('**')            -- include subdirs in search
o.selection = 'inclusive'      -- include last char in selection
o.mouse = 'a'                  -- enable mouse support
o.clipboard:append('unnamedplus') -- use system clipboard
o.modifiable = true            -- allow buffer modifications

-- cursor
o.guicursor = 'n-v-c:block,i-ci-ve:block,r-cr:hor20,o:hor50,a:blinkwait700-blinkoff400-blinkon250-Cursor/lCursor,sm:block-blinkwait175-blinkoff150-blinkon175' -- cursor blinking and settings

-- folding: requires treesitter available at runtime; safe fallback if not
o.foldmethod = 'expr'                          -- use expression for folding
o.foldexpr = 'v:lua.vim.treesitter.foldexpr()' -- use treesitter for folding
o.foldlevel = 99                               -- start with all folds open

-- splits
o.splitbelow = true            -- horizontal splits go below
o.splitright = true            -- vertical splits go right

-- completion & performance
o.wildmenu = true              -- tab completion
o.wildmode = 'longest:full,full' -- longest common match, then full list, cycle with Tab
o.diffopt:append('linematch:60') -- improve diff display
o.redrawtime = 10000           -- increase neovim redraw tolerance
o.maxmempattern = 20000        -- increase max memory

