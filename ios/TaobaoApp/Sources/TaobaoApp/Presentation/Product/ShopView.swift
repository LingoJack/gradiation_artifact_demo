import SwiftUI

struct ShopView: View {
    @StateObject private var shopStore = ShopStore.shared
    @StateObject private var favoriteStore = FavoriteStore.shared
    
    let shopId: String
    @State private var isFollowing = false
    @State private var showAllProducts = false
    @State private var selectedCategory: String = "全部"
    
    var shop: Shop? { shopStore.getShop(by: shopId) }
    
    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: 0) {
                    // Shop Header
                    if let shop = shop {
                        shopHeader(shop)
                    }
                    
                    // Shop Categories
                    if let shop = shop {
                        shopCategories(shop)
                    }
                    
                    // Products
                    shopProducts
                }
            }
            .background(Color.tbBackground)
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .principal) {
                    Text(shop?.name ?? "店铺")
                        .font(.tbBodyBold)
                }
            }
        }
    }
    
    // MARK: - Shop Header
    private func shopHeader(_ shop: Shop) -> some View {
        VStack(spacing: .tbSpacing16) {
            // Avatar and Name
            HStack(spacing: .tbSpacing12) {
                AsyncImage(url: URL(string: shop.avatar ?? "")) { phase in
                    switch phase {
                    case .success(let image):
                        image
                            .resizable()
                            .aspectRatio(contentMode: .fill)
                    default:
                        Color.tbDivider
                    }
                }
                .frame(width: 60, height: 60)
                .cornerRadius(30)
                
                VStack(alignment: .leading, spacing: 4) {
                    Text(shop.name)
                        .font(.tbHeadline)
                        .foregroundColor(.tbTextPrimary)
                    
                    HStack(spacing: .tbSpacing4) {
                        Image(systemName: "star.fill")
                            .font(.system(size: 10))
                            .foregroundColor(.tbPrice)
                        Text(shop.ratingText)
                            .font(.tbCaption)
                            .foregroundColor(.tbTextSecondary)
                        
                        Text("|")
                            .font(.tbCaption)
                            .foregroundColor(.tbTextTertiary)
                        
                        Text("\(shop.productCount)件商品")
                            .font(.tbCaption)
                            .foregroundColor(.tbTextSecondary)
                    }
                }
                
                Spacer()
                
                Button(action: { isFollowing.toggle() }) {
                    Text(isFollowing ? "已关注" : "关注")
                        .font(.tbCaption)
                        .foregroundColor(isFollowing ? .tbTextSecondary : .white)
                        .padding(.horizontal, .tbSpacing16)
                        .padding(.vertical, .tbSpacing8)
                        .background(isFollowing ? Color.tbBackground : Color.tbOrange)
                        .cornerRadius(16)
                }
            }
            
            // Stats
            HStack(spacing: 0) {
                statItem(title: "\(shop.followerCount / 10000)万", subtitle: "粉丝")
                Divider().frame(height: 30)
                statItem(title: "\(shop.salesCount / 10000)万", subtitle: "销量")
                Divider().frame(height: 30)
                statItem(title: "\(shop.ratingText)", subtitle: "评分")
            }
            .padding(.vertical, .tbSpacing12)
            .background(Color.tbBackground)
            .cornerRadius(.tbRadius8)
            
            // Tags
            ScrollView(.horizontal, showsIndicators: false) {
                HStack(spacing: .tbSpacing8) {
                    ForEach(shop.tags, id: \.self) { tag in
                        Text(tag)
                            .font(.tbCaption2)
                            .foregroundColor(.tbTextSecondary)
                            .padding(.horizontal, .tbSpacing8)
                            .padding(.vertical, 4)
                            .background(Color.tbBackground)
                            .cornerRadius(4)
                    }
                }
            }
        }
        .padding(.tbSpacing16)
        .background(Color.white)
    }
    
    private func statItem(title: String, subtitle: String) -> some View {
        VStack(spacing: 2) {
            Text(title)
                .font(.tbBodyBold)
                .foregroundColor(.tbTextPrimary)
            Text(subtitle)
                .font(.tbCaption2)
                .foregroundColor(.tbTextTertiary)
        }
        .frame(maxWidth: .infinity)
    }
    
    // MARK: - Shop Categories
    private func shopCategories(_ shop: Shop) -> some View {
        ScrollView(.horizontal, showsIndicators: false) {
            HStack(spacing: .tbSpacing8) {
                ForEach(shop.categories, id: \.self) { category in
                    Button(action: {
                        // Filter products by category
                        selectedCategory = category
                    }) {
                        Text(category)
                            .font(.tbCaption)
                            .foregroundColor(selectedCategory == category ? .white : .tbTextPrimary)
                            .padding(.horizontal, .tbSpacing16)
                            .padding(.vertical, .tbSpacing8)
                            .background(selectedCategory == category ? Color.tbOrange : Color.white)
                            .cornerRadius(16)
                    }
                }
            }
            .padding(.horizontal, .tbSpacing12)
            .padding(.vertical, .tbSpacing8)
        }
        .background(Color.tbBackground)
    }
    
    // MARK: - Shop Products
    private var shopProducts: some View {
        LazyVGrid(columns: [
            GridItem(.flexible(), spacing: .tbSpacing8),
            GridItem(.flexible(), spacing: .tbSpacing8)
        ], spacing: .tbSpacing8) {
            ForEach(MockDataService.shared.products.prefix(6)) { product in
                NavigationLink(value: product) {
                    ProductGridItem(product: product)
                }
                .buttonStyle(PlainButtonStyle())
            }
        }
        .padding(.tbSpacing12)
        .navigationDestination(for: Product.self) { product in
            ProductDetailView(product: product)
        }
    }
}

struct ProductGridItem: View {
    let product: Product
    
    var body: some View {
        VStack(alignment: .leading, spacing: .tbSpacing8) {
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
            .frame(height: 150)
            .clipped()
            .cornerRadius(.tbRadius4)
            
            Text(product.name)
                .font(.tbCaption)
                .foregroundColor(.tbTextPrimary)
                .lineLimit(2)
                .frame(height: 32, alignment: .top)
            
            HStack(alignment: .firstTextBaseline, spacing: 2) {
                Text("¥")
                    .font(.tbCaption)
                    .foregroundColor(.tbPrice)
                Text(String(format: "%.0f", product.price))
                    .font(.tbBodyBold)
                    .foregroundColor(.tbPrice)
            }
            
            HStack {
                if let discount = product.discount {
                    Text(discount)
                        .font(.system(size: 9))
                        .foregroundColor(.tbPrice)
                        .padding(.horizontal, 4)
                        .padding(.vertical, 2)
                        .background(Color.tbPrice.opacity(0.1))
                        .cornerRadius(2)
                }
                
                Spacer()
                
                Text("\(product.sales)人付款")
                    .font(.system(size: 10))
                    .foregroundColor(.tbTextTertiary)
            }
        }
        .padding(.tbSpacing8)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
}

#Preview {
    ShopView(shopId: "shop1")
}
