import SwiftUI

// MARK: - Color Theme
extension Color {
    // 主色调 - 淘宝橙
    static let tbOrange = Color(hex: "FF5000")
    static let tbOrangeLight = Color(hex: "FF7538")
    
    // 背景色
    static let tbBackground = Color(hex: "F5F5F5")
    static let tbCardBackground = Color.white
    
    // 文字颜色
    static let tbTextPrimary = Color(hex: "1A1A1A")
    static let tbTextSecondary = Color(hex: "666666")
    static let tbTextTertiary = Color(hex: "999999")
    static let tbTextPlaceholder = Color(hex: "CCCCCC")
    
    // 价格颜色
    static let tbPrice = Color(hex: "FF2D2D")
    
    // 状态颜色
    static let tbGreen = Color(hex: "00C853")
    static let tbRed = Color(hex: "FF2D2D")
    static let tbBlue = Color(hex: "1976D2")
    
    // 分割线
    static let tbDivider = Color(hex: "EEEEEE")
    
    // 标签背景
    static let tbTagBackground = Color(hex: "FFF0E6")
    
    init(hex: String) {
        let hex = hex.trimmingCharacters(in: CharacterSet.alphanumerics.inverted)
        var int: UInt64 = 0
        Scanner(string: hex).scanHexInt64(&int)
        let a, r, g, b: UInt64
        switch hex.count {
        case 3: // RGB (12-bit)
            (a, r, g, b) = (255, (int >> 8) * 17, (int >> 4 & 0xF) * 17, (int & 0xF) * 17)
        case 6: // RGB (24-bit)
            (a, r, g, b) = (255, int >> 16, int >> 8 & 0xFF, int & 0xFF)
        case 8: // ARGB (32-bit)
            (a, r, g, b) = (int >> 24, int >> 16 & 0xFF, int >> 8 & 0xFF, int & 0xFF)
        default:
            (a, r, g, b) = (1, 1, 1, 0)
        }
        
        self.init(
            .sRGB,
            red: Double(r) / 255,
            green: Double(g) / 255,
            blue: Double(b) / 255,
            opacity: Double(a) / 255
        )
    }
}
