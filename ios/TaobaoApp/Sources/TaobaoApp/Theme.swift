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

// MARK: - Font Theme
extension Font {
    static let tbLargeTitle = Font.system(size: 28, weight: .bold)
    static let tbTitle = Font.system(size: 20, weight: .semibold)
    static let tbTitle2 = Font.system(size: 18, weight: .semibold)
    static let tbTitle3 = Font.system(size: 16, weight: .medium)
    static let tbBody = Font.system(size: 14, weight: .regular)
    static let tbBodyBold = Font.system(size: 14, weight: .semibold)
    static let tbCaption = Font.system(size: 12, weight: .regular)
    static let tbCaption2 = Font.system(size: 10, weight: .regular)
    
    static let tbPriceLarge = Font.system(size: 24, weight: .bold)
    static let tbPrice = Font.system(size: 16, weight: .bold)
    static let tbPriceSmall = Font.system(size: 14, weight: .semibold)
}

// MARK: - Spacing
extension CGFloat {
    static let tbSpacing4: CGFloat = 4
    static let tbSpacing8: CGFloat = 8
    static let tbSpacing12: CGFloat = 12
    static let tbSpacing16: CGFloat = 16
    static let tbSpacing20: CGFloat = 20
    static let tbSpacing24: CGFloat = 24
}

// MARK: - Corner Radius
extension CGFloat {
    static let tbRadius4: CGFloat = 4
    static let tbRadius8: CGFloat = 8
    static let tbRadius12: CGFloat = 12
    static let tbRadius16: CGFloat = 16
}
