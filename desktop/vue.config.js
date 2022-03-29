const smDskIconURL = "https://assets.secman.dev/apps/icons/dsk.svg";

module.exports = {
  css: {
    loaderOptions: {
      sass: {
        prependData: `@import "@/styles/config/variables.scss";`,
      },
    },
  },
  pwa: {
    name: "Secman Desktop",
    themeColor: "#1163E6",
    msTileColor: "#000000",

    manifestOptions: {
      background_color: "#000000",
      icons: [
        {
          src: smDskIconURL,
          sizes: "192x192",
          type: "image/svg+xml",
        },
        {
          src: smDskIconURL,
          sizes: "512x512",
          type: "image/svg+xml",
        },
        {
          src: smDskIconURL,
          sizes: "192x192",
          type: "image/svg+xml",
          purpose: "maskable",
        },
        {
          src: smDskIconURL,
          sizes: "512x512",
          type: "image/svg+xml",
          purpose: "maskable",
        },
      ],
    },
  },
};
