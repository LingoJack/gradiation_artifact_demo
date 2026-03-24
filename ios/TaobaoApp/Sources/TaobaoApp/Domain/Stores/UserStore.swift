import Foundation
import Combine

// MARK: - User Store
class UserStore: ObservableObject {
    static let shared = UserStore()
    
    @Published var user: UserProfile?
    @Published var isLoggedIn: Bool = false
    @Published var token: String?
    
    // Convenience property for ProfileView
    var currentUser: User? {
        guard let profile = user else {
            // Return mock user for preview
            return User(
                id: "u1",
                name: "淘宝用户",
                avatar: "https://picsum.photos/seed/avatar/100/100",
                level: "VIP会员",
                points: 12580
            )
        }
        return User(
            id: profile.id,
            name: profile.nickname ?? profile.username,
            avatar: profile.avatar ?? "",
            level: "VIP会员",
            points: 12580
        )
    }
    
    // Orders for ProfileView
    var orders: [Order] { OrderStore.shared.orders }
    
    private init() {
        // Check for saved login state
        loadUser()
    }
    
    func login(user: UserProfile, token: String) {
        self.user = user
        self.token = token
        self.isLoggedIn = true
        saveUser()
    }
    
    func logout() {
        self.user = nil
        self.token = nil
        self.isLoggedIn = false
        clearUser()
    }
    
    func updateProfile(_ updates: [String: Any]) {
        guard var currentUser = user else { return }
        // Create updated user profile
        let updated = UserProfile(
            id: currentUser.id,
            username: (updates["username"] as? String) ?? currentUser.username,
            email: (updates["email"] as? String) ?? currentUser.email,
            phone: (updates["phone"] as? String) ?? currentUser.phone,
            avatar: (updates["avatar"] as? String) ?? currentUser.avatar,
            nickname: (updates["nickname"] as? String) ?? currentUser.nickname,
            gender: (updates["gender"] as? String) ?? currentUser.gender,
            birthday: (updates["birthday"] as? String) ?? currentUser.birthday,
            bio: (updates["bio"] as? String) ?? currentUser.bio,
            createdAt: currentUser.createdAt
        )
        self.user = updated
        saveUser()
    }
    
    private func saveUser() {
        if let user = user, let encoded = try? JSONEncoder().encode(user) {
            UserDefaults.standard.set(encoded, forKey: "saved_user")
        }
        if let token = token {
            UserDefaults.standard.set(token, forKey: "saved_token")
        }
        UserDefaults.standard.set(isLoggedIn, forKey: "is_logged_in")
    }
    
    private func loadUser() {
        if let data = UserDefaults.standard.data(forKey: "saved_user"),
           let user = try? JSONDecoder().decode(UserProfile.self, from: data) {
            self.user = user
        }
        self.token = UserDefaults.standard.string(forKey: "saved_token")
        self.isLoggedIn = UserDefaults.standard.bool(forKey: "is_logged_in")
    }
    
    private func clearUser() {
        UserDefaults.standard.removeObject(forKey: "saved_user")
        UserDefaults.standard.removeObject(forKey: "saved_token")
        UserDefaults.standard.set(false, forKey: "is_logged_in")
    }
}
