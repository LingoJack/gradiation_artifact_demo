import React from 'react';

export const Footer: React.FC = () => {
  return (
    <footer className="bg-gray-100 dark:bg-gray-900 border-t border-gray-200 dark:border-gray-800 mt-16">
      <div className="container py-10">
        {/* 主要链接区域 */}
        <div className="grid grid-cols-5 gap-8 mb-8">
          <div>
            <h3 className="text-base font-bold text-gray-800 dark:text-gray-200 mb-4">购物指南</h3>
            <ul className="space-y-2.5 text-sm text-gray-600 dark:text-gray-400">
              <li className="hover:text-orange-500 cursor-pointer transition">购物流程</li>
              <li className="hover:text-orange-500 cursor-pointer transition">会员介绍</li>
              <li className="hover:text-orange-500 cursor-pointer transition">常见问题</li>
              <li className="hover:text-orange-500 cursor-pointer transition">联系客服</li>
            </ul>
          </div>
          <div>
            <h3 className="text-base font-bold text-gray-800 dark:text-gray-200 mb-4">配送方式</h3>
            <ul className="space-y-2.5 text-sm text-gray-600 dark:text-gray-400">
              <li className="hover:text-orange-500 cursor-pointer transition">上门自提</li>
              <li className="hover:text-orange-500 cursor-pointer transition">配送服务</li>
              <li className="hover:text-orange-500 cursor-pointer transition">配送范围</li>
              <li className="hover:text-orange-500 cursor-pointer transition">配送时效</li>
            </ul>
          </div>
          <div>
            <h3 className="text-base font-bold text-gray-800 dark:text-gray-200 mb-4">支付方式</h3>
            <ul className="space-y-2.5 text-sm text-gray-600 dark:text-gray-400">
              <li className="hover:text-orange-500 cursor-pointer transition">货到付款</li>
              <li className="hover:text-orange-500 cursor-pointer transition">在线支付</li>
              <li className="hover:text-orange-500 cursor-pointer transition">分期付款</li>
              <li className="hover:text-orange-500 cursor-pointer transition">公司转账</li>
            </ul>
          </div>
          <div>
            <h3 className="text-base font-bold text-gray-800 dark:text-gray-200 mb-4">售后服务</h3>
            <ul className="space-y-2.5 text-sm text-gray-600 dark:text-gray-400">
              <li className="hover:text-orange-500 cursor-pointer transition">退换货政策</li>
              <li className="hover:text-orange-500 cursor-pointer transition">退换货流程</li>
              <li className="hover:text-orange-500 cursor-pointer transition">价格保护</li>
              <li className="hover:text-orange-500 cursor-pointer transition">退款说明</li>
            </ul>
          </div>
          <div>
            <h3 className="text-base font-bold text-gray-800 dark:text-gray-200 mb-4">关于我们</h3>
            <ul className="space-y-2.5 text-sm text-gray-600 dark:text-gray-400">
              <li className="hover:text-orange-500 cursor-pointer transition">关于淘宝</li>
              <li className="hover:text-orange-500 cursor-pointer transition">联系我们</li>
              <li className="hover:text-orange-500 cursor-pointer transition">加入我们</li>
              <li className="hover:text-orange-500 cursor-pointer transition">合作招商</li>
            </ul>
          </div>
        </div>

        {/* 分隔线 */}
        <div className="border-t border-gray-200 dark:border-gray-800 pt-6 mb-6">
          {/* 特色服务 */}
          <div className="flex items-center justify-center space-x-12 text-sm text-gray-600 dark:text-gray-400 mb-6">
            <div className="flex items-center space-x-2">
              <svg className="w-5 h-5 text-orange-500" fill="currentColor" viewBox="0 0 20 20">
                <path d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" />
              </svg>
              <span>正品保障</span>
            </div>
            <div className="flex items-center space-x-2">
              <svg className="w-5 h-5 text-orange-500" fill="currentColor" viewBox="0 0 20 20">
                <path d="M8 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zM15 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0z" />
                <path d="M3 4a1 1 0 00-1 1v10a1 1 0 001 1h1.05a2.5 2.5 0 014.9 0H10a1 1 0 001-1V5a1 1 0 00-1-1H3zM14 7a1 1 0 00-1 1v6.05A2.5 2.5 0 0115.95 16H17a1 1 0 001-1v-5a1 1 0 00-.293-.707l-2-2A1 1 0 0015 7h-1z" />
              </svg>
              <span>极速发货</span>
            </div>
            <div className="flex items-center space-x-2">
              <svg className="w-5 h-5 text-orange-500" fill="currentColor" viewBox="0 0 20 20">
                <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-12a1 1 0 10-2 0v4a1 1 0 00.293.707l2.828 2.829a1 1 0 101.415-1.415L11 9.586V6z" clipRule="evenodd" />
              </svg>
              <span>七天退换</span>
            </div>
            <div className="flex items-center space-x-2">
              <svg className="w-5 h-5 text-orange-500" fill="currentColor" viewBox="0 0 20 20">
                <path fillRule="evenodd" d="M2.166 4.999A11.954 11.954 0 0010 1.944 11.954 11.954 0 0017.834 5c.11.65.166 1.32.166 2.001 0 5.225-3.34 9.67-8 11.317C5.34 16.67 2 12.225 2 7c0-.682.057-1.35.166-2.001zm11.541 3.708a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clipRule="evenodd" />
              </svg>
              <span>安全支付</span>
            </div>
          </div>
        </div>

        {/* 版权信息 */}
        <div className="text-center">
          <p className="text-xs text-gray-400 dark:text-gray-500 mb-2">
            客服热线：400-123-4567 | 工作时间：周一至周日 9:00-21:00
          </p>
          <p className="text-xs text-gray-400 dark:text-gray-500">
            © 2024 淘宝克隆版 - 仅供学习使用
          </p>
        </div>
      </div>
    </footer>
  );
};
