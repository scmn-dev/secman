const smIconLink = "https://secman-assets.vercel.app/apps/icons/dsk.svg";

module.exports = {
  css: {
    loaderOptions: {
      sass: {
        prependData: `@import "@/styles/config/variables.scss";`,
      },
    },
  },
  pwa: {
    name: "Secman Hub",
    themeColor: "#1163E6",
    msTileColor: "#000000",

    manifestOptions: {
      background_color: "#000000",
      icons: [
        {
          src: smIconLink,
          sizes: "192x192",
          type: "image/svg+xml",
        },
        {
          src: smIconLink,
          sizes: "512x512",
          type: "image/svg+xml",
        },
        {
          src: smIconLink,
          sizes: "192x192",
          type: "image/svg+xml",
          purpose: "maskable",
        },
        {
          src: smIconLink,
          sizes: "512x512",
          type: "image/svg+xml",
          purpose: "maskable",
        },
      ],
    },
  },
};
