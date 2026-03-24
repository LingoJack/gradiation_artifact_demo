import SwiftUI

struct CheckoutView: View {
    @Environment(\.dismiss) private var dismiss
    @StateObject private var addressStore = AddressStore.shared
    @StateObject private var couponStore = CouponStore.shared
    @StateObject private var orderStore = OrderStore.shared
    @StateObject private var cartStore = CartStore.shared
    
    @State private var selectedAddress: Address?
    @State private var selectedCoupon: Coupon?
    @State private var remark: String = ""
    @State private var showAddressPicker = false
    @State private var showCouponPicker = false
    @State private var isProcessing = false
    @State private var orderCreated = false
    @State private var createdOrder: Order?
    
    let items: [CartItem]
    
    var subtotal: Double { items.reduce(0) { $0 + $1.subtotal } }
    var discount: Double { selectedCoupon?.value ?? 0 }
    var shippingFee: Double { subtotal >= 99 ? 0 : 10 }
    var total: Double { max(0, subtotal - discount + shippingFee) }
    
    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: .tbSpacing12) {
                    // Address Section
                    addressSection
                    
                    // Items Section
                    itemsSection
                    
                    // Coupon Section
                    couponSection
                    
                    // Remark Section
                    remarkSection
                    
                    // Price Breakdown
                    priceBreakdown
                }
                .padding(.tbSpacing12)
            }
            .background(Color.tbBackground)
            .navigationTitle("确认订单")
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .navigationBarLeading) {
                    Button("取消") { dismiss() }
                }
            }
            .safeAreaInset(edge: .bottom) {
                bottomBar
            }
            .sheet(isPresented: $showAddressPicker) {
                AddressPickerView(selectedAddress: $selectedAddress)
            }
            .sheet(isPresented: $showCouponPicker) {
                CouponPickerView(selectedCoupon: $selectedCoupon, subtotal: subtotal)
            }
            .fullScreenCover(isPresented: $orderCreated) {
                if let order = createdOrder {
                    PaymentView(order: order, onDismiss: { dismiss() })
                }
            }
        }
    }
    
    // MARK: - Address Section
    private var addressSection: some View {
        Button(action: { showAddressPicker = true }) {
            HStack(spacing: .tbSpacing12) {
                Image(systemName: "location.fill")
                    .font(.system(size: 20))
                    .foregroundColor(.tbOrange)
                
                if let address = selectedAddress ?? addressStore.defaultAddress {
                    VStack(alignment: .leading, spacing: 4) {
                        HStack {
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
                            .lineLimit(2)
                    }
                } else {
                    Text("请选择收货地址")
                        .font(.tbBody)
                        .foregroundColor(.tbTextTertiary)
                }
                
                Spacer()
                
                Image(systemName: "chevron.right")
                    .font(.system(size: 14))
                    .foregroundColor(.tbTextTertiary)
            }
            .padding(.tbSpacing16)
            .background(Color.white)
            .cornerRadius(.tbRadius8)
        }
    }
    
    // MARK: - Items Section
    private var itemsSection: some View {
        VStack(spacing: 0) {
            ForEach(items) { item in
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
                    
                    VStack(alignment: .leading, spacing: 4) {
                        Text(item.productName)
                            .font(.tbBody)
                            .foregroundColor(.tbTextPrimary)
                            .lineLimit(2)
                        
                        if let specName = item.specName, let specValue = item.specValue {
                            Text("\(specName): \(specValue)")
                                .font(.tbCaption)
                                .foregroundColor(.tbTextTertiary)
                        }
                        
                        Spacer()
                        
                        HStack {
                            Text("¥\(String(format: "%.0f", item.price))")
                                .font(.tbBodyBold)
                                .foregroundColor(.tbPrice)
                            
                            Spacer()
                            
                            Text("x\(item.quantity)")
                                .font(.tbCaption)
                                .foregroundColor(.tbTextTertiary)
                        }
                    }
                    .frame(maxHeight: .infinity)
                }
                .padding(.tbSpacing12)
                
                if item.id != items.last?.id {
                    Divider()
                        .padding(.horizontal, .tbSpacing12)
                }
            }
        }
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
    
    // MARK: - Coupon Section
    private var couponSection: some View {
        Button(action: { showCouponPicker = true }) {
            HStack {
                Text("优惠券")
                    .font(.tbBody)
                    .foregroundColor(.tbTextPrimary)
                
                Spacer()
                
                if let coupon = selectedCoupon {
                    Text("-¥\(Int(coupon.value))")
                        .font(.tbBodyBold)
                        .foregroundColor(.tbPrice)
                } else {
                    let available = couponStore.applicableCoupons(for: subtotal)
                    Text(available.isEmpty ? "暂无可用" : "\(available.count)张可用")
                        .font(.tbBody)
                        .foregroundColor(.tbTextTertiary)
                }
                
                Image(systemName: "chevron.right")
                    .font(.system(size: 14))
                    .foregroundColor(.tbTextTertiary)
            }
            .padding(.tbSpacing16)
            .background(Color.white)
            .cornerRadius(.tbRadius8)
        }
    }
    
    // MARK: - Remark Section
    private var remarkSection: some View {
        VStack(alignment: .leading, spacing: .tbSpacing8) {
            Text("订单备注")
                .font(.tbBodyBold)
                .foregroundColor(.tbTextPrimary)
            
            TextField("选填，请输入备注信息", text: $remark)
                .font(.tbBody)
                .padding(.tbSpacing12)
                .background(Color.tbBackground)
                .cornerRadius(.tbRadius4)
        }
        .padding(.tbSpacing16)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
    
    // MARK: - Price Breakdown
    private var priceBreakdown: some View {
        VStack(spacing: .tbSpacing8) {
            HStack {
                Text("商品金额")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                Spacer()
                Text("¥\(String(format: "%.2f", subtotal))")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextPrimary)
            }
            
            HStack {
                Text("运费")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                Spacer()
                Text(shippingFee > 0 ? "¥\(String(format: "%.2f", shippingFee))" : "免运费")
                    .font(.tbCaption)
                    .foregroundColor(shippingFee > 0 ? .tbTextPrimary : .tbGreen)
            }
            
            if discount > 0 {
                HStack {
                    Text("优惠")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextSecondary)
                    Spacer()
                    Text("-¥\(String(format: "%.2f", discount))")
                        .font(.tbCaption)
                        .foregroundColor(.tbPrice)
                }
            }
        }
        .padding(.tbSpacing16)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
    
    // MARK: - Bottom Bar
    private var bottomBar: some View {
        HStack(spacing: .tbSpacing16) {
            VStack(alignment: .leading, spacing: 2) {
                HStack(alignment: .firstTextBaseline, spacing: 2) {
                    Text("合计:")
                        .font(.tbBody)
                        .foregroundColor(.tbTextSecondary)
                    Text("¥")
                        .font(.tbSubheadline)
                        .foregroundColor(.tbPrice)
                    Text(String(format: "%.2f", total))
                        .font(.system(size: 24, weight: .bold))
                        .foregroundColor(.tbPrice)
                }
            }
            
            Spacer()
            
            Button(action: submitOrder) {
                Text("提交订单")
                    .font(.tbBodyBold)
                    .foregroundColor(.white)
                    .padding(.horizontal, .tbSpacing24)
                    .padding(.vertical, .tbSpacing12)
                    .background(Color.tbOrange)
                    .cornerRadius(20)
            }
            .disabled(isProcessing)
        }
        .padding(.horizontal, .tbSpacing16)
        .padding(.vertical, .tbSpacing12)
        .background(Color.white)
        .shadow(color: Color.black.opacity(0.08), radius: 4, y: -2)
    }
    
    private func submitOrder() {
        guard selectedAddress != nil || addressStore.defaultAddress != nil else {
            return
        }
        
        isProcessing = true
        
        // Create order
        let order = orderStore.createOrder(
            items: items,
            address: selectedAddress ?? addressStore.defaultAddress,
            coupon: selectedCoupon,
            remark: remark
        )
        
        // Remove items from cart
        for item in items {
            if let index = cartStore.items.firstIndex(where: { $0.id == item.id }) {
                cartStore.removeItem(at: index)
            }
        }
        
        createdOrder = order
        isProcessing = false
        orderCreated = true
    }
}

// MARK: - Address Picker View
struct AddressPickerView: View {
    @Environment(\.dismiss) private var dismiss
    @StateObject private var addressStore = AddressStore.shared
    @Binding var selectedAddress: Address?
    
    var body: some View {
        NavigationStack {
            List {
                ForEach(addressStore.addresses) { address in
                    Button(action: {
                        selectedAddress = address
                        dismiss()
                    }) {
                        HStack(spacing: .tbSpacing12) {
                            Image(systemName: selectedAddress?.id == address.id ? "checkmark.circle.fill" : "circle")
                                .foregroundColor(selectedAddress?.id == address.id ? .tbOrange : .tbTextTertiary)
                            
                            VStack(alignment: .leading, spacing: 4) {
                                HStack {
                                    Text(address.name)
                                        .font(.tbBodyBold)
                                    Text(address.phone)
                                        .font(.tbBody)
                                        .foregroundColor(.tbTextSecondary)
                                    
                                    if address.isDefault {
                                        Text("默认")
                                            .font(.tbCaption2)
                                            .foregroundColor(.tbOrange)
                                            .padding(.horizontal, 4)
                                            .padding(.vertical, 2)
                                            .background(Color.tbOrange.opacity(0.1))
                                            .cornerRadius(2)
                                    }
                                }
                                
                                Text(address.fullAddress)
                                    .font(.tbCaption)
                                    .foregroundColor(.tbTextSecondary)
                            }
                        }
                    }
                    .foregroundColor(.tbTextPrimary)
                }
            }
            .listStyle(.plain)
            .navigationTitle("选择地址")
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .navigationBarLeading) {
                    Button("取消") { dismiss() }
                }
            }
        }
    }
}

// MARK: - Coupon Picker View
struct CouponPickerView: View {
    @Environment(\.dismiss) private var dismiss
    @StateObject private var couponStore = CouponStore.shared
    @Binding var selectedCoupon: Coupon?
    let subtotal: Double
    
    private var availableCoupons: [Coupon] {
        couponStore.applicableCoupons(for: subtotal)
    }
    
    var body: some View {
        NavigationStack {
            ScrollView {
                LazyVStack(spacing: .tbSpacing12) {
                    ForEach(availableCoupons) { coupon in
                        Button(action: {
                            selectedCoupon = selectedCoupon?.id == coupon.id ? nil : coupon
                        }) {
                            HStack {
                                CouponCardView(coupon: coupon)
                                
                                Image(systemName: selectedCoupon?.id == coupon.id ? "checkmark.circle.fill" : "circle")
                                    .foregroundColor(selectedCoupon?.id == coupon.id ? .tbOrange : .tbTextTertiary)
                            }
                        }
                    }
                }
                .padding(.tbSpacing12)
            }
            .background(Color.tbBackground)
            .navigationTitle("选择优惠券")
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .navigationBarLeading) {
                    Button("取消") { dismiss() }
                }
                ToolbarItem(placement: .navigationBarTrailing) {
                    Button("确定") { dismiss() }
                }
            }
        }
    }
}

// MARK: - Payment View
struct PaymentView: View {
    @Environment(\.dismiss) private var dismiss
    @StateObject private var orderStore = OrderStore.shared
    let order: Order
    let onDismiss: () -> Void
    
    @State private var selectedMethod: PaymentMethod = .alipay
    @State private var isPaying = false
    
    enum PaymentMethod: Int, CaseIterable {
        case alipay = 0
        case wechat = 1
        
        var name: String {
            switch self {
            case .alipay: return "支付宝"
            case .wechat: return "微信支付"
            }
        }
        
        var icon: String {
            switch self {
            case .alipay: return "a.circle.fill"
            case .wechat: return "message.circle.fill"
            }
        }
        
        var color: Color {
            switch self {
            case .alipay: return .blue
            case .wechat: return .green
            }
        }
    }
    
    var body: some View {
        NavigationStack {
            VStack(spacing: .tbSpacing24) {
                // Order Amount
                VStack(spacing: .tbSpacing8) {
                    Text("支付金额")
                        .font(.tbBody)
                        .foregroundColor(.tbTextSecondary)
                    
                    HStack(alignment: .firstTextBaseline, spacing: 2) {
                        Text("¥")
                            .font(.tbTitle3)
                            .foregroundColor(.tbPrice)
                        Text(String(format: "%.2f", order.payAmount))
                            .font(.system(size: 48, weight: .bold))
                            .foregroundColor(.tbPrice)
                    }
                }
                .padding(.top, 40)
                
                // Payment Methods
                VStack(alignment: .leading, spacing: .tbSpacing12) {
                    Text("选择支付方式")
                        .font(.tbBodyBold)
                        .foregroundColor(.tbTextPrimary)
                    
                    ForEach(PaymentMethod.allCases, id: \.self) { method in
                        Button(action: { selectedMethod = method }) {
                            HStack(spacing: .tbSpacing12) {
                                Image(systemName: method.icon)
                                    .font(.system(size: 28))
                                    .foregroundColor(method.color)
                                
                                Text(method.name)
                                    .font(.tbBody)
                                    .foregroundColor(.tbTextPrimary)
                                
                                Spacer()
                                
                                Image(systemName: selectedMethod == method ? "checkmark.circle.fill" : "circle")
                                    .foregroundColor(selectedMethod == method ? .tbOrange : .tbTextTertiary)
                            }
                            .padding(.tbSpacing16)
                            .background(Color.white)
                            .cornerRadius(.tbRadius8)
                        }
                    }
                }
                
                Spacer()
                
                // Pay Button
                Button(action: pay) {
                    Text(isPaying ? "支付中..." : "立即支付")
                        .font(.tbBodyBold)
                        .foregroundColor(.white)
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, .tbSpacing16)
                        .background(isPaying ? Color.gray : Color.tbOrange)
                        .cornerRadius(24)
                }
                .disabled(isPaying)
                .padding(.horizontal, .tbSpacing24)
            }
            .background(Color.tbBackground)
            .navigationTitle("订单支付")
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .navigationBarLeading) {
                    Button("取消") {
                        dismiss()
                        onDismiss()
                    }
                }
            }
        }
    }
    
    private func pay() {
        isPaying = true
        
        // Simulate payment
        DispatchQueue.main.asyncAfter(deadline: .now() + 1.5) {
            orderStore.payOrder(order.id)
            isPaying = false
            dismiss()
            onDismiss()
        }
    }
}

#Preview {
    CheckoutView(items: [
        CartItem(
            id: "1",
            userId: "u1",
            productId: "p1",
            productName: "Apple iPhone 15 Pro Max 256GB",
            productImage: "https://picsum.photos/seed/iphone15/200/200",
            price: 9999,
            originalPrice: 11999,
            specId: "s1",
            specName: "颜色",
            specValue: "原色钛金属",
            quantity: 1
        )
    ])
}
