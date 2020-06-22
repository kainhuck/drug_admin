-- 查找　售出订单　最大　最小　平均值
CREATE PROCEDURE sale_mam_total_price(
	OUT min_price DECIMAL(8, 2),
	OUT max_price DECIMAL(8, 2),
	OUT avg_price DECIMAL(8, 2)
)
BEGIN
	SELECT MIN(total_price) INTO min_price FROM sale_drug_total_price;
	SELECT MAX(total_price) INTO max_price FROM sale_drug_total_price;
	SELECT AVG(total_price) INTO avg_price FROM sale_drug_total_price;
END//

-- 查找　买入订单　最大　最小　平均值
CREATE PROCEDURE buy_mam_total_price(
	OUT min_price DECIMAL(8, 2),
	OUT max_price DECIMAL(8, 2),
	OUT avg_price DECIMAL(8, 2)
)
BEGIN
	SELECT MIN(total_price) INTO min_price FROM buy_drug_total_price;
	SELECT MAX(total_price) INTO max_price FROM buy_drug_total_price;
	SELECT AVG(total_price) INTO avg_price FROM buy_drug_total_price;
END//

-- 查找某一个售出订单的药品种类
CREATE PROCEDURE sale_drug_type_num(
	IN order_nb INT,
	OUT num DECIMAL(8, 2)
)
BEGIN
	SELECT COUNT(*) INTO num
	FROM drug_sale_order_detail dsod
	WHERE dsod.order_id = order_nb;
END//

-- 查找某一个进货订单的药品种类
CREATE PROCEDURE buy_drug_type_num(
	IN order_nb INT,
	OUT num DECIMAL(8, 2)
)
BEGIN
	SELECT COUNT(*) INTO num
	FROM buy_order_with_supplier_drug bowsd
	WHERE bowsd.drug_buy_order_id = order_nb;
END//

-- 查找某一个员工处理的单子数量
CREATE PROCEDURE empolyee_deal_num(
	IN eid INT,
	OUT num DECIMAL(8, 2)
)
BEGIN
	SELECT SUM(sdl.num)
	FROM sale_drug_list sdl , drug_sale_order dso 
	WHERE dso.employee_id = eid AND dso.drug_sale_order_id = sdl.drug_sale_order_id ;
END//