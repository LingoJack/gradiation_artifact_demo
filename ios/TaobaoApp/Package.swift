// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "TaobaoApp",
    platforms: [
        .iOS(.v16)
    ],
    products: [
        .library(
            name: "TaobaoApp",
            targets: ["TaobaoApp"]),
    ],
    targets: [
        .target(
            name: "TaobaoApp",
            path: "Sources/TaobaoApp"
        ),
    ]
)
