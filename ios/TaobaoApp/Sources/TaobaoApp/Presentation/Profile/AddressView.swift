import SwiftUI

struct AddressView: View {
    @StateObject private var addressStore = AddressStore.shared
    @State private var showAddAddress = false
    
    var body: some View {
        ScrollView {
            LazyVStack(spacing: .tbSpacing12) {
                ForEach(addressStore.addresses) { address in
                    AddressCardView(address: address)
                }
            }
            .padding(.tbSpacing12)
        }
        .background(Color.tbBackground)
        .navigationTitle("地址管理")
        .navigationBarTitleDisplayMode(.inline)
        .toolbar {
            ToolbarItem(placement: .navigationBarTrailing) {
                Button(action: { showAddAddress = true }) {
                    Image(systemName: "plus")
                }
            }
        }
        .sheet(isPresented: $showAddAddress) {
            AddAddressView()
        }
    }
}

struct AddressCardView: View {
    let address: Address
    @StateObject private var addressStore = AddressStore.shared
    
    var body: some View {
        VStack(alignment: .leading, spacing: .tbSpacing12) {
            HStack {
                Text(address.name)
                    .font(.tbBodyBold)
                    .foregroundColor(.tbTextPrimary)
                
                Text(address.phone)
                    .font(.tbBody)
                    .foregroundColor(.tbTextSecondary)
                
                if address.isDefault {
                    Text("默认")
                        .font(.tbCaption2)
                        .foregroundColor(.tbOrange)
                        .padding(.horizontal, 4)
                        .padding(.vertical, 2)
                        .background(Color.tbOrange.opacity(0.1))
                        .cornerRadius(2)
                }
                
                Spacer()
            }
            
            Text(address.fullAddress)
                .font(.tbBody)
                .foregroundColor(.tbTextSecondary)
            
            HStack {
                Button(action: { addressStore.setDefault(address.id) }) {
                    HStack(spacing: 4) {
                        Image(systemName: address.isDefault ? "checkmark.circle.fill" : "circle")
                        Text("设为默认")
                    }
                    .font(.tbCaption)
                    .foregroundColor(address.isDefault ? .tbOrange : .tbTextTertiary)
                }
                
                Spacer()
                
                Button(action: { }) {
                    Text("编辑")
                        .font(.tbCaption)
                        .foregroundColor(.tbTextSecondary)
                }
                
                Button(action: { addressStore.deleteAddress(address.id) }) {
                    Text("删除")
                        .font(.tbCaption)
                        .foregroundColor(.tbPrice)
                }
            }
        }
        .padding(.tbSpacing16)
        .background(Color.white)
        .cornerRadius(.tbRadius8)
    }
}

struct AddAddressView: View {
    @Environment(\.dismiss) private var dismiss
    @StateObject private var addressStore = AddressStore.shared
    
    @State private var name: String = ""
    @State private var phone: String = ""
    @State private var province: String = ""
    @State private var city: String = ""
    @State private var district: String = ""
    @State private var detail: String = ""
    @State private var isDefault: Bool = false
    
    var body: some View {
        NavigationStack {
            Form {
                Section("收货人信息") {
                    TextField("姓名", text: $name)
                    TextField("手机号", text: $phone)
                        .keyboardType(.phonePad)
                }
                
                Section("地址信息") {
                    TextField("省/市/区", text: $province)
                    TextField("详细地址", text: $detail)
                }
                
                Section {
                    Toggle("设为默认地址", isOn: $isDefault)
                }
            }
            .navigationTitle("新增地址")
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .navigationBarLeading) {
                    Button("取消") { dismiss() }
                }
                ToolbarItem(placement: .navigationBarTrailing) {
                    Button("保存") {
                        saveAddress()
                    }
                    .disabled(name.isEmpty || phone.isEmpty || detail.isEmpty)
                }
            }
        }
    }
    
    private func saveAddress() {
        let address = Address(
            id: UUID().uuidString,
            userId: UserStore.shared.user?.id ?? "guest",
            name: name,
            phone: phone,
            province: province,
            city: city,
            district: district,
            detail: detail,
            isDefault: isDefault
        )
        addressStore.addAddress(address)
        dismiss()
    }
}

#Preview {
    NavigationStack {
        AddressView()
    }
}
