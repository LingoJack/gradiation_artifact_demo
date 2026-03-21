import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Check, Plus } from 'lucide-react';
import { useCartStore } from '../../store/useCartStore';

interface Address {
  id: string;
  receiver: string;
  phone: string;
  province: string;
  city: string;
  district: string;
  detail: string;
  isDefault: boolean;
}

// Mock 地址列表
const mockAddresses: Address[] = [
  {
    id: '1',
    receiver: '张三',
    phone: '13800138000',
    province: '北京市',
    city: '北京市',
    district: '朝阳区',
    detail: '某某街道某某小区1号楼101室',
    isDefault: true,
  },
  {
    id: '2',
    receiver: '李四',
    phone: '13900139000',
    province: '上海市',
    city: '上海市',
    district: '浦东新区',
    detail: '某某路某某大厦A座2001室',
    isDefault: false,
  },
  {
    id: '3',
    receiver: '王五',
    phone: '13700137000',
    province: '广东省',
    city: '深圳市',
    district: '南山区',
    detail: '科技园某某大厦3楼',
    isDefault: false,
  },
];

export const Checkout: React.FC = () => {
  const navigate = useNavigate();
  const { items, total, clearCart } = useCartStore();
  const [addresses] = useState<Address[]>(mockAddresses);
  const [selectedAddressId, setSelectedAddressId] = useState<string>(
    mockAddresses.find((a) => a.isDefault)?.id || mockAddresses[0].id
  );

  const selectedItems = items.filter((i) => i.selected);
  const selectedAddress = addresses.find((a) => a.id === selectedAddressId);

  const handleSubmit = () => {
    // Mock 创建订单 - 模拟后端返回的订单号
    const mockOrderId = `ORD${new Date().toISOString().slice(0,10).replace(/-/g, '')}${Math.floor(Math.random() * 10000).toString().padStart(4, '0')}`;
    alert(`订单创建成功！订单号：${mockOrderId}`);
    clearCart();
    navigate('/orders');
  };

  if (selectedItems.length === 0) {
    navigate('/cart');
    return null;
  }

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6 dark:text-white">确认订单</h1>

      <div className="grid grid-cols-3 gap-8">
        <div className="col-span-2 space-y-6">
          {/* 收货地址 */}
          <div className="glass-card rounded-xl p-6">
            <h2 className="text-lg font-bold mb-4 dark:text-white">收货地址</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {addresses.map((address) => (
                <div
                  key={address.id}
                  onClick={() => setSelectedAddressId(address.id)}
                  className={`relative p-4 rounded-lg border-2 cursor-pointer transition-all ${
                    selectedAddressId === address.id
                      ? 'border-primary bg-primary/5'
                      : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600'
                  }`}
                >
                  {selectedAddressId === address.id && (
                    <div className="absolute top-2 right-2 w-5 h-5 bg-primary rounded-full flex items-center justify-center">
                      <Check className="w-3 h-3 text-white" />
                    </div>
                  )}
                  <div className="space-y-2 text-sm">
                    <div className="flex items-center space-x-2">
                      <span className="font-medium dark:text-white">{address.receiver}</span>
                      <span className="dark:text-gray-300">{address.phone}</span>
                      {address.isDefault && (
                        <span className="px-1.5 py-0.5 bg-primary/10 text-primary text-xs rounded">
                          默认
                        </span>
                      )}
                    </div>
                    <p className="text-gray-600 dark:text-gray-400">
                      {address.province} {address.city} {address.district} {address.detail}
                    </p>
                  </div>
                </div>
              ))}
              {/* 新增地址按钮 */}
              <div
                onClick={() => navigate('/addresses')}
                className="p-4 rounded-lg border-2 border-dashed border-gray-300 dark:border-gray-600 cursor-pointer hover:border-primary dark:hover:border-primary transition-colors flex items-center justify-center"
              >
                <div className="text-center text-gray-400 dark:text-gray-500">
                  <Plus className="w-6 h-6 mx-auto mb-1" />
                  <span className="text-sm">新增地址</span>
                </div>
              </div>
            </div>
          </div>

          {/* 商品清单 */}
          <div className="glass-card rounded-xl p-6">
            <h2 className="text-lg font-bold mb-4 dark:text-white">商品清单</h2>
            <div className="space-y-4">
              {selectedItems.map((item) => (
                <div key={item.id} className="flex space-x-4">
                  <img
                    src={item.product.mainImage}
                    alt={item.product.name}
                    className="w-20 h-20 object-cover rounded bg-gray-100 dark:bg-gray-700"
                    onError={(e) => {
                      const target = e.target as HTMLImageElement;
                      target.src = 'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="80" height="80"%3E%3Crect fill="%23f3f4f6" width="80" height="80"/%3E%3Ctext fill="%239ca3af" x="50%25" y="50%25" text-anchor="middle" dy=".3em" font-size="10"%3E暂无图片%3C/text%3E%3C/svg%3E';
                    }}
                  />
                  <div className="flex-1">
                    <p className="text-sm dark:text-white">{item.product.name}</p>
                    {item.spec && (
                      <p className="text-xs text-gray-500 dark:text-gray-400 mt-1">
                        {item.spec.name}: {item.spec.value}
                      </p>
                    )}
                    <div className="flex items-center justify-between mt-2">
                      <span className="text-primary font-bold">
                        ¥{item.product.price}
                      </span>
                      <span className="text-sm text-gray-500 dark:text-gray-400">
                        x{item.quantity}
                      </span>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>

        {/* 右侧结算信息 */}
        <div className="glass-card rounded-xl p-6 h-fit sticky top-24">
          <h2 className="text-lg font-bold mb-4 dark:text-white">订单信息</h2>
          
          {/* 已选地址摘要 */}
          {selectedAddress && (
            <div className="mb-4 p-3 bg-gray-50 dark:bg-gray-700/50 rounded-lg text-sm">
              <p className="font-medium dark:text-white">
                {selectedAddress.receiver} {selectedAddress.phone}
              </p>
              <p className="text-gray-500 dark:text-gray-400 mt-1 text-xs">
                {selectedAddress.province} {selectedAddress.city} {selectedAddress.district}
              </p>
            </div>
          )}
          
          <div className="space-y-3 text-sm dark:text-gray-300">
            <div className="flex justify-between">
              <span>商品总额：</span>
              <span>¥{total.toFixed(2)}</span>
            </div>
            <div className="flex justify-between">
              <span>运费：</span>
              <span className="text-success">免运费</span>
            </div>
            <div className="flex justify-between">
              <span>优惠：</span>
              <span className="text-error">-¥0.00</span>
            </div>
            <div className="border-t border-gray-200 dark:border-gray-700 pt-3 flex justify-between font-bold text-base">
              <span>实付：</span>
              <span className="text-primary text-xl">¥{total.toFixed(2)}</span>
            </div>
          </div>
          <button
            onClick={handleSubmit}
            className="w-full mt-6 py-3 bg-primary text-white rounded-lg hover:bg-primary-hover"
          >
            提交订单
          </button>
        </div>
      </div>
    </div>
  );
};
