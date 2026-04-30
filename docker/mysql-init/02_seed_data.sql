USE appdb;

-- 设置客户端字符集为 utf8mb4，确保中文正确存储
SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;

-- 测试用户（密码：password123）
INSERT INTO users (username, password, nickname, phone, email, avatar, gender, status) VALUES
('testuser', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZRGdjGj/n3.QoBfFqBqFqBqFqBqFq', '测试用户', '13800138000', 'test@example.com', 'https://picsum.photos/100', 1, 1);

-- 店铺
INSERT INTO shops (name, description, avatar, cover_image, rating, sales, fans, location) VALUES
('官方旗舰店', '品质保证，正品直销', 'https://picsum.photos/200', 'https://picsum.photos/800', 4.90, 10000, 1000, '广东深圳'),
('潮流服饰店', '时尚前沿，潮流必备', 'https://picsum.photos/200', 'https://picsum.photos/800', 4.70, 5000, 500, '浙江杭州');

-- 商品
INSERT INTO products (name, category_id, price, main_image, images, description, stock, sales, status) VALUES
('iPhone 15 Pro Max', 18, 9999.00, 'https://picsum.photos/400?random=1', '["https://picsum.photos/400?random=1","https://picsum.photos/400?random=2"]', '最新款苹果手机，钛金属边框', 100, 500, 1),
('MacBook Pro 14寸', 19, 14999.00, 'https://picsum.photos/400?random=3', '["https://picsum.photos/400?random=3","https://picsum.photos/400?random=4"]', 'M3芯片，专业级笔记本', 50, 200, 1),
('时尚T恤', 12, 99.00, 'https://picsum.photos/400?random=5', '["https://picsum.photos/400?random=5","https://picsum.photos/400?random=6"]', '纯棉材质，舒适透气', 500, 1000, 1),
('连衣裙', 11, 199.00, 'https://picsum.photos/400?random=7', '["https://picsum.photos/400?random=7","https://picsum.photos/400?random=8"]', '优雅设计，气质出众', 300, 800, 1),
('运动鞋', 16, 399.00, 'https://picsum.photos/400?random=13', '["https://picsum.photos/400?random=13"]', '轻便舒适，运动必备', 200, 600, 1),
('休闲外套', 14, 299.00, 'https://picsum.photos/400?random=14', '["https://picsum.photos/400?random=14"]', '百搭外套，秋冬推荐', 150, 400, 1);

-- SKU
INSERT INTO product_skus (product_id, sku_code, spec_values, price, stock, image, status) VALUES
(1, 'IP15-BK-256', '{"颜色":"深空黑","容量":"256GB"}', 9999.00, 50, 'https://picsum.photos/400?random=1', 1),
(1, 'IP15-TI-512', '{"颜色":"原色钛金属","容量":"512GB"}', 11999.00, 30, 'https://picsum.photos/400?random=9', 1),
(2, 'MBP-SV-16', '{"颜色":"银色","内存":"16GB"}', 14999.00, 20, 'https://picsum.photos/400?random=3', 1),
(3, 'TSHIRT-W-M', '{"颜色":"白色","尺码":"M"}', 99.00, 100, 'https://picsum.photos/400?random=5', 1),
(3, 'TSHIRT-B-L', '{"颜色":"黑色","尺码":"L"}', 99.00, 100, 'https://picsum.photos/400?random=10', 1),
(5, 'SHOE-BK-42', '{"颜色":"黑色","尺码":"42"}', 399.00, 50, 'https://picsum.photos/400?random=13', 1),
(6, 'COAT-GY-L', '{"颜色":"灰色","尺码":"L"}', 299.00, 80, 'https://picsum.photos/400?random=14', 1);

-- Banner
INSERT INTO banners (title, image_url, link_url, sort_order, status) VALUES
('新品首发', 'https://picsum.photos/800?random=11', '/products/1', 1, 1),
('限时特惠', 'https://picsum.photos/800?random=12', '/products', 2, 1),
('品牌专区', 'https://picsum.photos/800?random=15', '/products', 3, 1);

-- 优惠券
INSERT INTO coupons (name, discount, min_spend, total, claimed, start_time, end_time, status) VALUES
('新人专享', 50.00, 100.00, 1000, 0, NOW(), DATE_ADD(NOW(), INTERVAL 30 DAY), 1),
('满减券', 100.00, 500.00, 500, 0, NOW(), DATE_ADD(NOW(), INTERVAL 60 DAY), 1),
('数码专享', 200.00, 1000.00, 200, 0, NOW(), DATE_ADD(NOW(), INTERVAL 15 DAY), 1);

-- 收货地址
INSERT INTO user_addresses (user_id, receiver_name, receiver_phone, province, city, district, detail_address, is_default) VALUES
(1, '张三', '13800138000', '广东省', '深圳市', '南山区', '科技园路1号', 1);
