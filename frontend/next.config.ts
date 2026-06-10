const basePath = process.env.NEXT_PUBLIC_BASE_PATH || "";

module.exports = {
  basePath,
  assetPrefix: basePath || "",
  images: { unoptimized: true },
  output: "standalone",
};
