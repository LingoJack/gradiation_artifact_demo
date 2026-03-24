import Foundation

// MARK: - User Model
struct User: Identifiable, Codable {
    let id: String
    let username: String
    let email: String
    let phone: String
    let avatar: String?
    let nickname: String?
    let gender: Gender?
    let birthday: String?
    let bio: String?
    let createdAt: String
    
    var displayName: String { nickname ?? username }
}

enum Gender: String, Codable {
    case male = "male"
    case female = "female"
    case other = "other"
    
    var displayText: String {
        switch self {
        case .male: return "男"
        case .female: return "女"
        case .other: return "其他"
        }
    }
}

struct Address: Identifiable, Codable {
    let id: String
    let userId: String
    let receiver: String
    let phone: String
    let province: String
    let city: String
    let district: String
    let detail: String
    let isDefault: Bool
    
    var fullAddress: String { "\(province)\(city)\(district)\(detail)" }
}

// MARK: - Order Model
struct Order: Identifiable, Codable {
    let id: String
    let userId: String
    let status: OrderStatus
    let items: [OrderItem]
    let totalPrice: Double
    let discount: Double
    let shippingFee: Double
    let paymentAmount: Double
    let address: Address?
    let paymentMethod: PaymentMethod?
    let createdAt: String
    let paidAt: String?
    let shippedAt: String?
    let completedAt: String?
    let canceledAt: String?
    let trackingNumber: String?
    
    var statusText: String { status.displayText }
    var itemCount: Int { items.reduce(0) { $0 + $1.quantity } }
}

struct OrderItem: Identifiable, Codable {
    let id: String
    let orderId: String
    let productId: String
    let productName: String
    let productImage: String
    let specId: String?
    let specName: String?
    let specValue: String?
    let price: Double
    let quantity: Int
    
    var subtotal: Double { price * Double(quantity) }
}

enum OrderStatus: String, Codable {
    case pending = "pending"
    case paid = "paid"
    case shipped = "shipped"
    case completed = "completed"
    case canceled = "canceled"
    
    var displayText: String {
        switch self {
        case .pending: return "待付款"
        case .paid: return "待发货"
        case .shipped: return "待收货"
        case .completed: return "已完成"
        case .canceled: return "已取消"
        }
    }
    
    var icon: String {
        switch self {
        case .pending: return "creditcard"
        case .paid: return "shippingbox"
        case .shipped: return "truck.box"
        case .completed: return "checkmark.circle"
        case .canceled: return "xmark.circle"
        }
    }
}

enum PaymentMethod: String, Codable {
    case alipay = "alipay"
    case wechat = "wechat"
    case bankCard = "bank_card"
    
    var displayText: String {
        switch self {
        case .alipay: return "支付宝"
        case .wechat: return "微信支付"
        case .bankCard: return "银行卡"
        }
    }
}

// MARK: - Cart Item
struct CartItem: Identifiable, Codable {
    let id: String
    let userId: String
    let productId: String
    let productName: String
    let productImage: String
    let price: Double
    let originalPrice: Double?
    let specId: String?
    let specName: String?
    let specValue: String?
    var quantity: Int
    var isSelected: Bool = true
    
    var subtotal: Double { price * Double(quantity) }
    
    var discount: String? {
        guard let original = originalPrice, original > price else { return nil }
        let ratio = Int((price / original) * 10)
        return "\(ratio)折"
    }
}
