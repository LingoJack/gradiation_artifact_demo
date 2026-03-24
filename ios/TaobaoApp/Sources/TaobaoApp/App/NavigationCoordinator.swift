import SwiftUI

// MARK: - Navigation Coordinator
class NavigationCoordinator: ObservableObject {
    static let shared = NavigationCoordinator()
    
    // Tab selection
    @Published var selectedTab: MainTabView.Tab = .home
    
    // Navigation paths for each tab
    @Published var homePath = NavigationPath()
    @Published var categoryPath = NavigationPath()
    @Published var cartPath = NavigationPath()
    @Published var messagePath = NavigationPath()
    @Published var profilePath = NavigationPath()
    
    // Sheet presentations
    @Published var showLogin: Bool = false
    @Published var showCheckout: Bool = false
    @Published var checkoutItems: [CartItem] = []
    
    private init() {}
    
    // MARK: - Tab Navigation
    func switchToTab(_ tab: MainTabView.Tab) {
        selectedTab = tab
    }
    
    func switchToCart() {
        selectedTab = .cart
    }
    
    func switchToHome() {
        selectedTab = .home
    }
    
    func switchToProfile() {
        selectedTab = .profile
    }
    
    // MARK: - Home Navigation
    func pushToProduct(_ product: Product) {
        homePath.append(product)
    }
    
    func pushToProduct(productId: String) {
        if let product = MockDataService.shared.products.first(where: { $0.id == productId }) {
            homePath.append(product)
        }
    }
    
    func pushToCategory(_ category: Category) {
        homePath.append(category)
    }
    
    func pushToProductList(category: Category? = nil, products: [Product]? = nil) {
        homePath.append(ProductListDestination(category: category, products: products))
    }
    
    func pushToSearch(query: String) {
        homePath.append(SearchDestination(query: query))
    }
    
    func popHome() {
        if !homePath.isEmpty {
            homePath.removeLast()
        }
    }
    
    func popToHomeRoot() {
        homePath = NavigationPath()
    }
    
    // MARK: - Cart Navigation
    func pushToCheckout(items: [CartItem]) {
        checkoutItems = items
        showCheckout = true
    }
    
    func popCart() {
        if !cartPath.isEmpty {
            cartPath.removeLast()
        }
    }
    
    // MARK: - Profile Navigation
    func pushToOrderList(statusFilter: OrderStatus? = nil) {
        profilePath.append(OrderListDestination(statusFilter: statusFilter))
    }
    
    func pushToFavorites() {
        profilePath.append(FavoritesDestination())
    }
    
    func pushToAddresses() {
        profilePath.append(AddressDestination())
    }
    
    func pushToCoupons() {
        profilePath.append(CouponsDestination())
    }
    
    func popProfile() {
        if !profilePath.isEmpty {
            profilePath.removeLast()
        }
    }
    
    // MARK: - Login
    func presentLogin() {
        showLogin = true
    }
    
    func dismissLogin() {
        showLogin = false
    }
}

// MARK: - Navigation Destinations
struct ProductListDestination: Hashable {
    let category: Category?
    let products: [Product]?
    
    func hash(into hasher: inout Hasher) {
        hasher.combine(category?.id)
    }
    
    static func == (lhs: ProductListDestination, rhs: ProductListDestination) -> Bool {
        lhs.category?.id == rhs.category?.id
    }
}

struct SearchDestination: Hashable {
    let query: String
}

struct OrderListDestination: Hashable {
    let statusFilter: OrderStatus?
}

struct FavoritesDestination: Hashable {}

struct AddressDestination: Hashable {}

struct CouponsDestination: Hashable {}
