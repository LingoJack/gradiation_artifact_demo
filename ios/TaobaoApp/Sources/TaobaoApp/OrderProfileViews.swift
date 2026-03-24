import SwiftUI

// MARK: - Order List View
struct OrderListView: View {
    let statusFilter: OrderStatus?
    
    @State private var orders: [Order] = []
    
    init(statusFilter: OrderStatus? = nil) {
        self.statusFilter = statusFilter
    }
    
    var body: some View {
        NavigationStack {
            Group {
                if filteredOrders.isEmpty {
                    EmptyStateView(
                        icon: "shippingbox",
                        title: "暂无订单",
                        subtitle: "快去选购心仪的商品吧"
                    )
                } else {
                    ScrollView {
                        LazyVStack(spacing: .tbSpacing12) {
                            ForEach(filteredOrders) { order in
                                OrderCardView(order: order)
                            }
                        }
                        .padding(.tbSpacing12)
                    }
                    .background(Color.tbBackground)
                }
            }
            .navigationTitle("我的订单")
            .navigationBarTitleDisplayMode(.inline)
            .onAppear {
                loadOrders()
            }
        }
    }
    
    private var filteredOrders: [Order] {
        guard let filter = statusFilter else { return orders }
        return orders.filter { $0.status == filter }
    }
    
    private func loadOrders() {
        // Mock orders
        orders = [
            Order(
                id: "o1",
                userId: "u1",
                status: .pending,
                items: [
                    OrderItem(
                        id: "oi1",
                        orderId: "o1",
                        productId: "p1",
                        productName: "Apple iPhone 15 Pro Max 256GB",
                        productImage: "https://picsum.photos/seed/iphone15/200/200",
                        specId: "s1",
                        specName: "颜色",
                        specValue: "原色钛金属",
                        price: 9999,
                        quantity: 1
                    )
                ],
                totalPrice: 9999,
                discount: 0,
                shippingFee: 0,
                paymentAmount: 9999,
                address: Address(
                    id: "a1",
                    userId: "u1",
                    receiver: "张三",
                    phone: "13800138000",
                    province: "浙江省",
                    city: "杭州市",
                    district: "西湖区",
                    detail: "某某街道某某小区1号楼",
                    isDefault: true
                ),
                paymentMethod: .alipay,
                createdAt: "2024-03-15 10:30:00",
                paidAt: nil,
                shippedAt: nil,
                completedAt: nil,
                canceledAt: nil,
                trackingNumber: nil
            ),
            Order(
                id: "o2",
                userId: "u1",
                status: .shipped,
                items: [
                    OrderItem(
                        id: "oi2",
                        orderId: "o2",
                        productId: "p2",
                        productName: "Apple AirPods Pro (第二代)",
                        productImage: "https://picsum.photos/seed/airpods/200/200",
                        specId: nil,
                        specName: nil,
                        specValue: nil,
                        price: 1799,
                        quantity: 1
                    )
                ],
                totalPrice: 1799,
                discount: 200,
                shippingFee: 0,
                paymentAmount: 1599,
                address: Address(
                    id: "a1",
                    userId: "u1",
                    receiver: "张三",
                    phone: "13800138000",
                    province: "浙江省",
                    city: "杭州市",
                    district: "西湖区",
                    detail: "某某街道某某小区1号楼",
                    isDefault: true
                ),
                paymentMethod: .alipay,
                createdAt: "2024-03-10 14:20:00",
                paidAt: "2024-03-10 14:22:00",
                shippedAt: "2024-03-11 09:00:00",
                completedAt: nil,
                canceledAt: nil,
                trackingNumber: "SF1234567890"
            )
        ]
    }
}

// MARK: - Order Card View
struct OrderCardView: View {
    let order: Order
    
    var body: some View {
        VStack(alignment: .leading, spacing: .tbSpacing12) {
            // Header
            HStack {
                Image(systemName: "shop")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                
                Text("官方旗舰店")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbTextPrimary)
                
                Spacer()
                
                Text(order.statusText)
                    .font(.tbBodyBold)
                    .foregroundColor(statusColor)
            }
            
            // Items
            ForEach(order.items) { item in
                HStack(spacing: .tbSpacing12) {
                    AsyncImage(url: URL(string: item.productImage)) { phase in
                        switch phase {
                        case .success(let image):
                            image
                                .resizable()
                                .aspectRatio(contentMode: .fill)
                        default:
                            Color.tbDivider
                        }
                    }
                    .frame(width: 80, height: 80)
                    .cornerRadius(.tbRadius8)
                    
                    VStack(alignment: .leading, spacing: .tbSpacing4) {
                        Text(item.productName)
                            .font(.tbBody)
                            .foregroundColor(.tbTextPrimary)
                            .lineLimit(2)
                        
                        if let specValue = item.specValue {
                            Text(specValue)
                                .font(.tbCaption)
                                .foregroundColor(.tbTextTertiary)
                        }
                        
                        Spacer()
                        
                        HStack {
                            Text("¥\(String(format: "%.0f", item.price))")
                                .font(.tbBody)
                                .foregroundColor(.tbTextPrimary)
                            
                            Spacer()
                            
                            Text("x\(item.quantity)")
                                .font(.tbCaption)
                                .foregroundColor(.tbTextTertiary)
                        }
                    }
                }
            }
            
            // Total
            HStack {
                Spacer()
                Text("共\(order.itemCount)件商品 实付款: ")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                
                Text("¥\(String(format: "%.2f", order.paymentAmount))")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbPrice)
            }
            
            // Actions
            HStack {
                Spacer()
                
                switch order.status {
                case .pending:
                    actionButton(title: "取消订单") { }
                    primaryActionButton(title: "去支付") { }
                case .paid:
                    actionButton(title: "联系客服") { }
                case .shipped:
                    actionButton(title: "查看物流") { }
                    primaryActionButton(title: "确认收货") { }
                case .completed:
                    actionButton(title: "再次购买") { }
                    actionButton(title: "评价") { }
                case .canceled:
                    actionButton(title: "删除订单") { }
                }
            }
        }
        .padding(.tbSpacing12)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
    
    private var statusColor: Color {
        switch order.status {
        case .pending: return .tbPrice
        case .paid: return .tbOrange
        case .shipped: return .blue
        case .completed: return .green
        case .canceled: return .tbTextTertiary
        }
    }
    
    private func actionButton(title: String, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            Text(title)
                .font(.tbCaption)
                .foregroundColor(.tbTextSecondary)
                .padding(.horizontal, .tbSpacing12)
                .padding(.vertical, .tbSpacing8)
                .background(Color.tbBackground)
                .cornerRadius(.tbRadius4)
        }
    }
    
    private func primaryActionButton(title: String, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            Text(title)
                .font(.tbCaption)
                .foregroundColor(.white)
                .padding(.horizontal, .tbSpacing12)
                .padding(.vertical, .tbSpacing8)
                .background(Color.tbOrange)
                .cornerRadius(.tbRadius4)
        }
    }
}

// MARK: - Profile View
struct ProfileView: View {
    @State private var user: User?
    
    private var mockUser: User {
        User(
            id: "u1",
            username: "mock_user",
            email: "user@example.com",
            phone: "138****8000",
            avatar: nil,
            nickname: "数码爱好者",
            gender: .male,
            birthday: "1990-01-01",
            bio: "热爱科技，享受生活",
            createdAt: "2023-01-01"
        )
    }
    
    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: .tbSpacing12) {
                    // User Header
                    userHeader
                        .padding(.top, .tbSpacing20)
                    
                    // Order Section
                    orderSection
                    
                    // Services Section
                    servicesSection
                    
                    // Settings Section
                    settingsSection
                }
                .padding(.tbSpacing12)
            }
            .background(Color.tbBackground)
            .navigationTitle("我的")
            .navigationBarTitleDisplayMode(.inline)
        }
    }
    
    // MARK: - User Header
    private var userHeader: some View {
        HStack(spacing: .tbSpacing16) {
            Image(systemName: "person.circle.fill")
                .font(.system(size: 60))
                .foregroundColor(.tbOrange)
            
            VStack(alignment: .leading, spacing: .tbSpacing4) {
                Text(mockUser.displayName)
                    .font(.tbTitle2)
                    .foregroundColor(.tbTextPrimary)
                
                Text("ID: \(mockUser.id)")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextTertiary)
                
                HStack(spacing: .tbSpacing8) {
                    Text("会员等级: VIP")
                        .font(.tbCaption)
                        .foregroundColor(.tbOrange)
                    
                    Image(systemName: "crown.fill")
                        .font(.tbCaption2)
                        .foregroundColor(.tbOrange)
                }
            }
            
            Spacer()
            
            Image(systemName: "chevron.right")
                .font(.tbBody)
                .foregroundColor(.tbTextTertiary)
        }
        .padding(.tbSpacing16)
        .background(Color.white)
        .cornerRadius(.tbRadius12)
    }
    
    // MARK: - Order Section
    private var orderSection: some View {
        VStack(spacing: 0) {
            HStack {
                Text("我的订单")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbTextPrimary)
                
                Spacer()
                
                Button(action: {}) {
                    HStack(spacing: 4) {
                        Text("全部订单")
                            .font(.tbCaption)
                        Image(systemName: "chevron.right")
                            .font(.tbCaption2)
                    }
                    .foregroundColor(.tbTextTertiary)
                }
            }
            .padding(.tbSpacing12)
            
            Divider()
                .padding(.horizontal, .tbSpacing12)
            
            HStack(spacing: 0) {
                orderStatusItem(icon: "creditcard", title: "待付款", count: 1)
                orderStatusItem(icon: "shippingbox", title: "待发货", count: 0)
                orderStatusItem(icon: "truck.box", title: "待收货", count: 1)
                orderStatusItem(icon: "star", title: "待评价", count: 0)
                orderStatusItem(icon: "arrow.uturn.backward", title: "退换/售后", count: 0)
            }
            .padding(.vertical, .tbSpacing16)
        }
        .background(Color.white)
        .cornerRadius(.tbRadius12)
    }
    
    private func orderStatusItem(icon: String, title: String, count: Int) -> some View {
        VStack(spacing: .tbSpacing8) {
            ZStack(alignment: .topTrailing) {
                Image(systemName: icon)
                    .font(.title2)
                    .foregroundColor(.tbTextSecondary)
                
                if count > 0 {
                    Text("\(count)")
                        .font(.system(size: 10, weight: .bold))
                        .foregroundColor(.white)
                        .padding(4)
                        .background(Color.tbPrice)
                        .clipShape(Circle())
                        .offset(x: 8, y: -8)
                }
            }
            .frame(height: 28)
            
            Text(title)
                .font(.tbCaption)
                .foregroundColor(.tbTextSecondary)
        }
        .frame(maxWidth: .infinity)
    }
    
    // MARK: - Services Section
    private var servicesSection: some View {
        VStack(alignment: .leading, spacing: 0) {
            Text("我的服务")
                .font(.tbBodyBold)
                .foregroundColor(.tbTextPrimary)
                .padding(.tbSpacing12)
            
            Divider()
                .padding(.horizontal, .tbSpacing12)
            
            LazyVGrid(columns: Array(repeating: GridItem(.flexible()), count: 4), spacing: .tbSpacing16) {
                serviceItem(icon: "heart", title: "我的收藏")
                serviceItem(icon: "clock", title: "浏览记录")
                serviceItem(icon: "map", title: "收货地址")
                serviceItem(icon: "ticket", title: "优惠券")
                serviceItem(icon: "gift", title: "我的积分")
                serviceItem(icon: "questionmark.circle", title: "帮助中心")
                serviceItem(icon: "headphones", title: "在线客服")
                serviceItem(icon: "gear", title: "设置")
            }
            .padding(.tbSpacing12)
        }
        .background(Color.white)
        .cornerRadius(.tbRadius12)
    }
    
    private func serviceItem(icon: String, title: String) -> some View {
        VStack(spacing: .tbSpacing8) {
            Image(systemName: icon)
                .font(.title3)
                .foregroundColor(.tbOrange)
            
            Text(title)
                .font(.tbCaption)
                .foregroundColor(.tbTextSecondary)
        }
        .frame(maxWidth: .infinity)
    }
    
    // MARK: - Settings Section
    private var settingsSection: some View {
        VStack(spacing: 0) {
            settingsItem(icon: "person", title: "个人资料")
            Divider().padding(.leading, 44)
            settingsItem(icon: "lock", title: "账号安全")
            Divider().padding(.leading, 44)
            settingsItem(icon: "bell", title: "消息通知")
            Divider().padding(.leading, 44)
            settingsItem(icon: "info.circle", title: "关于我们")
        }
        .background(Color.white)
        .cornerRadius(.tbRadius12)
    }
    
    private func settingsItem(icon: String, title: String) -> some View {
        Button(action: {}) {
            HStack(spacing: .tbSpacing12) {
                Image(systemName: icon)
                    .font(.tbBody)
                    .foregroundColor(.tbTextSecondary)
                    .frame(width: 20)
                
                Text(title)
                    .font(.tbBody)
                    .foregroundColor(.tbTextPrimary)
                
                Spacer()
                
                Image(systemName: "chevron.right")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextTertiary)
            }
            .padding(.tbSpacing12)
        }
    }
}

// MARK: - Preview
#Preview {
    ProfileView()
}
