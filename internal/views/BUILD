go_library(
    name = "views",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    visibility = ["PUBLIC"],
    deps = [
        "//internal/config",
        "//internal/filetree",
        "//internal/state",
        "//internal/terminal",
    ],
)
