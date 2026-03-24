import SwiftUI

// MARK: - Product Card View
struct ProductCardView: View {
    let product: Product
    var showSales: Bool = true
    
    var body: some View {
        VStack(alignment: .leading, spacing: .tbSpacing8) {
            // Product Image
            AsyncImage(url: URL(string: product.mainImage)) { phase in
                switch phase {
                case .success(let image):
                    image
                        .resizable()
                        .aspectRatio(contentMode: .fill)
                case .failure:
                    Image(systemName: "photo")
                        .resizable()
                        .aspectRatio(contentMode: .fit)
                        .foregroundColor(.tbTextTertiary)
                default:
                    Color.tbDivider
                }
            }
            .frame(height: 160)
            .clipped()
            .cornerRadius(.tbRadius8)
            
            // Product Info
            VStack(alignment: .leading, spacing: .tbSpacing4) {
                Text(product.name)
                    .font(.tbBody)
                    .foregroundColor(.tbTextPrimary)
                    .lineLimit(2)
                    .frame(height: 40, alignment: .top)
                
                HStack(alignment: .bottom, spacing: .tbSpacing4) {
                    Text("¥")
                        .font(.tbCaption)
                        .foregroundColor(.tbPrice)
                    Text(String(format: "%.0f", product.price))
                        .font(.tbPrice)
                        .foregroundColor(.tbPrice)
                    
                    if let originalPrice = product.originalPrice {
                        Text("¥\(String(format: "%.0f", originalPrice))")
                            .font(.tbCaption2)
                            .foregroundColor(.tbTextTertiary)
                            .strikethrough()
                    }
                    
                    Spacer()
                    
                    if showSales {
                        Text("\(product.sales)人付款")
                            .font(.tbCaption2)
                            .foregroundColor(.tbTextTertiary)
                    }
                }
                
                if let discount = product.discount {
                    HStack {
                        DiscountTag(text: discount)
                        Spacer()
                    }
                }
            }
        }
        .padding(.tbSpacing8)
        .background(Color.tbCardBackground)
        .cornerRadius(.tbRadius8)
    }
}

// MARK: - Discount Tag
struct DiscountTag: View {
    let text: String
    
    var body: some View {
        Text(text)
            .font(.tbCaption2)
            .foregroundColor(.tbOrange)
            .padding(.horizontal, .tbSpacing4)
            .padding(.vertical, 2)
            .background(Color.tbTagBackground)
            .cornerRadius(2)
    }
}

// MARK: - Search Bar
struct SearchBar: View {
    @Binding var text: String
    var placeholder: String = "搜索商品"
    var onSubmit: (() -> Void)? = nil
    
    var body: some View {
        HStack(spacing: .tbSpacing8) {
            HStack(spacing: .tbSpacing8) {
                Image(systemName: "magnifyingglass")
                    .foregroundColor(.tbTextTertiary)
                
                TextField(placeholder, text: $text)
                    .font(.tbBody)
                    .submitLabel(.search)
                    .onSubmit {
                        onSubmit?()
                    }
                
                if !text.isEmpty {
                    Button(action: { text = "" }) {
                        Image(systemName: "xmark.circle.fill")
                            .foregroundColor(.tbTextTertiary)
                    }
                }
            }
            .padding(.tbSpacing8)
            .background(Color.tbCardBackground)
            .cornerRadius(.tbRadius8)
        }
    }
}

// MARK: - Empty State View
struct EmptyStateView: View {
    let icon: String
    let title: String
    let subtitle: String?
    let actionTitle: String?
    let action: (() -> Void)?
    
    init(
        icon: String,
        title: String,
        subtitle: String? = nil,
        actionTitle: String? = nil,
        action: (() -> Void)? = nil
    ) {
        self.icon = icon
        self.title = title
        self.subtitle = subtitle
        self.actionTitle = actionTitle
        self.action = action
    }
    
    var body: some View {
        VStack(spacing: .tbSpacing16) {
            Image(systemName: icon)
                .font(.system(size: 60))
                .foregroundColor(.tbTextTertiary)
            
            Text(title)
                .font(.tbTitle3)
                .foregroundColor(.tbTextSecondary)
            
            if let subtitle = subtitle {
                Text(subtitle)
                    .font(.tbCaption)
                    .foregroundColor(.tbTextTertiary)
            }
            
            if let actionTitle = actionTitle, let action = action {
                Button(action: action) {
                    Text(actionTitle)
                        .font(.tbBodyBold)
                        .foregroundColor(.white)
                        .padding(.horizontal, .tbSpacing24)
                        .padding(.vertical, .tbSpacing12)
                        .background(Color.tbOrange)
                        .cornerRadius(.tbRadius8)
                }
            }
        }
    }
}

// MARK: - Loading View
struct LoadingView: View {
    var text: String = "加载中..."
    
    var body: some View {
        VStack(spacing: .tbSpacing12) {
            ProgressView()
                .progressViewStyle(CircularProgressViewStyle(tint: .tbOrange))
            
            Text(text)
                .font(.tbCaption)
                .foregroundColor(.tbTextTertiary)
        }
    }
}

// MARK: - Quantity Stepper
struct QuantityStepper: View {
    @Binding var quantity: Int
    var min: Int = 1
    var max: Int = 99
    
    var body: some View {
        HStack(spacing: 0) {
            Button(action: {
                if quantity > min { quantity -= 1 }
            }) {
                Image(systemName: "minus")
                    .font(.tbCaption)
                    .foregroundColor(quantity <= min ? .tbTextTertiary : .tbTextPrimary)
                    .frame(width: 28, height: 28)
            }
            .disabled(quantity <= min)
            
            Text("\(quantity)")
                .font(.tbBody)
                .foregroundColor(.tbTextPrimary)
                .frame(minWidth: 36)
            
            Button(action: {
                if quantity < max { quantity += 1 }
            }) {
                Image(systemName: "plus")
                    .font(.tbCaption)
                    .foregroundColor(quantity >= max ? .tbTextTertiary : .tbTextPrimary)
                    .frame(width: 28, height: 28)
            }
            .disabled(quantity >= max)
        }
        .background(Color.tbBackground)
        .cornerRadius(.tbRadius4)
    }
}

// MARK: - Price View
struct PriceView: View {
    let price: Double
    var originalPrice: Double? = nil
    var fontSize: Font = .tbPrice
    var showDiscount: Bool = true
    
    var body: some View {
        HStack(alignment: .bottom, spacing: .tbSpacing4) {
            Text("¥")
                .font(.tbCaption)
                .foregroundColor(.tbPrice)
            Text(String(format: "%.0f", price))
                .font(fontSize)
                .foregroundColor(.tbPrice)
            
            if let original = originalPrice, showDiscount {
                Text("¥\(String(format: "%.0f", original))")
                    .font(.tbCaption2)
                    .foregroundColor(.tbTextTertiary)
                    .strikethrough()
            }
        }
    }
}

// MARK: - Section Header
struct SectionHeader: View {
    let title: String
    var showMore: Bool = false
    var moreText: String = "更多"
    var onMore: (() -> Void)? = nil
    
    var body: some View {
        HStack {
            Text(title)
                .font(.tbTitle2)
                .foregroundColor(.tbTextPrimary)
            
            Spacer()
            
            if showMore {
                Button(action: { onMore?() }) {
                    HStack(spacing: 2) {
                        Text(moreText)
                            .font(.tbCaption)
                        Image(systemName: "chevron.right")
                            .font(.tbCaption2)
                    }
                    .foregroundColor(.tbTextTertiary)
                }
            }
        }
        .padding(.horizontal, .tbSpacing16)
        .padding(.vertical, .tbSpacing12)
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
    ScrollView {
        VStack(spacing: 20) {
            ProductCardView(product: MockDataService.shared.products[0])
                .frame(width: 180)
            
            SearchBar(text: .constant("iPhone"))
                .padding()
            
            PriceView(price: 9999, originalPrice: 11999)
            
            QuantityStepper(quantity: .constant(1))
        }
        .padding()
    }
}
