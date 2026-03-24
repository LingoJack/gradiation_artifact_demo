import SwiftUI

// MARK: - Category View
struct CategoryView: View {
    private let categories = MockDataService.shared.categories
    @State private var selectedCategory: Category?
    @StateObject private var coordinator = NavigationCoordinator.shared
    
    var body: some View {
        HStack(spacing: 0) {
            // Left Category List
            categoryList
            
            // Right Product Grid
            productGrid
        }
        .background(Color.tbBackground)
        .onAppear {
            if selectedCategory == nil {
                selectedCategory = categories.first
            }
        }
    }
    
    // MARK: - Category List
    private var categoryList: some View {
        ScrollView {
            VStack(spacing: 0) {
                ForEach(categories) { category in
                    Button(action: { selectedCategory = category }) {
                        HStack(spacing: 0) {
                            // Indicator
                            Rectangle()
                                .fill(selectedCategory?.id == category.id ? Color.tbOrange : Color.clear)
                                .frame(width: 3)
                            
                            Text(category.name)
                                .font(.tbBody)
                                .foregroundColor(selectedCategory?.id == category.id ? .tbOrange : .tbTextPrimary)
                                .frame(maxWidth: .infinity, alignment: .center)
                                .padding(.vertical, .tbSpacing16)
                        }
                        .background(selectedCategory?.id == category.id ? Color.white : Color.tbBackground)
                    }
                }
            }
        }
        .frame(width: 88)
        .background(Color.tbBackground)
    }
    
    // MARK: - Product Grid
    private var productGrid: some View {
        VStack(spacing: 0) {
            // Category Header
            if let category = selectedCategory {
                HStack {
                    Text(category.name)
                        .font(.tbBodyBold)
                        .foregroundColor(.tbTextPrimary)
                    Spacer()
                }
                .padding(.horizontal, .tbSpacing12)
                .padding(.vertical, .tbSpacing8)
                .background(Color.white)
            }
            
            // Product Grid
            ScrollView {
                LazyVGrid(columns: [
                    GridItem(.flexible(), spacing: .tbSpacing8),
                    GridItem(.flexible(), spacing: .tbSpacing8),
                    GridItem(.flexible(), spacing: .tbSpacing8)
                ], spacing: .tbSpacing8) {
                    ForEach(filteredProducts) { product in
                        NavigationLink(value: product) {
                            CategoryProductItem(product: product)
                        }
                        .buttonStyle(PlainButtonStyle())
                    }
                }
                .padding(.tbSpacing12)
            }
            .background(Color.white)
        }
    }
    
    private var filteredProducts: [Product] {
        guard let category = selectedCategory else {
            return MockDataService.shared.products
        }
        return MockDataService.shared.products.filter { $0.categoryId == category.id }
    }
}

// MARK: - Category Navigation View
struct CategoryNavigationView: View {
    @StateObject private var coordinator = NavigationCoordinator.shared
    
    var body: some View {
        NavigationStack(path: $coordinator.categoryPath) {
            CategoryView()
                .navigationDestination(for: Product.self) { product in
                    ProductDetailView(product: product)
                }
        }
    }
}

// MARK: - Category Product Item
struct CategoryProductItem: View {
    let product: Product
    
    var body: some View {
        VStack(alignment: .leading, spacing: .tbSpacing4) {
            AsyncImage(url: URL(string: product.mainImage)) { phase in
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
                        .overlay {
                            ProgressView()
                        }
                }
            }
            .frame(height: 80)
            .clipped()
            .cornerRadius(.tbRadius4)
            
            Text(product.name)
                .font(.system(size: 11))
                .foregroundColor(.tbTextPrimary)
                .lineLimit(2)
                .frame(height: 28, alignment: .top)
            
            HStack(alignment: .firstTextBaseline, spacing: 1) {
                Text("¥")
                    .font(.system(size: 10))
                    .foregroundColor(.tbPrice)
                Text(String(format: "%.0f", product.price))
                    .font(.system(size: 13, weight: .medium))
                    .foregroundColor(.tbPrice)
            }
        }
        .padding(.tbSpacing8)
        .background(Color.white)
        .cornerRadius(.tbRadius4)
        .shadow(color: Color.black.opacity(0.03), radius: 2, y: 1)
    }
}
