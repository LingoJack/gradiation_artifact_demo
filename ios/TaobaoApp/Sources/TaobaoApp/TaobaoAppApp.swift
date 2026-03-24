import SwiftUI

@main
struct TaobaoAppApp: App {
    var body: some Scene {
        WindowGroup {
            MainTabView()
        }
    }
}

// MARK: - Main Tab View
struct MainTabView: View {
    @State private var selectedTab: Tab = .home
    
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
            TabView(selection: $selectedTab) {
                HomeView()
                    .tag(Tab.home)
                
                CategoryView()
                    .tag(Tab.category)
                
                CartView()
                    .tag(Tab.cart)
                
                MessageView()
                    .tag(Tab.message)
                
                ProfileView()
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
                Button(action: { selectedTab = tab }) {
                    VStack(spacing: 4) {
                        ZStack(alignment: .topTrailing) {
                            Image(systemName: selectedTab == tab ? tab.iconFilled : tab.icon)
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
                    .foregroundColor(selectedTab == tab ? .tbOrange : .tbTextTertiary)
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

// MARK: - Category View
struct CategoryView: View {
    private let categories = MockDataService.shared.categories
    
    var body: some View {
        NavigationStack {
            HStack(spacing: 0) {
                // Category List
                ScrollView {
                    VStack(spacing: 0) {
                        ForEach(categories) { category in
                            Text(category.name)
                                .font(.tbBody)
                                .foregroundColor(.tbTextPrimary)
                                .frame(maxWidth: .infinity, alignment: .leading)
                                .padding(.horizontal, .tbSpacing16)
                                .padding(.vertical, .tbSpacing16)
                                .background(Color.tbBackground)
                        }
                    }
                }
                .frame(width: 90)
                .background(Color.tbBackground)
                
                // Products
                ScrollView {
                    LazyVGrid(columns: [
                        GridItem(.flexible(), spacing: .tbSpacing8),
                        GridItem(.flexible(), spacing: .tbSpacing8),
                        GridItem(.flexible(), spacing: .tbSpacing8)
                    ], spacing: .tbSpacing8) {
                        ForEach(MockDataService.shared.products) { product in
                            NavigationLink(value: product) {
                                VStack(spacing: .tbSpacing4) {
                                    AsyncImage(url: URL(string: product.mainImage)) { phase in
                                        switch phase {
                                        case .success(let image):
                                            image
                                                .resizable()
                                                .aspectRatio(contentMode: .fill)
                                        default:
                                            Color.tbDivider
                                        }
                                    }
                                    .frame(height: 80)
                                    .clipped()
                                    .cornerRadius(.tbRadius4)
                                    
                                    Text(product.name)
                                        .font(.tbCaption)
                                        .foregroundColor(.tbTextPrimary)
                                        .lineLimit(2)
                                        .frame(height: 32, alignment: .top)
                                    
                                    Text("¥\(String(format: "%.0f", product.price))")
                                        .font(.tbCaption)
                                        .foregroundColor(.tbPrice)
                                }
                                .padding(.tbSpacing8)
                                .background(Color.white)
                                .cornerRadius(.tbRadius4)
                            }
                            .buttonStyle(PlainButtonStyle())
                        }
                    }
                    .padding(.tbSpacing12)
                }
                .background(Color.white)
                .navigationDestination(for: Product.self) { product in
                    ProductDetailView(product: product)
                }
            }
            .navigationTitle("分类")
            .navigationBarTitleDisplayMode(.inline)
        }
    }
}

// MARK: - Message View
struct MessageView: View {
    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: .tbSpacing12) {
                    ForEach(0..<5) { index in
                        messageItem(
                            title: "系统消息",
                            content: "您的订单已发货，请注意查收",
                            time: "10:30"
                        )
                    }
                }
                .padding(.tbSpacing12)
            }
            .background(Color.tbBackground)
            .navigationTitle("消息")
            .navigationBarTitleDisplayMode(.inline)
        }
    }
    
    private func messageItem(title: String, content: String, time: String) -> some View {
        HStack(spacing: .tbSpacing12) {
            Image(systemName: "bell.circle.fill")
                .font(.system(size: 40))
                .foregroundColor(.tbOrange)
            
            VStack(alignment: .leading, spacing: 4) {
                HStack {
                    Text(title)
                        .font(.tbBodyBold)
                        .foregroundColor(.tbTextPrimary)
                    
                    Spacer()
                    
                    Text(time)
                        .font(.tbCaption2)
                        .foregroundColor(.tbTextTertiary)
                }
                
                Text(content)
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                    .lineLimit(1)
            }
        }
        .padding(.tbSpacing12)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
}

// MARK: - Preview
#Preview {
    MainTabView()
}
