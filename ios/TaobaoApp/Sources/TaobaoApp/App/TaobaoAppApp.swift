import SwiftUI

@main
struct TaobaoAppApp: App {
    @StateObject private var coordinator = NavigationCoordinator.shared
    @StateObject private var cartStore = CartStore.shared
    
    var body: some Scene {
        WindowGroup {
            MainTabView()
                .environmentObject(coordinator)
                .sheet(isPresented: $coordinator.showCheckout) {
                    CheckoutView(items: coordinator.checkoutItems)
                }
        }
    }
}
