import Foundation
import Combine

// MARK: - Address Store
class AddressStore: ObservableObject {
    static let shared = AddressStore()
    
    @Published var addresses: [Address] = []
    
    private init() {
        loadMockData()
    }
    
    var defaultAddress: Address? { addresses.first { $0.isDefault } }
    
    func addAddress(_ address: Address) {
        // If new address is default, remove default from others
        if address.isDefault {
            for i in 0..<addresses.count {
                if addresses[i].isDefault {
                    addresses[i] = Address(
                        id: addresses[i].id,
                        userId: addresses[i].userId,
                        name: addresses[i].name,
                        phone: addresses[i].phone,
                        province: addresses[i].province,
                        city: addresses[i].city,
                        district: addresses[i].district,
                        detail: addresses[i].detail,
                        isDefault: false
                    )
                }
            }
        }
        addresses.append(address)
    }
    
    func updateAddress(_ address: Address) {
        guard let index = addresses.firstIndex(where: { $0.id == address.id }) else { return }
        // If updated address is default, remove default from others
        if address.isDefault {
            for i in 0..<addresses.count where i != index {
                if addresses[i].isDefault {
                    addresses[i] = Address(
                        id: addresses[i].id,
                        userId: addresses[i].userId,
                        name: addresses[i].name,
                        phone: addresses[i].phone,
                        province: addresses[i].province,
                        city: addresses[i].city,
                        district: addresses[i].district,
                        detail: addresses[i].detail,
                        isDefault: false
                    )
                }
            }
        }
        addresses[index] = address
    }
    
    func deleteAddress(_ id: String) {
        addresses.removeAll { $0.id == id }
    }
    
    func setDefault(_ id: String) {
        for i in 0..<addresses.count {
            addresses[i] = Address(
                id: addresses[i].id,
                userId: addresses[i].userId,
                name: addresses[i].name,
                phone: addresses[i].phone,
                province: addresses[i].province,
                city: addresses[i].city,
                district: addresses[i].district,
                detail: addresses[i].detail,
                isDefault: addresses[i].id == id
            )
        }
    }
    
    private func loadMockData() {
        addresses = [
            Address(
                id: "a1",
                userId: "u1",
                name: "张三",
                phone: "13800138000",
                province: "浙江省",
                city: "杭州市",
                district: "西湖区",
                detail: "文三路 123 号",
                isDefault: true
            ),
            Address(
                id: "a2",
                userId: "u1",
                name: "李四",
                phone: "13900139000",
                province: "上海市",
                city: "上海市",
                district: "浦东新区",
                detail: "陆家嘴环路 1000 号",
                isDefault: false
            )
        ]
    }
}
