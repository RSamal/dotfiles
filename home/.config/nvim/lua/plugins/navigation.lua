-- ~/.config/nvim/lua/plugins/navigation.lua

vim.pack.add({
  { src = "https://github.com/mikavilpas/yazi.nvim" },
  { src = "https://github.com/folke/snacks.nvim" },
})

-- yazi file browser
require("yazi").setup({
  open_for_directories = true, -- Integrates yazi to open when you pass a directory to nvim
})

vim.keymap.set("n", "<leader>e", "<cmd>Yazi<cr>", {
  desc = "File Browser (Yazi)",
})

-- Snacks
require("snacks").setup({
  picker = { enabled = true },
  notifier = { enabled = true },
  input = { enabled = true },
})

vim.keymap.set("n", "<leader>f", function()
  Snacks.picker.files()
end, { desc = "Find Files" })

vim.keymap.set("n", "<leader>s", function()
  Snacks.picker.grep()
end, { desc = "Search Text" })

vim.keymap.set("n", "<leader>b", function()
  Snacks.picker.buffers()
end, { desc = "Buffers" })

vim.keymap.set("n", "gd", function()
  Snacks.picker.lsp_definitions()
end, { desc = "Goto Definition" })
