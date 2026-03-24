import SwiftUI

struct FavoritesView: View {
    @StateObject private var favoriteStore = FavoriteStore.shared
    @State private var isEditing = false
    
    var body: some View {
        Group {
            if favoriteStore.favorites.isEmpty {
                emptyView
            } else {
                favoriteList
            }
        }
        .navigationTitle("我的收藏")
        .navigationBarTitleDisplayMode(.inline)
        .toolbar {
            if !favoriteStore.favorites.isEmpty {
                Button(isEditing ? "完成" : "编辑") {
                    withAnimation { isEditing.toggle() }
                }
            }
        }
    }
    
    private var emptyView: some View {
        VStack(spacing: .tbSpacing16) {
            Image(systemName: "heart.slash")
                .font(.system(size: 60))
                .foregroundColor(.tbTextTertiary)
            
            Text("暂无收藏商品")
                .font(.tbBody)
                .foregroundColor(.tbTextSecondary)
            
            Text("快去收藏心仪的商品吧")
                .font(.tbCaption)
                .foregroundColor(.tbTextTertiary)
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .background(Color.tbBackground)
    }
    
    private var favoriteList: some View {
        ScrollView {
            LazyVStack(spacing: .tbSpacing12) {
                ForEach(Array(favoriteStore.favorites.enumerated()), id: \.element.id) { index, favorite in
                    FavoriteItemView(
                        favorite: favorite,
                        isEditing: isEditing,
                        onDelete: { favoriteStore.removeFavorite(at: index) }
                    )
                }
            }
            .padding(.tbSpacing12)
        }
        .background(Color.tbBackground)
    }
}

struct FavoriteItemView: View {
    let favorite: Favorite
    let isEditing: Bool
    let onDelete: () -> Void
    
    var body: some View {
        HStack(spacing: .tbSpacing12) {
            if isEditing {
                Button(action: onDelete) {
                    Image(systemName: "minus.circle.fill")
                        .font(.system(size: 24))
                        .foregroundColor(.red)
                }
            }
            
            AsyncImage(url: URL(string: favorite.product.mainImage)) { phase in
                switch phase {
                case .success(let image):
                    image
                        .resizable()
                        .aspectRatio(contentMode: .fill)
                default:
                    Color.tbDivider
                }
            }
            .frame(width: 100, height: 100)
            .cornerRadius(.tbRadius8)
            
            VStack(alignment: .leading, spacing: .tbSpacing8) {
                Text(favorite.product.name)
                    .font(.tbBody)
                    .foregroundColor(.tbTextPrimary)
                    .lineLimit(2)
                
                Spacer()
                
                HStack {
                    Text("¥\(String(format: "%.0f", favorite.product.price))")
                        .font(.tbHeadline)
                        .foregroundColor(.tbPrice)
                    
                    if let discount = favorite.product.discount {
                        Text(discount)
                            .font(.tbCaption2)
                            .foregroundColor(.white)
                            .padding(.horizontal, 4)
                            .padding(.vertical, 2)
                            .background(Color.tbPrice)
                            .cornerRadius(2)
                    }
                    
                    Spacer()
                    
                    Text("\(favorite.product.sales)人付款")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextTertiary)
                }
            }
            .frame(maxHeight: .infinity)
        }
        .padding(.tbSpacing12)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
}

#Preview {
    FavoritesView()
}
