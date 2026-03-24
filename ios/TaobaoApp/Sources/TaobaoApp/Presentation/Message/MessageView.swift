import SwiftUI

// MARK: - Message View
struct MessageView: View {
    var body: some View {
        ScrollView {
            VStack(spacing: .tbSpacing12) {
                ForEach(0..<5) { index in
                    messageItem(
                        title: "系统消息",
                        content: "您的订单已发货，请注意查收",
                        time: "10:30"
                    )
                }
            }
            .padding(.tbSpacing12)
        }
        .background(Color.tbBackground)
        .navigationTitle("消息")
        .navigationBarTitleDisplayMode(.inline)
    }
    
    private func messageItem(title: String, content: String, time: String) -> some View {
        HStack(spacing: .tbSpacing12) {
            Image(systemName: "bell.circle.fill")
                .font(.system(size: 40))
                .foregroundColor(.tbOrange)
            
            VStack(alignment: .leading, spacing: 4) {
                HStack {
                    Text(title)
                        .font(.tbBodyBold)
                        .foregroundColor(.tbTextPrimary)
                    
                    Spacer()
                    
                    Text(time)
                        .font(.tbCaption2)
                        .foregroundColor(.tbTextTertiary)
                }
                
                Text(content)
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                    .lineLimit(1)
            }
        }
        .padding(.tbSpacing12)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
}

// MARK: - Message Navigation View
struct MessageNavigationView: View {
    @StateObject private var coordinator = NavigationCoordinator.shared
    
    var body: some View {
        NavigationStack(path: $coordinator.messagePath) {
            MessageView()
        }
    }
}
