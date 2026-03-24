import Foundation
import Combine

// MARK: - Shop Store
class ShopStore: ObservableObject {
    static let shared = ShopStore()
    
    private let shops: [Shop]
    
    private init() {
        shops = [
            Shop(
                id: "shop1",
                name: "Apple官方旗舰店",
                avatar: "https://picsum.photos/seed/apple-shop/100/100",
                description: "Apple官方旗舰店，正品保障，全国联保",
                rating: 4.9,
                productCount: 156,
                followerCount: 1250000,
                salesCount: 580000,
                createdAt: "2015-01-01",
                categories: ["手机", "电脑", "数码配件"],
                tags: ["官方正品", "全国联保", "七天无理由"]
            ),
            Shop(
                id: "shop2",
                name: "华为官方旗舰店",
                avatar: "https://picsum.photos/seed/huawei-shop/100/100",
                description: "华为官方旗舰店，科技创新，品质保证",
                rating: 4.8,
                productCount: 234,
                followerCount: 980000,
                salesCount: 420000,
                createdAt: "2016-01-01",
                categories: ["手机", "电脑", "智能家居"],
                tags: ["官方正品", "正品保障", "以旧换新"]
            ),
            Shop(
                id: "shop3",
                name: "小米官方旗舰店",
                avatar: "https://picsum.photos/seed/xiaomi-shop/100/100",
                description: "小米官方旗舰店，为发烧而生",
                rating: 4.7,
                productCount: 312,
                followerCount: 850000,
                salesCount: 390000,
                createdAt: "2017-01-01",
                categories: ["手机", "智能家居", "生活电器"],
                tags: ["官方正品", "七天无理由"]
            )
        ]
    }
    
    func getShop(by id: String) -> Shop? { shops.first { $0.id == id } }
    
    func getShop(byName name: String) -> Shop? { shops.first { $0.name == name } }
}
