import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useCartStore } from '../../store/useCartStore';

export const Checkout: React.FC = () => {
  const navigate = useNavigate();
  const { items, total, clearCart } = useCartStore();
  const [address] = useState({
    receiver: '张三',
    phone: '13800138000',
    province: '北京市',
    city: '北京市',
    district: '朝阳区',
    detail: '某某街道某某小区1号楼',
  });

  const selectedItems = items.filter((i) => i.selected);

  const handleSubmit = () => {
    // Mock 创建订单
    alert('订单创建成功！订单号：' + Date.now());
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
            <div className="space-y-2 text-sm">
              <div className="flex items-center space-x-2">
                <span className="font-medium dark:text-white">{address.receiver}</span>
                <span className="dark:text-gray-300">{address.phone}</span>
              </div>
              <p className="text-gray-600 dark:text-gray-400">
                {address.province} {address.city} {address.district}{' '}
                {address.detail}
              </p>
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
                    className="w-20 h-20 object-cover rounded"
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
