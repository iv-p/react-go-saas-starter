const webpack = require("webpack");
const path = require("path");
const glob = require("glob");

const HtmlWebpackPlugin = require("html-webpack-plugin");
const { ESBuildPlugin, ESBuildMinifyPlugin } = require("esbuild-loader");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const CssMinimizerPlugin = require("css-minimizer-webpack-plugin");
const PurgecssPlugin = require("purgecss-webpack-plugin");
const CopyPlugin = require("copy-webpack-plugin");

const clientConfig = {
  entry: path.resolve(__dirname, "src/client.tsx"),
  module: {
    rules: [
      {
        test: /\.css$/,
        exclude: /node_modules/,
        use: [MiniCssExtractPlugin.loader, "css-loader", "postcss-loader"],
      },
      {
        test: /\.tsx?$/,
        loader: "esbuild-loader",
        options: {
          loader: "tsx",
          target: "es2015",
        },
      },
    ],
  },
  optimization: {
    minimize: true,
    minimizer: [
      new ESBuildMinifyPlugin({
        target: "es2015",
      }),
      new CssMinimizerPlugin({
        minimizerOptions: {
          preset: [
            "default",
            {
              discardComments: { removeAll: true },
            },
          ],
        },
      }),
    ],
    splitChunks: {
      cacheGroups: {
        default: false,
        vendors: false,
        styles: {
          name: "styles",
          test: /\.css$/,
          chunks: "all",
          enforce: true,
        },
        vendor: {
          chunks: "all",
          name: "vendor",
          test: /node_modules/,
        },
      },
    },
  },
  resolve: {
    extensions: [".tsx", ".ts", ".js", ".jsx", ".json", "css"],
  },
  plugins: [
    new CopyPlugin({
      patterns: [
        {
          from: path.resolve(__dirname, "public"),
          to: path.resolve(__dirname, "dist"),
        },
      ],
    }),
    new ESBuildPlugin(),
    new MiniCssExtractPlugin({
      filename: "[name].css",
    }),
    new HtmlWebpackPlugin({
      template: path.resolve(__dirname, "public", "index.html"),
      minify: true,
      hash: true,
    }),
    new webpack.ProvidePlugin({
      process: "process/browser",
      React: "react",
    }),
    new PurgecssPlugin({
      paths: glob.sync(`${path.join(__dirname, "src")}/**/*`, { nodir: true }),
    }),
  ],
};

const serverConfig = {
  target: "node",
  entry: path.resolve(__dirname, "server/index.tsx"),
  output: {
    filename: "server.js",
  },
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        loader: "esbuild-loader",
        options: {
          loader: "tsx",
          target: "es2015",
        },
      },
    ],
  },
  optimization: {
    minimize: false,
  },
  resolve: {
    extensions: [".tsx", ".ts", ".js", ".jsx", ".json"],
  },
  plugins: [
    new ESBuildPlugin(),
    new webpack.ProvidePlugin({
      React: "react",
    }),
  ],
};

module.exports = [clientConfig, serverConfig];
