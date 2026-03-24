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
