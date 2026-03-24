import SwiftUI

struct ProductDetailView: View {
    let product: Product
    
    @State private var selectedImageIndex: Int = 0
    @State private var quantity: Int = 1
    @State private var selectedSpec: ProductSpec?
    @State private var showSpecSheet: Bool = false
    @State private var showCartToast: Bool = false
    
    @Environment(\.dismiss) private var dismiss
    @ObservedObject private var cartStore = CartStore.shared
    
    private let dataService = MockDataService.shared
    
    var body: some View {
        ScrollView {
            VStack(spacing: 0) {
                // Image Gallery
                imageGallery
                
                // Product Info
                productInfoSection
                
                // Specs
                specsSection
                
                // Reviews
                reviewsSection
                
                // Shop Info
                shopInfoSection
                
                // Details
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
            specSelectionSheet
                .presentationDetents([.medium])
                .presentationDragIndicator(.visible)
        }
        .alert("已加入购物车", isPresented: $showCartToast) {
            Button("继续购物", role: .cancel) { }
            Button("去购物车") { /* Navigate to cart */ }
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
            // Price
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
            
            // Title
            Text(product.name)
                .font(.tbBody)
                .foregroundColor(.tbTextPrimary)
                .lineLimit(3)
            
            // Subtitle
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
                
                Button(action: {}) {
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
            
            Button(action: {}) {
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
            // Cart Button
            Button(action: {}) {
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
            
            // Add to Cart Button
            Button(action: {
                showSpecSheet = true
            }) {
                Text("加入购物车")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbOrange)
                    .frame(maxWidth: .infinity)
                    .padding(.vertical, .tbSpacing12)
                    .background(Color.tbTagBackground)
                    .cornerRadius(.tbRadius8)
            }
            
            // Buy Now Button
            Button(action: {
                // Handle buy now
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
        .padding(.horizontal, .tbSpacing16)
        .padding(.vertical, .tbSpacing8)
        .background(Color.white)
        .shadow(color: Color.black.opacity(0.05), radius: 4, y: -2)
    }
    
    // MARK: - Spec Selection Sheet
    private var specSelectionSheet: some View {
        VStack(alignment: .leading, spacing: .tbSpacing16) {
            // Product Summary
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
            
            // Specs
            if !product.specs.isEmpty {
                VStack(alignment: .leading, spacing: .tbSpacing8) {
                    Text(product.specs[0].name)
                        .font(.tbBodyBold)
                        .foregroundColor(.tbTextPrimary)
                    
                    FlowLayout(spacing: .tbSpacing8) {
                        ForEach(product.specs) { spec in
                            Button(action: {
                                selectedSpec = spec
                            }) {
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
            
            // Quantity
            HStack {
                Text("数量")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbTextPrimary)
                
                Spacer()
                
                QuantityStepper(quantity: $quantity, max: selectedSpec?.stock ?? product.stock)
            }
            
            Spacer()
            
            // Action Buttons
            HStack(spacing: .tbSpacing12) {
                Button(action: {
                    cartStore.addItem(product, spec: selectedSpec, quantity: quantity)
                    showSpecSheet = false
                    showCartToast = true
                }) {
                    Text("加入购物车")
                        .font(.tbBodyBold)
                        .foregroundColor(.tbOrange)
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, .tbSpacing12)
                        .background(Color.tbTagBackground)
                        .cornerRadius(.tbRadius8)
                }
                
                Button(action: {
                    // Handle buy now
                    showSpecSheet = false
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

// MARK: - Review Item View
struct ReviewItemView: View {
    let review: Review
    
    var body: some View {
        VStack(alignment: .leading, spacing: .tbSpacing8) {
            HStack(spacing: .tbSpacing8) {
                AsyncImage(url: URL(string: review.avatar ?? "")) { phase in
                    switch phase {
                    case .success(let image):
                        image
                            .resizable()
                            .aspectRatio(contentMode: .fill)
                    default:
                        Image(systemName: "person.circle.fill")
                            .resizable()
                            .foregroundColor(.tbTextTertiary)
                    }
                }
                .frame(width: 32, height: 32)
                .clipShape(Circle())
                
                VStack(alignment: .leading, spacing: 2) {
                    HStack(spacing: .tbSpacing4) {
                        Text(review.userName)
                            .font(.tbCaption)
                            .foregroundColor(.tbTextPrimary)
                        
                        if let badge = review.badge {
                            Text(badge)
                                .font(.system(size: 8))
                                .foregroundColor(.tbOrange)
                                .padding(.horizontal, 4)
                                .padding(.vertical, 2)
                                .background(Color.tbTagBackground)
                                .cornerRadius(2)
                        }
                    }
                    
                    HStack(spacing: 2) {
                        ForEach(0..<5, id: \.self) { index in
                            Image(systemName: index < review.rating ? "star.fill" : "star")
                                .font(.system(size: 10))
                                .foregroundColor(index < review.rating ? .yellow : .tbTextTertiary)
                        }
                    }
                }
                
                Spacer()
                
                Text(review.time ?? "")
                    .font(.tbCaption2)
                    .foregroundColor(.tbTextTertiary)
            }
            
            Text(review.content)
                .font(.tbCaption)
                .foregroundColor(.tbTextSecondary)
                .lineLimit(3)
            
            HStack(spacing: .tbSpacing16) {
                if let specs = review.specs {
                    Text(specs)
                        .font(.tbCaption2)
                        .foregroundColor(.tbTextTertiary)
                }
                
                Spacer()
                
                Button(action: {}) {
                    HStack(spacing: 4) {
                        Image(systemName: "hand.thumbsup")
                            .font(.tbCaption2)
                        Text("\(review.likes)")
                            .font(.tbCaption2)
                    }
                    .foregroundColor(.tbTextTertiary)
                }
            }
        }
        .padding(.tbSpacing12)
        .background(Color.tbBackground)
        .cornerRadius(.tbRadius8)
    }
}

// MARK: - Flow Layout
struct FlowLayout: Layout {
    var spacing: CGFloat = 8
    
    func sizeThatFits(proposal: ProposedViewSize, subviews: Subviews, cache: inout ()) -> CGSize {
        let result = FlowResult(in: proposal.width ?? 0, subviews: subviews, spacing: spacing)
        return result.size
    }
    
    func placeSubviews(in bounds: CGRect, proposal: ProposedViewSize, subviews: Subviews, cache: inout ()) {
        let result = FlowResult(in: bounds.width, subviews: subviews, spacing: spacing)
        for (index, subview) in subviews.enumerated() {
            subview.place(at: CGPoint(x: bounds.minX + result.positions[index].x,
                                      y: bounds.minY + result.positions[index].y),
                         proposal: .unspecified)
        }
    }
    
    struct FlowResult {
        var size: CGSize = .zero
        var positions: [CGPoint] = []
        
        init(in maxWidth: CGFloat, subviews: Subviews, spacing: CGFloat) {
            var x: CGFloat = 0
            var y: CGFloat = 0
            var rowHeight: CGFloat = 0
            
            for subview in subviews {
                let size = subview.sizeThatFits(.unspecified)
                
                if x + size.width > maxWidth, x > 0 {
                    x = 0
                    y += rowHeight + spacing
                    rowHeight = 0
                }
                
                positions.append(CGPoint(x: x, y: y))
                rowHeight = max(rowHeight, size.height)
                x += size.width + spacing
            }
            
            self.size = CGSize(width: maxWidth, height: y + rowHeight)
        }
    }
}

// MARK: - Preview
#Preview {
    NavigationStack {
        ProductDetailView(product: MockDataService.shared.products[0])
    }
}
