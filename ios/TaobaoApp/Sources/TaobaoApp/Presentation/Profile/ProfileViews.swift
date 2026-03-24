import SwiftUI

struct ProfileView: View {
    @StateObject private var coordinator = NavigationCoordinator.shared
    @StateObject private var userStore = UserStore.shared
    
    var body: some View {
        ScrollView {
            VStack(spacing: .tbSpacing12) {
                // User Header
                userHeader
                
                // Order Section
                orderSection
                
                // Tools Section
                toolsSection
                
                // Menu Section
                menuSection
            }
            .padding(.tbSpacing12)
        }
        .background(Color.tbBackground)
        .navigationTitle("我的淘宝")
        .navigationBarTitleDisplayMode(.inline)
    }
    
    // MARK: - User Header
    private var userHeader: some View {
        HStack(spacing: .tbSpacing16) {
            AsyncImage(url: URL(string: userStore.currentUser?.avatar ?? "")) { phase in
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
            .frame(width: 60, height: 60)
            .clipShape(Circle())
            
            VStack(alignment: .leading, spacing: .tbSpacing4) {
                Text(userStore.currentUser?.name ?? "未登录")
                    .font(.tbTitle3)
                    .foregroundColor(.tbTextPrimary)
                
                if let user = userStore.currentUser {
                    Text("会员等级: \(user.level)")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextSecondary)
                }
            }
            
            Spacer()
            
            Image(systemName: "chevron.right")
                .font(.tbBody)
                .foregroundColor(.tbTextTertiary)
        }
        .padding(.tbSpacing16)
        .background(Color.white)
        .cornerRadius(.tbRadius12)
    }
    
    // MARK: - Order Section
    private var orderSection: some View {
        VStack(spacing: 0) {
            HStack {
                Text("我的订单")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbTextPrimary)
                
                Spacer()
                
                Button(action: { coordinator.pushToOrderList() }) {
                    HStack(spacing: 4) {
                        Text("全部订单")
                            .font(.tbCaption)
                        Image(systemName: "chevron.right")
                            .font(.tbCaption2)
                    }
                    .foregroundColor(.tbTextTertiary)
                }
            }
            .padding(.tbSpacing16)
            
            HStack(spacing: 0) {
                orderItem(icon: "creditcard", title: "待付款", count: 2, status: .pending)
                orderItem(icon: "box", title: "待发货", count: 1, status: .paid)
                orderItem(icon: "truck.box", title: "待收货", count: 3, status: .shipped)
                orderItem(icon: "star", title: "待评价", count: 5, status: .completed)
                orderItem(icon: "arrowshape.turn.up.backward", title: "退换/售后", count: 0, status: nil)
            }
            .padding(.bottom, .tbSpacing16)
        }
        .background(Color.white)
        .cornerRadius(.tbRadius12)
    }
    
    private func orderItem(icon: String, title: String, count: Int, status: OrderStatus?) -> some View {
        Button(action: { coordinator.pushToOrderList(statusFilter: status) }) {
            VStack(spacing: .tbSpacing8) {
                ZStack(alignment: .topTrailing) {
                    Image(systemName: icon)
                        .font(.title2)
                        .foregroundColor(.tbTextPrimary)
                    
                    if count > 0 {
                        Text("\(count)")
                            .font(.system(size: 10, weight: .bold))
                            .foregroundColor(.white)
                            .padding(4)
                            .background(Color.tbPrice)
                            .clipShape(Circle())
                            .offset(x: 8, y: -8)
                    }
                }
                .frame(height: 24)
                
                Text(title)
                    .font(.tbCaption)
                    .foregroundColor(.tbTextPrimary)
            }
            .frame(maxWidth: .infinity)
        }
    }
    
    // MARK: - Tools Section
    private var toolsSection: some View {
        VStack(spacing: 0) {
            HStack {
                Text("我的服务")
                    .font(.tbBodyBold)
                    .foregroundColor(.tbTextPrimary)
                
                Spacer()
            }
            .padding(.tbSpacing16)
            
            LazyVGrid(columns: Array(repeating: GridItem(.flexible(), spacing: .tbSpacing16), count: 4), spacing: .tbSpacing16) {
                toolItem(icon: "heart", title: "收藏夹", action: { coordinator.pushToFavorites() })
                toolItem(icon: "map", title: "地址管理", action: { coordinator.pushToAddresses() })
                toolItem(icon: "ticket", title: "优惠券", action: { coordinator.pushToCoupons() })
                toolItem(icon: "clock", title: "浏览记录", action: { })
                toolItem(icon: "gift", title: "积分商城", action: { })
                toolItem(icon: "questionmark.circle", title: "帮助中心", action: { })
                toolItem(icon: "gearshape", title: "设置", action: { })
                toolItem(icon: "ellipsis", title: "更多", action: { })
            }
            .padding(.horizontal, .tbSpacing16)
            .padding(.bottom, .tbSpacing16)
        }
        .background(Color.white)
        .cornerRadius(.tbRadius12)
    }
    
    private func toolItem(icon: String, title: String, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            VStack(spacing: .tbSpacing8) {
                Image(systemName: icon)
                    .font(.title2)
                    .foregroundColor(.tbOrange)
                
                Text(title)
                    .font(.tbCaption)
                    .foregroundColor(.tbTextPrimary)
            }
            .frame(maxWidth: .infinity)
        }
    }
    
    // MARK: - Menu Section
    private var menuSection: some View {
        VStack(spacing: 0) {
            menuItem(icon: "bell", title: "消息通知", showBadge: true) {
                // TODO: Navigate to notifications
            }
            Divider().padding(.leading, 44)
            menuItem(icon: "shield", title: "账户安全") {
                // TODO: Navigate to security settings
            }
            Divider().padding(.leading, 44)
            menuItem(icon: "doc.text", title: "用户协议") {
                // TODO: Navigate to user agreement
            }
        }
        .background(Color.white)
        .cornerRadius(.tbRadius12)
    }
    
    private func menuItem(icon: String, title: String, showBadge: Bool = false, action: @escaping () -> Void = {}) -> some View {
        Button(action: action) {
            HStack(spacing: .tbSpacing12) {
                Image(systemName: icon)
                    .font(.title3)
                    .foregroundColor(.tbTextSecondary)
                    .frame(width: 20)
                
                Text(title)
                    .font(.tbBody)
                    .foregroundColor(.tbTextPrimary)
                
                Spacer()
                
                if showBadge {
                    Circle()
                        .fill(Color.tbPrice)
                        .frame(width: 8, height: 8)
                }
                
                Image(systemName: "chevron.right")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextTertiary)
            }
            .padding(.tbSpacing16)
        }
    }
}

// MARK: - Preview
#Preview {
    NavigationStack {
        ProfileView()
    }
}
