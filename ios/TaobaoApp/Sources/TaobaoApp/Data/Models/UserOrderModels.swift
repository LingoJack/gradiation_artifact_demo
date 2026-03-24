import Foundation

// MARK: - Order Model
struct Order: Identifiable, Codable, Hashable {
    let id: String
    let orderNo: String
    let userId: String
    let status: OrderStatus
    let items: [OrderItem]
    let address: Address?
    let subtotal: Double
    let discount: Double
    let shippingFee: Double
    let payAmount: Double
    let paymentMethod: String?
    let paymentTime: String?
    let shipTime: String?
    let receiveTime: String?
    let remark: String
    let createdAt: String
    let updatedAt: String
    
    var statusText: String { status.displayText }
    var itemCount: Int { items.reduce(0) { $0 + $1.quantity } }
}

struct OrderItem: Identifiable, Codable, Hashable {
    let id: String
    let orderId: String
    let productId: String
    let productName: String
    let productImage: String
    let specId: String?
    let specName: String?
    let price: Double
    let quantity: Int
    
    var subtotal: Double { price * Double(quantity) }
    
    init(from cartItem: CartItem) {
        self.id = UUID().uuidString
        self.orderId = ""
        self.productId = cartItem.productId
        self.productName = cartItem.productName
        self.productImage = cartItem.productImage
        self.specId = cartItem.specId
        self.specName = cartItem.specName
        self.price = cartItem.price
        self.quantity = cartItem.quantity
    }
    
    init(id: String, orderId: String, productId: String, productName: String, productImage: String, specId: String?, specName: String?, price: Double, quantity: Int) {
        self.id = id
        self.orderId = orderId
        self.productId = productId
        self.productName = productName
        self.productImage = productImage
        self.specId = specId
        self.specName = specName
        self.price = price
        self.quantity = quantity
    }
}

enum OrderStatus: String, Codable, CaseIterable {
    case pending = "pending"
    case paid = "paid"
    case shipped = "shipped"
    case completed = "completed"
    case cancelled = "cancelled"
    
    var displayText: String {
        switch self {
        case .pending: return "待付款"
        case .paid: return "待发货"
        case .shipped: return "待收货"
        case .completed: return "已完成"
        case .cancelled: return "已取消"
        }
    }
    
    var icon: String {
        switch self {
        case .pending: return "creditcard"
        case .paid: return "shippingbox"
        case .shipped: return "truck.box"
        case .completed: return "checkmark.circle"
        case .cancelled: return "xmark.circle"
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
