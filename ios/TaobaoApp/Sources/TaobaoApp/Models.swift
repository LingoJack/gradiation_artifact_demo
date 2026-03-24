import Foundation

// MARK: - Product Model
struct Product: Identifiable, Codable, Hashable {
    let id: String
    let categoryId: String
    let name: String
    let description: String
    let price: Double
    let originalPrice: Double?
    let stock: Int
    let sales: Int
    let mainImage: String
    let images: [String]
    let specs: [ProductSpec]
    let status: ProductStatus
    let createdAt: String
    let shopName: String?
    
    var discount: String? {
        guard let original = originalPrice, original > price else { return nil }
        let ratio = Int((price / original) * 10)
        return "\(ratio)折"
    }
    
    var savedAmount: Double? {
        guard let original = originalPrice else { return nil }
        return original - price
    }
}

struct ProductSpec: Identifiable, Codable, Hashable {
    let id: String
    let productId: String
    let name: String
    let value: String
    let stock: Int
    let price: Double
}

enum ProductStatus: String, Codable, Hashable {
    case active
    case inactive
}

struct Category: Identifiable, Codable {
    let id: String
    let name: String
    let parentId: String?
    let icon: String?
    let sortOrder: Int
}

struct Review: Identifiable, Codable {
    let id: String
    let userId: String
    let userName: String
    let productId: String
    let orderId: String
    let rating: Int
    let content: String
    let images: [String]
    let createdAt: String
    let avatar: String?
    let badge: String?
    let specs: String?
    let likes: Int
    let reply: String?
    let time: String?
}

struct Banner: Identifiable {
    let id: String
    let image: String
    let link: String
    let title: String
}
