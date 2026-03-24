import SwiftUI

// MARK: - Review Item View
struct ReviewItemView: View {
    let review: Review
    @State private var isLiked: Bool = false
    @State private var likeCount: Int
    
    init(review: Review) {
        self.review = review
        self._likeCount = State(initialValue: review.likes)
    }
    
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
                
                Button(action: toggleLike) {
                    HStack(spacing: 4) {
                        Image(systemName: isLiked ? "hand.thumbsup.fill" : "hand.thumbsup")
                            .font(.tbCaption2)
                        Text("\(likeCount)")
                            .font(.tbCaption2)
                    }
                    .foregroundColor(isLiked ? .tbOrange : .tbTextTertiary)
                }
            }
        }
        .padding(.tbSpacing12)
        .background(Color.tbBackground)
        .cornerRadius(.tbRadius8)
    }
    
    private func toggleLike() {
        if isLiked {
            likeCount -= 1
        } else {
            likeCount += 1
        }
        isLiked.toggle()
    }
}
