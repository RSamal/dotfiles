// Finicky v4 (~/.finicky.js) — set Finicky as default browser in
// System Settings > Desktop & Dock. Links from outside a browser route to
// ChromeNewWindow.app (~/Applications, an osacompile applet) which runs
// `open -na "Google Chrome" --args --new-window <url>` — so the link opens
// a NEW Chrome window on the CURRENT workspace and toe tiles it here,
// instead of Chrome focusing an old window on another workspace.
// (Finicky v4 has no `args` support — that's why the applet exists.)
export default {
  defaultBrowser: {
    name: "ChromeNewWindow",
    appType: "appName",
  },
};
