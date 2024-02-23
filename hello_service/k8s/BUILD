docker_image(
    name = "image",
    srcs= ["//hello_service"],
    dockerfile = "Dockerfile",
    base_image = "//common/docker:base",
    run_args = "-p 8000:8000",
    visibility = ["//k8s/hello_service:all"],
    # see https://please.build/codelabs/k8s/index.html#7
)

k8s_config(
    name = "k8s",
    srcs = [
        "deployment.yaml",
        "service.yaml",
    ],
    containers = [":image"],
)