import Foundation

// MARK: - User Model (for Profile display)
struct User: Identifiable, Codable, Hashable {
    let id: String
    let name: String
    let avatar: String
    let level: String
    let points: Int
}

// MARK: - Extended User Model
struct UserProfile: Identifiable, Codable {
    let id: String
    let username: String
    let email: String?
    let phone: String?
    let avatar: String?
    let nickname: String?
    let gender: String?
    let birthday: String?
    let bio: String?
    let createdAt: String
    
    var displayName: String { nickname ?? username }
    var initials: String {
        let first = username.first.map(String.init) ?? "U"
        return first.uppercased()
    }
}
