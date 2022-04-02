package(default_visibility = ["PUBLIC"])

go_module(
    name = "testify",
    install = [
        "assert",
        "require",
    ],
    module = "github.com/stretchr/testify",
    test_only = True,
    version = "v1.4.0",
    deps = [
        ":difflib",
        ":spew",
        ":yaml",
    ],
)

go_module(
    name = "spew",
    install = ["spew"],
    module = "github.com/davecgh/go-spew",
    version = "v1.1.1",
)

go_module(
    name = "difflib",
    install = ["difflib"],
    module = "github.com/pmezard/go-difflib",
    version = "v1.0.0",
)

go_module(
    name = "zap",
    install = [
        "",
        "internal/bufferpool",
        "internal/exit",
        "internal/color",
        "buffer",
        "zapcore",
    ],
    module = "go.uber.org/zap",
    version = "v1.14.1",
    deps = [
        ":atomic",
        ":multierr",
    ],
)

go_module(
    name = "x_crypto",
    install = [
        "ssh/terminal",
    ],
    module = "golang.org/x/crypto",
    version = "v0.0.0-20191011191535-87dc89f01550",
    deps = [
        "//third_party/x:sys",
    ],
)

# go_module(
#     name = "x_sys",
#     install = [
#         "unix",
#     ],
#     module = "golang.org/x/sys",
#     version = "v0.0.0-20200301040627-c5d0d7b4ec88",
# )

go_module(
    name = "x_text",
    install = [
        "width",
        "transform",
    ],
    module = "golang.org/x/text",
    version = "v0.3.0",
)

# go_module(
#     name = "x_tools",
#     module = "golang.org/x/tools",
#     version = "v0.0.0-20200708181441-6004c8539734",  # indirect
# )

go_module(
    name = "check",
    module = "gopkg.in/check.v1",
    version = "v1.0.0-20190902080502-41f04d3bba15",  # indirect
    deps = [
        ":pretty",
    ],
)

go_module(
    name = "yaml",
    module = "gopkg.in/yaml.v2",
    version = "v2.2.4",  # indirect
)

# go_module(
#     name = "tools",
#     module = "honnef.co/go/tools",
#     version = "v0.0.1-2020.1.4",  # indirect
# )

go_module(
    name = "atomic",
    module = "go.uber.org/atomic",
    version = "v1",
)

go_module(
    name = "multierr",
    module = "go.uber.org/multierr",
    version = "v1.8.0",
    deps = [
        ":atomic",
    ],
)

go_module(
    name = "pretty",
    module = "github.com/kr/pretty",
    version = "d928460c8d68bc4455b4a0664899a846e8327c49",
    deps = [
        ":go_internal",
        ":text",
    ],
)

go_module(
    name = "text",
    module = "github.com/kr/text",
    version = "v0.2.0",
)

go_module(
    name = "go_internal",
    install = [
        "fmtsort",
    ],
    module = "github.com/rogpeppe/go-internal",
    version = "v1.8.1",
)
