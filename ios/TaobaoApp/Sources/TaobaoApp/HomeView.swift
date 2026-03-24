import SwiftUI

// MARK: - Home View
struct HomeView: View {
    @State private var searchText: String = ""
    @State private var selectedCategory: Category?
    @State private var bannerIndex: Int = 0
    
    private let dataService = MockDataService.shared
    
    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: .tbSpacing12) {
                    // Search Bar
                    SearchBar(text: $searchText, placeholder: "搜索商品") {
                        // Handle search
                    }
                    .padding(.horizontal, .tbSpacing12)
                    .padding(.top, .tbSpacing8)
                    
                    // Banner Carousel
                    BannerCarousel(banners: dataService.banners, selectedIndex: $bannerIndex)
                        .frame(height: 140)
                        .padding(.horizontal, .tbSpacing12)
                    
                    // Categories
                    CategoryGrid(categories: Array(dataService.categories.prefix(8))) { category in
                        selectedCategory = category
                    }
                    .padding(.horizontal, .tbSpacing12)
                    
                    // Hot Products
                    VStack(spacing: .tbSpacing8) {
                        SectionHeader(title: "热卖推荐", showMore: true) {
                            // Navigate to product list
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
                            // Navigate to product list
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
            .navigationDestination(for: Product.self) { product in
                ProductDetailView(product: product)
            }
        }
    }
}

// MARK: - Banner Carousel
struct BannerCarousel: View {
    let banners: [Banner]
    @Binding var selectedIndex: Int
    
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
                .tag(index)
            }
        }
        .tabViewStyle(.page(indexDisplayMode: .automatic))
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
    HomeView()
}
