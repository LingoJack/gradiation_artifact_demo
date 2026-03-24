import SwiftUI

struct ProductDetailView: View {
    let product: Product
    
    @State private var selectedImageIndex: Int = 0
    @State private var quantity: Int = 1
    @State private var selectedSpec: ProductSpec?
    @State private var showSpecSheet: Bool = false
    @State private var showCartToast: Bool = false
    @State private var showBuyNowSheet: Bool = false
    @State private var showAllReviews: Bool = false
    @State private var likedReviews: Set<String> = []
    
    @Environment(\.dismiss) private var dismiss
    @ObservedObject private var cartStore = CartStore.shared
    @StateObject private var coordinator = NavigationCoordinator.shared
    
    private let dataService = MockDataService.shared
    
    var body: some View {
        ScrollView {
            VStack(spacing: 0) {
                imageGallery
                productInfoSection
                specsSection
                reviewsSection
                shopInfoSection
                detailsSection
            }
            .padding(.bottom, 60)
        }
        .background(Color.tbBackground)
        .navigationTitle("")
        .navigationBarTitleDisplayMode(.inline)
        .overlay(alignment: .bottom) {
            bottomBar
        }
        .sheet(isPresented: $showSpecSheet) {
            specSelectionSheet(mode: .addToCart)
                .presentationDetents([.medium])
                .presentationDragIndicator(.visible)
        }
        .sheet(isPresented: $showBuyNowSheet) {
            specSelectionSheet(mode: .buyNow)
                .presentationDetents([.medium])
                .presentationDragIndicator(.visible)
        }
        .alert("已加入购物车", isPresented: $showCartToast) {
            Button("继续购物", role: .cancel) { }
            Button("去购物车") {
                coordinator.switchToCart()
            }
        }
    }
    
    // MARK: - Image Gallery
    private var imageGallery: some View {
        TabView(selection: $selectedImageIndex) {
            ForEach(Array(product.images.enumerated()), id: \.offset) { index, _ in
                AsyncImage(url: URL(string: product.mainImage)) { phase in
                    switch phase {
                    case .success(let image):
                        image
                            .resizable()
                            .aspectRatio(contentMode: .fit)
                    case .failure:
                        Color.tbDivider
                            .overlay {
                                Image(systemName: "photo")
                                    .foregroundColor(.tbTextTertiary)
                            }
                    default:
                        Color.tbDivider
                    }
                }
                .tag(index)
            }
        }
        .frame(height: 375)
        .tabViewStyle(.page(indexDisplayMode: .automatic))
        .background(Color.white)
    }
    
    // MARK: - Product Info Section
    private var productInfoSection: some View {
        VStack(alignment: .leading, spacing: .tbSpacing12) {
            HStack(alignment: .bottom, spacing: .tbSpacing8) {
                Text("¥")
                    .font(.tbTitle3)
                    .foregroundColor(.tbPrice)
                Text(String(format: "%.0f", product.price))
                    .font(.tbPriceLarge)
                    .foregroundColor(.tbPrice)
                
                if let originalPrice = product.originalPrice {
                    Text("¥\(String(format: "%.0f", originalPrice))")
                        .font(.tbBody)
                        .foregroundColor(.tbTextTertiary)
                        .strikethrough()
                    
                    if let discount = product.discount {
                        DiscountTag(text: discount)
                    }
                }
            }
            
            Text(product.name)
                .font(.tbBody)
                .foregroundColor(.tbTextPrimary)
                .lineLimit(3)
            
            HStack(spacing: .tbSpacing16) {
                Text("销量 \(product.sales)")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextTertiary)
                
                Text("库存 \(product.stock)")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextTertiary)
                
                Spacer()
            }
        }
        .padding(.tbSpacing16)
        .background(Color.white)
    }
    
    // MARK: - Specs Section
    private var specsSection: some View {
        Button(action: { showSpecSheet = true }) {
            HStack {
                Text("规格")
                    .font(.tbBody)
                    .foregroundColor(.tbTextSecondary)
                
                Text(selectedSpec?.value ?? "请选择")
                    .font(.tbBody)
                    .foregroundColor(.tbTextPrimary)
                
                Spacer()
                
                Image(systemName: "chevron.right")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextTertiary)
            }
            .padding(.tbSpacing16)
            .background(Color.white)
            .padding(.top, .tbSpacing8)
        }
    }
    
    // MARK: - Reviews Section
    private var reviewsSection: some View {
        VStack(spacing: .tbSpacing12) {
            HStack {
                Text("用户评价")
                    .font(.tbTitle3)
                    .foregroundColor(.tbTextPrimary)
                
                Text("(\(dataService.getReviews(for: product.id).count))")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextTertiary)
                
                Spacer()
                
                Button(action: { showAllReviews = true }) {
                    HStack(spacing: 4) {
                        Text("查看全部")
                            .font(.tbCaption)
                        Image(systemName: "chevron.right")
                            .font(.tbCaption2)
                    }
                    .foregroundColor(.tbTextTertiary)
                }
            }
            
            let reviews = dataService.getReviews(for: product.id)
            if reviews.isEmpty {
                Text("暂无评价")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextTertiary)
                    .frame(maxWidth: .infinity)
                    .padding(.vertical, .tbSpacing20)
            } else {
                ForEach(reviews.prefix(2)) { review in
                    ReviewItemView(review: review)
                }
            }
        }
        .padding(.tbSpacing16)
        .background(Color.white)
        .padding(.top, .tbSpacing8)
    }
    
    // MARK: - Shop Info Section
    private var shopInfoSection: some View {
        HStack(spacing: .tbSpacing12) {
            Image(systemName: "shop")
                .font(.title2)
                .foregroundColor(.tbOrange)
                .frame(width: 44, height: 44)
                .background(Color.tbTagBackground)
                .cornerRadius(.tbRadius8)
            
            VStack(alignment: .leading, spacing: 4) {
                Text(product.shopName ?? "官方店铺")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbTextPrimary)
                
                HStack(spacing: .tbSpacing8) {
                    Text("评分 4.9")
                        .font(.tbCaption2)
                        .foregroundColor(.tbTextTertiary)
                    Text("粉丝 10万+")
                        .font(.tbCaption2)
                        .foregroundColor(.tbTextTertiary)
                }
            }
            
            Spacer()
            
            NavigationLink(value: ShopDestination(shopId: product.shopId ?? "")) {
                Text("进店")
                    .font(.tbCaption)
                    .foregroundColor(.tbOrange)
                    .padding(.horizontal, .tbSpacing12)
                    .padding(.vertical, .tbSpacing8)
                    .overlay(
                        RoundedRectangle(cornerRadius: .tbRadius8)
                            .stroke(Color.tbOrange, lineWidth: 1)
                    )
            }
        }
        .padding(.tbSpacing16)
        .background(Color.white)
        .padding(.top, .tbSpacing8)
    }
    
    // MARK: - Details Section
    private var detailsSection: some View {
        VStack(alignment: .leading, spacing: .tbSpacing12) {
            Text("商品详情")
                .font(.tbTitle3)
                .foregroundColor(.tbTextPrimary)
            
            Text(product.description)
                .font(.tbBody)
                .foregroundColor(.tbTextSecondary)
            
            ForEach(product.images, id: \.self) { _ in
                AsyncImage(url: URL(string: product.mainImage)) { phase in
                    switch phase {
                    case .success(let image):
                        image
                            .resizable()
                            .aspectRatio(contentMode: .fit)
                    default:
                        Color.tbDivider
                            .frame(height: 200)
                    }
                }
            }
        }
        .padding(.tbSpacing16)
        .background(Color.white)
        .padding(.top, .tbSpacing8)
    }
    
    // MARK: - Bottom Bar
    private var bottomBar: some View {
        HStack(spacing: .tbSpacing12) {
            Button(action: { coordinator.switchToHome() }) {
                VStack(spacing: 2) {
                    Image(systemName: "house")
                        .font(.title3)
                    Text("首页")
                        .font(.system(size: 10))
                }
                .foregroundColor(.tbTextSecondary)
            }
            .frame(width: 44)
            
            Button(action: { coordinator.switchToCart() }) {
                VStack(spacing: 2) {
                    ZStack(alignment: .topTrailing) {
                        Image(systemName: "cart")
                            .font(.title3)
                        
                        if cartStore.totalCount > 0 {
                            Text("\(cartStore.totalCount)")
                                .font(.system(size: 10, weight: .bold))
                                .foregroundColor(.white)
                                .padding(4)
                                .background(Color.tbPrice)
                                .clipShape(Circle())
                                .offset(x: 8, y: -8)
                        }
                    }
                    Text("购物车")
                        .font(.system(size: 10))
                }
                .foregroundColor(.tbTextSecondary)
            }
            .frame(width: 50)
            
            Button(action: { showSpecSheet = true }) {
                Text("加入购物车")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbOrange)
                    .frame(maxWidth: .infinity)
                    .padding(.vertical, .tbSpacing12)
                    .background(Color.tbTagBackground)
                    .cornerRadius(.tbRadius8)
            }
            
            Button(action: { showBuyNowSheet = true }) {
                Text("立即购买")
                    .font(.tbBodyBold)
                    .foregroundColor(.white)
                    .frame(maxWidth: .infinity)
                    .padding(.vertical, .tbSpacing12)
                    .background(Color.tbOrange)
                    .cornerRadius(.tbRadius8)
            }
        }
        .padding(.horizontal, .tbSpacing16)
        .padding(.vertical, .tbSpacing8)
        .background(Color.white)
        .shadow(color: Color.black.opacity(0.05), radius: 4, y: -2)
    }
    
    // MARK: - Spec Selection Sheet
    enum SheetMode {
        case addToCart
        case buyNow
    }
    
    private func specSelectionSheet(mode: SheetMode) -> some View {
        VStack(alignment: .leading, spacing: .tbSpacing16) {
            HStack(spacing: .tbSpacing12) {
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
                .frame(width: 80, height: 80)
                .cornerRadius(.tbRadius8)
                
                VStack(alignment: .leading, spacing: .tbSpacing4) {
                    PriceView(price: selectedSpec?.price ?? product.price)
                    
                    Text("库存: \(selectedSpec?.stock ?? product.stock)")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextTertiary)
                    
                    if let spec = selectedSpec {
                        Text("已选: \(spec.value)")
                            .font(.tbCaption)
                            .foregroundColor(.tbTextSecondary)
                    }
                }
                
                Spacer()
            }
            
            if !product.specs.isEmpty {
                VStack(alignment: .leading, spacing: .tbSpacing8) {
                    Text(product.specs[0].name)
                        .font(.tbBodyBold)
                        .foregroundColor(.tbTextPrimary)
                    
                    FlowLayout(spacing: .tbSpacing8) {
                        ForEach(product.specs) { spec in
                            Button(action: { selectedSpec = spec }) {
                                Text(spec.value)
                                    .font(.tbBody)
                                    .foregroundColor(selectedSpec?.id == spec.id ? .white : .tbTextPrimary)
                                    .padding(.horizontal, .tbSpacing12)
                                    .padding(.vertical, .tbSpacing8)
                                    .background(selectedSpec?.id == spec.id ? Color.tbOrange : Color.tbBackground)
                                    .cornerRadius(.tbRadius4)
                                    .overlay(
                                        RoundedRectangle(cornerRadius: .tbRadius4)
                                            .stroke(selectedSpec?.id == spec.id ? Color.tbOrange : Color.clear, lineWidth: 1)
                                    )
                            }
                        }
                    }
                }
            }
            
            HStack {
                Text("数量")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbTextPrimary)
                
                Spacer()
                
                QuantityStepper(quantity: $quantity, max: selectedSpec?.stock ?? product.stock)
            }
            
            Spacer()
            
            HStack(spacing: .tbSpacing12) {
                Button(action: {
                    cartStore.addItem(product, spec: selectedSpec, quantity: quantity)
                    
                    if mode == .addToCart {
                        showSpecSheet = false
                        showCartToast = true
                    } else {
                        showBuyNowSheet = false
                        let item = CartItem(
                            id: UUID().uuidString,
                            userId: "user1",
                            productId: product.id,
                            productName: product.name,
                            productImage: product.mainImage,
                            price: product.price,
                            originalPrice: product.originalPrice,
                            specId: selectedSpec?.id,
                            specName: selectedSpec?.name,
                            specValue: selectedSpec?.value,
                            quantity: quantity
                        )
                        coordinator.pushToCheckout(items: [item])
                    }
                }) {
                    Text(mode == .addToCart ? "加入购物车" : "加入购物车")
                        .font(.tbBodyBold)
                        .foregroundColor(.tbOrange)
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, .tbSpacing12)
                        .background(Color.tbTagBackground)
                        .cornerRadius(.tbRadius8)
                }
                
                Button(action: {
                    if mode == .buyNow {
                        showBuyNowSheet = false
                        let item = CartItem(
                            id: UUID().uuidString,
                            userId: "user1",
                            productId: product.id,
                            productName: product.name,
                            productImage: product.mainImage,
                            price: product.price,
                            originalPrice: product.originalPrice,
                            specId: selectedSpec?.id,
                            specName: selectedSpec?.name,
                            specValue: selectedSpec?.value,
                            quantity: quantity
                        )
                        coordinator.pushToCheckout(items: [item])
                    } else {
                        showSpecSheet = false
                        showBuyNowSheet = true
                    }
                }) {
                    Text("立即购买")
                        .font(.tbBodyBold)
                        .foregroundColor(.white)
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, .tbSpacing12)
                        .background(Color.tbOrange)
                        .cornerRadius(.tbRadius8)
                }
            }
        }
        .padding(.tbSpacing16)
    }
}

// MARK: - Shop Destination
struct ShopDestination: Hashable {
    let shopId: String
}

#Preview {
    NavigationStack {
        ProductDetailView(product: MockDataService.shared.products[0])
    }
}
