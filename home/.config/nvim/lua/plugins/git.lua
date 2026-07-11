vim.pack.add({
  { src = "https://github.com/nvim-lua/plenary.nvim" },
  { src = "https://github.com/sindrets/diffview.nvim" },
  { src = "https://github.com/NeogitOrg/neogit" },
  { src = "https://github.com/lewis6991/gitsigns.nvim" },
})

-- Neogit
require("neogit").setup({
  integrations = {
    diffview = true,
  },
})

vim.keymap.set("n", "<leader>g", function()
  require("neogit").open()
end, { desc = "Neogit" })

-- Gitsigns
vim.api.nvim_create_autocmd("BufWinEnter", {
  once = true,
  callback = function()
    require("gitsigns").setup({
      current_line_blame = true,
    })
  end,
})
