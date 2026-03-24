import SwiftUI

struct ProductListView: View {
    let category: Category?
    let initialProducts: [Product]
    
    @State private var products: [Product]
    @State private var sortOption: SortOption = .default
    @State private var showFilter: Bool = false
    
    init(category: Category? = nil, products: [Product]? = nil) {
        self.category = category
        self.initialProducts = products ?? MockDataService.shared.products
        _products = State(initialValue: self.initialProducts)
    }
    
    enum SortOption: String, CaseIterable {
        case `default` = "综合"
        case sales = "销量"
        case priceAsc = "价格升序"
        case priceDesc = "价格降序"
        
        var icon: String {
            switch self {
            case .default: return "list.bullet"
            case .sales: return "flame"
            case .priceAsc: return "arrow.up"
            case .priceDesc: return "arrow.down"
            }
        }
    }
    
    var body: some View {
        VStack(spacing: 0) {
            // Sort Bar
            sortBar
            
            // Product Grid
            ScrollView {
                LazyVGrid(columns: [
                    GridItem(.flexible(), spacing: .tbSpacing8),
                    GridItem(.flexible(), spacing: .tbSpacing8)
                ], spacing: .tbSpacing8) {
                    ForEach(products) { product in
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
        .navigationTitle(category?.name ?? "商品列表")
        .navigationBarTitleDisplayMode(.inline)
        .navigationDestination(for: Product.self) { product in
            ProductDetailView(product: product)
        }
    }
    
    // MARK: - Sort Bar
    private var sortBar: some View {
        HStack(spacing: 0) {
            ForEach(SortOption.allCases, id: \.self) { option in
                Button(action: {
                    sortOption = option
                    sortProducts(by: option)
                }) {
                    HStack(spacing: 4) {
                        Text(option.rawValue)
                            .font(.tbCaption)
                        
                        if option == .priceAsc || option == .priceDesc {
                            Image(systemName: sortOption == option ? "chevron.up.chevron.down" : "chevron.up.chevron.down")
                                .font(.system(size: 8))
                        }
                    }
                    .foregroundColor(sortOption == option ? .tbOrange : .tbTextSecondary)
                    .frame(maxWidth: .infinity)
                    .padding(.vertical, .tbSpacing12)
                }
            }
        }
        .background(Color.white)
    }
    
    // MARK: - Sort Products
    private func sortProducts(by option: SortOption) {
        withAnimation {
            switch option {
            case .default:
                products = initialProducts
            case .sales:
                products = initialProducts.sorted { $0.sales > $1.sales }
            case .priceAsc:
                products = initialProducts.sorted { $0.price < $1.price }
            case .priceDesc:
                products = initialProducts.sorted { $0.price > $1.price }
            }
        }
    }
}

// MARK: - Preview
#Preview {
    NavigationStack {
        ProductListView(category: MockDataService.shared.categories[0])
    }
}
