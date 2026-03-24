import SwiftUI

// MARK: - Main Tab View
struct MainTabView: View {
    @StateObject private var coordinator = NavigationCoordinator.shared
    
    enum Tab: Int {
        case home = 0
        case category = 1
        case cart = 2
        case message = 3
        case profile = 4
        
        var title: String {
            switch self {
            case .home: return "首页"
            case .category: return "分类"
            case .cart: return "购物车"
            case .message: return "消息"
            case .profile: return "我的"
            }
        }
        
        var icon: String {
            switch self {
            case .home: return "house"
            case .category: return "square.grid.2x2"
            case .cart: return "cart"
            case .message: return "message"
            case .profile: return "person"
            }
        }
        
        var iconFilled: String {
            switch self {
            case .home: return "house.fill"
            case .category: return "square.grid.2x2.fill"
            case .cart: return "cart.fill"
            case .message: return "message.fill"
            case .profile: return "person.fill"
            }
        }
    }
    
    var body: some View {
        ZStack(alignment: .bottom) {
            // Content
            TabView(selection: $coordinator.selectedTab) {
                HomeNavigationView()
                    .tag(Tab.home)
                
                CategoryNavigationView()
                    .tag(Tab.category)
                
                CartNavigationView()
                    .tag(Tab.cart)
                
                MessageNavigationView()
                    .tag(Tab.message)
                
                ProfileNavigationView()
                    .tag(Tab.profile)
            }
            .tabViewStyle(.page(indexDisplayMode: .never))
            
            // Custom Tab Bar
            tabBar
        }
        .ignoresSafeArea(.keyboard)
    }
    
    // MARK: - Tab Bar
    private var tabBar: some View {
        HStack(spacing: 0) {
            ForEach([Tab.home, .category, .cart, .message, .profile], id: \.self) { tab in
                Button(action: { coordinator.selectedTab = tab }) {
                    VStack(spacing: 4) {
                        ZStack(alignment: .topTrailing) {
                            Image(systemName: coordinator.selectedTab == tab ? tab.iconFilled : tab.icon)
                                .font(.system(size: 22))
                            
                            // Cart badge
                            if tab == .cart {
                                let count = CartStore.shared.totalCount
                                if count > 0 {
                                    Text("\(min(count, 99))")
                                        .font(.system(size: 9, weight: .bold))
                                        .foregroundColor(.white)
                                        .padding(4)
                                        .background(Color.tbPrice)
                                        .clipShape(Circle())
                                        .offset(x: 10, y: -6)
                                }
                            }
                            
                            // Message badge
                            if tab == .message {
                                Text("3")
                                    .font(.system(size: 9, weight: .bold))
                                    .foregroundColor(.white)
                                    .padding(4)
                                    .background(Color.tbPrice)
                                    .clipShape(Circle())
                                    .offset(x: 10, y: -6)
                            }
                        }
                        .frame(height: 24)
                        
                        Text(tab.title)
                            .font(.system(size: 10))
                    }
                    .foregroundColor(coordinator.selectedTab == tab ? .tbOrange : .tbTextTertiary)
                }
                .frame(maxWidth: .infinity)
            }
        }
        .padding(.top, 8)
        .padding(.bottom, 24)
        .background(Color.white)
        .shadow(color: Color.black.opacity(0.08), radius: 4, y: -2)
    }
}

// MARK: - Navigation Views
struct HomeNavigationView: View {
    @StateObject private var coordinator = NavigationCoordinator.shared
    
    var body: some View {
        NavigationStack(path: $coordinator.homePath) {
            HomeView()
                .navigationDestination(for: Product.self) { product in
                    ProductDetailView(product: product)
                }
                .navigationDestination(for: Category.self) { category in
                    ProductListView(category: category)
                }
                .navigationDestination(for: ProductListDestination.self) { destination in
                    ProductListView(category: destination.category, products: destination.products)
                }
                .navigationDestination(for: SearchDestination.self) { destination in
                    ProductListView(searchQuery: destination.query)
                }
        }
    }
}

struct CartNavigationView: View {
    @StateObject private var coordinator = NavigationCoordinator.shared
    
    var body: some View {
        NavigationStack(path: $coordinator.cartPath) {
            CartView()
                .navigationDestination(for: Product.self) { product in
                    ProductDetailView(product: product)
                }
        }
    }
}

struct ProfileNavigationView: View {
    @StateObject private var coordinator = NavigationCoordinator.shared
    
    var body: some View {
        NavigationStack(path: $coordinator.profilePath) {
            ProfileView()
                .navigationDestination(for: OrderListDestination.self) { destination in
                    OrderListView(statusFilter: destination.statusFilter)
                }
                .navigationDestination(for: Order.self) { order in
                    OrderDetailView(order: order)
                }
                .navigationDestination(for: FavoritesDestination.self) { _ in
                    FavoritesView()
                }
                .navigationDestination(for: AddressDestination.self) { _ in
                    AddressView()
                }
                .navigationDestination(for: CouponsDestination.self) { _ in
                    CouponsView()
                }
        }
    }
}


