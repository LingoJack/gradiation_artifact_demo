import React, { useState } from 'react';
import { Ticket, Gift, Coins, Clock } from 'lucide-react';
import { useCouponStore, type Coupon } from '../../store/useCouponStore';
import { useSpotlight } from '../../hooks/useSpotlight';

type TabType = 'available' | 'used' | 'expired';

export const Coupons: React.FC = () => {
  const { coupons, points, redPacket } = useCouponStore();
  const [activeTab, setActiveTab] = useState<TabType>('available');
  const cardSpotlight = useSpotlight();

  const filteredCoupons = coupons.filter((c) => {
    if (activeTab === 'available') return c.status === 'available';
    if (activeTab === 'used') return c.status === 'used';
    return c.status === 'expired';
  });

  const getCouponValue = (coupon: Coupon) => {
    if (coupon.type === 'fixed') {
      return `¥${coupon.value}`;
    }
    return `${100 - coupon.value}折`;
  };

  const getCouponCondition = (coupon: Coupon) => {
    return `满${coupon.minAmount}元可用`;
  };

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6 dark:text-white">我的优惠</h1>

      {/* 资产概览 */}
      <div
        ref={cardSpotlight.ref as React.RefObject<HTMLDivElement>}
        className="glass-liquid rounded-xl p-6 mb-6 relative overflow-hidden"
        style={cardSpotlight.spotlightStyle}
        {...cardSpotlight.handlers}
      >
        <div className="grid grid-cols-3 gap-6">
          <div className="text-center">
            <div className="w-12 h-12 bg-primary/10 rounded-full flex items-center justify-center mx-auto mb-2">
              <Ticket className="w-6 h-6 text-primary" />
            </div>
            <p className="text-2xl font-bold text-primary">
              {coupons.filter((c) => c.status === 'available').length}
            </p>
            <p className="text-sm text-gray-500 dark:text-gray-400">优惠券</p>
          </div>
          <div className="text-center">
            <div className="w-12 h-12 bg-red-100 dark:bg-red-900/30 rounded-full flex items-center justify-center mx-auto mb-2">
              <Gift className="w-6 h-6 text-red-500" />
            </div>
            <p className="text-2xl font-bold text-red-500">¥{redPacket}</p>
            <p className="text-sm text-gray-500 dark:text-gray-400">红包</p>
          </div>
          <div className="text-center">
            <div className="w-12 h-12 bg-amber-100 dark:bg-amber-900/30 rounded-full flex items-center justify-center mx-auto mb-2">
              <Coins className="w-6 h-6 text-amber-500" />
            </div>
            <p className="text-2xl font-bold text-amber-500">{points}</p>
            <p className="text-sm text-gray-500 dark:text-gray-400">积分</p>
          </div>
        </div>
      </div>

      {/* 优惠券列表 */}
      <div className="glass-card rounded-xl overflow-hidden">
        {/* 标签切换 */}
        <div className="flex border-b border-gray-200 dark:border-gray-700">
          <button
            onClick={() => setActiveTab('available')}
            className={`flex-1 py-4 text-center font-medium transition-colors ${
              activeTab === 'available'
                ? 'text-primary border-b-2 border-primary'
                : 'text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-300'
            }`}
          >
            可使用 ({coupons.filter((c) => c.status === 'available').length})
          </button>
          <button
            onClick={() => setActiveTab('used')}
            className={`flex-1 py-4 text-center font-medium transition-colors ${
              activeTab === 'used'
                ? 'text-primary border-b-2 border-primary'
                : 'text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-300'
            }`}
          >
            已使用 ({coupons.filter((c) => c.status === 'used').length})
          </button>
          <button
            onClick={() => setActiveTab('expired')}
            className={`flex-1 py-4 text-center font-medium transition-colors ${
              activeTab === 'expired'
                ? 'text-primary border-b-2 border-primary'
                : 'text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-300'
            }`}
          >
            已过期 ({coupons.filter((c) => c.status === 'expired').length})
          </button>
        </div>

        {/* 优惠券卡片列表 */}
        <div className="p-6">
          {filteredCoupons.length === 0 ? (
            <div className="text-center py-12">
              <Ticket className="w-16 h-16 text-gray-300 dark:text-gray-600 mx-auto mb-4" />
              <p className="text-gray-500 dark:text-gray-400">暂无优惠券</p>
            </div>
          ) : (
            <div className="space-y-4">
              {filteredCoupons.map((coupon) => (
                <div
                  key={coupon.id}
                  className={`relative flex overflow-hidden rounded-xl border ${
                    coupon.status === 'available'
                      ? 'border-primary/30 bg-gradient-to-r from-primary/5 to-transparent'
                      : 'border-gray-200 dark:border-gray-700 bg-gray-50 dark:bg-gray-800/50'
                  }`}
                >
                  {/* 左侧金额 */}
                  <div
                    className={`w-28 flex-shrink-0 flex flex-col items-center justify-center py-6 ${
                      coupon.status === 'available'
                        ? 'bg-primary/10'
                        : 'bg-gray-100 dark:bg-gray-700'
                    }`}
                  >
                    <span
                      className={`text-3xl font-bold ${
                        coupon.status === 'available' ? 'text-primary' : 'text-gray-400'
                      }`}
                    >
                      {getCouponValue(coupon)}
                    </span>
                    <span
                      className={`text-xs mt-1 ${
                        coupon.status === 'available'
                          ? 'text-primary/70'
                          : 'text-gray-400'
                      }`}
                    >
                      {getCouponCondition(coupon)}
                    </span>
                  </div>

                  {/* 右侧信息 */}
                  <div className="flex-1 p-4 flex items-center justify-between">
                    <div>
                      <p
                        className={`font-medium ${
                          coupon.status === 'available'
                            ? 'dark:text-white'
                            : 'text-gray-500 dark:text-gray-400'
                        }`}
                      >
                        {coupon.name}
                      </p>
                      <div
                        className={`flex items-center text-xs mt-2 ${
                          coupon.status === 'available'
                            ? 'text-gray-500 dark:text-gray-400'
                            : 'text-gray-400'
                        }`}
                      >
                        <Clock className="w-3 h-3 mr-1" />
                        <span>
                          {coupon.startTime} - {coupon.endTime}
                        </span>
                      </div>
                    </div>
                    {coupon.status === 'available' && (
                      <button className="px-4 py-2 bg-primary text-white text-sm rounded-lg hover:bg-primary-hover">
                        去使用
                      </button>
                    )}
                    {coupon.status === 'used' && (
                      <span className="px-4 py-2 bg-gray-200 dark:bg-gray-600 text-gray-500 dark:text-gray-400 text-sm rounded-lg">
                        已使用
                      </span>
                    )}
                    {coupon.status === 'expired' && (
                      <span className="px-4 py-2 bg-gray-200 dark:bg-gray-600 text-gray-500 dark:text-gray-400 text-sm rounded-lg">
                        已过期
                      </span>
                    )}
                  </div>

                  {/* 装饰圆点 */}
                  <div
                    className={`absolute left-28 top-0 w-4 h-4 rounded-full -translate-x-1/2 -translate-y-1/2 ${
                      coupon.status === 'available'
                        ? 'bg-gray-50 dark:bg-gray-800'
                        : 'bg-gray-50 dark:bg-gray-700'
                    }`}
                  />
                  <div
                    className={`absolute left-28 bottom-0 w-4 h-4 rounded-full -translate-x-1/2 translate-y-1/2 ${
                      coupon.status === 'available'
                        ? 'bg-gray-50 dark:bg-gray-800'
                        : 'bg-gray-50 dark:bg-gray-700'
                    }`}
                  />
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
    </div>
  );
};
