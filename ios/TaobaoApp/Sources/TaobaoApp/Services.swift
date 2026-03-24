import Foundation
import Combine

// MARK: - Cart Store
class CartStore: ObservableObject {
    static let shared = CartStore()
    
    @Published var items: [CartItem] = []
    @Published var isEditing: Bool = false
    
    private init() {
        loadMockData()
    }
    
    var totalCount: Int { items.reduce(0) { $0 + $1.quantity } }
    
    var selectedItems: [CartItem] { items.filter { $0.isSelected } }
    
    var selectedCount: Int { selectedItems.reduce(0) { $0 + $1.quantity } }
    
    var selectedTotalPrice: Double { selectedItems.reduce(0) { $0 + $1.subtotal } }
    
    var isAllSelected: Bool { !items.isEmpty && items.allSatisfy { $0.isSelected } }
    
    func addItem(_ product: Product, spec: ProductSpec?, quantity: Int = 1) {
        let specId = spec?.id
        if let index = items.firstIndex(where: { $0.productId == product.id && $0.specId == specId }) {
            items[index].quantity += quantity
        } else {
            let newItem = CartItem(
                id: UUID().uuidString,
                userId: "mock_user",
                productId: product.id,
                productName: product.name,
                productImage: product.mainImage,
                price: spec?.price ?? product.price,
                originalPrice: product.originalPrice,
                specId: spec?.id,
                specName: spec?.name,
                specValue: spec?.value,
                quantity: quantity
            )
            items.append(newItem)
        }
    }
    
    func removeItem(at index: Int) { items.remove(at: index) }
    
    func removeSelectedItems() { items.removeAll { $0.isSelected } }
    
    func updateQuantity(at index: Int, quantity: Int) {
        guard index < items.count else { return }
        if quantity <= 0 { removeItem(at: index) }
        else { items[index].quantity = quantity }
    }
    
    func toggleSelection(at index: Int) {
        guard index < items.count else { return }
        items[index].isSelected.toggle()
    }
    
    func toggleAllSelection() {
        let newStatus = !isAllSelected
        for i in 0..<items.count { items[i].isSelected = newStatus }
    }
    
    private func loadMockData() {
        items = [
            CartItem(
                id: "1",
                userId: "mock_user",
                productId: "p1",
                productName: "Apple iPhone 15 Pro Max 256GB 原色钛金属",
                productImage: "https://picsum.photos/seed/iphone/200/200",
                price: 9999,
                originalPrice: 11999,
                specId: "s1",
                specName: "颜色",
                specValue: "原色钛金属",
                quantity: 1
            ),
            CartItem(
                id: "2",
                userId: "mock_user",
                productId: "p2",
                productName: "Apple AirPods Pro (第二代)",
                productImage: "https://picsum.photos/seed/airpods/200/200",
                price: 1799,
                originalPrice: 1999,
                specId: nil,
                specName: nil,
                specValue: nil,
                quantity: 2
            )
        ]
    }
}

// MARK: - Mock Data Service
class MockDataService {
    static let shared = MockDataService()
    private init() {}
    
    let products: [Product] = [
        Product(
            id: "p1",
            categoryId: "c1",
            name: "Apple iPhone 15 Pro Max 256GB 原色钛金属 支持移动联通电信5G",
            description: "iPhone 15 Pro Max. 钛金属设计，A17 Pro 芯片。",
            price: 9999,
            originalPrice: 11999,
            stock: 100,
            sales: 8542,
            mainImage: "https://picsum.photos/seed/iphone15/400/400",
            images: [
                "https://picsum.photos/seed/iphone15/400/400",
                "https://picsum.photos/seed/iphone15-2/400/400"
            ],
            specs: [
                ProductSpec(id: "s1", productId: "p1", name: "颜色", value: "原色钛金属", stock: 50, price: 9999),
                ProductSpec(id: "s2", productId: "p1", name: "颜色", value: "蓝色钛金属", stock: 30, price: 9999)
            ],
            status: .active,
            createdAt: "2024-01-01",
            shopName: "Apple官方旗舰店"
        ),
        Product(
            id: "p2",
            categoryId: "c1",
            name: "Apple AirPods Pro (第二代) 配MagSafe充电盒",
            description: "AirPods Pro 具备主动降噪功能。",
            price: 1799,
            originalPrice: 1999,
            stock: 200,
            sales: 15632,
            mainImage: "https://picsum.photos/seed/airpods/400/400",
            images: ["https://picsum.photos/seed/airpods/400/400"],
            specs: [],
            status: .active,
            createdAt: "2024-01-15",
            shopName: "Apple官方旗舰店"
        ),
        Product(
            id: "p3",
            categoryId: "c2",
            name: "华为Mate 60 Pro 12GB+512GB 雅丹黑",
            description: "HUAWEI Mate 60 Pro，卫星通话。",
            price: 7999,
            originalPrice: 8999,
            stock: 50,
            sales: 3256,
            mainImage: "https://picsum.photos/seed/mate60/400/400",
            images: ["https://picsum.photos/seed/mate60/400/400"],
            specs: [
                ProductSpec(id: "s5", productId: "p3", name: "颜色", value: "雅丹黑", stock: 20, price: 7999),
                ProductSpec(id: "s6", productId: "p3", name: "颜色", value: "白沙银", stock: 15, price: 7999)
            ],
            status: .active,
            createdAt: "2024-02-01",
            shopName: "华为官方旗舰店"
        ),
        Product(
            id: "p4",
            categoryId: "c3",
            name: "小米14 Ultra 16GB+512GB 白色 徕卡光学",
            description: "小米14 Ultra，徕卡光学镜头。",
            price: 6499,
            originalPrice: 6999,
            stock: 80,
            sales: 8921,
            mainImage: "https://picsum.photos/seed/mi14/400/400",
            images: ["https://picsum.photos/seed/mi14/400/400"],
            specs: [],
            status: .active,
            createdAt: "2024-02-15",
            shopName: "小米官方旗舰店"
        ),
        Product(
            id: "p5",
            categoryId: "c4",
            name: "MacBook Pro 14英寸 M3 Pro芯片 18GB+512GB",
            description: "MacBook Pro M3 Pro 芯片。",
            price: 16999,
            originalPrice: 17999,
            stock: 30,
            sales: 2456,
            mainImage: "https://picsum.photos/seed/macbook/400/400",
            images: ["https://picsum.photos/seed/macbook/400/400"],
            specs: [],
            status: .active,
            createdAt: "2024-03-01",
            shopName: "Apple官方旗舰店"
        ),
        Product(
            id: "p6",
            categoryId: "c5",
            name: "戴森Dyson V15吸尘器 智能无绳吸尘器",
            description: "戴森V15 Detect智能无绳吸尘器。",
            price: 5990,
            originalPrice: 6990,
            stock: 60,
            sales: 1876,
            mainImage: "https://picsum.photos/seed/dyson/400/400",
            images: ["https://picsum.photos/seed/dyson/400/400"],
            specs: [],
            status: .active,
            createdAt: "2024-02-20",
            shopName: "戴森官方旗舰店"
        )
    ]
    
    let categories: [Category] = [
        Category(id: "c1", name: "手机", parentId: nil, icon: "iphone", sortOrder: 1),
        Category(id: "c2", name: "电脑", parentId: nil, icon: "laptopcomputer", sortOrder: 2),
        Category(id: "c3", name: "数码", parentId: nil, icon: "headphones", sortOrder: 3),
        Category(id: "c4", name: "家电", parentId: nil, icon: "house", sortOrder: 4),
        Category(id: "c5", name: "服饰", parentId: nil, icon: "tshirt", sortOrder: 5),
        Category(id: "c6", name: "美妆", parentId: nil, icon: "sparkles", sortOrder: 6),
        Category(id: "c7", name: "食品", parentId: nil, icon: "fork.knife", sortOrder: 7),
        Category(id: "c8", name: "家居", parentId: nil, icon: "bed.double", sortOrder: 8)
    ]
    
    let banners: [Banner] = [
        Banner(id: "b1", image: "https://picsum.photos/seed/banner1/800/300", link: "p1", title: "iPhone 15 Pro 新品上市"),
        Banner(id: "b2", image: "https://picsum.photos/seed/banner2/800/300", link: "p3", title: "华为Mate 60 卫星通话"),
        Banner(id: "b3", image: "https://picsum.photos/seed/banner3/800/300", link: "p5", title: "MacBook Pro M3")
    ]
    
    let reviews: [Review] = [
        Review(
            id: "r1",
            userId: "u1",
            userName: "数码爱好者",
            productId: "p1",
            orderId: "o1",
            rating: 5,
            content: "iPhone 15 Pro Max 用了一周，体验非常棒！",
            images: [],
            createdAt: "2024-03-10",
            avatar: nil,
            badge: "达人认证",
            specs: "原色钛金属",
            likes: 128,
            reply: nil,
            time: "3天前"
        )
    ]
    
    func getProduct(by id: String) -> Product? { products.first { $0.id == id } }
    
    func getReviews(for productId: String) -> [Review] { reviews.filter { $0.productId == productId } }
    
    func searchProducts(query: String) -> [Product] {
        products.filter { $0.name.localizedCaseInsensitiveContains(query) }
    }
}
