import React, { useState } from 'react';
import { Gift, Clock, Check, AlertCircle } from 'lucide-react';
import { useCouponStore } from '../../store/useCouponStore';
import { showToast } from '../../utils/toast';

export const Coupons: React.FC = () => {
  const { availableCoupons, userCoupons, claimCoupon } = useCouponStore();
  const [activeTab, setActiveTab] = useState<'available' | 'my'>('available');

  // 领取优惠券
  const handleClaim = (couponId: string) => {
    const success = claimCoupon(couponId);
    if (success) {
      showToast('领取成功！', 'success');
    } else {
      showToast('已经领取过该优惠券了', 'warning');
    }
  };

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6 dark:text-white">优惠券</h1>

      {/* Tabs */}
      <div className="flex space-x-4 mb-6">
        <button
          onClick={() => setActiveTab('available')}
          className={`px-6 py-3 rounded-xl font-medium transition-all ${
            activeTab === 'available'
              ? 'bg-primary text-white shadow-lg'
              : 'glass-card dark:text-white hover:bg-gray-100 dark:hover:bg-gray-800'
          }`}
        >
          可领取 ({availableCoupons.length})
        </button>
        <button
          onClick={() => setActiveTab('my')}
          className={`px-6 py-3 rounded-xl font-medium transition-all ${
            activeTab === 'my'
              ? 'bg-primary text-white shadow-lg'
              : 'glass-card dark:text-white hover:bg-gray-100 dark:hover:bg-gray-800'
          }`}
        >
          我的优惠券 ({userCoupons.length})
        </button>
      </div>

      {/* 优惠券列表 */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {activeTab === 'available' ? (
          // 可领取的优惠券
          availableCoupons.map((coupon) => (
            <div
              key={coupon.id}
              className="glass-card rounded-2xl overflow-hidden hover:shadow-xl transition-all"
            >
              <div className="flex">
                {/* 左侧金额 */}
                <div className="bg-gradient-to-br from-orange-500 to-red-500 text-white p-6 flex flex-col items-center justify-center min-w-[120px]">
                  <div className="flex items-baseline">
                    <span className="text-sm">¥</span>
                    <span className="text-3xl font-bold">{coupon.discount}</span>
                  </div>
                  <div className="text-xs opacity-80 mt-1">满{coupon.minSpend}可用</div>
                </div>

                {/* 右侧信息 */}
                <div className="flex-1 p-4 flex flex-col justify-between">
                  <div>
                    <h3 className="font-medium dark:text-white mb-1">{coupon.name}</h3>
                    <p className="text-xs text-gray-500 dark:text-gray-400 mb-2">
                      {coupon.scope === 'all' ? '全场通用' : `限${coupon.scope}品类使用`}
                    </p>
                    <div className="flex items-center text-xs text-gray-400 dark:text-gray-500">
                      <Clock className="w-3 h-3 mr-1" />
                      <span>{coupon.endTime} 到期</span>
                    </div>
                  </div>
                  <button
                    onClick={() => handleClaim(coupon.id)}
                    className="mt-3 w-full py-2 bg-gradient-to-r from-orange-500 to-red-500 text-white rounded-xl text-sm font-medium hover:shadow-lg transition-all flex items-center justify-center space-x-1"
                  >
                    <Gift className="w-4 h-4" />
                    <span>立即领取</span>
                  </button>
                </div>
              </div>
            </div>
          ))
        ) : (
          // 我的优惠券
          userCoupons.map((userCoupon) => {
            const coupon = availableCoupons.find((c) => c.id === userCoupon.couponId);
            if (!coupon) return null;

            const isUsed = userCoupon.status === 'used';
            const isExpired = new Date(coupon.endTime) < new Date();

            return (
              <div
                key={userCoupon.id}
                className={`glass-card rounded-2xl overflow-hidden ${
                  isUsed || isExpired ? 'opacity-60' : ''
                }`}
              >
                <div className="flex">
                  {/* 左侧金额 */}
                  <div
                    className={`p-6 flex flex-col items-center justify-center min-w-[120px] ${
                      isUsed || isExpired
                        ? 'bg-gray-400'
                        : 'bg-gradient-to-br from-orange-500 to-red-500'
                    } text-white`}
                  >
                    <div className="flex items-baseline">
                      <span className="text-sm">¥</span>
                      <span className="text-3xl font-bold">{coupon.discount}</span>
                    </div>
                    <div className="text-xs opacity-80 mt-1">满{coupon.minSpend}可用</div>
                  </div>

                  {/* 右侧信息 */}
                  <div className="flex-1 p-4 flex flex-col justify-between">
                    <div>
                      <div className="flex items-center justify-between mb-1">
                        <h3 className="font-medium dark:text-white">{coupon.name}</h3>
                        {isUsed && (
                          <span className="text-xs text-gray-500 dark:text-gray-400 flex items-center">
                            <Check className="w-3 h-3 mr-1" />
                            已使用
                          </span>
                        )}
                        {isExpired && !isUsed && (
                          <span className="text-xs text-red-500 flex items-center">
                            <AlertCircle className="w-3 h-3 mr-1" />
                            已过期
                          </span>
                        )}
                      </div>
                      <p className="text-xs text-gray-500 dark:text-gray-400 mb-2">
                        {coupon.scope === 'all' ? '全场通用' : `限${coupon.scope}品类使用`}
                      </p>
                      <div className="flex items-center text-xs text-gray-400 dark:text-gray-500">
                        <Clock className="w-3 h-3 mr-1" />
                        <span>{coupon.endTime} 到期</span>
                      </div>
                    </div>
                    {!isUsed && !isExpired && (
                      <button
                        onClick={() => showToast('结算时自动使用优惠券', 'info')}
                        className="mt-3 w-full py-2 border border-primary text-primary dark:text-orange-400 rounded-xl text-sm font-medium hover:bg-primary hover:text-white transition-all"
                      >
                        去使用
                      </button>
                    )}
                  </div>
                </div>
              </div>
            );
          })
        )}
      </div>

      {/* 空状态 */}
      {activeTab === 'available' && availableCoupons.length === 0 && (
        <div className="glass-card rounded-xl p-12 text-center">
          <p className="text-gray-500 dark:text-gray-400">暂无可领取的优惠券</p>
        </div>
      )}
      {activeTab === 'my' && userCoupons.length === 0 && (
        <div className="glass-card rounded-xl p-12 text-center">
          <p className="text-gray-500 dark:text-gray-400">暂无优惠券，快去领取吧！</p>
        </div>
      )}
    </div>
  );
};
