import React from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { Clock, Trash2, ArrowLeft, ShoppingBag } from 'lucide-react';
import { useBrowseHistoryStore } from '../../store/useBrowseHistoryStore';
import { showToast } from '../../utils/toast';

export const BrowseHistory: React.FC = () => {
  const navigate = useNavigate();
  const { items, removeItem, clearAll } = useBrowseHistoryStore();

  const handleClearAll = () => {
    clearAll();
    showToast('浏览历史已清空', 'success');
  };

  const handleRemove = (id: string, e: React.MouseEvent) => {
    e.preventDefault();
    e.stopPropagation();
    removeItem(id);
    showToast('已删除', 'success');
  };

  // Group by date
  const groupedByDate = items.reduce<Record<string, typeof items>>((acc, item) => {
    const date = new Date(item.visitedAt);
    const today = new Date();
    const yesterday = new Date(today);
    yesterday.setDate(yesterday.getDate() - 1);

    let dateKey: string;
    if (date.toDateString() === today.toDateString()) {
      dateKey = '今天';
    } else if (date.toDateString() === yesterday.toDateString()) {
      dateKey = '昨天';
    } else {
      dateKey = date.toLocaleDateString('zh-CN', {
        month: 'long',
        day: 'numeric',
      });
    }

    if (!acc[dateKey]) {
      acc[dateKey] = [];
    }
    acc[dateKey].push(item);
    return acc;
  }, {});

  return (
    <div className="container py-8">
      {/* Header */}
      <div className="flex items-center justify-between mb-6">
        <div className="flex items-center gap-3">
          <button
            onClick={() => navigate(-1)}
            className="p-2 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg transition-colors"
          >
            <ArrowLeft className="w-5 h-5 text-gray-600 dark:text-gray-400" />
          </button>
          <h1 className="text-2xl font-bold dark:text-white">浏览历史</h1>
          <span className="text-sm text-gray-400">({items.length}条)</span>
        </div>
        {items.length > 0 && (
          <button
            onClick={handleClearAll}
            className="flex items-center gap-1 text-sm text-gray-500 dark:text-gray-400 hover:text-red-500 dark:hover:text-red-400 transition-colors"
          >
            <Trash2 className="w-4 h-4" />
            清空全部
          </button>
        )}
      </div>

      {/* Empty state */}
      {items.length === 0 ? (
        <div className="glass-card rounded-xl p-16 text-center">
          <Clock className="w-20 h-20 text-gray-300 dark:text-gray-600 mx-auto mb-4" />
          <p className="text-gray-500 dark:text-gray-400 text-lg mb-2">暂无浏览记录</p>
          <p className="text-gray-400 dark:text-gray-500 text-sm mb-6">
            去逛逛，发现更多好物吧
          </p>
          <Link
            to="/products"
            className="inline-flex items-center gap-2 px-6 py-2.5 bg-primary text-white rounded-lg hover:bg-primary-hover transition-colors font-medium"
          >
            <ShoppingBag className="w-4 h-4" />
            去逛逛
          </Link>
        </div>
      ) : (
        /* Grouped by date */
        <div className="space-y-8">
          {Object.entries(groupedByDate).map(([dateKey, dateItems]) => (
            <div key={dateKey}>
              <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400 mb-3 flex items-center gap-2">
                <Clock className="w-4 h-4" />
                {dateKey}
              </h3>
              <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
                {dateItems.map((item) => (
                  <div
                    key={item.id}
                    className="glass-card rounded-xl overflow-hidden relative group"
                  >
                    <Link to={`/products/${item.id}`} className="block">
                      <div className="aspect-square bg-gray-100 dark:bg-gray-700 overflow-hidden">
                        <img
                          src={item.mainImage}
                          alt={item.name}
                          className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
                          onError={(e) => {
                            const target = e.target as HTMLImageElement;
                            target.src =
                              'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="200" height="200"%3E%3Crect fill="%23f3f4f6" width="200" height="200"/%3E%3Ctext fill="%239ca3af" x="50%25" y="50%25" text-anchor="middle" dy=".3em" font-size="12"%3E暂无图片%3C/text%3E%3C/svg%3E';
                          }}
                        />
                      </div>
                      <div className="p-3">
                        <p className="text-sm line-clamp-2 text-gray-800 dark:text-gray-200 group-hover:text-primary transition-colors">
                          {item.name}
                        </p>
                        <p className="text-primary font-bold mt-1">¥{item.price}</p>
                      </div>
                    </Link>
                    {/* Delete button */}
                    <button
                      onClick={(e) => handleRemove(item.id, e)}
                      className="absolute top-2 right-2 w-7 h-7 bg-black/40 hover:bg-black/60 rounded-full flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity"
                    >
                      <Trash2 className="w-3.5 h-3.5 text-white" />
                    </button>
                  </div>
                ))}
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};
