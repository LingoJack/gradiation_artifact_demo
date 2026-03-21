import React, { useState } from 'react';
import { Bell, Lock, Shield, Moon, Sun, Globe, ChevronRight } from 'lucide-react';
import { useTheme } from '../../hooks/useTheme';
import { useSpotlight } from '../../hooks/useSpotlight';

interface SettingItem {
  id: string;
  icon: React.ReactNode;
  title: string;
  description: string;
  action: 'toggle' | 'link' | 'select';
  value?: boolean | string;
}

interface SettingGroupProps {
  title: string;
  items: SettingItem[];
  onAction: (item: SettingItem) => void;
}

// 独立的设置组卡片，每个实例有自己的 spotlight
const SettingGroupCard: React.FC<SettingGroupProps> = ({ title, items, onAction }) => {
  const cardSpotlight = useSpotlight();

  return (
    <div
      ref={cardSpotlight.ref as React.RefObject<HTMLDivElement>}
      className="glass-liquid rounded-xl overflow-hidden relative"
      style={cardSpotlight.spotlightStyle}
      {...cardSpotlight.handlers}
    >
      <div className="px-6 py-4 border-b border-gray-200 dark:border-gray-700">
        <h2 className="font-bold dark:text-white">{title}</h2>
      </div>
      <div className="divide-y divide-gray-200 dark:divide-gray-700">
        {items.map((item) => (
          <div
            key={item.id}
            onClick={() => onAction(item)}
            className="flex items-center justify-between px-6 py-4 hover:bg-gray-50 dark:hover:bg-gray-700/50 cursor-pointer transition"
          >
            <div className="flex items-center space-x-4">
              <div className="text-gray-400 dark:text-gray-500">{item.icon}</div>
              <div>
                <p className="font-medium dark:text-white">{item.title}</p>
                <p className="text-sm text-gray-500 dark:text-gray-400">{item.description}</p>
              </div>
            </div>
            <div>
              {item.action === 'toggle' && (
                <button
                  className={`relative w-12 h-6 rounded-full transition-colors ${
                    item.value ? 'bg-primary' : 'bg-gray-300 dark:bg-gray-600'
                  }`}
                >
                  <span
                    className={`absolute top-1 w-4 h-4 bg-white rounded-full transition-transform ${
                      item.value ? 'left-7' : 'left-1'
                    }`}
                  />
                </button>
              )}
              {item.action === 'select' && (
                <span className="text-gray-500 dark:text-gray-400">{item.value}</span>
              )}
              {item.action === 'link' && (
                <ChevronRight className="w-5 h-5 text-gray-400 dark:text-gray-500" />
              )}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export const Settings: React.FC = () => {
  const { setTheme, isDark } = useTheme();
  const [settings, setSettings] = useState({
    notifications: true,
    emailNotify: true,
    smsNotify: false,
    language: 'zh-CN',
  });

  const handleToggle = (key: string) => {
    setSettings({
      ...settings,
      [key]: !settings[key as keyof typeof settings],
    });
  };

  const toggleTheme = () => {
    setTheme(isDark ? 'light' : 'dark');
  };

  const settingGroups = [
    {
      title: '外观设置',
      items: [
        {
          id: 'theme',
          icon: isDark ? <Moon className="w-5 h-5" /> : <Sun className="w-5 h-5" />,
          title: '深色模式',
          description: isDark ? '已开启' : '已关闭',
          action: 'toggle' as const,
          value: isDark,
        },
        {
          id: 'language',
          icon: <Globe className="w-5 h-5" />,
          title: '语言',
          description: '简体中文',
          action: 'select' as const,
          value: settings.language,
        },
      ],
    },
    {
      title: '通知设置',
      items: [
        {
          id: 'notifications',
          icon: <Bell className="w-5 h-5" />,
          title: '推送通知',
          description: '接收订单状态、促销活动等通知',
          action: 'toggle' as const,
          value: settings.notifications,
        },
        {
          id: 'emailNotify',
          icon: <Bell className="w-5 h-5" />,
          title: '邮件通知',
          description: '通过邮件接收重要通知',
          action: 'toggle' as const,
          value: settings.emailNotify,
        },
        {
          id: 'smsNotify',
          icon: <Bell className="w-5 h-5" />,
          title: '短信通知',
          description: '通过短信接收订单状态更新',
          action: 'toggle' as const,
          value: settings.smsNotify,
        },
      ],
    },
    {
      title: '安全设置',
      items: [
        {
          id: 'password',
          icon: <Lock className="w-5 h-5" />,
          title: '修改密码',
          description: '定期修改密码可以保护账户安全',
          action: 'link' as const,
        },
        {
          id: 'security',
          icon: <Shield className="w-5 h-5" />,
          title: '账户安全',
          description: '查看登录记录、绑定手机等',
          action: 'link' as const,
        },
      ],
    },
  ];

  const handleAction = (item: SettingItem) => {
    if (item.id === 'theme') {
      toggleTheme();
    } else if (item.action === 'toggle') {
      handleToggle(item.id);
    } else if (item.action === 'link') {
      // TODO: 跳转到对应页面
      alert('功能开发中...');
    }
  };

  return (
    <div className="container py-8">
      <h1 className="text-2xl font-bold mb-6 dark:text-white">账户设置</h1>

      <div className="space-y-6">
        {settingGroups.map((group) => (
          <SettingGroupCard
            key={group.title}
            title={group.title}
            items={group.items}
            onAction={handleAction}
          />
        ))}
      </div>

      {/* 关于信息 */}
      <div className="mt-8 text-center text-sm text-gray-500 dark:text-gray-400">
        <p>淘宝克隆版 v1.0.0</p>
        <p className="mt-1">仅供学习交流使用</p>
      </div>
    </div>
  );
};
