load("@io_bazel_rules_docker//container:container.bzl", "container_image")
load("@npm//webpack-cli:index.bzl", "webpack_cli")

webpack_cli(
    name = "bundle",
    outs = ["dist"],
    args = [
        "--config",
        "$(execpath webpack.config.js)",
        "-o",
        "$@",
        "--mode production",
    ],
    data = [
        "webpack.config.js",
        "postcss.config.js",
        "//ts:tsconfig.json",
        "@npm//:node_modules",
    ] + glob([
        "src/**",
        "server/**",
        "public/**",
    ]),
)

container_image(
    name = "image",
    base = "@node15//image",
    cmd = [
        "node",
        "server.js",
    ],
    directory = "/app",
    files = [
        ":bundle",
    ],
    workdir = "/app/dist",
)
