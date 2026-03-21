import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { Trash2 } from 'lucide-react';
import { useCartStore } from '../../store/useCartStore';
import { useSpotlight } from '../../hooks/useSpotlight';

export const Cart: React.FC = () => {
  const navigate = useNavigate();
  const { items, total, removeItem, updateQuantity, toggleSelect, toggleSelectAll } =
    useCartStore();
  const cartSpotlight = useSpotlight();
  const emptySpotlight = useSpotlight();

  const allSelected = items.length > 0 && items.every((i) => i.selected);

  const handleCheckout = () => {
    navigate('/checkout');
  };

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6">购物车</h1>

      {items.length === 0 ? (
        <div 
          ref={emptySpotlight.ref as React.RefObject<HTMLDivElement>}
          className="glass-liquid rounded-2xl p-12 text-center overflow-hidden relative"
          style={emptySpotlight.spotlightStyle}
          {...emptySpotlight.handlers}
        >
          <p className="text-gray-500 mb-4">购物车是空的</p>
          <Link to="/products" className="text-primary hover:underline">
            去逛逛
          </Link>
        </div>
      ) : (
        <div 
          ref={cartSpotlight.ref as React.RefObject<HTMLDivElement>}
          className="glass-cart rounded-2xl overflow-hidden relative"
          style={cartSpotlight.spotlightStyle}
          {...cartSpotlight.handlers}
        >
          {/* 表头 */}
          <div className="grid grid-cols-12 gap-4 p-4 border-b bg-gray-50 font-medium text-sm">
            <div className="col-span-1">
              <input
                type="checkbox"
                checked={allSelected}
                onChange={toggleSelectAll}
                className="w-4 h-4"
              />
            </div>
            <div className="col-span-5">商品信息</div>
            <div className="col-span-2 text-center">单价</div>
            <div className="col-span-2 text-center">数量</div>
            <div className="col-span-1 text-center">小计</div>
            <div className="col-span-1 text-center">操作</div>
          </div>

          {/* 商品列表 */}
          {items.map((item) => (
            <div key={item.id} className="grid grid-cols-12 gap-4 p-4 border-b items-center">
              <div className="col-span-1">
                <input
                  type="checkbox"
                  checked={item.selected}
                  onChange={() => toggleSelect(item.id)}
                  className="w-4 h-4"
                />
              </div>
              <div className="col-span-5 flex space-x-4">
                <img
                  src={item.product.mainImage}
                  alt={item.product.name}
                  className="w-20 h-20 object-cover rounded"
                />
                <div>
                  <Link
                    to={`/products/${item.productId}`}
                    className="text-sm hover:text-primary line-clamp-2"
                  >
                    {item.product.name}
                  </Link>
                  {item.spec && (
                    <p className="text-xs text-gray-500 mt-1">
                      {item.spec.name}: {item.spec.value}
                    </p>
                  )}
                </div>
              </div>
              <div className="col-span-2 text-center">¥{item.product.price}</div>
              <div className="col-span-2 flex items-center justify-center space-x-2">
                <button
                  onClick={() =>
                    updateQuantity(item.id, Math.max(1, item.quantity - 1))
                  }
                  className="w-6 h-6 border rounded flex items-center justify-center hover:bg-gray-100"
                >
                  -
                </button>
                <span className="w-8 text-center">{item.quantity}</span>
                <button
                  onClick={() =>
                    updateQuantity(
                      item.id,
                      Math.min(item.product.stock, item.quantity + 1)
                    )
                  }
                  className="w-6 h-6 border rounded flex items-center justify-center hover:bg-gray-100"
                >
                  +
                </button>
              </div>
              <div className="col-span-1 text-center font-bold text-primary">
                ¥{(item.product.price * item.quantity).toFixed(2)}
              </div>
              <div className="col-span-1 text-center">
                <button
                  onClick={() => removeItem(item.id)}
                  className="text-gray-400 hover:text-error"
                >
                  <Trash2 className="w-5 h-5" />
                </button>
              </div>
            </div>
          ))}

          {/* 结算栏 */}
          <div className="p-4 flex items-center justify-between">
            <div className="flex items-center space-x-4">
              <input
                type="checkbox"
                checked={allSelected}
                onChange={toggleSelectAll}
                className="w-4 h-4"
              />
              <span className="text-sm">全选</span>
            </div>
            <div className="flex items-center space-x-6">
              <div className="text-sm">
                已选择{' '}
                <span className="font-bold text-primary">
                  {items.filter((i) => i.selected).length}
                </span>{' '}
                件商品
              </div>
              <div className="text-sm">
                合计：
                <span className="text-2xl font-bold text-primary">
                  ¥{total.toFixed(2)}
                </span>
              </div>
              <button
                onClick={handleCheckout}
                disabled={items.filter((i) => i.selected).length === 0}
                className="px-12 py-3 bg-primary text-white rounded-lg hover:bg-primary-hover disabled:bg-gray-300 disabled:cursor-not-allowed"
              >
                去结算
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};
