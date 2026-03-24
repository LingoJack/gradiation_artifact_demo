import SwiftUI

struct CartView: View {
    @ObservedObject private var cartStore = CartStore.shared
    @StateObject private var coordinator = NavigationCoordinator.shared
    @State private var showDeleteConfirmation: Bool = false
    
    var body: some View {
        Group {
            if cartStore.items.isEmpty {
                EmptyStateView(
                    icon: "cart",
                    title: "购物车是空的",
                    subtitle: "快去挑选心仪的商品吧",
                    actionTitle: "去逛逛",
                    action: { coordinator.switchToHome() }
                )
            } else {
                VStack(spacing: 0) {
                    // Cart Items
                    ScrollView {
                        LazyVStack(spacing: .tbSpacing12) {
                            ForEach(Array(cartStore.items.enumerated()), id: \.element.id) { index, item in
                                CartItemRow(
                                    item: item,
                                    isEditing: cartStore.isEditing,
                                    onToggleSelection: { cartStore.toggleSelection(at: index) },
                                    onQuantityChange: { quantity in
                                        cartStore.updateQuantity(at: index, quantity: quantity)
                                    },
                                    onDelete: { cartStore.removeItem(at: index) },
                                    onProductTap: { coordinator.pushToProduct(productId: item.productId) }
                                )
                            }
                        }
                        .padding(.tbSpacing12)
                    }
                    .background(Color.tbBackground)
                    
                    // Bottom Bar
                    cartBottomBar
                }
            }
        }
        .navigationTitle("购物车")
        .navigationBarTitleDisplayMode(.inline)
        .toolbar {
            ToolbarItem(placement: .navigationBarTrailing) {
                Button(action: { cartStore.isEditing.toggle() }) {
                    Text(cartStore.isEditing ? "完成" : "编辑")
                        .font(.tbBody)
                        .foregroundColor(.tbTextPrimary)
                }
            }
        }
        .alert("确认删除", isPresented: $showDeleteConfirmation) {
            Button("取消", role: .cancel) { }
            Button("删除", role: .destructive) {
                cartStore.removeSelectedItems()
            }
        } message: {
            Text("确定要删除选中的商品吗？")
        }
    }
    
    // MARK: - Bottom Bar
    private var cartBottomBar: some View {
        HStack(spacing: .tbSpacing12) {
            // Select All
            Button(action: { cartStore.toggleAllSelection() }) {
                HStack(spacing: .tbSpacing8) {
                    Image(systemName: cartStore.isAllSelected ? "checkmark.circle.fill" : "circle")
                        .font(.title3)
                        .foregroundColor(cartStore.isAllSelected ? .tbOrange : .tbTextTertiary)
                    
                    Text("全选")
                        .font(.tbBody)
                        .foregroundColor(.tbTextPrimary)
                }
            }
            
            Spacer()
            
            if cartStore.isEditing {
                Button(action: {
                    if !cartStore.selectedItems.isEmpty {
                        showDeleteConfirmation = true
                    }
                }) {
                    Text("删除(\(cartStore.selectedItems.count))")
                        .font(.tbBodyBold)
                        .foregroundColor(cartStore.selectedItems.isEmpty ? .tbTextTertiary : .tbPrice)
                        .padding(.horizontal, .tbSpacing20)
                        .padding(.vertical, .tbSpacing12)
                        .background(cartStore.selectedItems.isEmpty ? Color.tbDivider : Color.tbTagBackground)
                        .cornerRadius(.tbRadius8)
                }
            } else {
                VStack(alignment: .trailing, spacing: 2) {
                    HStack(spacing: 4) {
                        Text("合计:")
                            .font(.tbCaption)
                            .foregroundColor(.tbTextSecondary)
                        Text("¥")
                            .font(.tbCaption)
                            .foregroundColor(.tbPrice)
                        Text(String(format: "%.2f", cartStore.selectedTotalPrice))
                            .font(.tbTitle2)
                            .foregroundColor(.tbPrice)
                    }
                    
                    Text("共\(cartStore.selectedCount)件商品")
                        .font(.tbCaption2)
                        .foregroundColor(.tbTextTertiary)
                }
                
                Button(action: {
                    if !cartStore.selectedItems.isEmpty {
                        coordinator.pushToCheckout(items: cartStore.selectedItems)
                    }
                }) {
                    Text("结算")
                        .font(.tbBodyBold)
                        .foregroundColor(.white)
                        .padding(.horizontal, .tbSpacing24)
                        .padding(.vertical, .tbSpacing12)
                        .background(cartStore.selectedItems.isEmpty ? Color.tbTextTertiary : Color.tbOrange)
                        .cornerRadius(.tbRadius8)
                }
                .disabled(cartStore.selectedItems.isEmpty)
            }
        }
        .padding(.horizontal, .tbSpacing16)
        .padding(.vertical, .tbSpacing12)
        .background(Color.white)
        .shadow(color: Color.black.opacity(0.05), radius: 4, y: -2)
    }
}

// MARK: - Cart Item Row
struct CartItemRow: View {
    let item: CartItem
    let isEditing: Bool
    let onToggleSelection: () -> Void
    let onQuantityChange: (Int) -> Void
    let onDelete: () -> Void
    let onProductTap: () -> Void
    
    @State private var quantity: Int
    
    init(
        item: CartItem,
        isEditing: Bool,
        onToggleSelection: @escaping () -> Void,
        onQuantityChange: @escaping (Int) -> Void,
        onDelete: @escaping () -> Void,
        onProductTap: @escaping () -> Void = {}
    ) {
        self.item = item
        self.isEditing = isEditing
        self.onToggleSelection = onToggleSelection
        self.onQuantityChange = onQuantityChange
        self.onDelete = onDelete
        self.onProductTap = onProductTap
        self._quantity = State(initialValue: item.quantity)
    }
    
    var body: some View {
        HStack(spacing: .tbSpacing12) {
            // Selection
            Button(action: onToggleSelection) {
                Image(systemName: item.isSelected ? "checkmark.circle.fill" : "circle")
                    .font(.title2)
                    .foregroundColor(item.isSelected ? .tbOrange : .tbTextTertiary)
            }
            
            // Product Image
            AsyncImage(url: URL(string: item.productImage)) { phase in
                switch phase {
                case .success(let image):
                    image
                        .resizable()
                        .aspectRatio(contentMode: .fill)
                case .failure:
                    Color.tbDivider
                        .overlay {
                            Image(systemName: "photo")
                                .foregroundColor(.tbTextTertiary)
                        }
                default:
                    Color.tbDivider
                }
            }
            .frame(width: 90, height: 90)
            .cornerRadius(.tbRadius8)
            .onTapGesture {
                onProductTap()
            }
            
            // Product Info
            VStack(alignment: .leading, spacing: .tbSpacing8) {
                Text(item.productName)
                    .font(.tbBody)
                    .foregroundColor(.tbTextPrimary)
                    .lineLimit(2)
                
                if let specValue = item.specValue {
                    Text(specValue)
                        .font(.tbCaption)
                        .foregroundColor(.tbTextTertiary)
                        .padding(.horizontal, .tbSpacing4)
                        .padding(.vertical, 2)
                        .background(Color.tbBackground)
                        .cornerRadius(2)
                }
                
                Spacer()
                
                HStack {
                    HStack(alignment: .bottom, spacing: 2) {
                        Text("¥")
                            .font(.tbCaption)
                            .foregroundColor(.tbPrice)
                        Text(String(format: "%.0f", item.price))
                            .font(.tbPriceSmall)
                            .foregroundColor(.tbPrice)
                    }
                    
                    if let discount = item.discount {
                        DiscountTag(text: discount)
                    }
                    
                    Spacer()
                    
                    if !isEditing {
                        QuantityStepper(quantity: $quantity, max: 99)
                            .onChange(of: quantity) { newValue in
                                onQuantityChange(newValue)
                            }
                    }
                }
            }
            
            if isEditing {
                Button(action: onDelete) {
                    Image(systemName: "trash")
                        .font(.title3)
                        .foregroundColor(.tbPrice)
                }
            }
        }
        .padding(.tbSpacing12)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
}

// MARK: - Preview
#Preview {
    NavigationStack {
        CartView()
    }
}
