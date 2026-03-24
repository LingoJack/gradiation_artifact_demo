import Foundation
import Combine

// MARK: - Order Store
class OrderStore: ObservableObject {
    static let shared = OrderStore()
    
    @Published var orders: [Order] = []
    
    private init() {
        loadMockData()
    }
    
    var pendingOrders: [Order] { orders.filter { $0.status == .pending } }
    var paidOrders: [Order] { orders.filter { $0.status == .paid } }
    var shippedOrders: [Order] { orders.filter { $0.status == .shipped } }
    var completedOrders: [Order] { orders.filter { $0.status == .completed } }
    
    func ordersByStatus(_ status: OrderStatus?) -> [Order] {
        guard let status = status else { return orders }
        return orders.filter { $0.status == status }
    }
    
    func createOrder(items: [CartItem], address: Address?, coupon: Coupon?, remark: String) -> Order {
        let orderItems = items.map { OrderItem(from: $0) }
        let subtotal = items.reduce(0) { $0 + $1.subtotal }
        let discount = coupon?.value ?? 0
        let order = Order(
            id: UUID().uuidString,
            orderNo: generateOrderNo(),
            userId: UserStore.shared.user?.id ?? "guest",
            status: .pending,
            items: orderItems,
            address: address,
            subtotal: subtotal,
            discount: discount,
            shippingFee: subtotal >= 99 ? 0 : 10,
            payAmount: subtotal - discount + (subtotal >= 99 ? 0 : 10),
            paymentMethod: nil,
            paymentTime: nil,
            shipTime: nil,
            receiveTime: nil,
            remark: remark,
            createdAt: formatDate(Date()),
            updatedAt: formatDate(Date())
        )
        orders.insert(order, at: 0)
        return order
    }
    
    func payOrder(_ orderId: String) {
        guard let index = orders.firstIndex(where: { $0.id == orderId }) else { return }
        let order = orders[index]
        orders[index] = Order(
            id: order.id,
            orderNo: order.orderNo,
            userId: order.userId,
            status: .paid,
            items: order.items,
            address: order.address,
            subtotal: order.subtotal,
            discount: order.discount,
            shippingFee: order.shippingFee,
            payAmount: order.payAmount,
            paymentMethod: "支付宝",
            paymentTime: formatDate(Date()),
            shipTime: nil,
            receiveTime: nil,
            remark: order.remark,
            createdAt: order.createdAt,
            updatedAt: formatDate(Date())
        )
    }
    
    func shipOrder(_ orderId: String) {
        guard let index = orders.firstIndex(where: { $0.id == orderId }) else { return }
        let order = orders[index]
        orders[index] = Order(
            id: order.id,
            orderNo: order.orderNo,
            userId: order.userId,
            status: .shipped,
            items: order.items,
            address: order.address,
            subtotal: order.subtotal,
            discount: order.discount,
            shippingFee: order.shippingFee,
            payAmount: order.payAmount,
            paymentMethod: order.paymentMethod,
            paymentTime: order.paymentTime,
            shipTime: formatDate(Date()),
            receiveTime: nil,
            remark: order.remark,
            createdAt: order.createdAt,
            updatedAt: formatDate(Date())
        )
    }
    
    func confirmReceive(_ orderId: String) {
        guard let index = orders.firstIndex(where: { $0.id == orderId }) else { return }
        let order = orders[index]
        orders[index] = Order(
            id: order.id,
            orderNo: order.orderNo,
            userId: order.userId,
            status: .completed,
            items: order.items,
            address: order.address,
            subtotal: order.subtotal,
            discount: order.discount,
            shippingFee: order.shippingFee,
            payAmount: order.payAmount,
            paymentMethod: order.paymentMethod,
            paymentTime: order.paymentTime,
            shipTime: order.shipTime,
            receiveTime: formatDate(Date()),
            remark: order.remark,
            createdAt: order.createdAt,
            updatedAt: formatDate(Date())
        )
    }
    
    func cancelOrder(_ orderId: String) {
        guard let index = orders.firstIndex(where: { $0.id == orderId }) else { return }
        let order = orders[index]
        orders[index] = Order(
            id: order.id,
            orderNo: order.orderNo,
            userId: order.userId,
            status: .cancelled,
            items: order.items,
            address: order.address,
            subtotal: order.subtotal,
            discount: order.discount,
            shippingFee: order.shippingFee,
            payAmount: order.payAmount,
            paymentMethod: order.paymentMethod,
            paymentTime: order.paymentTime,
            shipTime: order.shipTime,
            receiveTime: nil,
            remark: order.remark,
            createdAt: order.createdAt,
            updatedAt: formatDate(Date())
        )
    }
    
    private func generateOrderNo() -> String {
        let formatter = DateFormatter()
        formatter.dateFormat = "yyyyMMddHHmmss"
        let timestamp = formatter.string(from: Date())
        let random = String(format: "%04d", Int.random(in: 0...9999))
        return "\(timestamp)\(random)"
    }
    
    private func formatDate(_ date: Date) -> String {
        let formatter = DateFormatter()
        formatter.dateFormat = "yyyy-MM-dd HH:mm:ss"
        return formatter.string(from: date)
    }
    
    private func loadMockData() {
        orders = [
            Order(
                id: "o1",
                orderNo: "202403151234567890",
                userId: "u1",
                status: .pending,
                items: [
                    OrderItem(id: "oi1", orderId: "o1", productId: "p1", productName: "Apple iPhone 15 Pro Max 256GB", productImage: "https://picsum.photos/seed/iphone15/200/200", specId: "s1", specName: "原色钛金属", price: 9999, quantity: 1)
                ],
                address: Address(id: "a1", userId: "u1", name: "张三", phone: "13800138000", province: "浙江省", city: "杭州市", district: "西湖区", detail: "文三路 123 号", isDefault: true),
                subtotal: 9999,
                discount: 0,
                shippingFee: 0,
                payAmount: 9999,
                paymentMethod: nil,
                paymentTime: nil,
                shipTime: nil,
                receiveTime: nil,
                remark: "请尽快发货",
                createdAt: "2024-03-15 10:30:00",
                updatedAt: "2024-03-15 10:30:00"
            ),
            Order(
                id: "o2",
                orderNo: "202403141234567891",
                userId: "u1",
                status: .shipped,
                items: [
                    OrderItem(id: "oi2", orderId: "o2", productId: "p2", productName: "Apple AirPods Pro", productImage: "https://picsum.photos/seed/airpods/200/200", specId: nil, specName: nil, price: 1799, quantity: 2)
                ],
                address: Address(id: "a1", userId: "u1", name: "张三", phone: "13800138000", province: "浙江省", city: "杭州市", district: "西湖区", detail: "文三路 123 号", isDefault: true),
                subtotal: 3598,
                discount: 50,
                shippingFee: 0,
                payAmount: 3548,
                paymentMethod: "支付宝",
                paymentTime: "2024-03-14 09:00:00",
                shipTime: "2024-03-14 15:00:00",
                receiveTime: nil,
                remark: "",
                createdAt: "2024-03-14 08:30:00",
                updatedAt: "2024-03-14 15:00:00"
            ),
            Order(
                id: "o3",
                orderNo: "202403101234567892",
                userId: "u1",
                status: .completed,
                items: [
                    OrderItem(id: "oi3", orderId: "o3", productId: "p3", productName: "华为Mate 60 Pro", productImage: "https://picsum.photos/seed/mate60/200/200", specId: "s5", specName: "雅丹黑", price: 7999, quantity: 1)
                ],
                address: Address(id: "a1", userId: "u1", name: "张三", phone: "13800138000", province: "浙江省", city: "杭州市", district: "西湖区", detail: "文三路 123 号", isDefault: true),
                subtotal: 7999,
                discount: 100,
                shippingFee: 0,
                payAmount: 7899,
                paymentMethod: "支付宝",
                paymentTime: "2024-03-10 10:00:00",
                shipTime: "2024-03-10 16:00:00",
                receiveTime: "2024-03-12 14:00:00",
                remark: "",
                createdAt: "2024-03-10 09:30:00",
                updatedAt: "2024-03-12 14:00:00"
            )
        ]
    }
}
