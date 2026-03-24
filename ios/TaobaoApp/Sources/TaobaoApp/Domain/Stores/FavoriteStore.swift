import Foundation
import Combine

// MARK: - Favorite Store
class FavoriteStore: ObservableObject {
    static let shared = FavoriteStore()
    
    @Published var favorites: [Favorite] = []
    
    private init() {
        loadMockData()
    }
    
    var favoriteProductIds: Set<String> { Set(favorites.map { $0.productId }) }
    
    func isFavorite(_ productId: String) -> Bool {
        favoriteProductIds.contains(productId)
    }
    
    func toggleFavorite(_ product: Product) {
        if let index = favorites.firstIndex(where: { $0.productId == product.id }) {
            favorites.remove(at: index)
        } else {
            let favorite = Favorite(
                id: UUID().uuidString,
                userId: UserStore.shared.user?.id ?? "guest",
                productId: product.id,
                product: product,
                createdAt: ISO8601DateFormatter().string(from: Date())
            )
            favorites.append(favorite)
        }
    }
    
    func removeFavorite(at index: Int) {
        guard index < favorites.count else { return }
        favorites.remove(at: index)
    }
    
    private func loadMockData() {
        // Load some mock favorites from mock products
        let products = MockDataService.shared.products
        let sampleProducts = products.prefix(5)
        
        favorites = sampleProducts.map { product in
            Favorite(
                id: UUID().uuidString,
                userId: "u1",
                productId: product.id,
                product: product,
                createdAt: ISO8601DateFormatter().string(from: Date())
            )
        }
    }
}
