load("@io_bazel_rules_docker//container:container.bzl", "container_image")
load("@npm//webpack-cli:index.bzl", "webpack_cli")

webpack_cli(
    name = "bundle",
    outs = ["build"],
    args = [
        "./$(execpath src/index.tsx)",
        "--config",
        "$(execpath webpack.config.js)",
        "-o",
        "$@",
        "--mode production",
    ],
    data = [
        "webpack.config.js",
        "babel.config.json",
        "postcss.config.js",
        "//ts:tsconfig.json",
        "@npm//:node_modules",
    ] + glob([
        "src/**",
        "public/**",
    ]),
)

container_image(
    name = "nginx_base",
    base = "@nginx_stable//image",
    cmd = [
        "nginx",
        "-g",
        "daemon off;",
    ],
    directory = "/etc/nginx/conf.d/",
    files = [
        "default.conf",
    ],
    ports = [
        "80",
    ],
)

container_image(
    name = "image",
    base = ":nginx_base",
    directory = "/app",
    files = [
        ":bundle",
    ],
)
