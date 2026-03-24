import SwiftUI

struct LoginView: View {
    @Environment(\.dismiss) private var dismiss
    @StateObject private var userStore = UserStore.shared
    
    @State private var phone: String = ""
    @State private var code: String = ""
    @State private var isCodeSent = false
    @State private var countdown = 0
    @State private var isLoggingIn = false
    
    private let timer = Timer.publish(every: 1, on: .main, in: .common).autoconnect()
    
    var body: some View {
        NavigationStack {
            VStack(spacing: .tbSpacing32) {
                // Logo
                VStack(spacing: .tbSpacing16) {
                    Image(systemName: "bag.fill")
                        .font(.system(size: 60))
                        .foregroundColor(.tbOrange)
                    
                    Text("淘宝")
                        .font(.system(size: 32, weight: .bold))
                        .foregroundColor(.tbTextPrimary)
                }
                .padding(.top, 60)
                
                // Login Form
                VStack(spacing: .tbSpacing16) {
                    // Phone Input
                    VStack(alignment: .leading, spacing: .tbSpacing8) {
                        Text("手机号")
                            .font(.tbCaption)
                            .foregroundColor(.tbTextSecondary)
                        
                        HStack(spacing: .tbSpacing8) {
                            Text("+86")
                                .font(.tbBody)
                                .foregroundColor(.tbTextPrimary)
                            
                            TextField("请输入手机号", text: $phone)
                                .font(.tbBody)
                                .keyboardType(.phonePad)
                                .textContentType(.telephoneNumber)
                        }
                        .padding(.tbSpacing12)
                        .background(Color.tbBackground)
                        .cornerRadius(.tbRadius8)
                    }
                    
                    // Code Input
                    VStack(alignment: .leading, spacing: .tbSpacing8) {
                        Text("验证码")
                            .font(.tbCaption)
                            .foregroundColor(.tbTextSecondary)
                        
                        HStack(spacing: .tbSpacing8) {
                            TextField("请输入验证码", text: $code)
                                .font(.tbBody)
                                .keyboardType(.numberPad)
                                .textContentType(.oneTimeCode)
                            
                            Button(action: sendCode) {
                                Text(countdown > 0 ? "\(countdown)s" : "获取验证码")
                                    .font(.tbCaption)
                                    .foregroundColor(countdown > 0 ? .tbTextTertiary : .tbOrange)
                            }
                            .disabled(countdown > 0 || phone.count < 11)
                        }
                        .padding(.tbSpacing12)
                        .background(Color.tbBackground)
                        .cornerRadius(.tbRadius8)
                    }
                }
                .padding(.horizontal, .tbSpacing24)
                
                // Login Button
                Button(action: login) {
                    Text(isLoggingIn ? "登录中..." : "登录")
                        .font(.tbBodyBold)
                        .foregroundColor(.white)
                        .frame(maxWidth: .infinity)
                        .padding(.vertical, .tbSpacing16)
                        .background(phone.count == 11 && code.count == 6 ? Color.tbOrange : Color.tbOrange.opacity(0.5))
                        .cornerRadius(24)
                }
                .disabled(phone.count != 11 || code.count != 6 || isLoggingIn)
                .padding(.horizontal, .tbSpacing24)
                
                // Other Login Options
                VStack(spacing: .tbSpacing16) {
                    Text("其他登录方式")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextTertiary)
                    
                    HStack(spacing: .tbSpacing32) {
                        Button(action: { loginWithSMS() }) {
                            Image(systemName: "message.circle.fill")
                                .font(.system(size: 40))
                                .foregroundColor(.green)
                        }
                        
                        Button(action: { loginWithAlipay() }) {
                            Image(systemName: "a.circle.fill")
                                .font(.system(size: 40))
                                .foregroundColor(.blue)
                        }
                    }
                }
                
                Spacer()
                
                // Terms
                VStack(spacing: .tbSpacing8) {
                    Text("登录即表示同意")
                        .font(.tbCaption2)
                        .foregroundColor(.tbTextTertiary)
                    
                    HStack(spacing: 4) {
                        Button("《用户协议》") {
                            showTerms(type: "用户协议")
                        }
                        .foregroundColor(.tbOrange)
                        Text("和")
                            .foregroundColor(.tbTextTertiary)
                        Button("《隐私政策》") {
                            showTerms(type: "隐私政策")
                        }
                        .foregroundColor(.tbOrange)
                    }
                    .font(.tbCaption2)
                }
                .padding(.bottom, 20)
            }
            .background(Color.white)
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .navigationBarLeading) {
                    Button(action: { dismiss() }) {
                        Image(systemName: "xmark")
                            .foregroundColor(.tbTextSecondary)
                    }
                }
            }
            .onReceive(timer) { _ in
                if countdown > 0 {
                    countdown -= 1
                }
            }
        }
    }
    
    private func sendCode() {
        guard phone.count == 11 else { return }
        isCodeSent = true
        countdown = 60
    }
    
    private func login() {
        guard phone.count == 11, code.count == 6 else { return }
        
        isLoggingIn = true
        
        // Simulate login
        DispatchQueue.main.asyncAfter(deadline: .now() + 1) {
            let user = UserProfile(
                id: "u1",
                username: "user_\(phone.suffix(4))",
                email: nil,
                phone: phone,
                avatar: nil,
                nickname: "用户\(phone.suffix(4))",
                gender: nil,
                birthday: nil,
                bio: nil,
                createdAt: ISO8601DateFormatter().string(from: Date())
            )
            userStore.login(user: user, token: UUID().uuidString)
            isLoggingIn = false
            dismiss()
        }
    }
    
    private func loginWithSMS() {
        // 短信登录 - 直接使用当前表单
        phone = ""
        code = ""
    }
    
    private func loginWithAlipay() {
        // 支付宝登录 - 模拟授权
        DispatchQueue.main.asyncAfter(deadline: .now() + 0.5) {
            let user = UserProfile(
                id: "u1",
                username: "alipay_user",
                email: nil,
                phone: "138****8888",
                avatar: nil,
                nickname: "支付宝用户",
                gender: nil,
                birthday: nil,
                bio: nil,
                createdAt: ISO8601DateFormatter().string(from: Date())
            )
            userStore.login(user: user, token: UUID().uuidString)
            dismiss()
        }
    }
    
    private func showTerms(type: String) {
        // 显示协议 - 实际应用中应跳转到协议页面
        print("显示\(type)")
    }
}

#Preview {
    LoginView()
}
