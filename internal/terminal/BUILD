go_library(
    name = "terminal",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    visibility = ["PUBLIC"],
    deps = [
        "//third_party:x_crypto",
        "//third_party:x_text",
        "//third_party:zap",
        "//third_party/x:sys",
    ],
)
