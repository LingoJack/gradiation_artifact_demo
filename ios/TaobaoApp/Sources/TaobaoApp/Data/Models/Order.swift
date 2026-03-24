import Foundation

// MARK: - Coupon Model
struct Coupon: Identifiable, Codable {
    let id: String
    let name: String
    let type: CouponType
    let value: Double
    let minAmount: Double
    let startTime: String
    let endTime: String
    let status: CouponStatus
    let description: String?
    let scope: CouponScope?
    let productId: String?
    let categoryId: String?
    
    var isValid: Bool { status == .available }
    
    var discountText: String {
        switch type {
        case .fixed: return "¥\(Int(value))"
        case .percent: return "\(Int(value * 10))折"
        }
    }
    
    var conditionText: String {
        minAmount > 0 ? "满\(Int(minAmount))可用" : "无门槛"
    }
}

enum CouponType: String, Codable {
    case fixed = "fixed"
    case percent = "percent"
}

enum CouponStatus: String, Codable {
    case available = "available"
    case used = "used"
    case expired = "expired"
}

enum CouponScope: String, Codable {
    case all = "all"
    case product = "product"
    case category = "category"
}

// MARK: - Favorite Model
struct Favorite: Identifiable, Codable {
    let id: String
    let userId: String
    let productId: String
    let product: Product
    let createdAt: String
}

// MARK: - Address Model
struct Address: Identifiable, Codable, Hashable {
    let id: String
    let userId: String
    let name: String
    let phone: String
    let province: String
    let city: String
    let district: String
    let detail: String
    let isDefault: Bool
    
    var fullAddress: String {
        "\(province)\(city)\(district)\(detail)"
    }
}

// MARK: - Checkout Model
struct CheckoutItem: Identifiable {
    let id: String
    let cartItem: CartItem
    var isSelected: Bool = true
}

struct CheckoutOrder {
    var items: [CheckoutItem]
    var address: Address?
    var coupon: Coupon?
    var remark: String = ""
    
    var subtotal: Double { items.filter { $0.isSelected }.reduce(0) { $0 + $1.cartItem.subtotal } }
    var discount: Double { coupon.map { $0.value } ?? 0 }
    var total: Double { max(0, subtotal - discount) }
}
