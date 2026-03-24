import SwiftUI

// MARK: - Order List View
struct OrderListView: View {
    @StateObject private var orderStore = OrderStore.shared
    var statusFilter: OrderStatus?
    
    @State private var selectedStatus: OrderStatus?
    @State private var showOrderDetail: Order?
    
    init(statusFilter: OrderStatus? = nil) {
        self.statusFilter = statusFilter
        self._selectedStatus = State(initialValue: statusFilter)
    }
    
    var body: some View {
        VStack(spacing: 0) {
            // Status Tabs
            statusTabs
            
            // Order List
            if filteredOrders.isEmpty {
                emptyView
            } else {
                orderList
            }
        }
        .background(Color.tbBackground)
        .navigationTitle("我的订单")
        .navigationBarTitleDisplayMode(.inline)
    }
    
    // MARK: - Status Tabs
    private var statusTabs: some View {
        ScrollView(.horizontal, showsIndicators: false) {
            HStack(spacing: 0) {
                statusTab(title: "全部", status: nil)
                
                ForEach(OrderStatus.allCases, id: \.self) { status in
                    statusTab(title: status.displayText, status: status)
                }
            }
            .padding(.horizontal, .tbSpacing8)
        }
        .background(Color.white)
    }
    
    private func statusTab(title: String, status: OrderStatus?) -> some View {
        Button(action: { selectedStatus = status }) {
            VStack(spacing: .tbSpacing4) {
                Text(title)
                    .font(.tbCaption)
                    .foregroundColor(selectedStatus == status ? .tbOrange : .tbTextSecondary)
                    .padding(.horizontal, .tbSpacing12)
                    .padding(.vertical, .tbSpacing8)
                
                Rectangle()
                    .fill(selectedStatus == status ? Color.tbOrange : Color.clear)
                    .frame(height: 2)
            }
        }
    }
    
    // MARK: - Filtered Orders
    private var filteredOrders: [Order] {
        orderStore.ordersByStatus(selectedStatus)
    }
    
    // MARK: - Empty View
    private var emptyView: some View {
        VStack(spacing: .tbSpacing16) {
            Image(systemName: "tray")
                .font(.system(size: 60))
                .foregroundColor(.tbTextTertiary)
            
            Text("暂无订单")
                .font(.tbBody)
                .foregroundColor(.tbTextSecondary)
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity)
    }
    
    // MARK: - Order List
    private var orderList: some View {
        ScrollView {
            LazyVStack(spacing: .tbSpacing12) {
                ForEach(filteredOrders) { order in
                    NavigationLink(value: order) {
                        OrderCardView(order: order)
                    }
                    .buttonStyle(PlainButtonStyle())
                }
            }
            .padding(.tbSpacing12)
        }
        .navigationDestination(for: Order.self) { order in
            OrderDetailView(order: order)
        }
    }
}

// MARK: - Order Card View
struct OrderCardView: View {
    let order: Order
    @StateObject private var orderStore = OrderStore.shared
    
    var body: some View {
        VStack(alignment: .leading, spacing: .tbSpacing12) {
            // Header
            HStack {
                Image(systemName: "shop")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                
                Text("店铺名称")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextPrimary)
                
                Spacer()
                
                Text(order.statusText)
                    .font(.tbCaption)
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
                    .cornerRadius(.tbRadius4)
                    
                    VStack(alignment: .leading, spacing: .tbSpacing4) {
                        Text(item.productName)
                            .font(.tbCaption)
                            .foregroundColor(.tbTextPrimary)
                            .lineLimit(2)
                        
                        if let specName = item.specName {
                            Text(specName)
                                .font(.tbCaption2)
                                .foregroundColor(.tbTextTertiary)
                        }
                        
                        Spacer()
                        
                        HStack {
                            Text("¥\(String(format: "%.0f", item.price))")
                                .font(.tbCaption)
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
                Text("共\(order.itemCount)件商品 合计: ")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                
                Text("¥\(String(format: "%.0f", order.payAmount))")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbPrice)
            }
            
            // Actions
            HStack {
                Spacer()
                
                switch order.status {
                case .pending:
                    actionButton(title: "取消订单", isPrimary: false) {
                        orderStore.cancelOrder(order.id)
                    }
                    actionButton(title: "去支付", isPrimary: true) {
                        // TODO: Navigate to payment
                    }
                    
                case .shipped:
                    actionButton(title: "确认收货", isPrimary: true) {
                        orderStore.confirmReceive(order.id)
                    }
                    
                case .completed:
                    actionButton(title: "再次购买", isPrimary: true) {}
                    
                case .cancelled:
                    actionButton(title: "删除订单", isPrimary: false) {}
                    
                case .paid:
                    actionButton(title: "催发货", isPrimary: false) {}
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
        case .paid, .shipped: return .tbOrange
        case .completed: return .tbTextSecondary
        case .cancelled: return .tbTextTertiary
        }
    }
    
    private func actionButton(title: String, isPrimary: Bool, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            Text(title)
                .font(.tbCaption)
                .foregroundColor(isPrimary ? .white : .tbTextSecondary)
                .padding(.horizontal, .tbSpacing12)
                .padding(.vertical, .tbSpacing8)
                .background(isPrimary ? Color.tbOrange : Color.tbBackground)
                .cornerRadius(16)
        }
    }
}

// MARK: - Order Detail View
struct OrderDetailView: View {
    let order: Order
    @StateObject private var orderStore = OrderStore.shared
    @Environment(\.dismiss) private var dismiss
    
    var body: some View {
        ScrollView {
            VStack(spacing: .tbSpacing12) {
                // Status
                statusSection
                
                // Address
                if let address = order.address {
                    addressSection(address)
                }
                
                // Items
                itemsSection
                
                // Price Detail
                priceSection
                
                // Order Info
                orderInfoSection
            }
            .padding(.tbSpacing12)
        }
        .background(Color.tbBackground)
        .navigationTitle("订单详情")
        .navigationBarTitleDisplayMode(.inline)
    }
    
    // MARK: - Status Section
    private var statusSection: some View {
        HStack(spacing: .tbSpacing16) {
            Image(systemName: order.status.icon)
                .font(.title)
                .foregroundColor(.tbOrange)
            
            VStack(alignment: .leading, spacing: .tbSpacing4) {
                Text(order.statusText)
                    .font(.tbBodyBold)
                    .foregroundColor(.tbTextPrimary)
                
                if order.status == .pending {
                    Text("请在30分钟内完成支付")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextSecondary)
                } else if order.status == .shipped {
                    Text("包裹正在运输中")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextSecondary)
                }
            }
            
            Spacer()
        }
        .padding(.tbSpacing16)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
    
    // MARK: - Address Section
    private func addressSection(_ address: Address) -> some View {
        HStack(spacing: .tbSpacing12) {
            Image(systemName: "location.fill")
                .font(.tbBody)
                .foregroundColor(.tbOrange)
            
            VStack(alignment: .leading, spacing: .tbSpacing4) {
                HStack(spacing: .tbSpacing8) {
                    Text(address.name)
                        .font(.tbBodyBold)
                        .foregroundColor(.tbTextPrimary)
                    
                    Text(address.phone)
                        .font(.tbBody)
                        .foregroundColor(.tbTextSecondary)
                }
                
                Text(address.fullAddress)
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
            }
            
            Spacer()
        }
        .padding(.tbSpacing16)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
    
    // MARK: - Items Section
    private var itemsSection: some View {
        VStack(alignment: .leading, spacing: .tbSpacing12) {
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
                    .cornerRadius(.tbRadius4)
                    
                    VStack(alignment: .leading, spacing: .tbSpacing4) {
                        Text(item.productName)
                            .font(.tbCaption)
                            .foregroundColor(.tbTextPrimary)
                            .lineLimit(2)
                        
                        if let specName = item.specName {
                            Text(specName)
                                .font(.tbCaption2)
                                .foregroundColor(.tbTextTertiary)
                        }
                        
                        Spacer()
                        
                        HStack {
                            Text("¥\(String(format: "%.0f", item.price))")
                                .font(.tbCaption)
                                .foregroundColor(.tbTextPrimary)
                            
                            Spacer()
                            
                            Text("x\(item.quantity)")
                                .font(.tbCaption)
                                .foregroundColor(.tbTextTertiary)
                        }
                    }
                }
            }
            
            Divider()
            
            HStack {
                Spacer()
                
                Text("共\(order.itemCount)件商品")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
            }
        }
        .padding(.tbSpacing12)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
    
    // MARK: - Price Section
    private var priceSection: some View {
        VStack(spacing: .tbSpacing8) {
            HStack {
                Text("商品金额")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                
                Spacer()
                
                Text("¥\(String(format: "%.0f", order.subtotal))")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextPrimary)
            }
            
            if order.discount > 0 {
                HStack {
                    Text("优惠")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextSecondary)
                    
                    Spacer()
                    
                    Text("-¥\(String(format: "%.0f", order.discount))")
                        .font(.tbCaption)
                        .foregroundColor(.tbPrice)
                }
            }
            
            HStack {
                Text("运费")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                
                Spacer()
                
                Text(order.shippingFee > 0 ? "¥\(String(format: "%.0f", order.shippingFee))" : "免运费")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextPrimary)
            }
            
            Divider()
            
            HStack {
                Spacer()
                
                Text("实付款: ")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                
                Text("¥\(String(format: "%.0f", order.payAmount))")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbPrice)
            }
        }
        .padding(.tbSpacing12)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
    
    // MARK: - Order Info Section
    private var orderInfoSection: some View {
        VStack(spacing: .tbSpacing8) {
            HStack {
                Text("订单编号")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                
                Spacer()
                
                Text(order.orderNo)
                    .font(.tbCaption)
                    .foregroundColor(.tbTextPrimary)
            }
            
            HStack {
                Text("创建时间")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                
                Spacer()
                
                Text(order.createdAt)
                    .font(.tbCaption)
                    .foregroundColor(.tbTextPrimary)
            }
            
            if let paymentTime = order.paymentTime {
                HStack {
                    Text("付款时间")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextSecondary)
                    
                    Spacer()
                    
                    Text(paymentTime)
                        .font(.tbCaption)
                        .foregroundColor(.tbTextPrimary)
                }
            }
            
            if let shipTime = order.shipTime {
                HStack {
                    Text("发货时间")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextSecondary)
                    
                    Spacer()
                    
                    Text(shipTime)
                        .font(.tbCaption)
                        .foregroundColor(.tbTextPrimary)
                }
            }
            
            if !order.remark.isEmpty {
                HStack {
                    Text("备注")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextSecondary)
                    
                    Spacer()
                    
                    Text(order.remark)
                        .font(.tbCaption)
                        .foregroundColor(.tbTextPrimary)
                }
            }
        }
        .padding(.tbSpacing12)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
}

#Preview {
    OrderListView()
}
