import SwiftUI

// MARK: - Home View
struct HomeView: View {
    @StateObject private var coordinator = NavigationCoordinator.shared
    @State private var searchText: String = ""
    @State private var bannerIndex: Int = 0
    
    private let dataService = MockDataService.shared
    
    var body: some View {
        ScrollView {
            VStack(spacing: .tbSpacing12) {
                // Search Bar
                SearchBar(text: $searchText, placeholder: "搜索商品") {
                    if !searchText.isEmpty {
                        coordinator.pushToSearch(query: searchText)
                    }
                }
                .onSubmit {
                    if !searchText.isEmpty {
                        coordinator.pushToSearch(query: searchText)
                    }
                }
                .padding(.horizontal, .tbSpacing12)
                .padding(.top, .tbSpacing8)
                
                // Banner Carousel
                BannerCarousel(banners: dataService.banners, selectedIndex: $bannerIndex)
                    .frame(height: 140)
                    .padding(.horizontal, .tbSpacing12)
                
                // Categories
                CategoryGrid(categories: Array(dataService.categories.prefix(8))) { category in
                    coordinator.pushToCategory(category)
                }
                .padding(.horizontal, .tbSpacing12)
                
                // Hot Products
                VStack(spacing: .tbSpacing8) {
                    SectionHeader(title: "热卖推荐", showMore: true) {
                        coordinator.pushToProductList()
                    }
                    
                    LazyVGrid(columns: [
                        GridItem(.flexible(), spacing: .tbSpacing8),
                        GridItem(.flexible(), spacing: .tbSpacing8)
                    ], spacing: .tbSpacing8) {
                        ForEach(dataService.products.prefix(4)) { product in
                            NavigationLink(value: product) {
                                ProductCardView(product: product)
                            }
                            .buttonStyle(PlainButtonStyle())
                        }
                    }
                    .padding(.horizontal, .tbSpacing12)
                }
                
                // Recommendations
                VStack(spacing: .tbSpacing8) {
                    SectionHeader(title: "为你推荐", showMore: true) {
                        coordinator.pushToProductList()
                    }
                    
                    LazyVGrid(columns: [
                        GridItem(.flexible(), spacing: .tbSpacing8),
                        GridItem(.flexible(), spacing: .tbSpacing8)
                    ], spacing: .tbSpacing8) {
                        ForEach(dataService.products) { product in
                            NavigationLink(value: product) {
                                ProductCardView(product: product)
                            }
                            .buttonStyle(PlainButtonStyle())
                        }
                    }
                    .padding(.horizontal, .tbSpacing12)
                }
            }
            .padding(.bottom, .tbSpacing20)
        }
        .background(Color.tbBackground)
    }
}

// MARK: - Banner Carousel
struct BannerCarousel: View {
    let banners: [Banner]
    @Binding var selectedIndex: Int
    @StateObject private var coordinator = NavigationCoordinator.shared
    
    var body: some View {
        TabView(selection: $selectedIndex) {
            ForEach(Array(banners.enumerated()), id: \.element.id) { index, banner in
                AsyncImage(url: URL(string: banner.image)) { phase in
                    switch phase {
                    case .success(let image):
                        image
                            .resizable()
                            .aspectRatio(contentMode: .fill)
                    case .failure:
                        Color.tbDivider
                    default:
                        Color.tbDivider
                    }
                }
                .cornerRadius(.tbRadius12)
                .onTapGesture {
                    handleBannerTap(banner)
                }
                .tag(index)
            }
        }
        .tabViewStyle(.page(indexDisplayMode: .automatic))
    }
    
    private func handleBannerTap(_ banner: Banner) {
        // 根据 link 类型决定跳转方式
        if banner.link.hasPrefix("product://") {
            let productId = banner.link.replacingOccurrences(of: "product://", with: "")
            coordinator.pushToProduct(productId: productId)
        } else if banner.link.hasPrefix("category://") {
            let categoryId = banner.link.replacingOccurrences(of: "category://", with: "")
            if let category = MockDataService.shared.categories.first(where: { $0.id == categoryId }) {
                coordinator.pushToCategory(category)
            }
        } else if banner.link.hasPrefix("search://") {
            let query = banner.link.replacingOccurrences(of: "search://", with: "")
            coordinator.pushToSearch(query: query)
        }
    }
}

// MARK: - Category Grid
struct CategoryGrid: View {
    let categories: [Category]
    let onSelect: (Category) -> Void
    
    private let columns = Array(repeating: GridItem(.flexible(), spacing: .tbSpacing12), count: 4)
    
    var body: some View {
        LazyVGrid(columns: columns, spacing: .tbSpacing16) {
            ForEach(categories) { category in
                Button(action: { onSelect(category) }) {
                    VStack(spacing: .tbSpacing8) {
                        Image(systemName: category.icon ?? "square.grid.2x2")
                            .font(.system(size: 28))
                            .foregroundColor(.tbOrange)
                        
                        Text(category.name)
                            .font(.tbCaption)
                            .foregroundColor(.tbTextPrimary)
                            .lineLimit(1)
                    }
                    .frame(maxWidth: .infinity)
                    .padding(.vertical, .tbSpacing12)
                    .background(Color.tbCardBackground)
                    .cornerRadius(.tbRadius8)
                }
            }
        }
    }
}

// MARK: - Preview
#Preview {
    NavigationStack {
        HomeView()
    }
}
