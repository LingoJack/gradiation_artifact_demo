import React, { useState } from 'react';
import { ORDER_STATUS_TEXT, type OrderStatus } from '../../types/order';

// Mock 订单数据
const mockOrders = [
  {
    id: '1',
    orderNo: '202401200001',
    totalAmount: 258,
    payAmount: 258,
    status: 'pending' as OrderStatus,
    receiverName: '张三',
    receiverPhone: '13800138000',
    receiverAddress: '北京市朝阳区某某街道',
    items: [
      {
        id: 'i1',
        orderId: '1',
        productId: '1',
        productName: '时尚休闲连帽卫衣 男士秋季新款',
        productImage: 'https://via.placeholder.com/100x100?text=卫衣',
        specName: '黑色',
        price: 129,
        quantity: 2,
      },
    ],
    createdAt: '2024-01-20 10:30:00',
  },
  {
    id: '2',
    orderNo: '202401190002',
    totalAmount: 9999,
    payAmount: 9999,
    status: 'shipped' as OrderStatus,
    receiverName: '张三',
    receiverPhone: '13800138000',
    receiverAddress: '北京市朝阳区某某街道',
    items: [
      {
        id: 'i2',
        orderId: '2',
        productId: '2',
        productName: 'Apple iPhone 15 Pro Max 256GB',
        productImage: 'https://via.placeholder.com/100x100?text=iPhone',
        specName: '深空黑',
        price: 9999,
        quantity: 1,
      },
    ],
    createdAt: '2024-01-19 15:20:00',
    paidAt: '2024-01-19 15:21:00',
    shippedAt: '2024-01-20 09:00:00',
  },
];

export const OrderList: React.FC = () => {
  const [activeTab, setActiveTab] = useState<OrderStatus | 'all'>('all');

  const tabs: Array<{ key: OrderStatus | 'all'; label: string }> = [
    { key: 'all', label: '全部订单' },
    { key: 'pending', label: '待付款' },
    { key: 'paid', label: '待发货' },
    { key: 'shipped', label: '待收货' },
    { key: 'completed', label: '已完成' },
  ];

  const filteredOrders =
    activeTab === 'all'
      ? mockOrders
      : mockOrders.filter((o) => o.status === activeTab);

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6">我的订单</h1>

      {/* Tabs */}
      <div className="bg-white rounded-lg mb-6">
        <div className="flex border-b">
          {tabs.map((tab) => (
            <button
              key={tab.key}
              onClick={() => setActiveTab(tab.key)}
              className={`px-6 py-4 font-medium ${
                activeTab === tab.key
                  ? 'text-primary border-b-2 border-primary'
                  : 'text-gray-600 hover:text-gray-900'
              }`}
            >
              {tab.label}
            </button>
          ))}
        </div>
      </div>

      {/* 订单列表 */}
      <div className="space-y-4">
        {filteredOrders.length === 0 ? (
          <div className="bg-white rounded-lg p-12 text-center">
            <p className="text-gray-500">暂无订单</p>
          </div>
        ) : (
          filteredOrders.map((order) => (
            <div key={order.id} className="bg-white rounded-lg overflow-hidden">
              {/* 订单头部 */}
              <div className="bg-gray-50 px-6 py-3 flex items-center justify-between text-sm">
                <div className="flex items-center space-x-6">
                  <span>订单号：{order.orderNo}</span>
                  <span>{order.createdAt}</span>
                </div>
                <span className="font-bold text-primary">
                  {ORDER_STATUS_TEXT[order.status]}
                </span>
              </div>

              {/* 商品列表 */}
              <div className="p-6">
                {order.items.map((item) => (
                  <div key={item.id} className="flex space-x-4">
                    <img
                      src={item.productImage}
                      alt={item.productName}
                      className="w-20 h-20 object-cover rounded"
                    />
                    <div className="flex-1">
                      <p className="text-sm">{item.productName}</p>
                      {item.specName && (
                        <p className="text-xs text-gray-500 mt-1">
                          规格：{item.specName}
                        </p>
                      )}
                      <div className="flex items-center justify-between mt-2">
                        <span className="text-primary font-bold">
                          ¥{item.price}
                        </span>
                        <span className="text-sm text-gray-500">
                          x{item.quantity}
                        </span>
                      </div>
                    </div>
                  </div>
                ))}

                <div className="flex items-center justify-between mt-4 pt-4 border-t">
                  <div className="text-sm">
                    共 {order.items.length} 件商品 合计：
                    <span className="text-lg font-bold text-primary">
                      ¥{order.payAmount}
                    </span>
                  </div>
                  <div className="flex space-x-2">
                    {order.status === 'pending' && (
                      <>
                        <button className="px-4 py-2 border rounded hover:bg-gray-50">
                          取消订单
                        </button>
                        <button className="px-4 py-2 bg-primary text-white rounded hover:bg-primary-hover">
                          立即支付
                        </button>
                      </>
                    )}
                    {order.status === 'shipped' && (
                      <button className="px-4 py-2 bg-primary text-white rounded hover:bg-primary-hover">
                        确认收货
                      </button>
                    )}
                    {order.status === 'completed' && (
                      <button className="px-4 py-2 border rounded hover:bg-gray-50">
                        评价
                      </button>
                    )}
                  </div>
                </div>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
};
