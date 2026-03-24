import SwiftUI

struct CouponsView: View {
    @StateObject private var couponStore = CouponStore.shared
    @State private var selectedTab: Int = 0
    
    var body: some View {
        VStack(spacing: 0) {
            // Tab Bar
            HStack(spacing: 0) {
                ForEach(["可用", "已使用", "已过期"], id: \.self) { tab in
                    Button(action: {
                        selectedTab = ["可用", "已使用", "已过期"].firstIndex(of: tab) ?? 0
                    }) {
                        Text(tab)
                            .font(.tbBody)
                            .foregroundColor(selectedTab == ["可用", "已使用", "已过期"].firstIndex(of: tab) ? .tbOrange : .tbTextSecondary)
                            .frame(maxWidth: .infinity)
                            .padding(.vertical, .tbSpacing12)
                    }
                }
            }
            .background(Color.white)
            
            // Content
            ScrollView {
                LazyVStack(spacing: .tbSpacing12) {
                    let coupons: [Coupon] = {
                        switch selectedTab {
                        case 0: return couponStore.availableCoupons
                        case 1: return couponStore.usedCoupons
                        default: return couponStore.expiredCoupons
                        }
                    }()
                    
                    if coupons.isEmpty {
                        VStack(spacing: .tbSpacing16) {
                            Image(systemName: "ticket")
                                .font(.system(size: 60))
                                .foregroundColor(.tbTextTertiary)
                            
                            Text("暂无优惠券")
                                .font(.tbBody)
                                .foregroundColor(.tbTextSecondary)
                        }
                        .frame(maxWidth: .infinity)
                        .padding(.top, 100)
                    } else {
                        ForEach(coupons) { coupon in
                            CouponCardView(coupon: coupon)
                        }
                    }
                }
                .padding(.tbSpacing12)
            }
            .background(Color.tbBackground)
        }
        .navigationTitle("优惠券")
        .navigationBarTitleDisplayMode(.inline)
    }
}

struct CouponCardView: View {
    let coupon: Coupon
    
    var body: some View {
        HStack(spacing: 0) {
            // Left: Value
            VStack(spacing: 4) {
                HStack(alignment: .firstTextBaseline, spacing: 2) {
                    Text("¥")
                        .font(.tbCaption)
                    Text(Int(coupon.value).description)
                        .font(.system(size: 36, weight: .bold))
                }
                .foregroundColor(coupon.status == .available ? .tbPrice : .tbTextTertiary)
                
                Text("满\(Int(coupon.minAmount))可用")
                    .font(.tbCaption2)
                    .foregroundColor(coupon.status == .available ? .tbPrice : .tbTextTertiary)
            }
            .frame(width: 100)
            .padding(.vertical, .tbSpacing16)
            .background(coupon.status == .available ? Color.tbTagBackground : Color.tbDivider)
            
            // Right: Info
            VStack(alignment: .leading, spacing: .tbSpacing8) {
                Text(coupon.name)
                    .font(.tbBodyBold)
                    .foregroundColor(coupon.status == .available ? .tbTextPrimary : .tbTextTertiary)
                
                Text(coupon.description ?? "")
                    .font(.tbCaption)
                    .foregroundColor(.tbTextSecondary)
                
                Text("\(coupon.startTime) - \(coupon.endTime)")
                    .font(.tbCaption2)
                    .foregroundColor(.tbTextTertiary)
            }
            .padding(.tbSpacing12)
            
            Spacer()
        }
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
}

#Preview {
    NavigationStack {
        CouponsView()
    }
}
