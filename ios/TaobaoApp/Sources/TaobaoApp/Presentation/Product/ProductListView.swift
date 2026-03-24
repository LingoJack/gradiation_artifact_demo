import SwiftUI

struct ProductListView: View {
    var category: Category?
    var products: [Product]?
    var searchQuery: String?
    
    @State private var sortType: SortType = .default
    @State private var showFilter: Bool = false
    @StateObject private var coordinator = NavigationCoordinator.shared
    
    private var allProducts: [Product] {
        if let products = products {
            return products
        } else if let category = category {
            return MockDataService.shared.products.filter { $0.categoryId == category.id }
        } else if let query = searchQuery {
            return MockDataService.shared.products.filter { $0.name.localizedCaseInsensitiveContains(query) }
        } else {
            return MockDataService.shared.products
        }
    }
    
    private var filteredProducts: [Product] {
        switch sortType {
        case .default:
            return allProducts
        case .priceAsc:
            return allProducts.sorted { $0.price < $1.price }
        case .priceDesc:
            return allProducts.sorted { $0.price > $1.price }
        case .sales:
            return allProducts.sorted { $0.sales > $1.sales }
        }
    }
    
    var body: some View {
        VStack(spacing: 0) {
            // Search Bar
            HStack(spacing: .tbSpacing12) {
                HStack(spacing: .tbSpacing8) {
                    Image(systemName: "magnifyingglass")
                        .font(.tbBody)
                        .foregroundColor(.tbTextTertiary)
                    
                    Text(searchQuery ?? (category?.name ?? "搜索商品"))
                        .font(.tbBody)
                        .foregroundColor(searchQuery != nil ? .tbTextPrimary : .tbTextTertiary)
                    
                    Spacer()
                }
                .padding(.horizontal, .tbSpacing12)
                .padding(.vertical, .tbSpacing8)
                .background(Color.tbBackground)
                .cornerRadius(.tbRadius8)
                .onTapGesture {
                    coordinator.popHome()
                }
            }
            .padding(.horizontal, .tbSpacing12)
            .padding(.vertical, .tbSpacing8)
            .background(Color.white)
            
            // Sort Bar
            sortBar
            
            // Product Grid
            ScrollView {
                LazyVGrid(columns: [
                    GridItem(.flexible(), spacing: .tbSpacing8),
                    GridItem(.flexible(), spacing: .tbSpacing8)
                ], spacing: .tbSpacing8) {
                    ForEach(filteredProducts) { product in
                        NavigationLink(value: product) {
                            ProductCardView(product: product)
                        }
                        .buttonStyle(PlainButtonStyle())
                    }
                }
                .padding(.tbSpacing12)
            }
            .background(Color.tbBackground)
        }
        .navigationTitle(category?.name ?? (searchQuery != nil ? "搜索结果" : "商品列表"))
        .navigationBarTitleDisplayMode(.inline)
        .navigationDestination(for: Product.self) { product in
            ProductDetailView(product: product)
        }
    }
    
    // MARK: - Sort Bar
    private var sortBar: some View {
        HStack(spacing: 0) {
            ForEach(SortType.allCases, id: \.self) { sort in
                Button(action: { sortType = sort }) {
                    Text(sort.title)
                        .font(.tbBody)
                        .foregroundColor(sortType == sort ? .tbOrange : .tbTextSecondary)
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, .tbSpacing12)
                }
            }
            
            Button(action: { showFilter.toggle() }) {
                Image(systemName: "slider.horizontal.3")
                    .font(.tbBody)
                    .foregroundColor(.tbTextSecondary)
                    .padding(.trailing, .tbSpacing16)
            }
        }
        .background(Color.white)
    }
    
    enum SortType: CaseIterable {
        case `default`
        case sales
        case priceAsc
        case priceDesc
        
        var title: String {
            switch self {
            case .default: return "综合"
            case .sales: return "销量"
            case .priceAsc: return "价格↑"
            case .priceDesc: return "价格↓"
            }
        }
    }
}

// MARK: - Preview
#Preview {
    NavigationStack {
        ProductListView(category: MockDataService.shared.categories.first)
    }
}
