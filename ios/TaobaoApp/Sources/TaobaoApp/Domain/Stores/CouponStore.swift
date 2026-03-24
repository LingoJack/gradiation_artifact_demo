import Foundation
import Combine

// MARK: - Coupon Store
class CouponStore: ObservableObject {
    static let shared = CouponStore()
    
    @Published var coupons: [Coupon] = []
    
    private init() {
        loadMockData()
    }
    
    var availableCoupons: [Coupon] { coupons.filter { $0.status == .available } }
    var usedCoupons: [Coupon] { coupons.filter { $0.status == .used } }
    var expiredCoupons: [Coupon] { coupons.filter { $0.status == .expired } }
    
    func applicableCoupons(for amount: Double) -> [Coupon] {
        availableCoupons.filter { $0.minAmount <= amount }
    }
    
    func bestCoupon(for amount: Double) -> Coupon? {
        applicableCoupons(for: amount).max { $0.value < $1.value }
    }
    
    private func loadMockData() {
        coupons = [
            Coupon(
                id: "c1",
                name: "新人专享优惠券",
                type: .fixed,
                value: 50,
                minAmount: 199,
                startTime: "2024-01-01",
                endTime: "2024-12-31",
                status: .available,
                description: "新用户首单立减50元",
                scope: .all,
                productId: nil,
                categoryId: nil
            ),
            Coupon(
                id: "c2",
                name: "手机品类优惠券",
                type: .fixed,
                value: 100,
                minAmount: 1000,
                startTime: "2024-01-01",
                endTime: "2024-12-31",
                status: .available,
                description: "手机品类满1000减100",
                scope: .category,
                productId: nil,
                categoryId: "c1"
            ),
            Coupon(
                id: "c3",
                name: "苹果产品专享券",
                type: .fixed,
                value: 200,
                minAmount: 5000,
                startTime: "2024-01-01",
                endTime: "2024-12-31",
                status: .available,
                description: "苹果产品满5000减200",
                scope: .product,
                productId: "p1",
                categoryId: nil
            ),
            Coupon(
                id: "c4",
                name: "已使用的优惠券",
                type: .fixed,
                value: 30,
                minAmount: 100,
                startTime: "2024-01-01",
                endTime: "2024-12-31",
                status: .used,
                description: "已使用",
                scope: .all,
                productId: nil,
                categoryId: nil
            ),
            Coupon(
                id: "c5",
                name: "已过期优惠券",
                type: .fixed,
                value: 20,
                minAmount: 50,
                startTime: "2023-01-01",
                endTime: "2023-12-31",
                status: .expired,
                description: "已过期",
                scope: .all,
                productId: nil,
                categoryId: nil
            )
        ]
    }
}
