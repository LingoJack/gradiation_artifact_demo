# 前端设计文档

## 一、项目结构

```
frontend/
├── public/
│   └── favicon.ico
├── src/
│   ├── api/              # API 接口封装
│   │   ├── auth.ts       # 认证相关 API
│   │   ├── product.ts    # 商品相关 API
│   │   ├── cart.ts       # 购物车相关 API
│   │   ├── order.ts      # 订单相关 API
│   │   └── user.ts       # 用户相关 API
│   ├── components/       # 公共组件
│   │   ├── Layout/       # 布局组件
│   │   ├── Header/       # 头部组件
│   │   ├── Footer/       # 底部组件
│   │   ├── ProductCard/  # 商品卡片
│   │   ├── SearchBar/    # 搜索栏
│   │   └── Loading/      # 加载组件
│   ├── pages/            # 页面组件
│   │   ├── Home/         # 首页
│   │   ├── Login/        # 登录页
│   │   ├── Register/     # 注册页
│   │   ├── ProductList/  # 商品列表页
│   │   ├── ProductDetail/# 商品详情页
│   │   ├── Cart/         # 购物车页
│   │   ├── Order/        # 订单页
│   │   ├── OrderDetail/  # 订单详情页
│   │   ├── UserCenter/   # 用户中心页
│   │   └── Checkout/     # 结算页
│   ├── store/            # 状态管理
│   │   ├── useUserStore.ts    # 用户状态
│   │   ├── useCartStore.ts    # 购物车状态
│   │   └── useProductStore.ts # 商品状态
│   ├── hooks/            # 自定义 Hooks
│   │   ├── useAuth.ts    # 认证 Hook
│   │   └── useRequest.ts # 请求 Hook
│   ├── utils/            # 工具函数
│   │   ├── request.ts    # Axios 封装
│   │   ├── storage.ts    # 本地存储
│   │   └── format.ts     # 格式化工具
│   ├── types/            # TypeScript 类型定义
│   │   ├── user.ts
│   │   ├── product.ts
│   │   ├── cart.ts
│   │   └── order.ts
│   ├── styles/           # 全局样式
│   │   └── index.css
│   ├── App.tsx           # 根组件
│   ├── main.tsx          # 入口文件
│   └── vite-env.d.ts     # Vite 类型定义
├── index.html
├── package.json
├── tsconfig.json
├── tailwind.config.js
└── vite.config.ts
```

## 二、页面组件列表

### 2.1 主要页面
| 页面名称 | 路由路径 | 功能描述 |
|---------|---------|---------|
| 首页 | `/` | 轮播广告、分类导航、热门推荐 |
| 登录 | `/login` | 用户登录 |
| 注册 | `/register` | 用户注册 |
| 商品列表 | `/products` | 商品搜索、筛选、列表展示 |
| 商品详情 | `/products/:id` | 商品详细信息、规格选择、加入购物车 |
| 购物车 | `/cart` | 购物车管理、结算 |
| 结算 | `/checkout` | 订单确认、地址选择、支付 |
| 订单列表 | `/orders` | 我的订单（全部/待付款/待发货/待收货/待评价）|
| 订单详情 | `/orders/:id` | 订单详细信息、物流跟踪 |
| 用户中心 | `/user` | 个人信息、地址管理、收藏、浏览历史 |
| 搜索结果 | `/search` | 搜索结果展示 |

### 2.2 公共组件
| 组件名称 | 功能描述 |
|---------|---------|
| Layout | 页面布局容器（Header + Content + Footer）|
| Header | 顶部导航栏（Logo、搜索框、购物车、用户菜单）|
| Footer | 底部信息栏 |
| ProductCard | 商品卡片（图片、标题、价格、销量）|
| SearchBar | 搜索栏（输入框、搜索按钮）|
| CategoryNav | 分类导航组件 |
| Carousel | 轮播图组件 |
| Pagination | 分页组件 |
| Loading | 加载动画 |
| Empty | 空状态提示 |
| Modal | 模态框 |

## 三、状态管理方案

使用 **Zustand** 作为状态管理库（轻量级、易于使用）

### 3.1 用户状态 (useUserStore)
```typescript
interface UserState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  login: (user: User, token: string) => void;
  logout: () => void;
  updateUser: (user: User) => void;
}
```

### 3.2 购物车状态 (useCartStore)
```typescript
interface CartState {
  items: CartItem[];
  total: number;
  addItem: (item: CartItem) => void;
  removeItem: (id: string) => void;
  updateQuantity: (id: string, quantity: number) => void;
  toggleSelect: (id: string) => void;
  toggleSelectAll: () => void;
  clearCart: () => void;
}
```

### 3.3 商品状态 (useProductStore)
```typescript
interface ProductState {
  products: Product[];
  currentProduct: Product | null;
  loading: boolean;
  setProducts: (products: Product[]) => void;
  setCurrentProduct: (product: Product) => void;
}
```

## 四、UI 库选择

### 主要 UI 库
**Tailwind CSS** + **Headless UI** + **Lucide React Icons**

### 选择理由
- **Tailwind CSS**: 原子化 CSS，开发效率高，体积小
- **Headless UI**: 无样式的可访问性组件，可与 Tailwind 完美结合
- **Lucide React**: 现代化的图标库，图标丰富

### 备选方案
- **Ant Design**: 功能全面的组件库（如果需要更丰富的组件）

## 五、路由规划

使用 **React Router v6**

```typescript
const routes = [
  {
    path: '/',
    element: <Layout />,
    children: [
      { index: true, element: <Home /> },
      { path: 'login', element: <Login /> },
      { path: 'register', element: <Register /> },
      { path: 'products', element: <ProductList /> },
      { path: 'products/:id', element: <ProductDetail /> },
      { path: 'cart', element: <Cart />, auth: true },
      { path: 'checkout', element: <Checkout />, auth: true },
      { path: 'orders', element: <OrderList />, auth: true },
      { path: 'orders/:id', element: <OrderDetail />, auth: true },
      { path: 'user', element: <UserCenter />, auth: true },
      { path: 'search', element: <SearchResult /> },
    ]
  }
]
```

### 路由守卫
- 需要登录的页面添加 `auth: true` 标记
- 使用高阶组件检查登录状态，未登录跳转到登录页

## 六、API 接口定义

### 6.1 认证接口
```typescript
POST   /api/auth/register      // 用户注册
POST   /api/auth/login         // 用户登录
POST   /api/auth/logout        // 用户登出
GET    /api/auth/profile       // 获取用户信息
PUT    /api/auth/profile       // 更新用户信息
```

### 6.2 商品接口
```typescript
GET    /api/products           // 获取商品列表（支持分页、筛选）
GET    /api/products/:id       // 获取商品详情
GET    /api/products/hot       // 获取热门商品
GET    /api/products/search    // 搜索商品
GET    /api/categories         // 获取商品分类
```

### 6.3 购物车接口
```typescript
GET    /api/cart               // 获取购物车
POST   /api/cart               // 添加商品到购物车
PUT    /api/cart/:id           // 更新购物车商品
DELETE /api/cart/:id           // 删除购物车商品
```

### 6.4 订单接口
```typescript
POST   /api/orders             // 创建订单
GET    /api/orders             // 获取订单列表
GET    /api/orders/:id         // 获取订单详情
PUT    /api/orders/:id/cancel  // 取消订单
PUT    /api/orders/:id/pay     // 支付订单
PUT    /api/orders/:id/confirm // 确认收货
POST   /api/orders/:id/review  // 评价订单
```

### 6.5 用户接口
```typescript
GET    /api/user/addresses     // 获取收货地址列表
POST   /api/user/addresses     // 添加收货地址
PUT    /api/user/addresses/:id // 更新收货地址
DELETE /api/user/addresses/:id // 删除收货地址
GET    /api/user/favorites     // 获取收藏列表
POST   /api/user/favorites     // 添加收藏
DELETE /api/user/favorites/:id // 取消收藏
```

## 七、全局样式规范

### 7.1 主题色板
```css
/* 主色调 - 淘宝橙 */
--primary-color: #ff4400;
--primary-hover: #ff5722;
--primary-light: #fff3e0;

/* 辅助色 */
--success-color: #52c41a;
--warning-color: #faad14;
--error-color: #f5222d;
--info-color: #1890ff;

/* 中性色 */
--text-primary: #262626;
--text-secondary: #8c8c8c;
--text-disabled: #bfbfbf;
--border-color: #e8e8e8;
--background-color: #f5f5f5;
```

### 7.2 字体规范
```css
/* 字体家族 */
font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 
             'Hiragino Sans GB', 'Microsoft YaHei', sans-serif;

/* 字体大小 */
--font-size-xs: 12px;
--font-size-sm: 14px;
--font-size-base: 16px;
--font-size-lg: 18px;
--font-size-xl: 20px;
--font-size-2xl: 24px;
--font-size-3xl: 30px;
```

### 7.3 间距规范
```css
--spacing-xs: 4px;
--spacing-sm: 8px;
--spacing-md: 16px;
--spacing-lg: 24px;
--spacing-xl: 32px;
--spacing-2xl: 48px;
```

### 7.4 圆角规范
```css
--radius-sm: 4px;
--radius-md: 8px;
--radius-lg: 12px;
--radius-xl: 16px;
--radius-full: 9999px;
```

### 7.5 阴影规范
```css
--shadow-sm: 0 1px 2px rgba(0, 0, 0, 0.05);
--shadow-md: 0 4px 6px rgba(0, 0, 0, 0.1);
--shadow-lg: 0 10px 15px rgba(0, 0, 0, 0.1);
--shadow-xl: 0 20px 25px rgba(0, 0, 0, 0.1);
```

## 八、响应式断点

使用 Tailwind CSS 默认断点：

```css
sm: 640px   /* 手机横屏 */
md: 768px   /* 平板竖屏 */
lg: 1024px  /* 平板横屏/小型笔记本 */
xl: 1280px  /* 桌面显示器 */
2xl: 1536px /* 大型显示器 */
```

### 响应式设计原则
- **移动优先**: 默认样式针对移动端，使用断点逐步增强
- **关键断点**: 重点优化 `md` (768px) 和 `lg` (1024px) 两个断点
- **弹性布局**: 使用 Flexbox 和 Grid 实现弹性布局

## 九、图标系统

使用 **Lucide React** 图标库

### 常用图标
```typescript
import { 
  Search, ShoppingBag, User, Heart, 
  Menu, X, ChevronRight, Star,
  Plus, Minus, Trash2, Check,
  MapPin, Phone, Mail, Lock
} from 'lucide-react'
```

### 图标规范
- **大小**: 16px (sm), 20px (md), 24px (lg)
- **颜色**: 继承父元素文本颜色或使用主题色
- **线宽**: stroke-width = 2

## 十、组件库规范

### 10.1 命名规范
- **组件名**: PascalCase (如 `ProductCard`)
- **文件名**: PascalCase (如 `ProductCard.tsx`)
- **样式文件**: 与组件同名 (如 `ProductCard.css`)

### 10.2 组件结构
```typescript
// 组件目录结构
ComponentName/
├── index.tsx          # 组件主文件
├── ComponentName.tsx  # 组件实现
├── styles.css         # 组件样式
├── types.ts           # 类型定义
└── utils.ts           # 工具函数
```

### 10.3 Props 规范
```typescript
interface ComponentProps {
  // 必填属性放前面
  id: string;
  title: string;
  
  // 可选属性放后面
  className?: string;
  style?: React.CSSProperties;
  
  // 事件处理函数
  onClick?: (event: React.MouseEvent) => void;
}
```

### 10.4 样式规范
- 优先使用 Tailwind CSS utility classes
- 复杂样式使用 CSS Modules 或 styled-components
- 避免内联样式（动态样式除外）

## 十一、性能优化策略

### 11.1 代码分割
```typescript
// 路由级别懒加载
const ProductDetail = lazy(() => import('./pages/ProductDetail'))
const Checkout = lazy(() => import('./pages/Checkout'))
```

### 11.2 图片优化
- 使用 WebP 格式
- 响应式图片（srcset）
- 图片懒加载（Intersection Observer）
- 使用占位图避免布局抖动

### 11.3 列表优化
- 虚拟滚动（react-window）处理长列表
- 分页加载避免一次性渲染过多数据

### 11.4 缓存策略
- 使用 React Query 进行数据缓存
- 静态资源使用浏览器缓存
- Service Worker 缓存关键资源

## 十二、开发工具配置

### 12.1 ESLint 配置
```json
{
  "extends": [
    "eslint:recommended",
    "plugin:react/recommended",
    "plugin:react-hooks/recommended",
    "plugin:@typescript-eslint/recommended"
  ],
  "rules": {
    "react/react-in-jsx-scope": "off",
    "@typescript-eslint/no-unused-vars": "warn"
  }
}
```

### 12.2 Prettier 配置
```json
{
  "semi": false,
  "singleQuote": true,
  "tabWidth": 2,
  "trailingComma": "es5",
  "printWidth": 100
}
```

### 12.3 VSCode 配置
```json
{
  "editor.formatOnSave": true,
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true
  }
}
```
