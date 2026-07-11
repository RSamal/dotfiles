-- save by pressing Escape
vim.keymap.set('n', '<Esc>', ':w<CR>', { desc = 'Save' })
-- select all
vim.keymap.set('n', '<C-a>', 'ggVG', { desc = 'Select All' })
-- pasting over a selection no longer clobbers your clipboard
vim.cmd([[ xnoremap <expr> p 'pgv"'.v:register.'y' ]])

-- move between splits with ctrl + h/j/k/l (no <C-w> prefix).
-- these are the same keys vim-tmux-navigator uses, so if you add tmux later
-- the plugin makes them glide across nvim splits AND tmux panes seamlessly.
vim.keymap.set('n', '<C-h>', '<C-w>h', { desc = 'Go to left split' })
vim.keymap.set('n', '<C-j>', '<C-w>j', { desc = 'Go to below split' })
vim.keymap.set('n', '<C-k>', '<C-w>k', { desc = 'Go to above split' })
vim.keymap.set('n', '<C-l>', '<C-w>l', { desc = 'Go to right split' })
