// Finicky v4 (~/.finicky.js) — set Finicky as the default browser in
// System Settings › Desktop & Dock. Every link clicked outside a browser
// (Slack, Mail, ...) then opens Chrome with --new-window, so the window
// lands on the current workspace and toe tiles it next to the app,
// instead of Chrome focusing an existing window on another workspace.
export default {
  defaultBrowser: {
    name: "Google Chrome",
    args: ["--new-window"],
  },
};
